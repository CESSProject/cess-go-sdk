/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package process

import (
	"io"
	"os"
	"path/filepath"

	"github.com/CESSProject/cess-go-sdk/chain"
	"github.com/CESSProject/cess-go-sdk/config"
	"github.com/CESSProject/cess-go-sdk/core/crypte"
	"github.com/CESSProject/cess-go-sdk/core/erasure"
	"github.com/CESSProject/cess-go-sdk/core/hashtree"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/pkg/errors"
)

// FillAndCut fill and cut files
//
// Receive parameter:
//   - file: the file to be processed
//   - saveDir: segment save directory
//
// Return parameter:
//   - []string: segment list
//   - error: error message
func FillAndCut(file string, saveDir string) ([]string, error) {
	fstat, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	if fstat.IsDir() {
		return nil, errors.New("FillingAndCutting: not a file")
	}
	if fstat.Size() <= 0 {
		return nil, errors.New("FillingAndCutting: file is empty")
	}
	err = os.MkdirAll(saveDir, 0755)
	if err != nil {
		return nil, errors.Wrap(err, "FillingAndCutting")
	}

	segmentCount := fstat.Size() / config.SegmentSize
	if fstat.Size()%int64(config.SegmentSize) != 0 {
		segmentCount++
	}

	segment := make([]string, segmentCount)
	buf := make([]byte, config.SegmentSize)
	f, err := os.Open(file)
	if err != nil {
		return segment, errors.Wrap(err, "FillingAndCutting")
	}
	defer f.Close()

	var num int
	for i := int64(0); i < segmentCount; i++ {
		f.Seek(config.SegmentSize*i, 0)
		num, err = f.Read(buf)
		if err != nil && err != io.EOF {
			return segment, err
		}
		if num == 0 {
			return segment, errors.New("read file is empty")
		}
		if num < config.SegmentSize {
			if i+1 != segmentCount {
				return segment, errors.New("read file failed")
			}
			copy(buf[num:], make([]byte, config.SegmentSize-num))
		}
		hash, err := utils.CalcSHA256(buf)
		if err != nil {
			return segment, err
		}
		err = utils.WriteBufToFile(buf, filepath.Join(saveDir, hash))
		if err != nil {
			return segment, errors.Wrapf(err, "[WriteBufToFile]")
		}
		segment[i] = filepath.Join(saveDir, hash)
	}
	return segment, nil
}

// FillAndCutWithAESEncryption fill and cut files, then encrypt using AES algorithm
//
// Receive parameter:
//   - file: the file to be processed
//   - cipher: encryption and decryption cipher
//   - saveDir: segment save directory
//
// Return parameter:
//   - []string: segment list
//   - error: error message
func FillAndCutWithAESEncryption(file string, cipher string, saveDir string) ([]string, error) {
	fstat, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	if fstat.IsDir() {
		return nil, errors.New("FillAndCutWithEncryption: not a file")
	}
	if fstat.Size() <= 0 {
		return nil, errors.New("FillAndCutWithEncryption: file is empty")
	}
	err = os.MkdirAll(saveDir, 0755)
	if err != nil {
		return nil, errors.Wrap(err, "FillAndCutWithEncryption")
	}

	segmentSize := config.SegmentSize - 16
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
				return segment, errors.New("read file failed")
			}
			copy(buf[num:], make([]byte, segmentSize-num))
		}
		hash, err := utils.CalcSHA256(buf)
		if err != nil {
			return segment, err
		}
		err = utils.WriteBufToFile(buf, filepath.Join(saveDir, hash))
		if err != nil {
			return segment, errors.Wrapf(err, "[WriteBufToFile]")
		}
		segment[i] = filepath.Join(saveDir, hash)
	}
	segment_encrypted := make([]string, segmentCount)
	for i := int64(0); i < segmentCount; i++ {
		segment_encrypted[i], err = EncryptWithAES(segment[i], cipher, saveDir)
		if err != nil {
			return nil, err
		}
		os.Remove(segment[i])
	}
	return segment_encrypted, nil
}

