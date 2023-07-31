package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/config"
	"github.com/CESSProject/cess-go-sdk/core/sdk"
)

// Substrate well-known mnemonic:
//
//	https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
var MY_MNEMONIC = "bottom drive obey lake curtain smoke basket hold race lonely fit walk"

var RPC_ADDRS = []string{
	"wss://testnet-rpc0.cess.cloud/ws/",
	"wss://testnet-rpc1.cess.cloud/ws/",
	"wss://testnet-rpc2.cess.cloud/ws/",
}

const PublicGateway = "http://deoss-pub-gateway.cess.cloud/"

const UploadFile = "example.go"
const BucketName = "myBucket"

func main() {
	sdk, err := NewSDK()
	if err != nil {
		panic(err)
	}

	// upload file
	fid, err := sdk.StoreFile(PublicGateway, UploadFile, BucketName)
	if err != nil {
		panic(err)
	}
	log.Println("fid:", fid)

	// Retrieve file
	err = sdk.RetrieveFile(fid, fmt.Sprintf("download_%d", time.Now().UnixNano()))
	if err != nil {
		panic(err)
	}

	// upload object
	fid, err = sdk.StoreObject(PublicGateway, bytes.NewReader([]byte("test date")), BucketName)
	if err != nil {
		panic(err)
	}
	log.Println("fid:", fid)

	// Retrieve file
	err = sdk.RetrieveFile(fid, fmt.Sprintf("download_%d", time.Now().UnixNano()))
	if err != nil {
		panic(err)
	}
}

func NewSDK() (sdk.SDK, error) {
	return cess.New(
		context.Background(),
		config.CharacterName_Client,
		cess.ConnectRpcAddrs(RPC_ADDRS),
		cess.Mnemonic(MY_MNEMONIC),
		cess.TransactionTimeout(time.Second*10),
	)
}
