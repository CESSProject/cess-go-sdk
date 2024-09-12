/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package erasure

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/CESSProject/cess-go-sdk/chain"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/klauspost/reedsolomon"
)

// ReedSolomon uses reed-solomon algorithm to redundancy file
//
// Receive parameter:
//   - file: the file to be processed
//   - saveDir: save directory
//
// Return parameter:
//   - []string: processed data fragmentation
//   - error: error message.
func ReedSolomon(file string, saveDir string) ([]string, error) {
	var shardspath = make([]string, 0)
	fstat, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	if fstat.IsDir() {
		return nil, errors.New("not a file")
	}
	if fstat.Size() != chain.SegmentSize {
		return nil, errors.New("invalid size")
	}

	enc, err := reedsolomon.New(chain.DataShards, chain.ParShards)
	if err != nil {
		return shardspath, err
	}

	b, err := os.ReadFile(file)
	if err != nil {
		return shardspath, err
	}

	// Split the file into equally sized shards.
	shards, err := enc.Split(b)
	if err != nil {
		return shardspath, err
	}
	// Encode parity
	err = enc.Encode(shards)
	if err != nil {
		return shardspath, err
	}
	// Write out the resulting files.
	for _, shard := range shards {
		hash, err := utils.CalcSHA256(shard)
		if err != nil {
			return shardspath, err
		}
		newpath := filepath.Join(saveDir, hash)
		_, err = os.Stat(newpath)
		if err != nil {
			err = os.WriteFile(newpath, shard, 0755)
			if err != nil {
				return shardspath, err
			}
		}
		shardspath = append(shardspath, newpath)
	}
	return shardspath, nil
}

// Restore files from shards and save to outpath.
//
// Receive parameter:
//   - outpath: file save location.
//   - shardspath: file fragments.
//
// Return parameter:
//   - error: error message.
func RSRestore(outpath string, shardspath []string) error {
	_, err := os.Stat(outpath)
	if err == nil {
		return nil
	}

	enc, err := reedsolomon.New(chain.DataShards, chain.ParShards)
	if err != nil {
		return err
	}

	shards := make([][]byte, chain.DataShards+chain.ParShards)
	for k, v := range shardspath {
		shards[k], err = os.ReadFile(v)
		if err != nil {
			shards[k] = nil
		}
	}

	// Verify the shards
	ok, _ := enc.Verify(shards)
	if !ok {
		err = enc.Reconstruct(shards)
		if err != nil {
			return err
		}
		ok, err = enc.Verify(shards)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("invalid shards")
		}
	}

	f, err := os.Create(outpath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = enc.Join(f, shards, len(shards[0])*chain.DataShards)
	return err
}

// Restore files from shards and save to outpath.
//
// Receive parameter:
//   - outpath: file save location.
//   - sharddata: file fragments data.
//
// Return parameter:
//   - error: error message.
func RSRestoreData(outpath string, sharddata [][]byte) error {
	_, err := os.Stat(outpath)
	if err == nil {
		return nil
	}

	datashards, parshards := chain.DataShards, chain.ParShards

	enc, err := reedsolomon.New(datashards, parshards)
	if err != nil {
		return err
	}

	shards := sharddata

	// Verify the shards
	ok, _ := enc.Verify(shards)
	if !ok {
		err = enc.Reconstruct(shards)
		if err != nil {
			return err
		}
		ok, err = enc.Verify(shards)
		if !ok {
			return err
		}
	}
	f, err := os.Create(outpath)
	if err != nil {
		return err
	}
	defer f.Close()
	err = enc.Join(f, shards, len(shards[0])*datashards)
	return err
}
