/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/CESSProject/cess-go-sdk/core/crypte"
	"github.com/CESSProject/cess-go-sdk/core/erasure"
	"github.com/CESSProject/cess-go-sdk/core/hashtree"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	keyring "github.com/CESSProject/go-keyring"
	"github.com/btcsuite/btcutil/base58"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

// ProcessingData
func (c *chainClient) ProcessingData(file string) ([]pattern.SegmentDataInfo, string, error) {
	segmentPath, err := cutfile(file)
	if err != nil {
		if segmentPath != nil {
			for _, v := range segmentPath {
				os.Remove(v)
			}
		}
		return nil, "", errors.Wrapf(err, "[cutfile]")
	}

	var segmentDataInfo = make([]pattern.SegmentDataInfo, len(segmentPath))

	for i := 0; i < len(segmentPath); i++ {
		segmentDataInfo[i].SegmentHash = segmentPath[i]
		segmentDataInfo[i].FragmentHash, err = erasure.ReedSolomon(segmentPath[i])
		if err != nil {
			return segmentDataInfo, "", errors.Wrapf(err, "[ReedSolomon]")
		}
		os.Remove(segmentPath[i])
	}

	// calculate merkle root hash
	var hash string
	if len(segmentPath) == 1 {
		hash, err = hashtree.BuildSimpleMerkelRootHash(filepath.Base(segmentPath[0]))
		if err != nil {
			return nil, "", err
		}
	} else {
		hash, err = hashtree.BuildMerkelRootHash(ExtractSegmenthash(segmentPath))
		if err != nil {
			return nil, "", err
		}
	}

	return segmentDataInfo, hash, nil
}

// ProcessingData
func (c *chainClient) ShardedEncryptionProcessing(file string, cipher string) ([]pattern.SegmentDataInfo, string, error) {
	var err error
	var segmentPath []string
	if cipher != "" {
		segmentPath, err = cutFileWithEncryption(file)
		if err != nil {
			if segmentPath != nil {
				for _, v := range segmentPath {
					os.Remove(v)
				}
			}
			return nil, "", errors.Wrapf(err, "[cutFileWithEncryption]")
		}
		segmentPath, err = encryptedSegment(segmentPath, cipher)
		if err != nil {
			if segmentPath != nil {
				for _, v := range segmentPath {
					os.Remove(v)
				}
			}
			return nil, "", err
		}
	} else {
		segmentPath, err = cutfile(file)
		if err != nil {
			if segmentPath != nil {
				for _, v := range segmentPath {
					os.Remove(v)
				}
			}
			return nil, "", errors.Wrapf(err, "[cutFileWithEncryption]")
		}
	}

	var segmentDataInfo = make([]pattern.SegmentDataInfo, len(segmentPath))

	for i := 0; i < len(segmentPath); i++ {
		segmentDataInfo[i].SegmentHash = segmentPath[i]
		segmentDataInfo[i].FragmentHash, err = erasure.ReedSolomon(segmentPath[i])
		if err != nil {
			return segmentDataInfo, "", errors.Wrapf(err, "[ReedSolomon]")
		}
		os.Remove(segmentPath[i])
	}

	// calculate merkle root hash
	var hash string
	if len(segmentPath) == 1 {
		hash, err = hashtree.BuildSimpleMerkelRootHash(filepath.Base(segmentPath[0]))
		if err != nil {
			return nil, "", err
		}
	} else {
		hash, err = hashtree.BuildMerkelRootHash(ExtractSegmenthash(segmentPath))
		if err != nil {
			return nil, "", err
		}
	}

	return segmentDataInfo, hash, nil
}

func cutfile(file string) ([]string, error) {
	fstat, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	if fstat.IsDir() {
		return nil, errors.New("not a file")
	}
	if fstat.Size() == 0 {
		return nil, errors.New("empty file")
	}
	baseDir := filepath.Dir(file)
	segmentCount := fstat.Size() / pattern.SegmentSize
	if fstat.Size()%int64(pattern.SegmentSize) != 0 {
		segmentCount++
	}

	segment := make([]string, segmentCount)
	buf := make([]byte, pattern.SegmentSize)
	f, err := os.Open(file)
	if err != nil {
		return segment, err
	}
	defer f.Close()

	var num int
	for i := int64(0); i < segmentCount; i++ {
		f.Seek(pattern.SegmentSize*i, 0)
		num, err = f.Read(buf)
		if err != nil && err != io.EOF {
			return segment, err
		}
		if num == 0 {
			return segment, errors.New("read file is empty")
		}
		if num < pattern.SegmentSize {
			if i+1 != segmentCount {
				return segment, errors.New("read file err")
			}
			copy(buf[num:], []byte(utils.RandStr(pattern.SegmentSize-num)))
		}

		hash, err := utils.CalcSHA256(buf)
		if err != nil {
			return segment, err
		}

		err = utils.WriteBufToFile(buf, filepath.Join(baseDir, hash))
		if err != nil {
			return segment, errors.Wrapf(err, "[WriteBufToFile]")
		}
		segment[i] = filepath.Join(baseDir, hash)
	}
	return segment, nil
}

