/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/CESSProject/cess-go-sdk/core/erasure"
	"github.com/CESSProject/cess-go-sdk/core/hashtree"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/utils"
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

	// Calculate merkle hash tree
	hTree, err := hashtree.NewHashTree(ExtractSegmenthash(segmentDataInfo))
	if err != nil {
		return segmentDataInfo, "", err
	}

	return segmentDataInfo, hex.EncodeToString(hTree.MerkleRoot()), nil
}

func cutfile(file string) ([]string, error) {
	fstat, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	if fstat.IsDir() {
		return nil, errors.New("not a file")
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

func ExtractSegmenthash(segment []pattern.SegmentDataInfo) []string {
	var segmenthash = make([]string, len(segment))
	for i := 0; i < len(segment); i++ {
		segmenthash[i] = segment[i].SegmentHash
	}
	return segmenthash
}

func (c *ChainSDK) RedundancyRecovery(outpath string, shardspath []string) error {
	return erasure.ReedSolomonRestore(outpath, shardspath)
}

func (c *ChainSDK) StoreFile(owner []byte, file string, bucket string) (string, error) {
	if !c.enabledP2P {
		return "", errors.New("P2P network not enabled")
	}

	fstat, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	if fstat.IsDir() {
		return "", errors.New("not a file")
	}

	// verify mem availability
	mem, err := utils.GetSysMemAvailable()
	if err == nil {
		if fstat.Size() > int64(mem*80/100) {
			return "", errors.New("unsupported file size")
		}
	}

	if !utils.CompareSlice(owner, c.GetSignatureAccPulickey()) {
		return "", errors.New("account error")
	}

	if !utils.CheckBucketName(bucket) {
		return "", errors.New("invalid bucket name")
	}

	userInfo, err := c.QueryUserSpaceSt(owner)
	if err != nil {
		if err.Error() == pattern.ERR_Empty {
			return "", errors.New("no space in the account")
		}
		return "", errors.Wrapf(err, "[QueryUserSpaceSt]")
	}

	blockheight, err := c.QueryBlockHeight("")
	if err != nil {
		return "", errors.Wrapf(err, "[QueryBlockHeight]")
	}

	if userInfo.Deadline < (blockheight + 30) {
		return "", errors.Wrapf(err, "account space expires soon")
	}

	segmentDataInfo, roothash, err := c.ProcessingData(file)
	if err != nil {
		return "", errors.Wrapf(err, "[ProcessingData]")
	}

	var storageOrder pattern.StorageOrder

	_, err = c.QueryFileMetadata(roothash)
	if err == nil {
		return roothash, nil
	}

	storageOrder, err = c.QueryStorageOrder(roothash)
	if err != nil {
		if err.Error() != pattern.ERR_Empty {
			return roothash, errors.Wrapf(err, "[QueryStorageOrder]")
		}
		_, err = c.GenerateStorageOrder(roothash, segmentDataInfo, owner, fstat.Name(), bucket, uint64(fstat.Size()))
		if err != nil {
			return roothash, errors.Wrapf(err, "[GenerateStorageOrder]")
		}
		time.Sleep(pattern.BlockInterval)
	}

	// store data to storage node
	err = c.StorageData(roothash, segmentDataInfo, storageOrder.AssignedMiner)
	if err != nil {
		return roothash, errors.Wrapf(err, "[StorageData]")
	}
	return roothash, nil
}

func (c *ChainSDK) RetrieveFile(roothash, savepath string) error {
	if !c.enabledP2P {
		return errors.New("P2P network not enabled")
	}

	fmeta, err := c.QueryFileMetadata(roothash)
	if err != nil {
		return errors.Wrapf(err, "[QueryFileMetadata]")
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
			peerid, _ := peer.Decode(string(miner.PeerId[:]))
			fragmentpath := filepath.Join(baseDir, string(fragment.Hash[:]))
			err = c.ReadFileAction(peerid, roothash, string(fragment.Hash[:]), fragmentpath, pattern.FragmentSize)
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
