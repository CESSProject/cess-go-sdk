package main

import (
	"context"
	"fmt"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
)

var RPC_ADDRS = []string{
	//testnet
	"wss://testnet-rpc.cess.network/ws/",
}

func main() {
	sdk, err := cess.New(
		context.Background(),
		cess.ConnectRpcAddrs(RPC_ADDRS),
		cess.TransactionTimeout(time.Second*10),
	)
	if err != nil {
		panic(err)
	}
	defer sdk.Close()
	err = sdk.InitExtrinsicsName()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 2500000; i++ {
		filedata, err := sdk.ParseFileInBlock(uint64(i))
		if err != nil {
			fmt.Println(i, " ERR: ", err)
			continue
		}
		fmt.Println(i, " time: ", filedata.Timestamp)
		if len(filedata.UploadDecInfo) > 0 {
			fmt.Println(i, " uploadDecInfo:")
			for _, v := range filedata.UploadDecInfo {
				fmt.Println("    Owner: ", v.Owner)
				fmt.Println("    Fid: ", v.Fid)
			}
		}
		if len(filedata.DeleteFileInfo) > 0 {
			fmt.Println(i, " deleteFileInfo:")
			for _, v := range filedata.DeleteFileInfo {
				fmt.Println("    Owner: ", v.Owner)
				fmt.Println("    Fid: ", v.Fid)
			}
		}
		if len(filedata.StorageCompleted) > 0 {
			fmt.Println(i, " StorageCompleted:")
			for _, v := range filedata.StorageCompleted {
				fmt.Println("    StorageCompleted fid: ", v)
			}
		}
	}
	return

	blockData, err := sdk.ParseBlockData(2616537)
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
	fmt.Println("Punishment:")
	for _, v := range blockData.Punishment {
		fmt.Println("    Punishment miner: ", v.From)
		fmt.Println("    Punishment to: ", v.To)
		fmt.Println("    Punishment Amount: ", v.Amount)
		fmt.Println("    Punishment ExtrinsicHash: ", v.ExtrinsicHash)
		fmt.Println("    Punishment ExtrinsicName: ", v.ExtrinsicName)
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
	fmt.Println("MintTerritory:")
	for _, v := range blockData.MintTerritory {
		fmt.Println("    MintTerritory Account: ", v.Account)
		fmt.Println("    MintTerritory token: ", v.TerritoryToken)
		fmt.Println("    MintTerritory name: ", v.TerritoryName)
		fmt.Println("    MintTerritory size: ", v.TerritorySize)
		fmt.Println("    MintTerritory ExtrinsicHash: ", v.ExtrinsicHash)
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
