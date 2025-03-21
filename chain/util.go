/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/AstaFrode/go-substrate-rpc-client/v4/types"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
	"github.com/vedhavyas/go-subkey/v2"
	"golang.org/x/crypto/blake2b"
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

func RrscAppPublicToByte(public RrscAppPublic) types.Bytes {
	var result = make(types.Bytes, RrscAppPublicLen)
	for i := 0; i < RrscAppPublicLen; i++ {
		result[i] = byte(public[i])
	}
	return result
}

// H160ToSS58 convert Eth account to polkadot account
//   - origin: eth account
//   - chain_id: chain id, CESS chain id is 11330
//
// Return:
//   - string: polkadot account
//   - error: error message
func H160ToSS58(origin string, chain_id uint16) (string, error) {
	origin = strings.TrimPrefix(origin, "0x")
	decode_origin, err := hex.DecodeString(origin)
	if err != nil {
		log.Error("[CESS-GO-SDK][H160 convert to SS58] Error")
		return "", err
	}

	temp_data := []byte("evm:")
	data := append(temp_data, decode_origin...)

	hashed := blake2b.Sum256(data)

	new_acc := subkey.SS58Encode(hashed[:], chain_id)

	return new_acc, nil
}

var re = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)

func CheckBucketName(name string) bool {
	if len(name) < int(MinBucketNameLength) || len(name) > int(MaxBucketNameLength) {
		return false
	}

	if !re.MatchString(name) {
		return false
	}

	if strings.Contains(name, " ") {
		return false
	}

	if strings.Count(name, ".") > 2 {
		return false
	}

	if byte(name[0]) == byte('.') ||
		byte(name[0]) == byte('-') ||
		byte(name[0]) == byte('_') ||
		byte(name[len(name)-1]) == byte('.') ||
		byte(name[len(name)-1]) == byte('-') ||
		byte(name[len(name)-1]) == byte('_') {
		return false
	}

	if utils.IsIPv4(name) {
		return false
	}

	if utils.IsIPv6(name) {
		return false
	}

	return true
}
