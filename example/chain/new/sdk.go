/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"context"
	"fmt"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/utils"
)

// Substrate well-known mnemonic:
//
//   - cXgaee2N8E77JJv9gdsGAckv1Qsf3hqWYf7NL4q6ZuQzuAUtB
//   - https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
var MY_MNEMONIC = "bottom drive obey lake curtain smoke basket hold race lonely fit walk"

var RPC_ADDRS = []string{
	//testnet
	"wss://testnet-rpc.cess.network/ws/",
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

	allminer, err := sdk.QueryAllMiner(2241971)
	if err != nil {
		panic(err)
	}
	fmt.Println("len: ", len(allminer))
	//fmt.Println(sdk.GetCurrentRpcAddr())
	return

	err = sdk.InitExtrinsicsName()
	if err != nil {
		panic(err)
	}

	fmt.Println(sdk.SystemVersion())
	fmt.Println(sdk.GetCurrentRpcAddr())
	fmt.Println(sdk.SystemProperties())

	puk, err := utils.ParsingPublickey("cXhRaGVf6XeiSiK2GULgy12muovp5ysR1ViaiVRNa6mvQXQHn")
	if err != nil {
		panic(err)
	}

	fmt.Println(sdk.QueryMinerItems(puk, 2241971))
}
