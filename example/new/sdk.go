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
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// Substrate well-known mnemonic:
//
//   - cXgaee2N8E77JJv9gdsGAckv1Qsf3hqWYf7NL4q6ZuQzuAUtB
//   - https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
var MY_MNEMONIC = "bottom drive obey lake curtain smoke basket hold race lonely fit walk"

var RPC_ADDRS = []string{
	//devnet
	"wss://devnet-rpc.cess.cloud/ws/",
	//testnet
	//"wss://testnet-rpc0.cess.cloud/ws/",
	//"wss://testnet-rpc1.cess.cloud/ws/",
	//"wss://testnet-rpc2.cess.cloud/ws/",
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
	defer sdk.Close()

	fmt.Println(sdk.SystemVersion())
	fmt.Println(sdk.InitExtrinsicsName())
	fmt.Println(sdk.GetCurrentRpcAddr())

	fmt.Println(sdk.QueryValidatorsCount(-1))
	return

	pk, err := utils.ParsingPublickey("cXiKthh2dyY1taTydtdxiqQwXY1HKZcXvYGmjS2UmuPi2qNDS")
	if err != nil {
		panic(err)
	}
	fmt.Println(sdk.QueryCurrentCounters(pk, -1))
	return

	pk, err = utils.ParsingPublickey("cXfnLrW67qKkTn7DPXfh1SykwLSmqzrui2D81RVooUSV4e5VK")
	if err != nil {
		panic(err)
	}
	nominatorData, err := sdk.QueryeNominators(pk, 1507967)
	if err != nil {
		panic(err)
	}
	fmt.Println("len(nominatorData.Targets): ", len(nominatorData.Targets))
	for _, v := range nominatorData.Targets {
		fmt.Println(utils.EncodePublicKeyAsCessAccount(v[:]))
	}
	return

	pk, err = utils.ParsingPublickey("cXik7GNf8qYgt6TtGajELHN8QRjd9iy4pd6soPnjcccsenSuh")
	if err != nil {
		panic(err)
	}
	dd, err := sdk.QueryeErasStakers(432, pk)
	if err != nil {
		panic(err)
	}
	to := big.Int(dd.Total)
	ac := big.Int(dd.Own)
	fmt.Printf("total: %v\n", to.String())
	fmt.Printf("own: %v\n", ac.String())
	for _, v := range dd.Others {
		bg := big.Int(v.Value)
		fmt.Println(utils.EncodePublicKeyAsCessAccount(v.Who[:]))
		fmt.Printf("value: %v\n", bg.String())
	}
	return

	blockhash, err := sdk.GetSubstrateAPI().RPC.Chain.GetBlockHash(180)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sdk.RetrieveAllEventFromBlock(blockhash))

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

	var eve = types.EventRecords{}
	_ = eve

	// for _, e := range eve.System_ExtrinsicFailed {
	// 	//if IsApplyExtrinsic
	// 	//e.Phase.AsApplyExtrinsic
	// }

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
	fmt.Println(sdk.RetrieveAllEventFromBlock(blockhash))
}
