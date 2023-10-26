/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/sdk"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/retriever"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/state"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/xxhash"
)

type chainClient struct {
	lock           *sync.Mutex
	api            *gsrpc.SubstrateAPI
	chainState     *atomic.Bool
	metadata       *types.Metadata
	runtimeVersion *types.RuntimeVersion
	eventRetriever retriever.EventRetriever
	keyEvents      types.StorageKey
	genesisHash    types.Hash
	keyring        signature.KeyringPair
	rpcAddr        []string
	currentRpcAddr string
	packingTime    time.Duration
	tokenSymbol    string
	networkEnv     string
	signatureAcc   string
	name           string
}

var _ sdk.SDK = (*chainClient)(nil)

var globalTransport = &http.Transport{
	DisableKeepAlives: true,
}

func NewChainClient(
	ctx context.Context,
	serviceName string,
	rpcs []string,
	mnemonic string,
	t time.Duration,
) (*chainClient, error) {
	var (
		err         error
		chainClient = &chainClient{
			lock:        new(sync.Mutex),
			chainState:  new(atomic.Bool),
			rpcAddr:     rpcs,
			packingTime: t,
			name:        serviceName,
		}
	)

	log.SetOutput(io.Discard)
	for i := 0; i < len(rpcs); i++ {
		chainClient.api, err = gsrpc.NewSubstrateAPI(rpcs[i])
		if err == nil {
			chainClient.currentRpcAddr = rpcs[i]
			break
		}
	}
	log.SetOutput(os.Stdout)
	if err != nil {
		return nil, err
	}

	if chainClient.api == nil {
		return nil, pattern.ERR_RPC_CONNECTION
	}

	chainClient.SetChainState(true)

	chainClient.metadata, err = chainClient.api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}
	chainClient.genesisHash, err = chainClient.api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, err
	}
	chainClient.runtimeVersion, err = chainClient.api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return nil, err
	}
	chainClient.keyEvents, err = types.CreateStorageKey(chainClient.metadata, pattern.SYSTEM, pattern.EVENTS, nil)
	if err != nil {
		return nil, err
	}
	chainClient.eventRetriever, err = retriever.NewDefaultEventRetriever(state.NewEventProvider(chainClient.api.RPC.State), chainClient.api.RPC.State)
	if err != nil {
		return nil, err
	}
	if mnemonic != "" {
		chainClient.keyring, err = signature.KeyringPairFromSecret(mnemonic, 0)
		if err != nil {
			return nil, err
		}
		chainClient.signatureAcc, err = utils.EncodePublicKeyAsCessAccount(chainClient.keyring.PublicKey)
		if err != nil {
			return nil, err
		}
	}
	properties, err := chainClient.SysProperties()
	if err != nil {
		return nil, err
	}
	chainClient.tokenSymbol = string(properties.TokenSymbol)

	chainClient.networkEnv, err = chainClient.SysChain()
	if err != nil {
		return nil, err
	}

	return chainClient, nil
}

func (c *chainClient) Reconnect() error {
	var err error
	if c.api != nil {
		if c.api.Client != nil {
			c.api.Client.Close()
			c.api.Client = nil
		}
		c.api = nil
	}

	c.api, c.metadata, c.runtimeVersion, c.keyEvents, c.eventRetriever, c.genesisHash, c.currentRpcAddr, err = reconnectChainSDK(c.rpcAddr)
	if err != nil {
		return err
	}
	c.SetChainState(true)
	return nil
}

func (c *chainClient) GetSdkName() string {
	return c.name
}

func (c *chainClient) GetCurrentRpcAddr() string {
	return c.currentRpcAddr
}

func (c *chainClient) SetSdkName(name string) {
	c.name = name
}

func (c *chainClient) SetChainState(state bool) {
	c.chainState.Store(state)
}

func (c *chainClient) GetChainState() bool {
	return c.chainState.Load()
}

func (c *chainClient) GetSignatureAcc() string {
	return c.signatureAcc
}

func (c *chainClient) GetKeyEvents() types.StorageKey {
	return c.keyEvents
}

func (c *chainClient) GetSignatureAccPulickey() []byte {
	return c.keyring.PublicKey
}

func (c *chainClient) GetSubstrateAPI() *gsrpc.SubstrateAPI {
	return c.api
}

func (c *chainClient) GetMetadata() *types.Metadata {
	return c.metadata
}

func (c *chainClient) GetTokenSymbol() string {
	return c.tokenSymbol
}

func (c *chainClient) GetNetworkEnv() string {
	return c.networkEnv
}

func (c *chainClient) GetURI() string {
	return c.keyring.URI
}

func (c *chainClient) Sign(msg []byte) ([]byte, error) {
	return signature.Sign(msg, c.keyring.URI)
}

func (c *chainClient) Verify(msg []byte, sig []byte) (bool, error) {
	return signature.Verify(msg, sig, c.keyring.URI)
}

func reconnectChainSDK(rpcs []string) (
	*gsrpc.SubstrateAPI,
	*types.Metadata,
	*types.RuntimeVersion,
	types.StorageKey,
	retriever.EventRetriever,
	types.Hash,
	string,
	error,
) {
	var err error
	var rpcAddr string
	var api *gsrpc.SubstrateAPI

	defer log.SetOutput(os.Stdout)
	log.SetOutput(io.Discard)
	for i := 0; i < len(rpcs); i++ {
		api, err = gsrpc.NewSubstrateAPI(rpcs[i])
		if err != nil {
			continue
		}
		rpcAddr = rpcs[i]
	}
	if api == nil {
		return nil, nil, nil, nil, nil, types.Hash{}, rpcAddr, pattern.ERR_RPC_CONNECTION
	}
	var metadata *types.Metadata
	var runtimeVer *types.RuntimeVersion
	var keyEvents types.StorageKey
	var genesisHash types.Hash
	var eventRetriever retriever.EventRetriever

	metadata, err = api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, nil, nil, nil, nil, types.Hash{}, rpcAddr, pattern.ERR_RPC_CONNECTION
	}
	genesisHash, err = api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, nil, nil, nil, nil, types.Hash{}, rpcAddr, pattern.ERR_RPC_CONNECTION
	}
	runtimeVer, err = api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return nil, nil, nil, nil, nil, types.Hash{}, rpcAddr, pattern.ERR_RPC_CONNECTION
	}
	keyEvents, err = types.CreateStorageKey(metadata, pattern.SYSTEM, pattern.EVENTS, nil)
	if err != nil {
		return nil, nil, nil, nil, nil, types.Hash{}, rpcAddr, pattern.ERR_RPC_CONNECTION
	}
	eventRetriever, err = retriever.NewDefaultEventRetriever(state.NewEventProvider(api.RPC.State), api.RPC.State)
	if err != nil {
		return nil, nil, nil, nil, nil, types.Hash{}, rpcAddr, pattern.ERR_RPC_CONNECTION
	}
	return api, metadata, runtimeVer, keyEvents, eventRetriever, genesisHash, rpcAddr, err
}

func createPrefixedKey(pallet, method string) []byte {
	return append(xxhash.New128([]byte(pallet)).Sum(nil), xxhash.New128([]byte(method)).Sum(nil)...)
}
