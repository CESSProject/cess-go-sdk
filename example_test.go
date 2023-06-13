/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdkgo_test

import (
	"testing"
	"context"
	"log"
	"time"

	p2pgo "github.com/CESSProject/p2p-go"
	cess "github.com/CESSProject/sdk-go"
	"github.com/CESSProject/sdk-go/config"
)

// Substrate well-known mnemonic:
//
//	https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
var MNEMONIC = "bottom drive obey lake curtain smoke basket hold race lonely fit walk" //Alice
var AliceAddr = "cXjmuHdBk4J3Zyt2oGodwGegNFaTFPcfC48PZ9NMmcUFzF6cc"

var testnets = []string{
	"wss://testnet-rpc0.cess.cloud/ws/",
	"wss://testnet-rpc1.cess.cloud/ws/",
}

var localNode = []string{
	"ws://localhost:9944",
}

var P2pBootstrapNodes = []string{
	"/ip4/221.122.79.2/tcp/10010/p2p/12D3KooWHY6BRu2MtG9SempACgYCcGHRSEai2ZkWY3E4VKDYrqh9",
	"/ip4/45.77.47.184/tcp/10010/p2p/12D3KooWBW5YSqJtABaaTmMZ1ByARcsTtCmmfB6na5HBEuUoKkLM",
	"/ip4/221.122.79.3/tcp/10010/p2p/12D3KooWAdyc4qPWFHsxMtXvSrm7CXNFhUmKPQdoXuKQXki69qBo",
}

// Tmp files will be downloaded
var testWorkspacePath = "/tmp"

const P2pCommunicationPort = 4001

// To run these examples, please run a [CESS node](https://github.com/cessProject/cess) locally
//   as well.

func TestNewClient(t *testing.T) {
	cli, err := cess.New(
		config.CharacterName_Client,
		cess.ConnectRpcAddrs(localNode),
		cess.Mnemonic(MNEMONIC),
		cess.TransactionTimeout(time.Duration(time.Second*15)),
	)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.FailNow()
	}

	blockhright, _ := cli.QueryBlockHeight("")
	log.Printf("Successfully created SDK client, latest block height: %d\n", blockhright)
}

func TestRegisterDeOSS(t *testing.T) {
	cli, err := cess.New(
		config.CharacterName_Deoss,
		cess.ConnectRpcAddrs(localNode),
		cess.Mnemonic(MNEMONIC),
		cess.TransactionTimeout(time.Duration(time.Second*15)),
	)
	if err != nil {
		t.Errorf("Error: %s", err)
		t.FailNow()
	}

	p2p, err := p2pgo.New(
		context.Background(),
		p2pgo.ListenPort(P2pCommunicationPort),
		p2pgo.Workspace(testWorkspacePath),
		p2pgo.BootPeers(P2pBootstrapNodes),
	)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.FailNow()
	}

	txhash, _, err := cli.Register(cli.GetRoleName(), p2p.GetPeerPublickey(), "", 0)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.FailNow()
	}

	log.Printf("Deoss registration successful, transaction hash is %s\n", txhash)
}

func TestRegisterStorageNode(t *testing.T) {
	cli, err := cess.New(
		config.CharacterName_Bucket,
		cess.ConnectRpcAddrs(localNode),
		cess.Mnemonic(MNEMONIC),
		cess.TransactionTimeout(time.Duration(time.Second*10)),
	)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.FailNow()
	}

	p2p, err := p2pgo.New(
		context.Background(),
		p2pgo.ListenPort(P2pCommunicationPort),
		p2pgo.Workspace(testWorkspacePath),
		p2pgo.BootPeers(P2pBootstrapNodes),
	)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.FailNow()
	}

	txhash, _, err := cli.Register(cli.GetRoleName(), p2p.GetPeerPublickey(), AliceAddr, 2000)

	if err != nil {
		t.Errorf("Error: %s", err)
		t.FailNow()
	}

	log.Printf("Storage node registration successful, transaction hash is %s\n", txhash)
}
