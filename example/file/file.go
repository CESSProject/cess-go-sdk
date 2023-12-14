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
	"io"
	"log"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/core/sdk"
)

// Substrate well-known mnemonic:
//
//	https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
var MY_MNEMONIC = "bottom drive obey lake curtain smoke basket hold race lonely fit walk"

var RPC_ADDRS = []string{
	//devnet
	"wss://devnet-rpc.cess.cloud/ws/",
	//testnet
	"wss://testnet-rpc0.cess.cloud/ws/",
	"wss://testnet-rpc1.cess.cloud/ws/",
	"wss://testnet-rpc2.cess.cloud/ws/",
}

const PublicGateway = "http://deoss-pub-gateway.cess.cloud/"
const PublicGatewayAccount = "cXhwBytXqrZLr1qM5NHJhCzEMckSTzNKw17ci2aHft6ETSQm9"
const UploadFile = "file.go"
const BucketName = "myBucket"

func main() {
	// 1. new sdk
	sdk, err := NewSDK()
	if err != nil {
		panic(err)
	}

	// 2. buy space
	_, err = sdk.BuySpace(1)
	if err != nil {
		panic(err)
	}

	// 3. authorize space to deoss
	_, err = sdk.AuthorizeSpace(PublicGatewayAccount)
	if err != nil {
		panic(err)
	}

	// 4. upload file to deoss
	fid, err := sdk.StoreFile(PublicGateway, UploadFile, BucketName)
	if err != nil {
		panic(err)
	}
	log.Println("fid:", fid)

	// 5. downloag file from deoss
	err = sdk.RetrieveFile(PublicGateway, fid, fmt.Sprintf("download_%d", time.Now().UnixNano()))
	if err != nil {
		panic(err)
	}

	// 6. upload object to deoss
	fid, err = sdk.StoreObject(PublicGateway, bytes.NewReader([]byte("test date")), BucketName)
	if err != nil {
		panic(err)
	}
	log.Println("fid:", fid)

	// 7. download object from deoss
	body, err := sdk.RetrieveObject(PublicGateway, fid)
	if err != nil {
		panic(err)
	}
	defer body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func NewSDK() (sdk.SDK, error) {
	return cess.New(
		context.Background(),
		cess.ConnectRpcAddrs(RPC_ADDRS),
		cess.Mnemonic(MY_MNEMONIC),
		cess.TransactionTimeout(time.Second*10),
	)
}
