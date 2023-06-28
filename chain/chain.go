/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"io"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/sdk"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	"github.com/CESSProject/p2p-go/core"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/xxhash"
)

type ChainSDK struct {
	core.P2P
	lock            *sync.Mutex
	api             *gsrpc.SubstrateAPI
	chainState      *atomic.Bool
	metadata        *types.Metadata
	runtimeVersion  *types.RuntimeVersion
	keyEvents       types.StorageKey
	genesisHash     types.Hash
	keyring         signature.KeyringPair
	rpcAddr         []string
	timeForBlockOut time.Duration
	tokenSymbol     string
	signatureAcc    string
	name            string
}

var _ sdk.SDK = (*ChainSDK)(nil)

func NewChainSDK(name string, rpcs []string, mnemonic string, t time.Duration) (*ChainSDK, error) {
	var (
		err      error
		chainSDK = &ChainSDK{}
	)

	defer log.SetOutput(os.Stdout)
	log.SetOutput(io.Discard)

	for i := 0; i < len(rpcs); i++ {
		chainSDK.api, err = gsrpc.NewSubstrateAPI(rpcs[i])
		if err == nil {
			break
		}
	}

	if err != nil || chainSDK.api == nil {
		return nil, err
	}

	chainSDK.metadata, err = chainSDK.api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}
	chainSDK.genesisHash, err = chainSDK.api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, err
	}
	chainSDK.runtimeVersion, err = chainSDK.api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return nil, err
	}
	chainSDK.keyEvents, err = types.CreateStorageKey(chainSDK.metadata, pattern.SYSTEM, pattern.EVENTS, nil)
	if err != nil {
		return nil, err
	}
	if mnemonic != "" {
		chainSDK.keyring, err = signature.KeyringPairFromSecret(mnemonic, 0)
		if err != nil {
			return nil, err
		}
		chainSDK.signatureAcc, err = utils.EncodePublicKeyAsCessAccount(chainSDK.keyring.PublicKey)
		if err != nil {
			return nil, err
		}
	}
	properties, err := chainSDK.SysProperties()
	if err != nil {
		return nil, err
	}
	chainSDK.tokenSymbol = string(properties.TokenSymbol)
	chainSDK.lock = new(sync.Mutex)
	chainSDK.chainState = &atomic.Bool{}
	chainSDK.timeForBlockOut = t
	chainSDK.rpcAddr = rpcs
	chainSDK.SetChainState(true)
	chainSDK.name = name

	return chainSDK, nil
}

func (c *ChainSDK) Reconnect() error {
	var err error
	if c.api != nil {
		if c.api.Client != nil {
			c.api.Client.Close()
			c.api.Client = nil
		}
		c.api = nil
	}

	c.api, c.metadata, c.runtimeVersion, c.keyEvents, c.genesisHash, err = reconnectChainSDK(c.rpcAddr)
	if err != nil {
		return err
	}
	c.SetChainState(true)
	return nil
}

func (c *ChainSDK) SetChainState(state bool) {
	c.chainState.Store(state)
}

func (c *ChainSDK) GetChainState() bool {
	return c.chainState.Load()
}

func (c *ChainSDK) GetSignatureAcc() string {
	return c.signatureAcc
}

func (c *ChainSDK) GetKeyEvents() types.StorageKey {
	return c.keyEvents
}

func (c *ChainSDK) GetSignatureAccPulickey() []byte {
	return c.keyring.PublicKey
}

func (c *ChainSDK) GetSubstrateAPI() *gsrpc.SubstrateAPI {
	return c.api
}

func (c *ChainSDK) GetMetadata() *types.Metadata {
	return c.metadata
}

func (c *ChainSDK) GetTokenSymbol() string {
	return c.tokenSymbol
}

func (c *ChainSDK) GetRoleName() string {
	return c.name
}

func (c *ChainSDK) Sign(msg []byte) ([]byte, error) {
	return signature.Sign(msg, c.keyring.URI)
}

func (c *ChainSDK) Verify(msg []byte, sig []byte) (bool, error) {
	return signature.Verify(msg, sig, c.keyring.URI)
}

func reconnectChainSDK(rpcs []string) (
	*gsrpc.SubstrateAPI,
	*types.Metadata,
	*types.RuntimeVersion,
	types.StorageKey,
	types.Hash,
	error,
) {
	var err error
	var api *gsrpc.SubstrateAPI

	defer log.SetOutput(os.Stdout)
	log.SetOutput(io.Discard)
	for i := 0; i < len(rpcs); i++ {
		api, err = gsrpc.NewSubstrateAPI(rpcs[i])
		if err != nil {
			continue
		}
	}
	if api == nil {
		return nil, nil, nil, nil, types.Hash{}, pattern.ERR_RPC_CONNECTION
	}
	var metadata *types.Metadata
	var runtimeVer *types.RuntimeVersion
	var keyEvents types.StorageKey
	var genesisHash types.Hash

	metadata, err = api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, nil, nil, nil, types.Hash{}, pattern.ERR_RPC_CONNECTION
	}
	genesisHash, err = api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, nil, nil, nil, types.Hash{}, pattern.ERR_RPC_CONNECTION
	}
	runtimeVer, err = api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return nil, nil, nil, nil, types.Hash{}, pattern.ERR_RPC_CONNECTION
	}
	keyEvents, err = types.CreateStorageKey(metadata, pattern.SYSTEM, pattern.EVENTS, nil)
	if err != nil {
		return nil, nil, nil, nil, types.Hash{}, pattern.ERR_RPC_CONNECTION
	}

	return api, metadata, runtimeVer, keyEvents, genesisHash, err
}

func createPrefixedKey(pallet, method string) []byte {
	return append(xxhash.New128([]byte(pallet)).Sum(nil), xxhash.New128([]byte(method)).Sum(nil)...)
}
