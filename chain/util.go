/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

func BytesToFileHash(val []byte) (FileHash, error) {
	if len(val) != FileHashLen {
		return FileHash{}, errors.New("[BytesToFileHash] invalid length")
	}
	var filehash FileHash
	for k, v := range val {
		filehash[k] = types.U8(v)
	}
	return filehash, nil
}

func BytesToWorkPublickey(val []byte) (WorkerPublicKey, error) {
	if len(val) != WorkerPublicKeyLen {
		return WorkerPublicKey{}, errors.New("[BytesToWorkPublickey] invalid length")
	}
	var pubkey WorkerPublicKey
	for k, v := range val {
		pubkey[k] = types.U8(v)
	}
	return pubkey, nil
}

func BytesToPoISKeyInfo(g, n []byte) (PoISKeyInfo, error) {
	if len(g) != PoISKeyLen || len(n) != PoISKeyLen {
		return PoISKeyInfo{}, errors.New("[BytesToPoISKeyInfo] invalid length")
	}
	var poisKey PoISKeyInfo
	for i := 0; i < PoISKeyLen; i++ {
		poisKey.G[i] = types.U8(g[i])
		poisKey.N[i] = types.U8(n[i])
	}
	return poisKey, nil
}
func BytesToBloomFilter(val []byte) (BloomFilter, error) {
	if len(val) != BloomFilterLen {
		return BloomFilter{}, errors.New("[BytesToBloomFilter] invalid length")
	}
	var bloomfilter BloomFilter
	for i := 0; i < BloomFilterLen; i++ {
		bloomfilter[i] = types.U64(val[i])
	}
	return bloomfilter, nil
}

func IsWorkerPublicKeyAllZero(puk WorkerPublicKey) bool {
	for i := 0; i < WorkerPublicKeyLen; i++ {
		if puk[i] != 0 {
			return false
		}
	}
	return true
}
