/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"io"
	"log"
	"math/big"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/CESSProject/sdk-go/core/utils"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/xxhash"
)

type Chain interface {
	// QueryBlockHeight queries the block height corresponding to the block hash,
	// If the blockhash is empty, query the latest block height.
	QueryBlockHeight(blockhash string) (uint32, error)

	// QueryNodeSynchronizationSt returns the synchronization status of the current node.
	QueryNodeSynchronizationSt() (bool, error)

	// QueryNodeConnectionSt queries the connection status of the node.
	QueryNodeConnectionSt() bool

	// QueryDeoss queries deoss information.
	QueryDeoss(pubkey []byte) ([]byte, error)

	// QuaryAuthorizedAcc queries the account authorized by puk.
	QuaryAuthorizedAcc(puk []byte) (types.AccountID, error)

	// QueryBucketInfo queries the information of the "bucketname" bucket of the puk.
	QueryBucketInfo(puk []byte, bucketname string) (BucketInfo, error)

	// QueryBucketList queries all buckets of the puk.
	QueryBucketList(puk []byte) ([]types.Bytes, error)

	// QueryFileMetaInfo queries the metadata of the roothash file.
	QueryFileMetadata(roothash string) (FileMetadata, error)

	// QueryStorageMiner queries storage node information.
	QueryStorageMiner(puk []byte) (MinerInfo, error)

	// QuerySminerList queries the accounts of all storage miners.
	QuerySminerList() ([]types.AccountID, error)

	// QueryAccountInfo query account information.
	QueryAccountInfo(puk []byte) (types.AccountInfo, error)

	// QueryStorageOrder query storage order information.
	QueryStorageOrder(roothash string) (StorageOrder, error)

	// QueryPendingReplacements queries the amount of idle data that can be replaced.
	QueryPendingReplacements(puk []byte) (uint32, error)

	// QueryUserSpaceInfo queries the space information purchased by the user.
	QueryUserSpaceInfo(puk []byte) (UserSpaceInfo, error)

	// QuerySpacePricePerGib query space price per GiB.
	QuerySpacePricePerGib() (string, error)

	// QueryChallengeSnapshot query challenge information snapshot.
	QueryChallengeSnapshot() (ChallengeSnapShot, error)

	// QueryTeePodr2Puk queries the public key of the TEE.
	QueryTeePodr2Puk() ([]byte, error)

	// QueryTeePeerID queries the peerid of the Tee worker.
	QueryTeePeerID(puk []byte) ([]byte, error)

	// QueryTeeInfoList queries the information of all tee workers.
	QueryTeeInfoList() ([]TeeWorkerInfo, error)

	//
	QueryAssignedProof() ([][]ProofAssignmentInfo, error)

	//
	QueryTeeAssignedProof(puk []byte) ([]ProofAssignmentInfo, error)

	// Register is used to register OSS or BUCKET roles.
	Register(role string, puk []byte, income string, pledge uint64) (string, error)

	// UpdateAddress updates the address of oss or sminer.
	UpdateAddress(role, multiaddr string) (string, error)

	// UpdateIncomeAcc update income account.
	UpdateIncomeAcc(puk []byte) (string, error)

	// CreateBucket creates a bucket for puk.
	CreateBucket(puk []byte, bucketname string) (string, error)

	// DeleteBucket deletes buckets for puk.
	DeleteBucket(puk []byte, bucketname string) (string, error)

	// DeleteFile deletes files for puk.
	DeleteFile(puk []byte, roothash []string) (string, []FileHash, error)

	// UploadDeclaration creates a storage order.
	UploadDeclaration(roothash string, dealinfo []SegmentList, user UserBrief) (string, error)

	// SubmitIdleMetadata Submit idle file metadata.
	SubmitIdleMetadata(teeAcc []byte, idlefiles []IdleMetadata) (string, error)

	// SubmitFileReport submits a stored file report.
	SubmitFileReport(roothash []FileHash) (string, []FileHash, error)

	// ReplaceIdleFiles replaces idle files.
	ReplaceIdleFiles(roothash []FileHash) (string, []FileHash, error)

	// IncreaseStakes increase stakes.
	IncreaseStakes(tokens *big.Int) (string, error)

	// Exit exit the cess network.
	Exit(role string) (string, error)

	// ClaimRewards is used to claim rewards
	ClaimRewards() (string, error)

	// Withdraw is used to withdraw staking
	Withdraw() (string, error)

	//
	ReportProof(idlesigma, servicesigma string) (string, error)

	// ExtractAccountPuk extracts the public key of the account,
	// and returns its own public key if the account is empty.
	ExtractAccountPuk(account string) ([]byte, error)

	// GetSignatureAcc returns the signature account.
	GetSignatureAcc() string

	// GetSignatureURI to get the private key of the signing account
	GetSignatureURI() string

	// GetSubstrateAPI returns Substrate API
	GetSubstrateAPI() *gsrpc.SubstrateAPI

	// GetChainState returns chain node state
	GetChainState() bool

	//
	SetChainState(state bool)

	//
	Reconnect() error

	//
	GetMetadata() *types.Metadata

	//
	GetKeyEvents() types.StorageKey

	//
	SysProperties() (SysProperties, error)

	//
	SyncState() (SysSyncState, error)

	//
	SysVersion() (string, error)

	//
	NetListening() (bool, error)
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
	tokenSymbol     string
}

