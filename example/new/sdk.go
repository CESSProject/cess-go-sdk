/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/utils"
)

// Substrate well-known mnemonic:
//
//   - cXgaee2N8E77JJv9gdsGAckv1Qsf3hqWYf7NL4q6ZuQzuAUtB
//   - https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
var MY_MNEMONIC = "bottom drive obey lake curtain smoke basket hold race lonely fit walk"

var RPC_ADDRS = []string{
	//devnet
	"wss://devnet-rpc.cess.cloud/ws/",

	//testnet
	//"wss://testnet-rpc.cess.cloud/ws/",
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
	defer sdk.Close()

	err = sdk.InitExtrinsicsName()
	if err != nil {
		panic(err)
	}

	fmt.Println(sdk.SystemVersion())
	fmt.Println(sdk.GetCurrentRpcAddr())

	fmt.Println(sdk.QueryRewardMap(sdk.GetSignatureAccPulickey(), -1))
	return
	puk, err := utils.ParsingPublickey("cXfg2SYcq85nyZ1U4ccx6QnAgSeLQB8aXZ2jstbw9CPGSmhXY")
	if err != nil {
		panic(err)
	}

	fmt.Println(sdk.QueryeErasStakersOverview(6, puk))
	return

	puk, err = utils.ParsingPublickey("cXfg2SYcq85nyZ1U4ccx6QnAgSeLQB8aXZ2jstbw9CPGSmhXY")
	if err != nil {
		panic(err)
	}
	result, err := sdk.QueryeAllErasStakersPaged(6, puk)
	if err != nil {
		panic(err)
	}
	pagetotal_bg := big.Int(result[0].PageTotal)
	fmt.Printf("pagetotal: %v\n", pagetotal_bg.String())
	for i := 0; i < len(result); i++ {
		for _, v := range result[i].Others {
			bg := big.Int(v.Value)
			fmt.Println(utils.EncodePublicKeyAsCessAccount(v.Who[:]))
			fmt.Printf("value: %v\n", bg.String())
		}
	}
	return

	pk, err := utils.ParsingPublickey("cXiKthh2dyY1taTydtdxiqQwXY1HKZcXvYGmjS2UmuPi2qNDS")
	if err != nil {
		panic(err)
	}
	fmt.Println(sdk.QueryCurrentCounters(pk, -1))
	return

	pk, err = utils.ParsingPublickey("cXfnLrW67qKkTn7DPXfh1SykwLSmqzrui2D81RVooUSV4e5VK")
	if err != nil {
		panic(err)
	}
	nominatorData, err := sdk.QueryeNominators(pk, 1507967)
	if err != nil {
		panic(err)
	}
	fmt.Println("len(nominatorData.Targets): ", len(nominatorData.Targets))
	for _, v := range nominatorData.Targets {
		fmt.Println(utils.EncodePublicKeyAsCessAccount(v[:]))
	}
	return

	pk, err = utils.ParsingPublickey("cXik7GNf8qYgt6TtGajELHN8QRjd9iy4pd6soPnjcccsenSuh")
	if err != nil {
		panic(err)
	}
	dd, err := sdk.QueryeErasStakers(432, pk)
	if err != nil {
		panic(err)
	}
	to := big.Int(dd.Total)
	ac := big.Int(dd.Own)
	fmt.Printf("total: %v\n", to.String())
	fmt.Printf("own: %v\n", ac.String())
	for _, v := range dd.Others {
		bg := big.Int(v.Value)
		fmt.Println(utils.EncodePublicKeyAsCessAccount(v.Who[:]))
		fmt.Printf("value: %v\n", bg.String())
	}
}
