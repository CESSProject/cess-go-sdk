/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdk

import (
	"io"
	"math/big"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/p2p-go/core"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// CESS Go SDK Interface Description
type SDK interface {
	// References libp2p: https://github.com/libp2p/go-libp2p
	core.P2P

	// Audit-State

	// Query Challenge Expired Block Height Interface
	QueryChallengeExpiration() (uint32, error)

	// QueryChallengeSnapshot query challenge information snapshot.
	QueryChallengeSnapshot() (pattern.ChallengeSnapShot, error)
	QueryChallengeSt() (pattern.ChallengeSnapshot, error)
	// QueryChallenge queries puk's challenge information.
	QueryChallenge(puk []byte) (pattern.ChallengeInfo, error)
	// QueryAssignedProof queries all assigned proof information.
	QueryAssignedProof() ([][]pattern.ProofAssignmentInfo, error)
	// ProofAssignmentInfo queries the proof information assigned to puk.
	QueryTeeAssignedProof(puk []byte) ([]pattern.ProofAssignmentInfo, error)
	//
	QueryChallenge_V2() (pattern.ChallengeInfo_V2, error)
	//
	QueryUnverifiedIdleProof() (pattern.IdleProofInfo, error)
	//
	QueryUnverifiedServiceProof() (pattern.ServiceProofInfo, error)

	// Audit-Extrinsics

	// ReportProof is used to report proof data.
	ReportProof(idlesigma, servicesigma string) (string, error)
	//
	SubmitIdleProof(idleProve []types.U8) (string, error)
	//
	SubmitServiceProof(serviceProof []types.U8) (string, error)
	//
	SubmitIdleProofResult(puk []byte, result types.Bool, signature pattern.TeeSignature) (string, error)
	//
	SubmitServiceProofResult(puk []byte, result types.Bool, signature pattern.TeeSignature) (string, error)

	// Filebank-State

	// QueryBucketInfo queries the information of the "bucketname" bucket of the puk.
	QueryBucketInfo(puk []byte, bucketname string) (pattern.BucketInfo, error)
	// QueryStorageOrder query storage order information.
	QueryStorageOrder(roothash string) (pattern.StorageOrder, error)
	// QueryFileMetadata queries the metadata of the roothash file.
	QueryFileMetadata(roothash string) (pattern.FileMetadata, error)
	// QueryFillerMap queries filler information.
	QueryFillerMap(filehash string) (pattern.IdleMetadata, error)
	// QueryPendingReplacements queries the amount of idle data that can be replaced.
	QueryPendingReplacements(puk []byte) (uint32, error)
	// QueryRestoralOrder queries a restore order info.
	QueryRestoralOrder(roothash string) (pattern.RestoralOrderInfo, error)
	QueryRestoralOrderList() ([]pattern.RestoralOrderInfo, error)
	// QueryRestoralTargetList for recovery information on all exiting miners.
	QueryRestoralTarget(puk []byte) (pattern.RestoralTargetInfo, error)
	QueryRestoralTargetList() ([]pattern.RestoralTargetInfo, error)
	// QueryBucketList queries all buckets of the puk.
	// QueryAllBucketName queries all bucket names as strings.
	QueryBucketList(puk []byte) ([]types.Bytes, error)
	QueryAllBucketName(puk []byte) ([]string, error)

	// Filebank-Extrinsics

	// ClaimRestoralNoExistOrder is used to receive recovery orders from exiting miners.
	ClaimRestoralNoExistOrder(puk []byte, rootHash, restoralFragmentHash string) (string, error)
	// ClaimRestoralOrder is used to collect restoration orders.
	ClaimRestoralOrder(fragmentHash string) (string, error)
	// CreateBucket creates a bucket for puk.
	CreateBucket(puk []byte, bucketname string) (string, error)
	// DeleteBucket deletes buckets for puk.
	DeleteBucket(puk []byte, bucketname string) (string, error)
	// DeleteFile deletes files for puk.
	DeleteFile(puk []byte, roothash []string) (string, []pattern.FileHash, error)
	// DeleteFiller deletes an idle file.
	DeleteFiller(filehash string) (string, error)
	// GenerateRestoralOrder generates data for restoration orders.
	GenerateRestoralOrder(rootHash, fragmentHash string) (string, error)
	// Withdraw is used to withdraw staking.
	Withdraw() (string, error)
	// ReplaceIdleFiles replaces idle files.
	ReplaceIdleFiles(roothash []pattern.FileHash) (string, []pattern.FileHash, error)
	ReplaceFile(roothash []string) (string, []string, error)
	// RestoralComplete reports order recovery completion.
	RestoralComplete(restoralFragmentHash string) (string, error)
	// SubmitFileReport submits a stored file report.
	SubmitFileReport(roothash []pattern.FileHash) (string, []pattern.FileHash, error)
	ReportFiles(roothash []string) (string, []string, error)
	// UploadDeclaration creates a storage order.
	UploadDeclaration(roothash string, dealinfo []pattern.SegmentList, user pattern.UserBrief, filesize uint64) (string, error)
	// GenerateStorageOrder for generating storage orders
	GenerateStorageOrder(roothash string, segment []pattern.SegmentDataInfo, owner []byte, filename, buckname string, filesize uint64) (string, error)
	// SubmitIdleMetadata Submit idle file metadata.
	SubmitIdleMetadata(teeAcc []byte, idlefiles []pattern.IdleMetadata) (string, error)
	SubmitIdleFile(teeAcc []byte, idlefiles []pattern.IdleFileMeta) (string, error)
	// CertIdleSpace
	CertIdleSpace(idleSignInfo pattern.SpaceProofInfo, sign pattern.TeeSignature) (string, error)
	// ReplaceIdleSpace
	ReplaceIdleSpace(idleSignInfo pattern.SpaceProofInfo, sign pattern.TeeSignature) (string, error)

	// Oss-State

	// QuaryAuthorizedAcc queries the account authorized by puk.
	// QuaryAuthorizedAccount query account in string form.
	QuaryAuthorizedAcc(puk []byte) (types.AccountID, error)
	QuaryAuthorizedAccount(puk []byte) (string, error)
	// QueryDeossPeerPublickey queries deoss peer public key.
	QueryDeossPeerPublickey(pubkey []byte) ([]byte, error)
	// QueryDeossPeerIdList queries peerid of all deoss.
	QueryDeossPeerIdList() ([]string, error)
	// CheckSpaceUsageAuthorization checks if the puk is authorized to itself
	CheckSpaceUsageAuthorization(puk []byte) (bool, error)

	// Oss-Extrinsics

	// RegisterOrUpdateDeoss register or update deoss information
	RegisterOrUpdateDeoss(peerId []byte) (string, error)
	// ExitDeoss exit deoss
	ExitDeoss() (string, error)
	// AuthorizeSpace authorizes space to oss
	AuthorizeSpace(ossAccount string) (string, error)
	// UnAuthorizeSpace cancels space authorization
	UnAuthorizeSpace() (string, error)

	// Sminer-State

	// QueryStorageMiner queries storage node information.
	QueryStorageMiner(puk []byte) (pattern.MinerInfo, error)
	// QueryStorageMiner queries storage node information.
	QueryStorageMiner_V2(puk []byte) (pattern.MinerInfo_V2, error)
	// QuerySminerList queries the accounts of all storage miners.
	QuerySminerList() ([]types.AccountID, error)
	// QueryStorageNodeReward queries reward information for puk account.
	QueryStorageNodeReward(puk []byte) (pattern.MinerReward, error)
	QuaryStorageNodeRewardInfo(puk []byte) (pattern.RewardsType, error)

	//
	Expenders() (pattern.ExpendersInfo, error)

	// Sminer-Extrinsics

	// RegisterOrUpdateSminer register or update sminer information
	RegisterOrUpdateSminer(peerId []byte, earnings string, pledge uint64) (string, string, error)
	//
	RegisterOrUpdateSminer_V2(peerId []byte, earnings string, pledge uint64, poisKey pattern.PoISKeyInfo, sign pattern.TeeSignature) (string, string, error)
	// ExitSminer exit mining
	ExitSminer() (string, error)
	// UpdateEarningsAcc update earnings account.
	UpdateEarningsAcc(puk []byte) (string, error)
	UpdateEarningsAccount(earnings string) (string, error)
	// UpdateSminerPeerId update miner peerid
	UpdateSminerPeerId(peerid pattern.PeerId) (string, error)
	// IncreaseStakingAmount increase staking amount.
	IncreaseStakingAmount(tokens *big.Int) (string, error)
	IncreaseStorageNodeStakingAmount(token string) (string, error)
	// ClaimRewards is used to claim rewards.
	ClaimRewards() (string, error)

	// StorageHandler-State

	// QueryUserSpaceInfo queries the space information purchased by the user.
	QueryUserSpaceInfo(puk []byte) (pattern.UserSpaceInfo, error)
	QueryUserSpaceSt(puk []byte) (pattern.UserSpaceSt, error)
	// QuerySpacePricePerGib query space price per GiB.
	QuerySpacePricePerGib() (string, error)

	// StorageHandler-Extrinsics

	// BuySpace for purchasing space.
	BuySpace(count uint32) (string, error)
	// ExpansionSpace for expansion space.
	ExpansionSpace(count uint32) (string, error)
	// RenewalSpace is used to extend the validity of space.
	RenewalSpace(days uint32) (string, error)

	// TeeWorker-State

	// QueryTeePodr2Puk queries the public key of the TEE.
	QueryTeePodr2Puk() ([]byte, error)
	// QueryTeePeerID queries the peerid of the Tee worker.
	QueryTeePeerID(puk []byte) ([]byte, error)
	// QueryTeeInfoList queries the information of all tee workers.
	QueryTeeInfoList() ([]pattern.TeeWorkerInfo, error)
	QueryTeeWorkerList() ([]pattern.TeeWorkerSt, error)

	// System

	// QueryBlockHeight queries the block height corresponding to the block hash,
	// If the blockhash is empty, query the latest block height.
	QueryBlockHeight(blockhash string) (uint32, error)
	// QueryNodeSynchronizationSt returns the synchronization status of the current node.
	QueryNodeSynchronizationSt() (bool, error)
	// QueryAccountInfo query account information.
	QueryAccountInfo(puk []byte) (types.AccountInfo, error)
	// GetTokenSymbol returns the token symbol
	GetTokenSymbol() string
	// GetNetworkEnv returns the network environment
	GetNetworkEnv() string
	// SysProperties returns the system properties.
	SysProperties() (pattern.SysProperties, error)
	// SyncState returns the system sync state.
	SyncState() (pattern.SysSyncState, error)
	// SysVersion returns the system version.
	SysVersion() (string, error)
	// NetListening returns whether the current node is listening.
	NetListening() (bool, error)

	// Other

	// GetSignatureAcc returns the signature account.
	GetSignatureAcc() string
	// GetSignatureAccPulickey returns the signature account public key.
	GetSignatureAccPulickey() []byte
	// GetSubstrateAPI returns Substrate API.
	GetSubstrateAPI() *gsrpc.SubstrateAPI
	// GetChainState returns chain node state.
	GetChainState() bool
	// SetChainState sets the state of the chain node.
	SetChainState(state bool)
	// GetSdkName return sdk name
	GetSdkName() string
	// SetSdkName set sdk name
	SetSdkName(name string)
	// Reconnect for reconnecting chains.
	Reconnect() error
	// GetMetadata returns the metadata of the chain.
	GetMetadata() *types.Metadata
	// GetKeyEvents returns the events storage key.
	GetKeyEvents() types.StorageKey
	// GetURI returns URI.
	GetURI() string
	// Sign returns the signature of the msg with the private key of the signing account.
	Sign(msg []byte) ([]byte, error)
	// Verify the signature of the msg with the public key of the signing account.
	Verify(msg []byte, sig []byte) (bool, error)
	// EnabledP2P returns the p2p enable status
	EnabledP2P() bool

	// Process the file according to CESS specifications.
	//
	// Receive parameter:
	//   - file: the file to be processed.
	// Return parameter:
	//   - segmentDataInfo: segment and fragment information of the file.
	//   - string: [fid] unique identifier for the file.
	//   - error: error message.
	ProcessingData(file string) ([]pattern.SegmentDataInfo, string, error)

	// Upload file to CESS gateway.
	//
	// Receive parameter:
	//   - url: the address of the gateway.
	//   - file: uploaded file.
	//   - bucket: the bucket for storing files, it will be created automatically.
	// Return parameter:
	//   - string: [fid] unique identifier for the file.
	//   - error: error message.
	// Preconditions:
	//    1. Account requires purchasing space, refer to [BuySpace] interface.
	//    2. Authorize the space usage rights of the account to the gateway account,
	//    refer to the [AuthorizeSpace] interface.
	//    3. Make sure the name of the bucket is legal, use the [CheckBucketName] method to check.
	// Explanation:
	//    - Account refers to the account where you configured mnemonic when creating an SDK.
	//    - CESS public gateway address: [http://deoss-pub-gateway.cess.cloud/]
	//    - CESS public gateway account: [cXhwBytXqrZLr1qM5NHJhCzEMckSTzNKw17ci2aHft6ETSQm9]
	StoreFile(url, file, bucket string) (string, error)

	// Upload object to CESS gateway.
	//
	// Receive parameter:
	//   - url: the address of the gateway.
	//   - reader: strings, byte data, file streams, network streams, etc.
	//   - bucket: the bucket for storing object, it will be created automatically.
	// Return parameter:
	//   - string: [fid] unique identifier for the file.
	//   - error: error message.
	// Preconditions:
	//    1. Account requires purchasing space, refer to [BuySpace] interface.
	//    2. Authorize the space usage rights of the account to the gateway account,
	//    refer to the [AuthorizeSpace] interface.
	//    3. Make sure the name of the bucket is legal, use the [CheckBucketName] method to check.
	// Explanation:
	//    - Account refers to the account where you configured mnemonic when creating an SDK.
	//    - CESS public gateway address: [http://deoss-pub-gateway.cess.cloud/]
	//    - CESS public gateway account: [cXhwBytXqrZLr1qM5NHJhCzEMckSTzNKw17ci2aHft6ETSQm9]
	StoreObject(url string, reader io.Reader, bucket string) (string, error)

	// Download file from CESS public gateway.
	//
	// Receive parameter:
	//   - url: the address of the gateway.
	//   - fid: unique identifier for the file.
	//   - savepath: file save location.
	// Return parameter:
	//   - error: error message.
	RetrieveFile(url, fid, savepath string) error

	// Download object from CESS gateway.
	//
	// Receive parameter:
	//   - url: the address of the gateway.
	//   - fid: unique identifier for the file.
	// Return parameter:
	//   - Reader: object stream.
	//   - error: error message.
	RetrieveObject(url, fid string) (io.ReadCloser, error)

	// Restore files from shards and save to out.
	//
	// Receive parameter:
	//   - out: file save location.
	//   - shards: file fragments.
	// Return parameter:
	//   - error: error message.
	RedundancyRecovery(out string, shards []string) error
}