func NewChainClient(rpcAddr []string, secret string, t time.Duration) (Chain, error) {
	var (
		err error
		cli = &chainClient{}
	)

	defer log.SetOutput(os.Stdout)
	log.SetOutput(io.Discard)

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
	cli.keyEvents, err = types.CreateStorageKey(cli.metadata, SYSTEM, EVENTS, nil)
	if err != nil {
		return nil, err
	}
	if secret != "" {
		cli.keyring, err = signature.KeyringPairFromSecret(secret, 0)
		if err != nil {
			return nil, err
		}
	}
	properties, err := cli.SysProperties()
	if err != nil {
		return nil, err
	}
	cli.tokenSymbol = string(properties.TokenSymbol)
	cli.lock = new(sync.Mutex)
	cli.chainState = &atomic.Bool{}
	cli.chainState.Store(true)
	cli.timeForBlockOut = t
	cli.rpcAddr = rpcAddr
	cli.SetChainState(true)
	return cli, nil
}

func (c *chainClient) Reconnect() error {
	var err error
	if c.api.Client != nil {
		c.api.Client.Close()
	}
	c.api = nil
	c.api, err = reconnectChainClient(c.rpcAddr)
	if err != nil {
		return err
	}
	c.metadata, err = c.api.RPC.State.GetMetadataLatest()
	if err != nil {
		return err
	}
	return nil
}

func (c *chainClient) SetChainState(state bool) {
	c.chainState.Store(state)
}

func (c *chainClient) GetChainState() bool {
	return c.chainState.Load()
}

// QueryNodeConnectionSt
func (c *chainClient) QueryNodeConnectionSt() bool {
	return c.GetChainState()
}

func (c *chainClient) NewAccountId(pubkey []byte) types.AccountID {
	acc, _ := types.NewAccountID(pubkey)
	return *acc
}

func (c *chainClient) GetSignatureAcc() string {
	acc, _ := utils.EncodePublicKeyAsCessAccount(c.keyring.PublicKey)
	return acc
}

func (c *chainClient) GetKeyEvents() types.StorageKey {
	return c.keyEvents
}

// ExtractAccountPublickey
func (c *chainClient) ExtractAccountPuk(account string) ([]byte, error) {
	if account != "" {
		return utils.ParsingPublickey(account)
	}
	return c.keyring.PublicKey, nil
}

func (c *chainClient) GetSignatureURI() string {
	return c.keyring.URI
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

func reconnectChainClient(rpcAddr []string) (*gsrpc.SubstrateAPI, error) {
	var err error
	var api *gsrpc.SubstrateAPI
	defer log.SetOutput(os.Stdout)
	log.SetOutput(io.Discard)
	for i := 0; i < len(rpcAddr); i++ {
		api, err = gsrpc.NewSubstrateAPI(rpcAddr[i])
		if err == nil {
			return api, nil
		}
	}
	return api, err
}

func createPrefixedKey(pallet, method string) []byte {
	return append(xxhash.New128([]byte(pallet)).Sum(nil), xxhash.New128([]byte(method)).Sum(nil)...)
}
