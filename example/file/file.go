/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/CESSProject/cess-go-sdk/core/process"
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
	// upload file to gateway
	fid, err := process.StoreFile(PublicGateway, UploadFile, BucketName, MY_MNEMONIC)
	if err != nil {
		panic(err)
	}
	log.Println("fid:", fid)

	// downloag file from gateway
	err = process.RetrieveFile(PublicGateway, fid, MY_MNEMONIC, fmt.Sprintf("download_%d", time.Now().UnixNano()))
	if err != nil {
		panic(err)
	}

	// upload object to gateway
	fid, err = process.StoreObject(PublicGateway, BucketName, MY_MNEMONIC, bytes.NewReader([]byte("test date")))
	if err != nil {
		panic(err)
	}
	log.Println("fid:", fid)

	// download object from gateway
	body, err := process.RetrieveObject(PublicGateway, fid, MY_MNEMONIC)
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
