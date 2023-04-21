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

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type Chain interface {
	// QueryBlockHeight queries the block height corresponding to the block hash.
	// If the blockhash is empty, query the latest block height.
	QueryBlockHeight(blockhash string) (uint32, error)

	// ExtractAccountPuk extracts the public key of the account,
	// and returns its own public key if the account is empty.
	ExtractAccountPuk(account string) ([]byte, error)

	// QueryNodeSynchronizationSt returns the synchronization status of the current node.
	QueryNodeSynchronizationSt() (bool, error)

	// QueryNodeConnectionSt queries the connection status of the node.
	QueryNodeConnectionSt() bool

	// QueryStorageMiner queries storage node information.
	QueryStorageMiner(puk []byte) (MinerInfo, error)

	// QueryDeoss queries deoss information.
	QueryDeoss(puk []byte) (string, error)

	// QuaryAuthorizedAcc queries the account authorized by puk
	QuaryAuthorizedAcc(puk []byte) (types.AccountID, error)

	// QueryBucketInfo queries the information of the "bucketname" bucket of the puk
	QueryBucketInfo(puk []byte, bucketname string) (BucketInfo, error)

	// QueryBucketList queries all buckets of the puk
	QueryBucketList(puk []byte) ([]types.Bytes, error)

	// QueryFileMetaInfo queries the metadata of the roothash file
	QueryFileMetadata(roothash string) (FileMetadata, error)

	// Getallstorageminer is used to obtain the AccountID of all miners
	GetAllStorageMiner() ([]types.AccountID, error)
	// GetCessAccount is used to get the account in cess chain format
	GetCessAccount() (string, error)
	// GetAccountInfo is used to get account information
	GetAccountInfo(pkey []byte) (types.AccountInfo, error)

	// Register is used to register OSS or BUCKET roles
	Register(name, multiaddr string, income string, pledge uint64) (string, error)
	// Update is used to update the communication address of the scheduling service
	UpdateAddress(name, multiaddr string) (string, error)
	//
	UpdateIncomeAccount(pubkey []byte) (string, error)
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
	//
	Exit(role string) (string, error)
	//
	QuerySpacePricePerGib() (string, error)
	//
	QueryNetSnapShot() (NetSnapShot, error)
	//
	QueryTeePodr2Puk() ([]byte, error)
	//
	QueryTeeWorker(pubkey []byte) ([]byte, error)
	//
	QueryTeeWorkerList() ([]TeeWorkerInfo, error)
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
