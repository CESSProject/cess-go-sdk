/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/CESSProject/sdk-go/core/erasure"
	"github.com/CESSProject/sdk-go/core/rule"
)

func (c *Cli) GetFile(roothash, dir string) (string, error) {
	var (
		segmentspath = make([]string, 0)
	)
	userfile := filepath.Join(dir, roothash)
	_, err := os.Stat(userfile)
	if err == nil {
		return userfile, nil
	}
	os.MkdirAll(dir, rule.DirMode)
	f, err := os.Create(userfile)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fmeta, err := c.Chain.GetFileMetaInfo(roothash)
	if err != nil {
		return "", err
	}

	defer func(basedir string) {
		for _, segment := range fmeta.SegmentList {
			os.Remove(filepath.Join(basedir, string(segment.Hash[:])))
			for _, fragment := range segment.FragmentList {
				os.Remove(filepath.Join(basedir, string(fragment.Hash[:])))
			}
		}
	}(dir)

	for _, segment := range fmeta.SegmentList {
		fragmentpaths := make([]string, 0)
		for _, fragment := range segment.FragmentList {
			miner, err := c.Chain.QueryStorageMiner(fragment.Miner[:])
			if err != nil {
				return "", err
			}
			peerid, err := c.AddMultiaddrToPearstore(string(miner.Ip), time.Hour)
			if err != nil {
				return "", err
			}
			fragmentpath := filepath.Join(dir, string(fragment.Hash[:]))
			err = c.Protocol.ReadFileAction(peerid, roothash, string(fragment.Hash[:]), fragmentpath, rule.FragmentSize)
			if err != nil {
				continue
			}
			fragmentpaths = append(fragmentpaths, fragmentpath)
			segmentpath := filepath.Join(dir, string(segment.Hash[:]))
			if len(fragmentpaths) >= rule.DataShards {
				err = erasure.ReedSolomon_Restore(segmentpath, fragmentpaths)
				if err != nil {
					return "", err
				}
				segmentspath = append(segmentspath, segmentpath)
				break
			}
		}
	}

	if len(segmentspath) != len(fmeta.SegmentList) {
		return "", fmt.Errorf("Download failed")
	}
	var writecount = 0
	for i := 0; i < len(fmeta.SegmentList); i++ {
		for j := 0; j < len(segmentspath); j++ {
			if string(fmeta.SegmentList[i].Hash[:]) == filepath.Base(segmentspath[j]) {
				buf, err := os.ReadFile(segmentspath[j])
				if err != nil {
					return "", err
				}
				f.Write(buf)
				writecount++
				break
			}
		}
	}
	if writecount != len(fmeta.SegmentList) {
		return "", fmt.Errorf("Write failed")
	}
	return userfile, nil
}
