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

func (n *ChainSDK) RedundancyRecovery(outpath string, shardspath []string) error {
	return erasure.ReedSolomonRestore(outpath, shardspath)
}
