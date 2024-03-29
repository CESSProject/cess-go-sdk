/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"errors"
	"fmt"
	"log"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/centrifuge/go-substrate-rpc-client/v4/xxhash"
	"github.com/decred/base58"
	"golang.org/x/crypto/blake2b"
)

var RPC_ADDRS = []string{
	//testnet
	"wss://testnet-rpc0.cess.cloud/ws/",
	"wss://testnet-rpc1.cess.cloud/ws/",
	"wss://testnet-rpc2.cess.cloud/ws/",
}

func main() {
	accounts, err := GetAllAccountInfoFromBlock(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(accounts), " : ", accounts)
}

func GetAllAccountInfoFromBlock(block int) ([]types.AccountInfo, error) {
	var (
		err error
		c   *gsrpc.SubstrateAPI
	)
	for i := 0; i < len(RPC_ADDRS); i++ {
		c, err = gsrpc.NewSubstrateAPI(RPC_ADDRS[i])
		if err == nil {
			break
		}
	}
	if c == nil {
		log.Fatal("all rpc addresses are unavailable")
	}

	defer c.Client.Close()

	var data []types.AccountInfo

	key := createPrefixedKey("System", "Account")

	// get all account information from the latest block
	if block < 0 {
		keys, err := c.RPC.State.GetKeysLatest(key)
		if err != nil {
			return nil, err
		}
		set, err := c.RPC.State.QueryStorageAtLatest(keys)
		if err != nil {
			return nil, err
		}
		for _, elem := range set {
			for _, change := range elem.Changes {
				var storageData types.AccountInfo
				if err := codec.Decode(change.StorageData, &storageData); err != nil {
					fmt.Println("Decode StorageData:", err)
					continue
				}
				var storageKey types.AccountID
				if err := codec.Decode(change.StorageKey, &storageKey); err != nil {
					fmt.Println("Decode StorageKey:", err)
					continue
				}
				fmt.Println(encodePublicKeyAsAccount(storageKey[:]))
				data = append(data, storageData)
			}
		}
		return data, nil
	}

	// get all account information from the block

	blockhash, err := c.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return data, err
	}

	keys, err := c.RPC.State.GetKeys(key, blockhash)
	if err != nil {
		return nil, err
	}
	set, err := c.RPC.State.QueryStorageAt(keys, blockhash)
	if err != nil {
		return nil, err
	}
	for _, elem := range set {
		for _, change := range elem.Changes {
			if change.HasStorageData {
				var storageData types.AccountInfo
				if err := codec.Decode(change.StorageData, &storageData); err != nil {
					fmt.Println("Decode StorageData err:", err)
					continue
				}
				var storageKey types.AccountID
				if err := codec.Decode(change.StorageKey, &storageKey); err != nil {
					fmt.Println("Decode StorageKey err:", err)
					continue
				}
				fmt.Println(encodePublicKeyAsAccount(storageKey[:]))
				data = append(data, storageData)
			}
		}
	}
	return data, nil
}

func createPrefixedKey(pallet, method string) []byte {
	return append(xxhash.New128([]byte(pallet)).Sum(nil), xxhash.New128([]byte(method)).Sum(nil)...)
}

func encodePublicKeyAsAccount(publicKey []byte) (string, error) {
	if len(publicKey) != 32 {
		return "", errors.New("invalid public key")
	}
	payload := appendBytes([]byte{0x50, 0xac}, publicKey)
	input := appendBytes([]byte{0x53, 0x53, 0x35, 0x38, 0x50, 0x52, 0x45}, payload)
	ck := blake2b.Sum512(input)
	checkum := ck[:2]
	address := base58.Encode(appendBytes(payload, checkum))
	if address == "" {
		return address, errors.New("public key encoding failed")
	}
	return address, nil
}

func appendBytes(data1, data2 []byte) []byte {
	if data2 == nil {
		return data1
	}
	return append(data1, data2...)
}
