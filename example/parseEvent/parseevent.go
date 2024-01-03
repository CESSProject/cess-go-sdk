package main

import (
	"context"
	"fmt"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// Substrate well-known mnemonic:
//
//	https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
//
// var MY_MNEMONIC = "bottom drive obey lake curtain smoke basket hold race lonely fit walk"//
var MY_MNEMONIC = "goose never chase despair dice phone penalty inside runway release cruise gain"

var RPC_ADDRS = []string{
	//devnet
	"wss://devnet-rpc.cess.cloud/ws/",
	//testnet
	// "wss://testnet-rpc0.cess.cloud/ws/",
	// "wss://testnet-rpc1.cess.cloud/ws/",
	// "wss://testnet-rpc2.cess.cloud/ws/",
}

func main() {
	sdk, err := cess.New(
		context.Background(),
		cess.ConnectRpcAddrs(RPC_ADDRS),
		cess.Mnemonic(MY_MNEMONIC),
		cess.TransactionTimeout(time.Second*10),
	)
	if err != nil {
		panic(err)
	}
	// RetrieveEvent_FileBank_CalculateReport
	bhash, err := sdk.GetSubstrateAPI().RPC.Chain.GetBlockHash(72745)
	if err != nil {
		panic(err)
	}
	//hash, err := codec.HexDecodeString("0x5f8be5b640bf0eedd07fb886ff0f94091dae8192a4596766e878a531e697693f")
	chash, err := types.NewHashFromHexString("0x5f8be5b640bf0eedd07fb886ff0f94091dae8192a4596766e878a531e697693f")
	if err != nil {
		panic(err)
	}
	fmt.Println(bhash)
	fmt.Println(chash)

	fmt.Println(sdk.RetrieveEvent_FileBank_CalculateReport(chash))
}
