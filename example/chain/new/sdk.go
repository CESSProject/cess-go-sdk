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
var MY_MNEMONIC = "success dirt biology roof episode clarify inflict unit average damage escape tackle" //"bottom drive obey lake curtain smoke basket hold race lonely fit walk"

var RPC_ADDRS = []string{
	//testnet
	"wss://devnet-rpc.cess.network/ws/",
}

func main() {
	sdk, err := cess.New(
		context.Background(),
		cess.ConnectRpcAddrs(RPC_ADDRS),
		//cess.Mnemonic(MY_MNEMONIC),
		cess.TransactionTimeout(time.Second*10),
	)
	if err != nil {
		panic(err)
	}
	defer sdk.Close()

	err = sdk.InitExtrinsicsName()
	if err != nil {
		panic(err)
	}

	fmt.Println(sdk.SystemVersion())
	fmt.Println(sdk.GetCurrentRpcAddr())
	fmt.Println(sdk.SystemProperties())

	puk, err := utils.ParsingPublickey("cXjqBVG3JZhNBAbwpaRDt9ct2TtZcq4R8Zp4DwhCMvhPjskrn")
	if err != nil {
		panic("cXjqBVG3JZhNBAbwpaRDt9ct2TtZcq4R8Zp4DwhCMvhPjskrn")
	}
	fmt.Println(sdk.QueryChallengeSnapShot(puk, 42436))

	return
	fmt.Println(sdk.TransferToken("cXkdXokcMa32BAYkmsGjhRGA2CYmLUN2pq69U8k9taXsQPHGp", "100000000000000000000"))
}
