/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"errors"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

func BytesToFileHash(val []byte) (pattern.FileHash, error) {
	if len(val) != pattern.FileHashLen {
		return pattern.FileHash{}, errors.New("[BytesToFileHash] invalid length")
	}
	var filehash pattern.FileHash
	for k, v := range val {
		filehash[k] = types.U8(v)
	}
	return filehash, nil
}

func BytesToWorkPublickey(val []byte) (pattern.WorkerPublicKey, error) {
	if len(val) != pattern.WorkerPublicKeyLen {
		return pattern.WorkerPublicKey{}, errors.New("[BytesToWorkPublickey] invalid length")
	}
	var pubkey pattern.WorkerPublicKey
	for k, v := range val {
		pubkey[k] = types.U8(v)
	}
	return pubkey, nil
}

func BytesToPoISKeyInfo(g, n []byte) (pattern.PoISKeyInfo, error) {
	if len(g) != pattern.PoISKeyLen || len(n) != pattern.PoISKeyLen {
		return pattern.PoISKeyInfo{}, errors.New("[BytesToPoISKeyInfo] invalid length")
	}
	var poisKey pattern.PoISKeyInfo
	for i := 0; i < pattern.PoISKeyLen; i++ {
		poisKey.G[i] = types.U8(g[i])
		poisKey.N[i] = types.U8(n[i])
	}
	return poisKey, nil
}
func BytesToBloomFilter(val []byte) (pattern.BloomFilter, error) {
	if len(val) != pattern.BloomFilterLen {
		return pattern.BloomFilter{}, errors.New("[BytesToBloomFilter] invalid length")
	}
	var bloomfilter pattern.BloomFilter
	for i := 0; i < pattern.BloomFilterLen; i++ {
		bloomfilter[i] = types.U64(val[i])
	}
	return bloomfilter, nil
}

func IsWorkerPublicKeyAllZero(puk pattern.WorkerPublicKey) bool {
	for i := 0; i < pattern.WorkerPublicKeyLen; i++ {
		if puk[i] != 0 {
			return false
		}
	}
	return true
}
