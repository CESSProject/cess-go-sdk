/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/sdk"
	"github.com/CESSProject/cess-go-sdk/utils"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/retriever"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/state"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/xxhash"
	"github.com/mr-tron/base58"
	"github.com/pkg/errors"
	"github.com/vedhavyas/go-subkey/sr25519"
)

type ChainClient struct {
	lock           *sync.Mutex
	chainStLock    *sync.Mutex
	txTicker       *time.Ticker
	api            *gsrpc.SubstrateAPI
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
	chainState     bool
}

var _ sdk.SDK = (*ChainClient)(nil)

var globalTransport = &http.Transport{
	DisableKeepAlives: true,
}

func NewEmptyChainClient() *ChainClient {
	return &ChainClient{}
}

func NewChainClient(
	ctx context.Context,
	serviceName string,
	rpcs []string,
	mnemonic string,
	t time.Duration,
) (*ChainClient, error) {
	var (
		err         error
		chainClient = &ChainClient{
			lock:        new(sync.Mutex),
			chainStLock: new(sync.Mutex),
			txTicker:    time.NewTicker(pattern.BlockInterval),
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

func (c *ChainClient) ReconnectRPC() error {
	var err error
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.GetChainState() {
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
		c.keyEvents,
		c.eventRetriever,
		c.genesisHash,
		c.currentRpcAddr, err = reconnectChainSDK(c.currentRpcAddr, c.rpcAddr)
	if err != nil {
		return err
	}
	c.SetChainState(true)
	return nil
}

func (c *ChainClient) GetSDKName() string {
	return c.name
}

func (c *ChainClient) GetCurrentRpcAddr() string {
	return c.currentRpcAddr
}

func (c *ChainClient) SetSDKName(name string) {
	c.name = name
}

func (c *ChainClient) SetChainState(state bool) {
	c.chainStLock.Lock()
	c.chainState = state
	c.chainStLock.Unlock()
}

func (c *ChainClient) GetChainState() bool {
	c.chainStLock.Lock()
	st := c.chainState
	c.chainStLock.Unlock()
	return st
}

func (c *ChainClient) GetSignatureAcc() string {
	return c.signatureAcc
}

func (c *ChainClient) GetKeyEvents() types.StorageKey {
	return c.keyEvents
}

func (c *ChainClient) GetSignatureAccPulickey() []byte {
	return c.keyring.PublicKey
}

func (c *ChainClient) GetSubstrateAPI() *gsrpc.SubstrateAPI {
	return c.api
}

func (c *ChainClient) GetMetadata() *types.Metadata {
	return c.metadata
}

func (c *ChainClient) GetTokenSymbol() string {
	return c.tokenSymbol
}

func (c *ChainClient) GetNetworkEnv() string {
	return c.networkEnv
}

func (c *ChainClient) GetURI() string {
	return c.keyring.URI
}

func (c *ChainClient) Sign(msg []byte) ([]byte, error) {
	return signature.Sign(msg, c.keyring.URI)
}

func (c *ChainClient) Verify(msg []byte, sig []byte) (bool, error) {
	return signature.Verify(msg, sig, c.keyring.URI)
}

func reconnectChainSDK(oldRpc string, rpcs []string) (
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

func (c *ChainClient) VerifyPolkaSignatureWithJS(account, msg, signature string) (bool, error) {
	if len(msg) == 0 {
		return false, errors.New("msg is empty")
	}

	pkey, err := utils.ParsingPublickey(account)
	if err != nil {
		return false, err
	}

	pub, err := sr25519.Scheme{}.FromPublicKey(pkey)
	if err != nil {
		return false, err
	}

	sign_bytes, err := hex.DecodeString(strings.TrimPrefix(signature, "0x"))
	if err != nil {
		return false, err
	}
	message := fmt.Sprintf("<Bytes>%s</Bytes>", msg)
	ok := pub.Verify([]byte(message), sign_bytes)
	return ok, nil
}

func (c *ChainClient) VerifyPolkaSignatureWithBase58(account, msg, signature string) (bool, error) {
	if len(msg) == 0 {
		return false, errors.New("msg is empty")
	}

	pkey, err := utils.ParsingPublickey(account)
	if err != nil {
		return false, err
	}

	pub, err := sr25519.Scheme{}.FromPublicKey(pkey)
	if err != nil {
		return false, err
	}

	sign_bytes, err := base58.Decode(signature)
	if err != nil {
		return false, err
	}
	message := fmt.Sprintf("<Bytes>%s</Bytes>", msg)
	ok := pub.Verify([]byte(message), sign_bytes)
	return ok, nil
}

func (c *ChainClient) Close() {
	if c.api.Client != nil {
		c.api.Client.Close()
	}
}
