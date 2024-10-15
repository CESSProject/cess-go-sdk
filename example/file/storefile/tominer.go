/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"time"

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

const UploadFile = "file_name"
const TerritoryName = "territory_name"

var WantMiner = []string{}

func main() {
	fid, err := process.StoreFileToMiners(
		UploadFile,
		MY_MNEMONIC,
		TerritoryName,
		time.Second*15,
		RPC_ADDRS,
		WantMiner,
	)
	fmt.Println("fid: ", fid)
	fmt.Println("err: ", err)
}