// Redundancy calculate redundancy for files
//
// Receive parameter:
//   - segment: the file to be processed
//   - saveDir: fragment save directory
//
// Return parameter:
//   - []chain.SegmentDataInfo: segment info
//   - string: fid
//   - error: error message
func Redundancy(segment []string, saveDir string) ([]chain.SegmentDataInfo, string, error) {
	var (
		err         error
		segmentInfo = make([]chain.SegmentDataInfo, len(segment))
	)
	for i := 0; i < len(segment); i++ {
		segmentInfo[i].SegmentHash = filepath.Base(segment[i])
		segmentInfo[i].FragmentHash, err = erasure.ReedSolomon(segment[i], saveDir)
		if err != nil {
			return segmentInfo, "", errors.Wrap(err, "[ReedSolomon]")
		}
		os.Remove(segment[i])
	}
	// calculate merkle root hash
	var hash string
	if len(segment) == 1 {
		hash, err = hashtree.BuildSimpleMerkelRootHash(filepath.Base(segment[0]))
		if err != nil {
			return nil, "", errors.Wrap(err, "[BuildSimpleMerkelRootHash]")
		}
	} else {
		hash, err = hashtree.BuildMerkelRootHash(ExtractSegmenthash(segment))
		if err != nil {
			return nil, "", errors.Wrap(err, "[BuildMerkelRootHash]")
		}
	}
	return segmentInfo, hash, nil
}

// EncryptWithAES encrypt a file with AES
//
// Receive parameter:
//   - file: file
//   - cipher: encryption and decryption cipher
//   - saveDir: encrypted file save directory
//
// Return parameter:
//   - string: encrypted file
//   - error: error message
func EncryptWithAES(file string, cipher string, saveDir string) (string, error) {
	buf, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	buf, err = crypte.AesCbcEncrypt(buf, []byte(cipher))
	if err != nil {
		return "", err
	}
	hash, err := utils.CalcSHA256(buf)
	if err != nil {
		return "", err
	}
	encryptedPath := filepath.Join(saveDir, hash)
	err = os.WriteFile(encryptedPath, buf, 0755)
	if err != nil {
		return "", err
	}
	return encryptedPath, nil
}

// BatchEncryptWithAES encrypt a batch of files with AES
//
// Receive parameter:
//   - files: file list
//   - cipher: encryption and decryption cipher
//   - saveDir: encrypted files save directory
//
// Return parameter:
//   - string: encrypted files
//   - error: error message
func BatchEncryptWithAES(files []string, cipher string, saveDir string) ([]string, error) {
	var err error
	var hash string
	var buf []byte
	var encryptedFiles = make([]string, len(files))
	for k, v := range files {
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
		encryptedPath := filepath.Join(saveDir, hash)
		err = os.WriteFile(encryptedPath, buf, 0755)
		if err != nil {
			return nil, err
		}
		encryptedFiles[k] = encryptedPath
	}
	for _, v := range files {
		os.Remove(v)
	}
	return encryptedFiles, nil
}

// FullProcessing perform full process processing on the file
//
// Receive parameter:
//   - file: the file to be processed
//   - cipher: encryption and decryption cipher
//   - saveDir: saved directory after processing
//
// Return parameter:
//   - []segmentDataInfo: segment and fragment information of the file
//   - string: [fid] unique identifier for the file
//   - error: error message
func FullProcessing(file string, cipher string, savedir string) ([]chain.SegmentDataInfo, string, error) {
	var err error
	var segmentList []string
	if cipher != "" {
		segmentList, err = FillAndCutWithAESEncryption(file, cipher, savedir)
		if err != nil {
			for _, v := range segmentList {
				os.Remove(v)
			}
			return nil, "", err
		}
		return Redundancy(segmentList, savedir)
	}
	segmentList, err = FillAndCut(file, savedir)
	if err != nil {
		for _, v := range segmentList {
			os.Remove(v)
		}
		return nil, "", err
	}
	return Redundancy(segmentList, savedir)
}

func ExtractSegmenthash(segment []string) []string {
	var segmenthash = make([]string, len(segment))
	for i := 0; i < len(segment); i++ {
		segmenthash[i] = filepath.Base(segment[i])
	}
	return segmenthash
}
