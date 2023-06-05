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
		config.CharacterName_Client,
		ConnectRpcAddrs([]string{"wss://testnet-rpc0.cess.cloud/ws/", "wss://testnet-rpc1.cess.cloud/ws/"}),
		Mnemonic("xxx xxx ... xxx"),
		TransactionTimeout(time.Duration(time.Second*10)),
	)
	if err != nil {
		panic(err)
	}
}

func Example_RegisterDeoss() {
	cli, err := New(
		config.CharacterName_Deoss,
		ConnectRpcAddrs([]string{"wss://testnet-rpc0.cess.cloud/ws/", "wss://testnet-rpc1.cess.cloud/ws/"}),
		Mnemonic("xxx xxx ... xxx"),
		TransactionTimeout(time.Duration(time.Second*10)),
	)
	if err != nil {
		panic(err)
	}
	txhash, _, err := cli.Register(cli.GetCharacterName(), cli.GetSignatureAccPulickey(), "", 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deoss registration successful, transaction hash is %s\n", txhash)
}

func Example_RegisterStorageNode() {
	cli, err := New(
		config.CharacterName_Bucket,
		ConnectRpcAddrs([]string{"wss://testnet-rpc0.cess.cloud/ws/", "wss://testnet-rpc1.cess.cloud/ws/"}),
		Mnemonic("xxx xxx ... xxx"),
		TransactionTimeout(time.Duration(time.Second*10)),
	)
	if err != nil {
		panic(err)
	}
	txhash, _, err := cli.Register(cli.GetCharacterName(), cli.GetSignatureAccPulickey(), "cXxxx...xxx", 100000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Storage node registration successful, transaction hash is %s\n", txhash)
}
