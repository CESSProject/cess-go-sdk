/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"

	"github.com/CESSProject/cess-go-sdk/core/process"
)

// Substrate well-known mnemonic:
//
//	https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
var MY_MNEMONIC = "bottom drive obey lake curtain smoke basket hold race lonely fit walk"

var RPC_ADDRS = []string{
	//testnet
	"wss://testnet-rpc.cess.network/ws/",
}

const (
	FID    = ""
	CIPHER = ""
	DIR    = "."
)

func main() {
	err := process.RetrieveFileFromMiners(RPC_ADDRS, MY_MNEMONIC, FID, CIPHER, DIR)
	fmt.Println("err: ", err)
}
