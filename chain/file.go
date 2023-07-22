/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/CESSProject/cess-go-sdk/core/erasure"
	"github.com/CESSProject/cess-go-sdk/core/hashtree"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	keyring "github.com/CESSProject/go-keyring"
	"github.com/btcsuite/btcutil/base58"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/libp2p/go-libp2p/core/peer"
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

func (c *chainClient) GenerateStorageOrder(roothash string, segment []pattern.SegmentDataInfo, owner []byte, filename, buckname string, filesize uint64) (string, error) {
	var err error
	var segmentList = make([]pattern.SegmentList, len(segment))
	var user pattern.UserBrief
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
	acc, err := types.NewAccountID(owner)
	if err != nil {
		return "", err
	}
	user.User = *acc
	user.BucketName = types.NewBytes([]byte(buckname))
	user.FileName = types.NewBytes([]byte(filename))
	return c.UploadDeclaration(roothash, segmentList, user, filesize)
}

func ExtractSegmenthash(segment []string) []string {
	var segmenthash = make([]string, len(segment))
	for i := 0; i < len(segment); i++ {
		segmenthash[i] = filepath.Base(segment[i])
	}
	return segmenthash
}

func (c *chainClient) RedundancyRecovery(outpath string, shardspath []string) error {
	return erasure.ReedSolomonRestore(outpath, shardspath)
}

// StoreFile
func (c *chainClient) StoreFile(file, bucket string) (string, error) {
	return c.UploadtoGateway(pattern.PublicDeoss, file, bucket)
}

func (c *chainClient) RetrieveFile(roothash, savepath string) error {
	return c.DownloadFromGateway(pattern.PublicDeoss, roothash, savepath)
}

func (c *chainClient) UploadtoGateway(url, uploadfile, bucketName string) (string, error) {
	fstat, err := os.Stat(uploadfile)
	if err != nil {
		return "", err
	}

	if fstat.IsDir() {
		return "", errors.New("not a file")
	}

	if fstat.Size() == 0 {
		return "", errors.New("empty file")
	}

	if !utils.CheckBucketName(bucketName) {
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

	file, err := os.Open(uploadfile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(formFile, file)
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

	req.Header.Set("BucketName", bucketName)
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

	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		if len(respbody) > 0 {
			return "", errors.New(string(respbody))
		}
		return "", errors.New("Deoss service failure, please retry or contact administrator.")
	}

	return strings.TrimPrefix(strings.TrimSuffix(string(respbody), "\""), "\""), nil
}

func (c *chainClient) DownloadFromGateway(url, roothash, savepath string) error {
	fstat, err := os.Stat(savepath)
	if err == nil {
		if fstat.IsDir() {
			savepath = filepath.Join(savepath, roothash)
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

	req, err := http.NewRequest(http.MethodGet, url+roothash, nil)
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

func (c *chainClient) StorageData(roothash string, segment []pattern.SegmentDataInfo, minerTaskList []pattern.MinerTaskList) error {
	if !c.enabledP2P {
		return errors.New("P2P network not enabled")
	}

	var err error

	// query all assigned miner multiaddr
	peerids, err := c.QueryAssignedMinerPeerId(minerTaskList)
	if err != nil {
		return errors.Wrapf(err, "[QueryAssignedMinerPeerId]")
	}

	basedir := filepath.Dir(segment[0].FragmentHash[0])
	for i := 0; i < len(peerids); i++ {
		for j := 0; j < len(minerTaskList[i].Hash); j++ {
			err = c.WriteFileAction(peerids[j], roothash, filepath.Join(basedir, string(minerTaskList[i].Hash[j][:])))
			if err != nil {
				return errors.Wrapf(err, "[WriteFileAction]")
			}
		}
	}

	return nil
}

func (c *chainClient) FindPeers() map[string]peer.AddrInfo {
	var peerMap = make(map[string]peer.AddrInfo, 100)
	timeOut := time.NewTicker(time.Second * 10)
	defer timeOut.Stop()
	c.RouteTableFindPeers(0)
	for {
		select {
		case peer, ok := <-c.GetDiscoveredPeers():
			if !ok {
				return peerMap
			}
			if len(peer.Responses) == 0 {
				break
			}
			for _, v := range peer.Responses {
				peerMap[v.ID.Pretty()] = *v
			}
		case <-timeOut.C:
			return peerMap
		}
	}
}

func (c *chainClient) QueryAssignedMinerPeerId(minerTaskList []pattern.MinerTaskList) ([]peer.ID, error) {
	var peerids = make([]peer.ID, len(minerTaskList))
	for i := 0; i < len(minerTaskList); i++ {
		minerInfo, err := c.QueryStorageMiner(minerTaskList[i].Account[:])
		if err != nil {
			return peerids, err
		}
		peerids[i], _ = peer.Decode(string(minerInfo.PeerId[:]))
	}
	return peerids, nil
}
