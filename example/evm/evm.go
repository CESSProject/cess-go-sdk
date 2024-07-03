package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/chain"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// Substrate well-known mnemonic:
//
//	https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
var MY_MNEMONIC = "head achieve piano online exhaust bulk trust vote inflict room keen maximum"

var RPC_ADDRS = []string{
	//testnet
	"wss://testnet-rpc.cess.cloud/ws/",
}

func main() {
	// 1. new sdk
	sdk, err := NewSDK()
	if err != nil {
		panic(err)
	}

	var (
		source       types.H160
		target       types.H160
		input        types.Bytes
		value        types.U256
		gasLimit     types.U64
		maxFeePerGas types.U256
		accessList   []chain.AccessInfo
	)

	s_h160, err := hex.DecodeString("1e3e1c69dfbd27d398e92da4844a9abdc2786ac0")
	if err != nil {
		log.Fatalln(err)
	}
	source = types.NewH160(s_h160)

	t_h160, err := hex.DecodeString("7352188979857675C3aD1AA6662326ebD6DDBf6d")
	if err != nil {
		log.Fatalln(err)
	}
	target = types.NewH160(t_h160)

	input_string, err := hex.DecodeString("a9059cbb00000000000000000000000085cdaca43a76c8ab769b974c2cf7306980742a310000000000000000000000000000000000000000000000000de0b6b3a7640000")
	input = input_string

	value = types.NewU256(*big.NewInt(0))

	gasLimit = 3000000

	maxFeePerGas = types.NewU256(*big.NewInt(500000000))

	block_hash, err := sdk.SendEvmCall(source, target, input, value, gasLimit, maxFeePerGas, accessList)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", block_hash)
}

func NewSDK() (*chain.ChainClient, error) {
	return cess.New(
		context.Background(),
		cess.ConnectRpcAddrs(RPC_ADDRS),
		cess.Mnemonic(MY_MNEMONIC),
		cess.TransactionTimeout(time.Second*10),
	)
}
