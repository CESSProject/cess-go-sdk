/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdkgo_test

import (
	"context"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/config"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	p2pgo "github.com/CESSProject/p2p-go"
)

const DEFAULT_WAIT_TIME = time.Second * 15
const P2P_PORT = 4001
const TMP_DIR = "/tmp"

func TestMain(m *testing.M) {
	// Change the following `.env.testnet` to `.env.local` to run test against a local node.
	// If you run these examples using localNode, please run a
	//   [CESS node](https://github.com/cessProject/cess) locally as well.
	godotenv.Load(".env.testnet")
	os.Exit(m.Run())
}

func TestNewClient(t *testing.T) {
	_, err := cess.New(
		context.Background(),
		config.CharacterName_Client,
		cess.ConnectRpcAddrs(strings.Split(os.Getenv("RPC_ADDRS"), " ")),
		cess.Mnemonic(os.Getenv("MY_MNEMONIC")),
		cess.TransactionTimeout(time.Duration(DEFAULT_WAIT_TIME)),
	)
	assert.NoError(t, err)
}

func Example_register_deoss() {
	cli, err := cess.New(
		context.Background(),
		config.CharacterName_Deoss,
		cess.ConnectRpcAddrs(strings.Split(os.Getenv("RPC_ADDRS"), " ")),
		cess.Mnemonic(os.Getenv("MY_MNEMONIC")),
		cess.TransactionTimeout(time.Duration(DEFAULT_WAIT_TIME)),
	)
	if err != nil {
		log.Fatalf("err: %v", err.Error())
	}

	var bootnodes = make([]string, 0)

	for _, v := range strings.Split(os.Getenv("BOOTSTRAP_NODES"), " ") {
		temp, err := utils.ParseMultiaddrs(v)
		if err != nil {
			continue
		}
		bootnodes = append(bootnodes, temp...)
	}

	p2p, err := p2pgo.New(
		context.Background(),
		p2pgo.ListenPort(P2P_PORT),
		p2pgo.Workspace(TMP_DIR),
		p2pgo.BootPeers(bootnodes),
	)
	if err != nil {
		log.Fatalf("err: %v", err.Error())
	}

	_, _, err = cli.Register(cli.GetRoleName(), p2p.GetPeerPublickey(), "", 0)
	if err != nil {
		log.Fatalf("err: %v", err.Error())
	}
}

func Example_register_storage_node() {
	cli, err := cess.New(
		context.Background(),
		config.CharacterName_Bucket,
		cess.ConnectRpcAddrs(strings.Split(os.Getenv("RPC_ADDRS"), " ")),
		cess.Mnemonic(os.Getenv("MY_MNEMONIC")),
		cess.TransactionTimeout(time.Duration(DEFAULT_WAIT_TIME)),
	)
	if err != nil {
		log.Fatalf("err: %v", err.Error())
	}

	var bootnodes = make([]string, 0)

	for _, v := range strings.Split(os.Getenv("BOOTSTRAP_NODES"), " ") {
		temp, err := utils.ParseMultiaddrs(v)
		if err != nil {
			continue
		}
		bootnodes = append(bootnodes, temp...)
	}

	p2p, err := p2pgo.New(
		context.Background(),
		p2pgo.ListenPort(P2P_PORT),
		p2pgo.Workspace(TMP_DIR),
		p2pgo.BootPeers(bootnodes),
	)
	if err != nil {
		log.Fatalf("err: %v", err.Error())
	}

	_, _, err = cli.Register(cli.GetRoleName(), p2p.GetPeerPublickey(), os.Getenv("MY_ADDR"), 0)
	if err != nil {
		log.Fatalf("err: %v", err.Error())
	}
}
