/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/CESSProject/sdk-go/core/chain"
	"github.com/CESSProject/sdk-go/core/erasure"
	"github.com/CESSProject/sdk-go/core/hashtree"
	"github.com/CESSProject/sdk-go/core/rule"
	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type SegmentInfo struct {
	SegmentHash  string
	FragmentHash []string
}

func (c *Cli) PutFile(owner []byte, segmentInfo []SegmentInfo, roothash, filename, bucketname string) (uint8, error) {
	var err error
	var storageOrder chain.StorageOrder

	_, err = c.Chain.QueryFileMetadata(roothash)
	if err == nil {
		return 0, nil
	}

	for i := 0; i < 3; i++ {
		storageOrder, err = c.Chain.QueryStorageOrder(roothash)
		if err != nil {
			if err.Error() == chain.ERR_Empty {
				err = c.GenerateStorageOrder(roothash, segmentInfo, owner, filename, bucketname)
				if err != nil {
					return 0, err
				}
			}
			time.Sleep(rule.BlockInterval)
			continue
		}
		break
	}
	if err != nil {
		return 0, err
	}

	// store fragment to storage
	err = c.StorageData(roothash, segmentInfo, storageOrder.AssignedMiner)
	if err != nil {
		return 0, err
	}
	return uint8(storageOrder.Count), nil
}

func (c *Cli) ProcessingData(path string) ([]SegmentInfo, string, error) {
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
		return nil, "", errors.New("Not a file")
	}

	baseDir := filepath.Dir(path)
	segmentCount = fstat.Size() / rule.SegmentSize
	if fstat.Size()%int64(rule.SegmentSize) != 0 {
		segmentCount++
	}

	segment := make([]SegmentInfo, segmentCount)

	buf := make([]byte, rule.SegmentSize)

	f, err = os.Open(path)
	if err != nil {
		return segment, "", err
	}
	defer f.Close()

	for i := int64(0); i < segmentCount; i++ {
		f.Seek(rule.SegmentSize*i, 0)
		num, err = f.Read(buf)
		if err != nil && err != io.EOF {
			return segment, "", err
		}
		if num == 0 {
			break
		}
		if num < rule.SegmentSize {
			if i+1 != segmentCount {
				return segment, "", fmt.Errorf("Error reading %s", path)
			}
			remainbuf := make([]byte, rule.SegmentSize-num)
			copy(buf[num:], remainbuf)
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

func (c *Cli) GenerateStorageOrder(roothash string, segment []SegmentInfo, owner []byte, filename, buckname string) error {
	var err error
	var segmentList = make([]chain.SegmentList, len(segment))
	var user chain.UserBrief
	for i := 0; i < len(segment); i++ {
		hash := filepath.Base(segment[i].SegmentHash)
		for k := 0; k < len(hash); k++ {
			segmentList[i].SegmentHash[k] = types.U8(hash[k])
		}
		segmentList[i].FragmentHash = make([]chain.FileHash, len(segment[i].FragmentHash))
		for j := 0; j < len(segment[i].FragmentHash); j++ {
			hash := filepath.Base(segment[i].FragmentHash[j])
			for k := 0; k < len(hash); k++ {
				segmentList[i].FragmentHash[j][k] = types.U8(hash[k])
			}
		}
	}
	acc, err := types.NewAccountID(owner)
	if err != nil {
		return err
	}
	user.User = *acc
	user.BucketName = types.NewBytes([]byte(buckname))
	user.FileName = types.NewBytes([]byte(filename))
	_, err = c.Chain.UploadDeclaration(roothash, segmentList, user)
	return err
}

func ExtractSegmenthash(segment []SegmentInfo) []string {
	var segmenthash = make([]string, len(segment))
	for i := 0; i < len(segment); i++ {
		segmenthash[i] = segment[i].SegmentHash
	}
	return segmenthash
}

func (c *Cli) StorageData(roothash string, segment []SegmentInfo, minerTaskList []chain.MinerTaskList) error {
	var err error

	// query all assigned miner multiaddr
	multiaddrs, err := c.QueryAssignedMiner(minerTaskList)
	if err != nil {
		return err
	}

	basedir := filepath.Dir(segment[0].FragmentHash[0])
	for i := 0; i < len(multiaddrs); i++ {
		peerid, err := c.Protocol.AddMultiaddrToPearstore(multiaddrs[i], 0)
		if err != nil {
			return err
		}

		for j := 0; j < len(minerTaskList[i].Hash); j++ {
			err = c.Protocol.WriteFileAction(peerid, roothash, filepath.Join(basedir, string(minerTaskList[i].Hash[j][:])))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Cli) QueryAssignedMiner(minerTaskList []chain.MinerTaskList) ([]string, error) {
	var multiaddrs = make([]string, len(minerTaskList))
	for i := 0; i < len(minerTaskList); i++ {
		minerInfo, err := c.Chain.QueryStorageMiner(minerTaskList[i].Account[:])
		if err != nil {
			return multiaddrs, err
		}
		multiaddrs[i] = string(minerInfo.PeerPuk[:])
	}
	return multiaddrs, nil
}
