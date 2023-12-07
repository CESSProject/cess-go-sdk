package main

import (
	"context"

	cess "github.com/CESSProject/cess-go-sdk"
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

func main() {
	sdk, err := cess.New(
		context.Background(),
		cess.ConnectRpcAddrs(RPC_ADDRS),
		cess.Mnemonic(MY_MNEMONIC),
	)
	if err != nil {
		panic(err)
	}
	// Upload file to oss
	sdk.StoreFile("http://deoss-pub-gateway.cess.cloud/", "test.log", "test-bucket")
}
