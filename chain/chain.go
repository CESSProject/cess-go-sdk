/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/sdk"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	p2pgo "github.com/CESSProject/p2p-go"
	"github.com/CESSProject/p2p-go/core"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/xxhash"
)

type Sdk struct {
	*core.Node
	lock           *sync.Mutex
	api            *gsrpc.SubstrateAPI
	chainState     *atomic.Bool
	metadata       *types.Metadata
	runtimeVersion *types.RuntimeVersion
	keyEvents      types.StorageKey
	genesisHash    types.Hash
	keyring        signature.KeyringPair
	rpcAddr        []string
	packingTime    time.Duration
	tokenSymbol    string
	networkEnv     string
	signatureAcc   string
	name           string
	enabledP2P     bool
}

var _ sdk.SDK = (*Sdk)(nil)

var globalTransport = &http.Transport{
	DisableKeepAlives: true,
}

func NewSDK(
	ctx context.Context,
	serviceName string,
	rpcs []string,
	mnemonic string,
	t time.Duration,
	workspace string,
	p2pPort int,
	bootnodes []string,
	protocolPrefix string,
) (*Sdk, error) {
	var (
		ok  bool
		err error
		sdk = &Sdk{
			lock:        new(sync.Mutex),
			chainState:  new(atomic.Bool),
			rpcAddr:     rpcs,
			packingTime: t,
			name:        serviceName,
		}
	)

	if !core.FreeLocalPort(uint32(p2pPort)) {
		return nil, fmt.Errorf("port [%d] is in use", p2pPort)
	}

	log.SetOutput(io.Discard)
	for i := 0; i < len(rpcs); i++ {
		sdk.api, err = gsrpc.NewSubstrateAPI(rpcs[i])
		if err == nil {
			break
		}
	}
	log.SetOutput(os.Stdout)
	if err != nil {
		return nil, err
	}

	if sdk.api == nil {
		return nil, pattern.ERR_RPC_CONNECTION
	}

	sdk.SetChainState(true)

	sdk.metadata, err = sdk.api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}
	sdk.genesisHash, err = sdk.api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, err
	}
	sdk.runtimeVersion, err = sdk.api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return nil, err
	}
	sdk.keyEvents, err = types.CreateStorageKey(sdk.metadata, pattern.SYSTEM, pattern.EVENTS, nil)
	if err != nil {
		return nil, err
	}
	if mnemonic != "" {
		sdk.keyring, err = signature.KeyringPairFromSecret(mnemonic, 0)
		if err != nil {
			return nil, err
		}
		sdk.signatureAcc, err = utils.EncodePublicKeyAsCessAccount(sdk.keyring.PublicKey)
		if err != nil {
			return nil, err
		}
	}
	properties, err := sdk.SysProperties()
	if err != nil {
		return nil, err
	}
	sdk.tokenSymbol = string(properties.TokenSymbol)

	sdk.networkEnv, err = sdk.SysChain()
	if err != nil {
		return nil, err
	}

	if workspace != "" && p2pPort > 0 {
		p2p, err := p2pgo.New(
			ctx,
			p2pgo.ListenPort(p2pPort),
			p2pgo.Workspace(filepath.Join(workspace, sdk.GetSignatureAcc(), sdk.GetSdkName())),
			p2pgo.BootPeers(bootnodes),
			p2pgo.ProtocolPrefix(protocolPrefix),
		)
		if err != nil {
			return nil, err
		}
		sdk.Node, ok = p2p.(*core.Node)
		if !ok {
			return nil, errors.New("invalid p2p type")
		}
		sdk.enabledP2P = true
	}

	return sdk, nil
}

func (c *Sdk) Reconnect() error {
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

func (c *Sdk) GetSdkName() string {
	return c.name
}

func (c *Sdk) SetSdkName(name string) {
	c.name = name
}

func (c *Sdk) SetChainState(state bool) {
	c.chainState.Store(state)
}

func (c *Sdk) GetChainState() bool {
	return c.chainState.Load()
}

func (c *Sdk) GetSignatureAcc() string {
	return c.signatureAcc
}

func (c *Sdk) GetKeyEvents() types.StorageKey {
	return c.keyEvents
}

func (c *Sdk) GetSignatureAccPulickey() []byte {
	return c.keyring.PublicKey
}

func (c *Sdk) GetSubstrateAPI() *gsrpc.SubstrateAPI {
	return c.api
}

func (c *Sdk) GetMetadata() *types.Metadata {
	return c.metadata
}

func (c *Sdk) GetTokenSymbol() string {
	return c.tokenSymbol
}

func (c *Sdk) GetNetworkEnv() string {
	return c.networkEnv
}

func (c *Sdk) GetURI() string {
	return c.keyring.URI
}

func (c *Sdk) Sign(msg []byte) ([]byte, error) {
	return signature.Sign(msg, c.keyring.URI)
}

func (c *Sdk) Verify(msg []byte, sig []byte) (bool, error) {
	return signature.Verify(msg, sig, c.keyring.URI)
}

func (c *Sdk) EnabledP2P() bool {
	return c.enabledP2P
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
