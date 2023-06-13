/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdkgo_test

import (
	"context"
	"testing"
	"time"

	p2pgo "github.com/CESSProject/p2p-go"
	cess "github.com/CESSProject/sdk-go"
	"github.com/CESSProject/sdk-go/config"
	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/stretchr/testify/assert"
)

// Substrate well-known mnemonic:
//
//	https://github.com/substrate-developer-hub/substrate-developer-hub.github.io/issues/613
var Test_Account = "cXkdXokcMa32BAYkmsGjhRGA2CYmLUN2pq69U8k9taXsQPHGp"                              //Alice
var Test_Account_Mnemonic = "bottom drive obey lake curtain smoke basket hold race lonely fit walk" //Alice

var testnets = []string{
	"wss://devnet-rpc.cess.cloud/ws/",
	"wss://testnet-rpc0.cess.cloud/ws/",
	"wss://testnet-rpc1.cess.cloud/ws/",
}

var localNode = []string{
	"ws://localhost:9944",
}

var Test_BootstrapNodes = []string{
	"_dnsaddr.bootstrap-kldr.cess.cloud",
}

// Tmp files will be downloaded
var Test_WorkspacePath_Deoss = "/tmp/deoss"
var Test_WorkspacePath_Bucket = "/tmp/bucket"

const Test_ListeningPort = 4001

// If you run these examples using localNode, please run a [CESS node](https://github.com/cessProject/cess) locally
//  as well.

func TestNewClient(t *testing.T) {
	_, err := cess.New(
		config.CharacterName_Client,
		cess.ConnectRpcAddrs(testnets),
		cess.Mnemonic(Test_Account_Mnemonic),
		cess.TransactionTimeout(time.Duration(time.Second*15)),
	)
	assert.NoError(t, err)
}

func TestRegisterDeOSS(t *testing.T) {
	cli, err := cess.New(
		config.CharacterName_Deoss,
		cess.ConnectRpcAddrs(testnets),
		cess.Mnemonic(Test_Account_Mnemonic),
		cess.TransactionTimeout(time.Duration(time.Second*15)),
	)
	assert.NoError(t, err)

	var bootnodes = make([]string, 0)

	for _, v := range Test_BootstrapNodes {
		temp, err := utils.ParseMultiaddrs(v)
		if err != nil {
			continue
		}
		bootnodes = append(bootnodes, temp...)
	}

	p2p, err := p2pgo.New(
		context.Background(),
		p2pgo.ListenPort(Test_ListeningPort),
		p2pgo.Workspace(Test_WorkspacePath_Deoss),
		p2pgo.BootPeers(bootnodes),
	)
	assert.NoError(t, err)

	_, _, err = cli.Register(cli.GetRoleName(), p2p.GetPeerPublickey(), "", 0)
	assert.NoError(t, err)
}

func TestRegisterStorageNode(t *testing.T) {
	cli, err := cess.New(
		config.CharacterName_Bucket,
		cess.ConnectRpcAddrs(testnets),
		cess.Mnemonic(Test_Account_Mnemonic),
		cess.TransactionTimeout(time.Duration(time.Second*10)),
	)
	assert.NoError(t, err)

	var bootnodes = make([]string, 0)

	for _, v := range Test_BootstrapNodes {
		temp, err := utils.ParseMultiaddrs(v)
		if err != nil {
			continue
		}
		bootnodes = append(bootnodes, temp...)
	}

	p2p, err := p2pgo.New(
		context.Background(),
		p2pgo.ListenPort(Test_ListeningPort),
		p2pgo.Workspace(Test_WorkspacePath_Bucket),
		p2pgo.BootPeers(bootnodes),
	)
	assert.NoError(t, err)

	_, _, err = cli.Register(cli.GetRoleName(), p2p.GetPeerPublickey(), Test_Account, 0)
	assert.NoError(t, err)
}
