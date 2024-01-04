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

	"github.com/CESSProject/cess-go-sdk/core/crypte"
	"github.com/CESSProject/cess-go-sdk/core/erasure"
	"github.com/CESSProject/cess-go-sdk/core/hashtree"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/pkg/errors"
)

// Process the file according to CESS specifications.
//
// Receive parameter:
//   - file: the file to be processed.
//
// Return parameter:
//   - segmentDataInfo: segment and fragment information of the file.
//   - string: [fid] unique identifier for the file.
//   - error: error message.
func ProcessingData(file string) ([]pattern.SegmentDataInfo, string, error) {
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

// Process the file according to CESS specifications.
//
// Receive parameter:
//   - file: the file to be processed.
//   - cipher: encryption and decryption keys.
//
// Return parameter:
//   - segmentDataInfo: segment and fragment information of the file.
//   - string: [fid] unique identifier for the file.
//   - error: error message.
func ShardedEncryptionProcessing(file string, cipher string) ([]pattern.SegmentDataInfo, string, error) {
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
			copy(buf[num:], make([]byte, pattern.SegmentSize-num, pattern.SegmentSize-num))
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
			copy(buf[num:], make([]byte, segmentSize-num, segmentSize-num))
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

func ExtractSegmenthash(segment []string) []string {
	var segmenthash = make([]string, len(segment))
	for i := 0; i < len(segment); i++ {
		segmenthash[i] = filepath.Base(segment[i])
	}
	return segmenthash
}
