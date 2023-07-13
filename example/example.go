package main

import (
	"context"
	"fmt"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/config"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
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

const UploadFile = "example.go"
const DownloadFile = "download_file"
const BucketName = "myBucket"
const FileHash = "3ea772e68cf615260916dc94f501c43da78f6fdc15dc20e722e5284aca612a92"

func main() {
	StoreFile()
	RetrieveFile()
}

func RetrieveFile() {
	sdk, err := cess.New(
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
	if err != nil {
		panic(err)
	}

	fmt.Println(sdk.RetrieveFile(FileHash, DownloadFile))
}

func StoreFile() {
	sdk, err := cess.New(
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
	if err != nil {
		panic(err)
	}

	fmt.Println(sdk.StoreFile(UploadFile, BucketName))
}

func CreateBucket() {
	sdk, err := cess.New(
		context.Background(),
		config.CharacterName_Client,
		cess.ConnectRpcAddrs(RPC_ADDRS),
		cess.Mnemonic(MY_MNEMONIC),
		cess.TransactionTimeout(time.Second*10),
		cess.ProtocolPrefix(config.TestnetProtocolPrefix),
	)
	if err != nil {
		panic(err)
	}

	keyringPair, err := signature.KeyringPairFromSecret(MY_MNEMONIC, 0)

	if !utils.CheckBucketName(BucketName) {
		panic("invalid bucket name")
	}

	fmt.Println(sdk.CreateBucket(keyringPair.PublicKey, BucketName))
}
