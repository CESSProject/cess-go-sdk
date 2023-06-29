/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package hashtree

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"

	"github.com/CESSProject/cess-go-sdk/core/utils"
	"github.com/cbergoon/merkletree"
)

// HashTreeContent implements the Content interface provided by merkletree
// and represents the content stored in the tree.
type HashTreeContent struct {
	x string
}

// CalculateHash hashes the values of a HashTreeContent
func (t HashTreeContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.x)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (t HashTreeContent) Equals(other merkletree.Content) (bool, error) {
	return t.x == other.(HashTreeContent).x, nil
}

// NewHashTree build file to build hash tree
func NewHashTree(chunkPath []string) (*merkletree.MerkleTree, error) {
	if len(chunkPath) == 0 {
		return nil, errors.New("Empty data")
	}
	var list = make([]merkletree.Content, len(chunkPath))
	for i := 0; i < len(chunkPath); i++ {
		f, err := os.Open(chunkPath[i])
		if err != nil {
			return nil, err
		}
		temp, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}
		f.Close()
		list[i] = HashTreeContent{x: string(temp)}
	}

	//Create a new Merkle Tree from the list of Content
	return merkletree.NewTree(list)
}

// BuildMerkelRootHash
func BuildMerkelRootHash(segmentHash []string) (string, error) {
	if len(segmentHash) == 0 {
		return "", errors.New("empty segment hash")
	}

	if len(segmentHash) == 1 {
		return segmentHash[0], nil
	}

	var hashlist = make([]string, 0)
	for i := 0; i < len(segmentHash); i = i + 2 {
		if (i + 1) >= len(segmentHash) {
			b, err := hex.DecodeString(segmentHash[i])
			if err != nil {
				return "", err
			}
			hash, err := utils.CalcSHA256(append(b, b...))
			if err != nil {
				return "", err
			}
			hashlist = append(hashlist, hash)
		} else {
			b1, err := hex.DecodeString(segmentHash[i])
			if err != nil {
				return "", err
			}
			b2, err := hex.DecodeString(segmentHash[i+1])
			if err != nil {
				return "", err
			}
			hash, err := utils.CalcSHA256(append(b1, b2...))
			if err != nil {
				return "", err
			}
			hashlist = append(hashlist, hash)
		}
	}
	return BuildMerkelRootHash(hashlist)
}

// BuildSimpleMerkelRootHash
func BuildSimpleMerkelRootHash(segmentHash string) (string, error) {
	b, err := hex.DecodeString(segmentHash)
	if err != nil {
		return "", err
	}
	return utils.CalcSHA256(append(b, b...))
}
