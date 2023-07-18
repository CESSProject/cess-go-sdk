package main

import (
	"context"
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

var Workspace = "/cess"
var Port = 4001
var Bootstrap = []string{
	"_dnsaddr.boot-kldr-testnet.cess.cloud",
}

const UploadFile = "test.log"
const DownloadFile = "download_file"
const BucketName = "myBucket"

func main() {
	sdk, err := NewSDK()
	if err != nil {
		panic(err)
	}
	fid, err := StoreFile(sdk, UploadFile, BucketName)
	if err != nil {
		panic(err)
	}
	log.Println("fid:", fid)
	err = RetrieveFile(sdk, fid, DownloadFile)
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
		cess.Workspace(Workspace),
		cess.P2pPort(Port),
		cess.Bootnodes(Bootstrap),
		cess.ProtocolPrefix(config.TestnetProtocolPrefix),
	)
}

func RetrieveFile(sdk sdk.SDK, fid, DownloadFile string) error {
	return sdk.RetrieveFile(fid, DownloadFile)
}

func StoreFile(sdk sdk.SDK, uploadFile, bucketName string) (string, error) {
	return sdk.StoreFile(uploadFile, bucketName)
}