func cutFileWithEncryption(file string) ([]string, error) {
	fstat, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	if fstat.IsDir() {
		return nil, errors.New("not a file")
	}
	if fstat.Size() == 0 {
		return nil, errors.New("empty file")
	}
	baseDir := filepath.Dir(file)
	segmentSize := pattern.SegmentSize - 16
	segmentCount := fstat.Size() / int64(segmentSize)
	if fstat.Size()%int64(segmentSize) != 0 {
		segmentCount++
	}

	segment := make([]string, segmentCount)
	buf := make([]byte, segmentSize)
	f, err := os.Open(file)
	if err != nil {
		return segment, err
	}
	defer f.Close()

	var num int
	for i := int64(0); i < segmentCount; i++ {
		f.Seek(int64(segmentSize)*i, 0)
		num, err = f.Read(buf)
		if err != nil && err != io.EOF {
			return segment, err
		}
		if num == 0 {
			return segment, errors.New("read file is empty")
		}
		if num < segmentSize {
			if i+1 != segmentCount {
				return segment, errors.New("read file err")
			}
			copy(buf[num:], []byte(utils.RandStr(segmentSize-num)))
		}

		hash, err := utils.CalcSHA256(buf)
		if err != nil {
			return segment, err
		}

		err = utils.WriteBufToFile(buf, filepath.Join(baseDir, hash))
		if err != nil {
			return segment, errors.Wrapf(err, "[WriteBufToFile]")
		}
		segment[i] = filepath.Join(baseDir, hash)
	}
	return segment, nil
}

func encryptedSegment(segments []string, cipher string) ([]string, error) {
	var err error
	var hash string
	var buf []byte
	var encryptedSegmentPath string
	var encryptedSegments = make([]string, len(segments))
	for k, v := range segments {
		buf, err = os.ReadFile(v)
		if err != nil {
			return nil, err
		}
		buf, err = crypte.AesCbcEncrypt(buf, []byte(cipher))
		if err != nil {
			return nil, err
		}
		hash, err = utils.CalcSHA256(buf)
		if err != nil {
			return nil, err
		}
		encryptedSegmentPath = filepath.Join(filepath.Dir(v), hash)
		err = os.WriteFile(encryptedSegmentPath, buf, 0755)
		if err != nil {
			return nil, err
		}
		encryptedSegments[k] = encryptedSegmentPath
	}
	for _, v := range segments {
		os.Remove(v)
	}
	return encryptedSegments, nil
}

func (c *chainClient) GenerateStorageOrder(
	roothash string,
	segment []pattern.SegmentDataInfo,
	owner []byte,
	filename string,
	buckname string,
	filesize uint64,
) (string, error) {
	var err error
	var segmentList = make([]pattern.SegmentList, len(segment))
	var user pattern.UserBrief
	var assignedData = make([][]pattern.FileHash, len(segment))
	for i := 0; i < len(segment); i++ {
		hash := filepath.Base(segment[i].SegmentHash)
		for k := 0; k < len(hash); k++ {
			segmentList[i].SegmentHash[k] = types.U8(hash[k])
		}
		segmentList[i].FragmentHash = make([]pattern.FileHash, len(segment[i].FragmentHash))
		for j := 0; j < len(segment[i].FragmentHash); j++ {
			hash := filepath.Base(segment[i].FragmentHash[j])
			for k := 0; k < len(hash); k++ {
				segmentList[i].FragmentHash[j][k] = types.U8(hash[k])
			}
		}
	}
	for i := 0; i < len(segmentList); i++ {
		assignedData[i] = make([]pattern.FileHash, len(segmentList[i].FragmentHash))
		for j := 0; j < len(segmentList[i].FragmentHash); j++ {
			assignedData[i][j] = segmentList[i].FragmentHash[j]
		}
	}

	acc, err := types.NewAccountID(owner)
	if err != nil {
		return "", err
	}
	user.User = *acc
	user.BucketName = types.NewBytes([]byte(buckname))
	user.FileName = types.NewBytes([]byte(filename))
	return c.UploadDeclaration(roothash, segmentList, assignedData, user, filesize)
}

func ExtractSegmenthash(segment []string) []string {
	var segmenthash = make([]string, len(segment))
	for i := 0; i < len(segment); i++ {
		segmenthash[i] = filepath.Base(segment[i])
	}
	return segmenthash
}

