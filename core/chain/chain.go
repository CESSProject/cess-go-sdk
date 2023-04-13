/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.
	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/CESSProject/sdk-go/core/utils"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type Chain interface {
	// Getpublickey returns its own public key
	GetPublicKey() []byte
	// Getpublickey returns its own private key
	GetMnemonicSeed() string
	// NewAccountId returns the account id
	NewAccountId(pubkey []byte) types.AccountID
	// GetSyncStatus returns whether the block is being synchronized
	GetSyncStatus() (bool, error)
	// GetChainStatus returns chain status
	GetChainStatus() bool
	// Getstorageminerinfo is used to get the details of the miner
	QueryStorageMiner(pkey []byte) (MinerInfo, error)
	//
	QueryDeoss(pubkey []byte) (string, error)
	// Getallstorageminer is used to obtain the AccountID of all miners
	GetAllStorageMiner() ([]types.AccountID, error)
	// GetFileMetaInfo is used to get the meta information of the file
	GetFileMetaInfo(fid string) (FileMetaInfo, error)
	// GetCessAccount is used to get the account in cess chain format
	GetCessAccount() (string, error)
	// GetAccountInfo is used to get account information
	GetAccountInfo(pkey []byte) (types.AccountInfo, error)
	//
	IsGrantor(pubkey []byte) (bool, error)

	// GetBucketList is used to obtain all buckets of the user
	GetBucketList(owner_pkey []byte) ([]types.Bytes, error)
	// GetBucketInfo is used to query bucket details
	GetBucketInfo(owner_pkey []byte, name string) (BucketInfo, error)
	// GetGrantor is used to query the user's space grantor
	GetGrantor(pkey []byte) (types.AccountID, error)
	// GetState is used to obtain OSS status information
	GetState(pubkey []byte) (string, error)
	// Register is used to register OSS or BUCKET roles
	Register(name, multiaddr string, income string, pledge uint64) (string, error)
	// Update is used to update the communication address of the scheduling service
	Update(name, multiaddr string) (string, error)
	// CreateBucket is used to create a bucket for users
	CreateBucket(owner_pkey []byte, name string) (string, error)
	// DeleteBucket is used to delete buckets created by users
	DeleteBucket(owner_pkey []byte, name string) (string, error)
	//
	DeleteFile(owner_pkey []byte, filehash string) (string, FileHash, error)
	//
	UploadDeclaration(filehash string, dealinfo []SegmentList, user UserBrief) (string, error)
	//
	GetStorageOrder(roothash string) (StorageOrder, error)
	//
	SubmitIdleFile(idlefiles []IdleMetaInfo) (string, error)
	//
	SubmitFileReport(roothash []FileHash) (string, []FileHash, error)
	//
	ReplaceFile(roothash []FileHash) (string, []FileHash, error)
	//
	QueryPendingReplacements(owner_pkey []byte) (types.U32, error)
	//
	QueryUserSpaceInfo(pubkey []byte) (UserSpaceInfo, error)
	//
	IncreaseStakes(tokens *big.Int) (string, error)
}

type chainClient struct {
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
}

func NewChainClient(rpcAddr []string, secret string, t time.Duration) (Chain, error) {
	var (
		err error
		cli = &chainClient{}
	)

	for i := 0; i < len(rpcAddr); i++ {
		cli.api, err = gsrpc.NewSubstrateAPI(rpcAddr[i])
		if err == nil {
			break
		}
	}

	if err != nil || cli.api == nil {
		return nil, err
	}

	cli.metadata, err = cli.api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}
	cli.genesisHash, err = cli.api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return nil, err
	}
	cli.runtimeVersion, err = cli.api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return nil, err
	}
	cli.keyEvents, err = types.CreateStorageKey(
		cli.metadata,
		SYSTEM,
		EVENTS,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if secret != "" {
		cli.keyring, err = signature.KeyringPairFromSecret(secret, 0)
		if err != nil {
			return nil, err
		}
	}
	cli.lock = new(sync.Mutex)
	cli.chainState = &atomic.Bool{}
	cli.chainState.Store(true)
	cli.timeForBlockOut = t
	cli.rpcAddr = rpcAddr
	return cli, nil
}

func (c *chainClient) IsChainClientOk() bool {
	err := healthchek(c.api)
	if err != nil {
		c.api = nil
		cli, err := reconnectChainClient(c.rpcAddr)
		if err != nil {
			return false
		}
		c.api = cli
		c.metadata, err = c.api.RPC.State.GetMetadataLatest()
		if err != nil {
			return false
		}
		return true
	}
	return true
}

func (c *chainClient) SetChainState(state bool) {
	c.chainState.Store(state)
}

func (c *chainClient) GetChainState() bool {
	return c.chainState.Load()
}

func (c *chainClient) NewAccountId(pubkey []byte) types.AccountID {
	acc, _ := types.NewAccountID(pubkey)
	return *acc
}

func reconnectChainClient(rpcAddr []string) (*gsrpc.SubstrateAPI, error) {
	var err error
	var api *gsrpc.SubstrateAPI
	for i := 0; i < len(rpcAddr); i++ {
		api, err = gsrpc.NewSubstrateAPI(rpcAddr[i])
		if err == nil {
			break
		}
	}
	return api, err
}

func healthchek(a *gsrpc.SubstrateAPI) error {
	defer func() { recover() }()
	_, err := a.RPC.System.Health()
	return err
}

func (c *chainClient) KeepConnect() {
	tick := time.NewTicker(time.Second * 20)
	select {
	case <-tick.C:
		healthchek(c.api)
	}
}

// VerifyGrantor is used to verify whether the right to use the space is authorized
func (c *chainClient) IsGrantor(pubkey []byte) (bool, error) {
	var (
		err     error
		grantor types.AccountID
	)

	grantor, err = c.GetGrantor(pubkey)
	if err != nil {
		if err.Error() == ERR_Empty {
			return false, nil
		}
		return false, err
	}
	account_chain, _ := utils.EncodePublicKeyAsCessAccount(grantor[:])
	account_local, _ := c.GetCessAccount()
	return account_chain == account_local, nil
}
