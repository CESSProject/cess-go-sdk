/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/vedhavyas/go-subkey/scale"
)

// Substrate well-known mnemonic:
//
//	https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
var MY_MNEMONIC = "bottom drive obey lake curtain smoke basket hold race lonely fit walk"

var RPC_ADDRS = []string{
	//devnet
	"wss://devnet-rpc.cess.cloud/ws/",
	//testnet
	// "wss://testnet-rpc0.cess.cloud/ws/",
	// "wss://testnet-rpc1.cess.cloud/ws/",
	// "wss://testnet-rpc2.cess.cloud/ws/",
}

func main() {
	sdk, err := cess.New(
		context.Background(),
		cess.ConnectRpcAddrs(RPC_ADDRS),
		cess.Mnemonic(MY_MNEMONIC),
		cess.TransactionTimeout(time.Second*10),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(sdk.ChainVersion())
	blockhash, err := sdk.GetSubstrateAPI().RPC.Chain.GetBlockHash(140471)
	if err != nil {
		log.Fatalln(err)
	}

	// header, err := sdk.GetSubstrateAPI().RPC.Chain.GetHeader(blockhash)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("------------block header-----------")
	// fmt.Println(blockhash.Hex())
	// fmt.Println(header.ParentHash.Hex())
	// fmt.Println(header.ExtrinsicsRoot.Hex())
	// fmt.Println(header.Number)
	// fmt.Println(header.StateRoot.Hex())
	// for i := 0; i < len(header.Digest); i++ {
	// 	fmt.Println(i, ": ", header.Digest[i])
	// }
	// fmt.Println("------------block-----------")
	//

	block, err := sdk.GetSubstrateAPI().RPC.Chain.GetBlock(blockhash)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("number of Extrinsics: ", len(block.Block.Extrinsics))

	callIndex, err := sdk.GetMetadata().FindCallIndex("Timestamp.set")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("callIndex.MethodIndex:", callIndex.MethodIndex)
	fmt.Println("callIndex.SectionIndex:", callIndex.SectionIndex)
	timestamp := new(big.Int)
	for _, extrinsic := range block.Block.Extrinsics {
		if extrinsic.Method.CallIndex != callIndex {
			continue
		}
		timeDecoder := scale.NewDecoder(bytes.NewReader(extrinsic.Method.Args))
		timestamp, err = timeDecoder.DecodeUintCompact()
		if err != nil {
			log.Fatalln(err)
		}
		break
	}
	msec := timestamp.Int64()
	time := time.Unix(msec/1e3, (msec%1e3)*1e6)
	fmt.Println(msec)
	fmt.Println(time)
	fmt.Println(sdk.RetrieveAllEvent(blockhash))
}