func (c *chainClient) StoreFile(url, file, bucket string) (string, error) {
	fstat, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	if fstat.IsDir() {
		return "", errors.New("not a file")
	}

	if fstat.Size() == 0 {
		return "", errors.New("empty file")
	}

	if !utils.CheckBucketName(bucket) {
		return "", errors.New("invalid bucket name")
	}

	kr, _ := keyring.FromURI(c.GetURI(), keyring.NetSubstrate{})

	// sign message
	message := utils.GetRandomcode(16)
	sig, _ := kr.Sign(kr.SigningContext([]byte(message)))

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	//
	formFile, err := writer.CreateFormFile("file", fstat.Name())
	if err != nil {
		return "", err
	}

	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(formFile, f)
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("BucketName", bucket)
	req.Header.Set("Account", c.GetSignatureAcc())
	req.Header.Set("Message", message)
	req.Header.Set("Signature", base58.Encode(sig[:]))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		if len(respbody) > 0 {
			return "", errors.New(string(respbody))
		}
		return "", errors.New(fmt.Sprintf("upload failed, code: %d", resp.StatusCode))
	}

	return strings.TrimPrefix(strings.TrimSuffix(string(respbody), "\""), "\""), nil
}

func (c *chainClient) StoreObject(url string, reader io.Reader, bucket string) (string, error) {
	if !utils.CheckBucketName(bucket) {
		return "", errors.New("invalid bucket name")
	}

	kr, _ := keyring.FromURI(c.GetURI(), keyring.NetSubstrate{})

	// sign message
	message := utils.GetRandomcode(16)
	sig, _ := kr.Sign(kr.SigningContext([]byte(message)))

	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return "", err
	}

	req.Header.Set("BucketName", bucket)
	req.Header.Set("Account", c.GetSignatureAcc())
	req.Header.Set("Message", message)
	req.Header.Set("Signature", base58.Encode(sig[:]))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		if len(respbody) > 0 {
			return "", errors.New(string(respbody))
		}
		return "", errors.New(fmt.Sprintf("upload failed, code: %d", resp.StatusCode))
	}

	return strings.TrimPrefix(strings.TrimSuffix(string(respbody), "\""), "\""), nil
}

func (c *chainClient) RetrieveFile(url, fid, savepath string) error {
	fstat, err := os.Stat(savepath)
	if err == nil {
		if fstat.IsDir() {
			savepath = filepath.Join(savepath, fid)
		}
		if fstat.Size() > 0 {
			return nil
		}
	}

	if url == "" {
		return errors.New("empty url")
	}

	if url[len(url)-1] != byte(47) {
		url += "/"
	}

	f, err := os.Create(savepath)
	if err != nil {
		return err
	}
	defer f.Close()

	req, err := http.NewRequest(http.MethodGet, url+fid, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Operation", "download")

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed")
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (c *chainClient) RetrieveObject(url, fid string) (io.ReadCloser, error) {
	if url == "" {
		return nil, errors.New("empty url")
	}

	if url[len(url)-1] != byte(47) {
		url += "/"
	}

	req, err := http.NewRequest(http.MethodGet, url+fid, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Operation", "download")

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Retrieve failed, code: %d", resp.StatusCode))
	}

	return resp.Body, nil
}

// func (c *chainClient) StorageData(roothash string, segment []pattern.SegmentDataInfo, minerTaskList []pattern.MinerTaskList) error {
// 	if !c.enabledP2P {
// 		return errors.New("P2P network not enabled")
// 	}

// 	var err error

// 	// query all assigned miner multiaddr
// 	peerids, err := c.QueryAssignedMinerPeerId(minerTaskList)
// 	if err != nil {
// 		return errors.Wrapf(err, "[QueryAssignedMinerPeerId]")
// 	}

// 	basedir := filepath.Dir(segment[0].FragmentHash[0])
// 	for i := 0; i < len(peerids); i++ {
// 		for j := 0; j < len(minerTaskList[i].Hash); j++ {
// 			err = c.WriteFileAction(peerids[j], roothash, filepath.Join(basedir, string(minerTaskList[i].Hash[j][:])))
// 			if err != nil {
// 				return errors.Wrapf(err, "[WriteFileAction]")
// 			}
// 		}
// 	}

// 	return nil
// }

// func (c *chainClient) FindPeers() map[string]peer.AddrInfo {
// 	var peerMap = make(map[string]peer.AddrInfo, 100)
// 	timeOut := time.NewTicker(time.Second * 10)
// 	defer timeOut.Stop()
// 	c.RouteTableFindPeers(0)
// 	for {
// 		select {
// 		case peer, ok := <-c.GetDiscoveredPeers():
// 			if !ok {
// 				return peerMap
// 			}
// 			if len(peer.Responses) == 0 {
// 				break
// 			}
// 			for _, v := range peer.Responses {
// 				peerMap[v.ID.Pretty()] = *v
// 			}
// 		case <-timeOut.C:
// 			return peerMap
// 		}
// 	}
// }

// func (c *chainClient) QueryAssignedMinerPeerId(minerTaskList []pattern.MinerTaskList) ([]peer.ID, error) {
// 	var peerids = make([]peer.ID, len(minerTaskList))
// 	for i := 0; i < len(minerTaskList); i++ {
// 		minerInfo, err := c.QueryStorageMiner(minerTaskList[i].Account[:])
// 		if err != nil {
// 			return peerids, err
// 		}
// 		peerids[i], _ = peer.Decode(string(minerInfo.PeerId[:]))
// 	}
// 	return peerids, nil
// }
