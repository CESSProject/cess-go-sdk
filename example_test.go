/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdkgo_test

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/config"
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
