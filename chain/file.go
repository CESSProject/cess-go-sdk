/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/CESSProject/cess-go-sdk/core/erasure"
	"github.com/CESSProject/cess-go-sdk/core/hashtree"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
)

// ProcessingData
func (c *ChainSDK) ProcessingData(path string) ([]pattern.SegmentDataInfo, string, error) {
	var (
		err          error
		f            *os.File
		fstat        fs.FileInfo
		segmentCount int64
		num          int
	)

	fstat, err = os.Stat(path)
	if err != nil {
		return nil, "", err
	}
	if fstat.IsDir() {
		return nil, "", errors.New("not a file")
	}

	baseDir := filepath.Dir(path)
	segmentCount = fstat.Size() / pattern.SegmentSize
	if fstat.Size()%int64(pattern.SegmentSize) != 0 {
		segmentCount++
	}

	segment := make([]pattern.SegmentDataInfo, segmentCount)

	buf := make([]byte, pattern.SegmentSize)

	f, err = os.Open(path)
	if err != nil {
		return segment, "", err
	}
	defer f.Close()

	for i := int64(0); i < segmentCount; i++ {
		f.Seek(pattern.SegmentSize*i, 0)
		num, err = f.Read(buf)
		if err != nil && err != io.EOF {
			return segment, "", err
		}
		if num == 0 {
			break
		}
		if num < pattern.SegmentSize {
			if i+1 != segmentCount {
				return segment, "", fmt.Errorf("Error reading %s", path)
			}
			copy(buf[num:], []byte(utils.RandStr(pattern.SegmentSize-num)))
		}

		hash, err := utils.CalcSHA256(buf)
		if err != nil {
			return segment, "", err
		}

		segmentPath := filepath.Join(baseDir, hash)
		_, err = os.Stat(segmentPath)
		if err != nil {
			fsegment, err := os.Create(segmentPath)
			if err != nil {
				return segment, "", err
			}
			_, err = fsegment.Write(buf)
			if err != nil {
				fsegment.Close()
				return segment, "", err
			}
			err = fsegment.Sync()
			if err != nil {
				fsegment.Close()
				return segment, "", err
			}
			fsegment.Close()
		}

		segment[i].SegmentHash = segmentPath
		segment[i].FragmentHash, err = erasure.ReedSolomon(segmentPath)
		if err != nil {
			return segment, "", err
		}
	}

	segmenthash := ExtractSegmenthash(segment)

	// Calculate merkle hash tree
	hTree, err := hashtree.NewHashTree(segmenthash)
	if err != nil {
		return segment, "", err
	}

	return segment, hex.EncodeToString(hTree.MerkleRoot()), err
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

func (c *ChainSDK) StoreFile(roothash string, segmentInfo []pattern.SegmentDataInfo, user pattern.UserInfo) error {
	var err error
	var storageOrder pattern.StorageOrder

	_, err = c.QueryFileMetadata(roothash)
	if err == nil {
		return nil
	}

	pubkey, err := utils.ParsingPublickey(user.UserAccount)
	if err != nil {
		return errors.Wrapf(err, "[ParsingPublickey]")
	}

	storageOrder, err = c.QueryStorageOrder(roothash)
	if err != nil {
		if err.Error() == pattern.ERR_Empty {
			_, err = c.GenerateStorageOrder(roothash, segmentInfo, pubkey, user.FileName, user.BucketName, user.FileSize)
			if err != nil {
				return errors.Wrapf(err, "[GenerateStorageOrder]")
			}
		}
	}

	// store data to storage node
	err = c.StorageData(roothash, segmentInfo, storageOrder.AssignedMiner)
	if err != nil {
		return errors.Wrapf(err, "[StorageData]")
	}
	return nil
}

func (c *ChainSDK) RetrieveFile(roothash, savepath string) error {
	var (
		segmentspath = make([]string, 0)
	)

	var userfile = savepath
	var f *os.File
	fstat, err := os.Stat(savepath)
	if err != nil {
		f, err = os.Create(savepath)
		if err != nil {
			return errors.Wrapf(err, "[os.Create]")
		}
	} else {
		if fstat.IsDir() {
			userfile = filepath.Join(savepath, roothash)
		}
		f, err = os.Create(userfile)
		if err != nil {
			return errors.Wrapf(err, "[os.Create]")
		}
	}
	defer f.Close()

	fmeta, err := c.QueryFileMetadata(roothash)
	if err != nil {
		return errors.Wrapf(err, "[QueryFileMetadata]")
	}

	var baseDir = filepath.Dir(userfile)

	defer func(basedir string) {
		for _, segment := range fmeta.SegmentList {
			os.Remove(filepath.Join(basedir, string(segment.Hash[:])))
			for _, fragment := range segment.FragmentList {
				os.Remove(filepath.Join(basedir, string(fragment.Hash[:])))
			}
		}
	}(baseDir)

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
