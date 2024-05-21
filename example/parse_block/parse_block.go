package main

import (
	"context"
	"fmt"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/chain"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// Substrate well-known mnemonic:
//
//	https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
var MY_MNEMONIC = "bottom drive obey lake curtain smoke basket hold race lonely fit walk" //

var RPC_ADDRS = []string{
	//devnet
	"wss://devnet-rpc.cess.cloud/ws/",
	//testnet
	// "wss://testnet-rpc0.cess.cloud/ws/",
	// "wss://testnet-rpc1.cess.cloud/ws/",
	// "wss://testnet-rpc2.cess.cloud/ws/",
	// "wss://testnet-rpc3.cess.cloud/ws/",
}

type MyEvent struct {
	Sminer_FaucetTopUpMoney []chain.Event_FaucetTopUpMoney
	types.EventRecords
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
	defer sdk.Close()
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

	blockData, err := sdk.ParseBlockData(36590)
	if err != nil {
		fmt.Println("ERR: ", err)
		return
	}
	fmt.Println("extrinsics:")
	for k, v := range blockData.Extrinsics {
		fmt.Println("  ", k, ": ", v.Name)
		fmt.Println("    Singer: ", v.Signer)
		fmt.Println("    Hash: ", v.Hash)
		fmt.Println("    FeePaid: ", v.FeePaid)
		fmt.Println("    Events: ", v.Events)
	}
	fmt.Println("uploadDecInfo:")
	for _, v := range blockData.UploadDecInfo {
		fmt.Println("    Owner: ", v.Owner)
		fmt.Println("    Fid: ", v.Fid)
	}
	fmt.Println("deleteFileInfo:")
	for _, v := range blockData.DeleteFileInfo {
		fmt.Println("    Owner: ", v.Owner)
		fmt.Println("    Fid: ", v.Fid)
	}
	fmt.Println("createBucketInfo:")
	for _, v := range blockData.CreateBucketInfo {
		fmt.Println("    Owner: ", v.Owner)
		fmt.Println("    BucketName: ", v.BucketName)
	}
	fmt.Println("DeleteBucketInfo:")
	for _, v := range blockData.DeleteBucketInfo {
		fmt.Println("    Owner: ", v.Owner)
		fmt.Println("    BucketName: ", v.BucketName)
	}
	fmt.Println("GenChallenge:")
	for _, v := range blockData.GenChallenge {
		fmt.Println("    GenChallenge miner: ", v)
	}
	fmt.Println("SubmitIdleProve:")
	for _, v := range blockData.SubmitIdleProve {
		fmt.Println("    SubmitIdleProve miner: ", v)
	}
	fmt.Println("SubmitServiceProve:")
	for _, v := range blockData.SubmitServiceProve {
		fmt.Println("    SubmitServiceProve miner: ", v)
	}
	fmt.Println("SubmitIdleResult:")
	for _, v := range blockData.SubmitIdleResult {
		fmt.Println("    SubmitIdleResult miner: ", v.Miner)
		fmt.Println("    SubmitIdleResult miner result: ", v.Result)
	}
	fmt.Println("SubmitServiceResult:")
	for _, v := range blockData.SubmitServiceResult {
		fmt.Println("    SubmitServiceResult miner: ", v.Miner)
		fmt.Println("    SubmitServiceResult miner result: ", v.Result)
	}
	fmt.Println("MinerRegPoiskeys:")
	for _, v := range blockData.MinerRegPoiskeys {
		fmt.Println("    MinerRegPoiskeys miner: ", v.Miner)
	}
	fmt.Println("GatewayReg:")
	for _, v := range blockData.GatewayReg {
		fmt.Println("    GatewayReg account: ", v.Account)
	}
	fmt.Println("StorageCompleted:")
	for _, v := range blockData.StorageCompleted {
		fmt.Println("    StorageCompleted fid: ", v)
	}
	fmt.Println("EraPaid:")
	fmt.Println("    EraPaid valu: ", blockData.EraPaid.HaveValue)
	fmt.Println("    EraPaid EraIndex: ", blockData.EraPaid.EraIndex)
	fmt.Println("    EraPaid ValidatorPayout: ", blockData.EraPaid.ValidatorPayout)
	fmt.Println("    EraPaid Remainder: ", blockData.EraPaid.Remainder)

	fmt.Println("StakingPayouts:")
	for _, v := range blockData.StakingPayouts {
		fmt.Println("    StakingPayouts EraIndex: ", v.EraIndex)
		fmt.Println("    StakingPayouts ClaimedAcc: ", v.ClaimedAcc)
		fmt.Println("    StakingPayouts Amount: ", v.Amount)
		fmt.Println("    StakingPayouts ExtrinsicHash: ", v.ExtrinsicHash)
	}
	fmt.Println("Unbonded:")
	for _, v := range blockData.Unbonded {
		fmt.Println("    Unbonded Account: ", v.Account)
		fmt.Println("    StakingPayouts Amount: ", v.Amount)
		fmt.Println("    Unbonded ExtrinsicHash: ", v.ExtrinsicHash)
	}

	fmt.Println("system events: ", blockData.SysEvents)
	fmt.Println("transfer info: ", blockData.TransferInfo)
	fmt.Println("minerReg info: ", blockData.MinerReg)
	fmt.Println("newAccounts info: ", blockData.NewAccounts)
	fmt.Println("blockhash: ", blockData.BlockHash)
	fmt.Println("preHash: ", blockData.PreHash)
	fmt.Println("extHash: ", blockData.ExtHash)
	fmt.Println("stHash: ", blockData.StHash)
	fmt.Println("timpstamp: ", blockData.Timestamp)
	fmt.Println("allGasFee: ", blockData.AllGasFee)
	fmt.Println("IsNewEra: ", blockData.IsNewEra)
}
