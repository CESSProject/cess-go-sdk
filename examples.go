/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdkgo

import (
	"fmt"
	"time"

	"github.com/CESSProject/sdk-go/config"
)

func Example_newClient() {
	_, err := New(
		config.DefaultName,
		ConnectRpcAddrs([]string{""}),
		Mnemonic("xxx xxx ... xxx"),
		TransactionTimeout(time.Duration(time.Second*10)),
	)
	if err != nil {
		panic(err)
	}
}

func Example_registerOss() {
	cli, err := New(
		"oss",
		ConnectRpcAddrs([]string{""}),
		Mnemonic(""),
		TransactionTimeout(time.Duration(time.Second*10)),
	)
	if err != nil {
		panic(err)
	}
	txhash, _, err := cli.Register("oss", nil, "", 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("OSS registration successful, transaction hash is %s\n", txhash)
}

func Example_registerMiner() {
	cli, err := New(
		"bucket",
		ConnectRpcAddrs([]string{""}),
		Mnemonic(""),
		TransactionTimeout(time.Duration(time.Second*10)),
	)
	if err != nil {
		panic(err)
	}
	txhash, _, err := cli.Register("bucket", nil, "cXxxx...xxx", 100000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Miner registration successful, transaction hash is %s\n", txhash)
}
