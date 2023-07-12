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
}

var Workspace = "/cess"
var Port = 4001
var Bootstrap = []string{
	"_dnsaddr.sjc-1.bootstrap-kldr.cess.cloud",
}

const File = "/home/test_download"
const BucketName = "myBucket"
const FileHash = "c158d7008e94d3af61033b6861aa4f35a4c2b829c7e97224fcbb54618de55945"

func main() {
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
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(sdk.RetrieveFile(FileHash, File))
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
	)
	if err != nil {
		panic(err)
	}

	keyringPair, err := signature.KeyringPairFromSecret(MY_MNEMONIC, 0)

	if !utils.CheckBucketName(BucketName) {
		panic("invalid bucket name")
	}

	fmt.Println(sdk.StoreFile(keyringPair.Address, File, BucketName))
}

func CreateBucket() {
	sdk, err := cess.New(
		context.Background(),
		config.CharacterName_Client,
		cess.ConnectRpcAddrs(RPC_ADDRS),
		cess.Mnemonic(MY_MNEMONIC),
		cess.TransactionTimeout(time.Second*10),
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
