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
func (c *ChainSDK) ProcessingData(file string) ([]pattern.SegmentDataInfo, string, error) {
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

func (c *ChainSDK) GenerateStorageOrder(roothash string, segment []pattern.SegmentDataInfo, owner []byte, filename, buckname string, filesize uint64) (string, error) {
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

func (c *ChainSDK) RedundancyRecovery(outpath string, shardspath []string) error {
	return erasure.ReedSolomonRestore(outpath, shardspath)
}

// StoreFile
func (c *ChainSDK) StoreFile(file, bucket string) (string, error) {
	c.AuthorizeSpace(pattern.PublicDeossAccount)
	return c.UploadtoGateway(pattern.PublicDeoss, c.GetSignatureAcc(), file, bucket)
}

func (c *ChainSDK) RetrieveFile(roothash, savepath string) error {
	err := c.DownloadFromGateway(pattern.PublicDeoss, roothash, savepath)
	if err == nil {
		return nil
	}

	if !c.enabledP2P {
		return errors.New("P2P network not enabled")
	}

	fmeta, err := c.QueryFileMetadata(roothash)
	if err != nil {
		if err.Error() != pattern.ERR_Empty {
			return errors.Wrapf(err, "[QueryFileMetadata]")
		}
		return errors.New("Not Found")
	}

	var userfile = savepath
	var f *os.File
	fstat, err := os.Stat(userfile)
	if err != nil {
		f, err = os.Create(userfile)
		if err != nil {
			return errors.Wrapf(err, "[os.Create]")
		}
	} else {
		if fstat.IsDir() {
			userfile = filepath.Join(savepath, roothash)
		}
		if fstat.Size() == fmeta.FileSize.Int64() {
			return nil
		}
		f, err = os.Create(userfile)
		if err != nil {
			return errors.Wrapf(err, "[os.Create]")
		}
	}
	defer f.Close()

	var baseDir = filepath.Dir(userfile)
	defer func(basedir string) {
		for _, segment := range fmeta.SegmentList {
			os.Remove(filepath.Join(basedir, string(segment.Hash[:])))
			for _, fragment := range segment.FragmentList {
				os.Remove(filepath.Join(basedir, string(fragment.Hash[:])))
			}
		}
	}(baseDir)

	var segmentspath = make([]string, 0)
	for _, segment := range fmeta.SegmentList {
		fragmentpaths := make([]string, 0)
		for _, fragment := range segment.FragmentList {
			miner, err := c.QueryStorageMiner(fragment.Miner[:])
			if err != nil {
				return errors.Wrapf(err, "[QueryStorageMiner]")
			}

			addr, err := c.DHTFindPeer(base58.Encode([]byte(string(miner.PeerId[:]))))
			if err != nil {
				return errors.Wrapf(err, "[DHTFindPeer]")
			}

			err = c.Connect(c.GetCtxQueryFromCtxCancel(), addr)
			if err != nil {
				return errors.Wrapf(err, "[Connect]")
			}

			fragmentpath := filepath.Join(baseDir, string(fragment.Hash[:]))
			err = c.ReadFileAction(addr.ID, roothash, string(fragment.Hash[:]), fragmentpath, pattern.FragmentSize)
			if err != nil {
				continue
			}

			fragmentpaths = append(fragmentpaths, fragmentpath)
			segmentpath := filepath.Join(baseDir, string(segment.Hash[:]))
			if len(fragmentpaths) >= pattern.DataShards {
				err = c.RedundancyRecovery(segmentpath, fragmentpaths)
				if err != nil {
					return errors.Wrapf(err, "[RedundancyRecovery]")
				}
				segmentspath = append(segmentspath, segmentpath)
				break
			}
		}
	}

	if len(segmentspath) != len(fmeta.SegmentList) {
		return errors.New("retrieve failed")
	}
	var writecount = 0
	for i := 0; i < len(fmeta.SegmentList); i++ {
		for j := 0; j < len(segmentspath); j++ {
			if string(fmeta.SegmentList[i].Hash[:]) == filepath.Base(segmentspath[j]) {
				buf, err := os.ReadFile(segmentspath[j])
				if err != nil {
					return errors.Wrapf(err, "[ReadFile]")
				}
				if (writecount + 1) >= len(fmeta.SegmentList) {
					f.Write(buf[:(fmeta.FileSize.Uint64() - uint64(writecount*pattern.SegmentSize))])
				} else {
					f.Write(buf)
				}
				writecount++
				break
			}
		}
	}
	if writecount != len(fmeta.SegmentList) {
		return errors.New("retrieve failed")
	}
	return nil
}

func (c *ChainSDK) UploadtoGateway(url, account, uploadfile, bucketName string) (string, error) {
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

	if account != c.GetSignatureAcc() {
		return "", errors.New("account error")
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
	_, err = io.Copy(formFile, file)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Account", account)
	req.Header.Set("Message", message)
	req.Header.Set("Signature", base58.Encode(sig[:]))
	req.Header.Set("Content-Type", "multipart/form-data; boundary=<calculated when request is sent>")

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

	return string(respbody), nil
}

func (c *ChainSDK) DownloadFromGateway(url, roothash, savepath string) error {
	_, err := os.Stat(savepath)
	if err == nil {
		return nil
	}

	f, err := os.Create(savepath)
	if err != nil {
		return err
	}
	defer f.Close()

	req, err := http.NewRequest(http.MethodGet, filepath.Join(url, roothash), nil)
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

	_, err = io.Copy(f, req.Body)
	if err != nil {
		return err
	}

	return nil
}

func (c *ChainSDK) StorageData(roothash string, segment []pattern.SegmentDataInfo, minerTaskList []pattern.MinerTaskList) error {
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

func (c *ChainSDK) QueryAssignedMinerPeerId(minerTaskList []pattern.MinerTaskList) ([]peer.ID, error) {
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
