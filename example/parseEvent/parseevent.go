package main

import (
	"context"
	"fmt"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
)

// Substrate well-known mnemonic:
//
//	https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
var MY_MNEMONIC = "bottom drive obey lake curtain smoke basket hold race lonely fit walk" //

var RPC_ADDRS = []string{
	//devnet
	//"wss://devnet-rpc.cess.cloud/ws/",
	//testnet
	"wss://testnet-rpc0.cess.cloud/ws/",
	"wss://testnet-rpc1.cess.cloud/ws/",
	"wss://testnet-rpc2.cess.cloud/ws/",
}

func main() {
	sdk, err := cess.New(
		context.Background(),
		cess.ConnectRpcAddrs(RPC_ADDRS),
		//cess.Mnemonic(MY_MNEMONIC),
		cess.TransactionTimeout(time.Second*10),
	)
	if err != nil {
		panic(err)
	}
	sdk.InitExtrinsicsName()
	// RetrieveEvent_FileBank_CalculateReport
	// bhash, err := sdk.GetSubstrateAPI().RPC.Chain.GetBlockHash(298712)
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Println(sdk.RetrieveAllEventFromBlock(bhash))
	fmt.Println(" --------- ")
	sysEvents, extrinsics, transferInfo, blockhash, preHash, extHash, stHash, t, err := sdk.RetrieveBlock(312732)
	if err != nil {
		panic(err)
	}
	fmt.Println("system events: ", sysEvents)
	fmt.Println("extrinsics:")
	for k, v := range extrinsics {
		fmt.Println("  ", k, ": ", v.Name)
		fmt.Println("    Singer: ", v.Signer)
		fmt.Println("    FeePaid: ", v.FeePaid)
		fmt.Println("    Events: ", v.Events)
	}
	fmt.Println("transfer info: ", transferInfo)
	fmt.Println("blockhash: ", blockhash)
	fmt.Println("preHash: ", preHash)
	fmt.Println("extHash: ", extHash)
	fmt.Println("stHash: ", stHash)
	fmt.Println("timpstamp: ", t)
	//fmt.Println(sdk.RetrieveEvent_FileBank_CalculateReport(bhash))
}
