package main

import (
	"context"
	"fmt"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/core/event"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
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

type MyEvent struct {
	Sminer_FaucetTopUpMoney []event.Event_FaucetTopUpMoney
	types.EventRecords
}

func main() {
	// fmt.Println(time.Now().UTC().Truncate(24 * time.Hour).Add(time.Second).UnixMilli())
	// fmt.Println(time.UnixMilli(1706699208000).UTC().Truncate(24 * time.Hour).Add(time.Hour * 24).UnixMilli())
	// fmt.Println(time.Now().UTC())
	// fmt.Println(time.Now().UTC().Hour())
	// fmt.Println(time.Now().UTC().Minute())
	// return
	sdk, err := cess.New(
		context.Background(),
		cess.ConnectRpcAddrs(RPC_ADDRS),
		//cess.Mnemonic(MY_MNEMONIC),
		cess.TransactionTimeout(time.Second*10),
	)
	if err != nil {
		panic(err)
	}

	// fmeta, err := sdk.QueryFileMetadataByBlock("bf7e61cf8abe365dc30e525be5058fd3f502245322300d76fe169c9292c6ba48", 2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, segment := range fmeta.SegmentList {
	// 	for _, fragment := range segment.FragmentList {
	// 		fmt.Println(utils.EncodePublicKeyAsCessAccount(fragment.Miner[:]))
	// 	}
	// }
	// return
	sdk.InitExtrinsicsName()

	// RetrieveEvent_FileBank_CalculateReport
	// bhash, err := sdk.GetSubstrateAPI().RPC.Chain.GetBlockHash(713)
	// if err != nil {
	// 	panic(err)
	// }
	// var data types.StorageDataRaw
	// key, err := types.CreateStorageKey(sdk.GetMetadata(), "System", "Events", nil, nil)
	// if err != nil {
	// 	panic(err)
	// }
	// ok, err := sdk.GetSubstrateAPI().RPC.State.GetStorage(key, &data, bhash)
	// if err != nil {
	// 	panic(err)
	// }

	// if ok {
	// 	events := MyEvent{}
	// 	err = types.EventRecordsRaw(data).DecodeEventRecords(sdk.GetMetadata(), &events)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	for _, e := range events.Balances_Transfer {
	// 		fmt.Printf("Balances:Transfer:: (phase=%#v)\n", e.Phase)
	// 		fmt.Printf("\t%v, %v, %v\n", e.From, e.To, e.Value)
	// 	}
	// }

	//fmt.Println(sdk.RetrieveAllEventFromBlock(bhash))

	sysEvents, extrinsics, transferInfo, blockhash, preHash, extHash, stHash, t, err := sdk.RetrieveBlock(11921)
	if err != nil {
		panic(err)
	}
	fmt.Println(" --------- ")

	fmt.Println("extrinsics:")
	for k, v := range extrinsics {
		fmt.Println("  ", k, ": ", v.Name)
		fmt.Println("    Singer: ", v.Signer)
		fmt.Println("    FeePaid: ", v.FeePaid)
		fmt.Println("    Events: ", v.Events)
	}
	fmt.Println("system events: ", sysEvents)
	fmt.Println("transfer info: ", transferInfo)
	fmt.Println("blockhash: ", blockhash)
	fmt.Println("preHash: ", preHash)
	fmt.Println("extHash: ", extHash)
	fmt.Println("stHash: ", stHash)
	fmt.Println("timpstamp: ", t)
	//fmt.Println(sdk.RetrieveEvent_FileBank_CalculateReport(bhash))
}
