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
	"os"
	"sync"
	"time"

	"github.com/CESSProject/cess-go-sdk/utils"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/retriever"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/state"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/xxhash"
)

type ChainClient struct {
	lock           *sync.Mutex
	chainStLock    *sync.Mutex
	txTicker       *time.Ticker
	api            *gsrpc.SubstrateAPI
	metadata       *types.Metadata
	runtimeVersion *types.RuntimeVersion
	eventRetriever retriever.EventRetriever
	genesisHash    types.Hash
	keyring        signature.KeyringPair
	rpcAddr        []string
	currentRpcAddr string
	packingTime    time.Duration
	tokenSymbol    string
	networkEnv     string
	signatureAcc   string
	name           string
	rpcState       bool
}

var _ Chainer = (*ChainClient)(nil)

// NewChainClient creates a chainclient
//   - ctx: context
//   - name: customised name, can be empty
//   - rpcs: rpc addresses
//   - mnemonic: account mnemonic, can be empty
//   - t: waiting time for transaction packing, default is 30 seconds
//
// Return:
//   - *ChainClient: chain client
//   - error: error message
func NewChainClient(ctx context.Context, name string, rpcs []string, mnemonic string, t time.Duration) (*ChainClient, error) {
	var (
		err         error
		chainClient = &ChainClient{
			lock:        new(sync.Mutex),
			chainStLock: new(sync.Mutex),
			txTicker:    time.NewTicker(BlockInterval),
			rpcAddr:     rpcs,
			packingTime: t,
			name:        name,
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
		return nil, ERR_RPC_CONNECTION
	}

	chainClient.SetRpcState(true)

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
	properties, err := chainClient.SystemProperties()
	if err != nil {
		return nil, err
	}
	chainClient.tokenSymbol = string(properties.TokenSymbol)

	chainClient.networkEnv, err = chainClient.SystemChain()
	if err != nil {
		return nil, err
	}

	return chainClient, nil
}

// GetSDKName get sdk name
func (c *ChainClient) GetSDKName() string {
	return c.name
}

// GetCurrentRpcAddr get the current rpc address being used
func (c *ChainClient) GetCurrentRpcAddr() string {
	return c.currentRpcAddr
}

// SetChainState set the rpc connection status flag,
// when the rpc connection is normal, set it to true,
// otherwise set it to false.
func (c *ChainClient) SetRpcState(state bool) {
	c.chainStLock.Lock()
	c.rpcState = state
	c.chainStLock.Unlock()
}

// GetRpcState get the rpc connection status flag
//   - true: connection is normal
//   - false: connection failed
func (c *ChainClient) GetRpcState() bool {
	c.chainStLock.Lock()
	st := c.rpcState
	c.chainStLock.Unlock()
	return st
}

// GetSignatureAcc get your current account address
//
// Note:
//   - make sure you fill in mnemonic when you create the chain client
func (c *ChainClient) GetSignatureAcc() string {
	return c.signatureAcc
}

// GetSignatureAccPulickey get your current account public key
//
// Note:
//   - make sure you fill in mnemonic when you create the chain client
func (c *ChainClient) GetSignatureAccPulickey() []byte {
	return c.keyring.PublicKey
}

// GetSubstrateAPI get substrate api
func (c *ChainClient) GetSubstrateAPI() *gsrpc.SubstrateAPI {
	return c.api
}

// GetMetadata get chain metadata
func (c *ChainClient) GetMetadata() *types.Metadata {
	return c.metadata
}

// GetTokenSymbol get token symbol
func (c *ChainClient) GetTokenSymbol() string {
	return c.tokenSymbol
}

// GetNetworkEnv get network env
func (c *ChainClient) GetNetworkEnv() string {
	return c.networkEnv
}

// GetURI get the mnemonic for your current account
func (c *ChainClient) GetURI() string {
	return c.keyring.URI
}

// Sign with the mnemonic of your current account
func (c *ChainClient) Sign(msg []byte) ([]byte, error) {
	return signature.Sign(msg, c.keyring.URI)
}

// Verify the signature with your current account's mnemonic
func (c *ChainClient) Verify(msg []byte, sig []byte) (bool, error) {
	return signature.Verify(msg, sig, c.keyring.URI)
}

// ReconnectRpc reconnect rpc
func (c *ChainClient) ReconnectRpc() error {
	var err error
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.GetRpcState() {
		return nil
	}
	if c.api != nil {
		if c.api.Client != nil {
			c.api.Client.Close()
			c.api.Client = nil
		}
		c.api = nil
	}
	c.api,
		c.metadata,
		c.runtimeVersion,
		c.eventRetriever,
		c.genesisHash,
		c.currentRpcAddr, err = reconnectRpc(c.currentRpcAddr, c.rpcAddr)
	if err != nil {
		return err
	}
	c.SetRpcState(true)
	return nil
}

func reconnectRpc(oldRpc string, rpcs []string) (
	*gsrpc.SubstrateAPI,
	*types.Metadata,
	*types.RuntimeVersion,
	retriever.EventRetriever,
	types.Hash,
	string,
	error,
) {
	var err error
	var rpcAddr string
	var api *gsrpc.SubstrateAPI
	var rpcaddrs = make([]string, 0)
	utils.RandSlice(rpcs)
	for i := 0; i < len(rpcs); i++ {
		if rpcs[i] != oldRpc {
			rpcaddrs = append(rpcaddrs, rpcs[i])
		}
	}
	rpcaddrs = append(rpcaddrs, oldRpc)
	defer log.SetOutput(os.Stdout)
	log.SetOutput(io.Discard)
	for i := 0; i < len(rpcaddrs); i++ {
		if oldRpc == rpcaddrs[i] {
			continue
		}
		api, err = gsrpc.NewSubstrateAPI(rpcaddrs[i])
		if err != nil {
			continue
		}
		rpcAddr = rpcaddrs[i]
	}
	if api == nil {
		return nil, nil, nil, nil, types.Hash{}, rpcAddr, ERR_RPC_CONNECTION
	}
	var metadata *types.Metadata
	var runtimeVer *types.RuntimeVersion
	var genesisHash types.Hash
	var eventRetriever retriever.EventRetriever

	metadata, err = api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, nil, nil, nil, types.Hash{}, rpcAddr, ERR_RPC_CONNECTION
	}
	genesisHash, err = api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, nil, nil, nil, types.Hash{}, rpcAddr, ERR_RPC_CONNECTION
	}
	runtimeVer, err = api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return nil, nil, nil, nil, types.Hash{}, rpcAddr, ERR_RPC_CONNECTION
	}
	eventRetriever, err = retriever.NewDefaultEventRetriever(state.NewEventProvider(api.RPC.State), api.RPC.State)
	if err != nil {
		return nil, nil, nil, nil, types.Hash{}, rpcAddr, ERR_RPC_CONNECTION
	}
	return api, metadata, runtimeVer, eventRetriever, genesisHash, rpcAddr, err
}

func CreatePrefixedKey(pallet, method string) []byte {
	return append(xxhash.New128([]byte(pallet)).Sum(nil), xxhash.New128([]byte(method)).Sum(nil)...)
}

// close chain client
func (c *ChainClient) Close() {
	if c.api != nil {
		if c.api.Client != nil {
			c.api.Client.Close()
		}
		c.api = nil
	}
}
