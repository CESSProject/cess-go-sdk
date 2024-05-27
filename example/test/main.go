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
)

// Substrate well-known mnemonic:
//
//   - https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
//   - cXgaee2N8E77JJv9gdsGAckv1Qsf3hqWYf7NL4q6ZuQzuAUtB
var MY_MNEMONIC = "bottom drive obey lake curtain smoke basket hold race lonely fit walk"

var RPC_ADDRS = []string{
	//testnet
	"wss://testnet-rpc0.cess.cloud/ws/",
	"wss://testnet-rpc1.cess.cloud/ws/",
	"wss://testnet-rpc2.cess.cloud/ws/",
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

	// get sdk name
	fmt.Println(sdk.GetSDKName())

	// get the current rpc address being used
	fmt.Println(sdk.GetCurrentRpcAddr())

	// get the rpc connection status flag
	//   - true: connection is normal
	//   - false: connection failed
	fmt.Println(sdk.GetRpcState())

	// get your current account address
	//   - make sure you fill in mnemonic when you create the sdk client
	fmt.Println(sdk.GetSignatureAcc())

	// get your current account public key
	//   - make sure you fill in mnemonic when you create the sdk client
	fmt.Println(sdk.GetSignatureAccPulickey())

	// get substrate api
	fmt.Println(sdk.GetSubstrateAPI())

	// get the mnemonic for your current account
	fmt.Println(sdk.GetURI())

	// get token symbol
	fmt.Println(sdk.GetTokenSymbol())

	// get network environment
	fmt.Println(sdk.GetNetworkEnv())
}
