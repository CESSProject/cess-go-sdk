/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdk

import (
	"io"
	"math/big"

	"github.com/CESSProject/cess-go-sdk/core/event"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// CESS Go SDK Interface Description
type SDK interface {
	// Audit-State

	// QueryChallengeVerifyExpiration Query Challenge Expiration Block High
	QueryChallengeVerifyExpiration() (uint32, error)
	// QueryChallengeInfo queries accountID's challenge information
	//   Tip: accountID can only be a storage node account
	QueryChallengeInfo(accountID []byte) (bool, pattern.ChallengeInfo, error)

	// Audit-Extrinsics

	// SubmitIdleProof submits proof of idle data to the chain
	//   Tip: This method can only be used for storage nodes
	SubmitIdleProof(idleProve []types.U8) (string, error)
	// SubmitServiceProof submits proof of service data to the chain
	//   Tip: This method can only be used for storage nodes
	SubmitServiceProof(serviceProof []types.U8) (string, error)
	// SubmitIdleProofResult submits the proof verification results of idle data to the chain
	//   Tip: This method can only be used for storage nodes
	SubmitIdleProofResult(totalProofHash []types.U8, front, rear types.U64, accumulator pattern.Accumulator, result types.Bool, signature pattern.TeeSignature, tee_acc []byte) (string, error)
	// SubmitServiceProofResult submits the proof verification results of service data to the chain
	//   Tip: This method can only be used for storage nodes
	SubmitServiceProofResult(result types.Bool, signature pattern.TeeSignature, bloomFilter pattern.BloomFilter, tee_acc []byte) (string, error)

	// Filebank-State

	// QueryBucketInfo query the bucket information of the accountID
	QueryBucketInfo(accountID []byte, bucketName string) (pattern.BucketInfo, error)
	// QueryAllBucket query all buckets of the accountID
	QueryAllBucket(accountID []byte) ([]types.Bytes, error)
	// QueryAllBucketString query all bucket names as string of the accountID
	QueryAllBucketString(accountID []byte) ([]string, error)
	// QueryStorageOrder query storage order information.
	QueryStorageOrder(fid string) (pattern.StorageOrder, error)
	// QueryFileMetadata queries the metadata of the roothash file.
	QueryFileMetadata(fid string) (pattern.FileMetadata, error)
	// QueryFileMetadataByBlock queries the metadata of the roothash file.
	QueryFileMetadataByBlock(fid string, block uint64) (pattern.FileMetadata, error)
	// QueryRestoralOrder queries a restore order info.
	QueryRestoralOrder(roothash string) (pattern.RestoralOrderInfo, error)
	QueryRestoralOrderList() ([]pattern.RestoralOrderInfo, error)

	// Filebank-Extrinsics

	// ClaimRestoralNoExistOrder is used to receive recovery orders from exiting miners.
	ClaimRestoralNoExistOrder(accountID []byte, fid, restoralFragmentHash string) (string, error)
	// ClaimRestoralOrder is used to collect restoration orders.
	ClaimRestoralOrder(fragmentHash string) (string, error)
	// CreateBucket creates a bucket for accountID
	//   For details on bucket naming rules, see:
	// https://app.gitbook.com/o/kiTNX10jBU59sjnYZbiH/s/G1ekWsjn9OlGH381wiK2/get-started/deoss-gateway/step-1-create-a-bucket#naming-conventions-for-a-bucket
	CreateBucket(accountID []byte, bucketName string) (string, error)
	// DeleteBucket deletes buckets for accountID
	//   Tip: Only empty buckets can be deleted
	DeleteBucket(accountID []byte, bucketName string) (string, error)
	// DeleteFile deletes files for accountID
	DeleteFile(accountID []byte, fid []string) (string, []pattern.FileHash, error)
	// GenerateRestoralOrder generates data for restoration orders.
	GenerateRestoralOrder(fid, fragmentHash string) (string, error)
	// RestoralComplete reports order recovery completion.
	RestoralComplete(restoralFragmentHash string) (string, error)
	// SubmitFileReport submits a stored file report.
	SubmitFileReport(index types.U8, roothash pattern.FileHash) (string, error)
	ReportFile(index uint8, roothash string) (string, error)
	// UploadDeclaration creates a storage order.
	UploadDeclaration(filehash string, dealinfo []pattern.SegmentList, user pattern.UserBrief, filesize uint64) (string, error)
	// GenerateStorageOrder for generating storage orders
	GenerateStorageOrder(roothash string, segment []pattern.SegmentDataInfo, owner []byte, filename, buckname string, filesize uint64) (string, error)
	// CertIdleSpace
	CertIdleSpace(idleSignInfo pattern.SpaceProofInfo, sign pattern.TeeSignature, teeWorkAcc string) (string, error)
	// ReplaceIdleSpace
	ReplaceIdleSpace(idleSignInfo pattern.SpaceProofInfo, sign pattern.TeeSignature, teeWorkAcc string) (string, error)
	// ReportTagCalculated
	ReportTagCalculated(teeSig pattern.TeeSignature, tagSigInfo pattern.TagSigInfo) (string, error)

	// Oss-State

	// QueryAuthorizedAccountIDs queries all DeOSS accountIDs authorized by accountID
	QueryAuthorizedAccountIDs(accountID []byte) ([]types.AccountID, error)
	// QueryAuthorizedAccounts queries all DeOSS accounts authorized by accountID
	QueryAuthorizedAccounts(accountID []byte) ([]string, error)
	// QueryDeOSSInfo Query the DeOSS information registered by accountID account
	QueryDeOSSInfo(accountID []byte) (pattern.OssInfo, error)
	// QueryAllDeOSSInfo queries all deoss information
	QueryAllDeOSSInfo() ([]pattern.OssInfo, error)
	// QueryAllDeOSSPeerId queries peerID of all DeOSS
	QueryAllDeOSSPeerId() ([]string, error)

	// Oss-Extrinsics

	// RegisterDeOSS register as deoss node
	RegisterDeOSS(peerId []byte, domain string) (string, error)
	// UpdateDeOSS updates the peerID and domain of deoss
	//   Tip: This method can only be used for DeOSS nodes
	UpdateDeOSS(peerId string, domain string) (string, error)
	// ExitDeoss exit deoss
	//   Tip: This method can only be used for DeOSS nodes
	ExitDeOSS() (string, error)
	// AuthorizeSpace authorizes space to oss account
	//   Tip: accountID can only be a DeOSS node account
	AuthorizeSpace(accountID string) (string, error)
	// UnAuthorizeSpace cancels space authorization
	//   Tip: accountID can only be a DeOSS node account
	UnAuthorizeSpace(accountID string) (string, error)

	// Sminer-State

	// QueryExpenders queries expenders information
	QueryExpenders() (pattern.ExpendersInfo, error)
	// QueryStorageMiner queries storage node information.
	QueryStorageMiner(accountID []byte) (pattern.MinerInfo, error)
	// QueryAllSminerAccount queries the accounts of all storage miners.
	QueryAllSminerAccount() ([]types.AccountID, error)
	// QueryRewardsMap queries rewardsMap for accountID
	//   Tip: accountID can only be a storage node account
	QueryRewardsMap(accountID []byte) (pattern.MinerReward, error)
	// QueryRewards queries rewards for accountID
	//   Tip: accountID can only be a storage node account
	QueryRewards(accountID []byte) (pattern.RewardsType, error)
	// QueryStorageMinerStakingStartBlock
	QueryStorageMinerStakingStartBlock(puk []byte) (types.U32, error)
	// QueryRestoralTarget queries the space recovery information of the exited storage node
	//   Tip: accountID can only be a storage node account
	QueryRestoralTarget(accountID []byte) (pattern.RestoralTargetInfo, error)
	// QueryRestoralTargetList queries the space recovery information of all exited storage nodes
	QueryRestoralTargetList() ([]pattern.RestoralTargetInfo, error)
	// QueryPendingReplacements queries the amount of idle data that can be replaced
	//   Tip: accountID can only be a storage node account
	QueryPendingReplacements(accountID []byte) (types.U128, error)

	// Sminer-Extrinsics

	//
	IncreaseDeclarationSpace(tibCount uint32) (string, error)
	// RegisterSminer register sminer information
	RegisterSminer(earnings string, peerId []byte, pledge uint64, tib_count uint32) (string, error)
	// RegisterSminerAssignStaking
	RegisterSminerAssignStaking(beneficiaryAcc string, peerId []byte, stakingAcc string, tib_count uint32) (string, error)
	// RegisterSminerPOISKey register the pois key of sminer
	RegisterSminerPOISKey(poisKey pattern.PoISKeyInfo, sign pattern.TeeSignature, teeWorkAcc string) (string, error)
	// ExitSminer exit mining
	ExitSminer(miner string) (string, error)
	// UpdateEarningsAcc update earnings account.
	UpdateEarningsAcc(puk []byte) (string, error)
	UpdateEarningsAccount(earnings string) (string, error)
	// UpdateSminerPeerId update miner peerid
	UpdateSminerPeerId(peerid pattern.PeerId) (string, error)
	// IncreaseStakingAmount increase staking amount.
	IncreaseStakingAmount(miner string, tokens *big.Int) (string, error)
	IncreaseStorageNodeStakingAmount(miner string, token string) (string, error)
	// ClaimRewards is used to claim rewards.
	ClaimRewards() (string, error)
	// Withdraw is used to withdraw staking.
	Withdraw() (string, error)

	// Staking-State

	// QueryValidatorCount queries the count of all validator
	QueryValidatorCount() (uint32, error)

	// StorageHandler-State

	// QueryUserSpaceInfo queries the space information purchased by the user.
	QueryUserSpaceInfo(puk []byte) (pattern.UserSpaceInfo, error)
	QueryUserSpaceSt(puk []byte) (pattern.UserSpaceSt, error)
	// QuerySpacePricePerGib query space price per GiB.
	QuerySpacePricePerGib() (string, error)
	// QueryTotalIdleSpace query total idle space
	QueryTotalIdleSpace() (uint64, error)
	// QueryTotalServiceSpace query total service space
	QueryTotalServiceSpace() (uint64, error)
	// QueryPurchasedSpace query purchased space
	QueryPurchasedSpace() (uint64, error)

	// StorageHandler-Extrinsics

	// BuySpace for purchasing space.
	BuySpace(count uint32) (string, error)
	// ExpansionSpace for expansion space.
	ExpansionSpace(count uint32) (string, error)
	// RenewalSpace is used to extend the validity of space.
	RenewalSpace(days uint32) (string, error)

	// TeeWorker-State

	// QueryTeeWorkerMap queries the information of the Tee worker.
	QueryTeeWorkerMap(accountID []byte) (pattern.TeeWorkerMap, error)
	// QueryTeeInfo queries the information of the Tee worker.
	QueryTeeInfo(accountID []byte) (pattern.TeeInfo, error)
	// QueryTeePodr2Puk queries the public key of the TEE.
	QueryTeePodr2Puk() ([]byte, error)
	// QueryAllTeeWorkerMap queries the information of all tee workers.
	QueryAllTeeWorkerMap() ([]pattern.TeeWorkerMap, error)
	// QueryAllTeeInfo queries the information of all tee workers.
	QueryAllTeeInfo() ([]pattern.TeeInfo, error)

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
	ChainVersion() (string, error)
	// NetListening returns whether the current node is listening.
	NetListening() (bool, error)
	//
	DecodeEventNameFromBlock(block uint64) ([]string, error)
	//
	DecodeEventNameFromBlockhash(blockhash types.Hash) ([]string, error)

	// TransferToken to dest.
	//
	// Receive parameter:
	//   - dest: target account.
	//   - amount: transfer amount.
	// Return parameter:
	//   - string: transaction hash.
	//   - string: target account.
	//   - error: error message.
	TransferToken(dest string, amount uint64) (string, string, error)

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
	GetSDKName() string
	// GetCurrentRpcAddr return current rpc address
	GetCurrentRpcAddr() string
	// SetSdkName set sdk name
	SetSDKName(name string)
	// ReconnectRPC for reconnecting chains.
	ReconnectRPC() error
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
	// Verify Polka Signature With JavaScript
	VerifyPolkaSignatureWithJS(account, msg, signature string) (bool, error)
	// Verify Polka Signature With Base58
	VerifyPolkaSignatureWithBase58(account, msg, signature string) (bool, error)

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

	// retrieve event

	//
	RetrieveEvent_Audit_SubmitIdleProof(blockhash types.Hash) (event.Event_SubmitIdleProof, error)
	//
	RetrieveEvent_Audit_SubmitServiceProof(blockhash types.Hash) (event.Event_SubmitServiceProof, error)
	//
	RetrieveEvent_Audit_SubmitIdleVerifyResult(blockhash types.Hash) (event.Event_SubmitIdleVerifyResult, error)
	//
	RetrieveEvent_Audit_SubmitServiceVerifyResult(blockhash types.Hash) (event.Event_SubmitServiceVerifyResult, error)
	//
	RetrieveEvent_Oss_OssUpdate(blockhash types.Hash) (event.Event_OssUpdate, error)
	//
	RetrieveEvent_Oss_OssRegister(blockhash types.Hash) (event.Event_OssRegister, error)
	//
	RetrieveEvent_Oss_OssDestroy(blockhash types.Hash) (event.Event_OssDestroy, error)
	//
	RetrieveEvent_Oss_Authorize(blockhash types.Hash) (event.Event_Authorize, error)
	//
	RetrieveEvent_Oss_CancelAuthorize(blockhash types.Hash) (event.Event_CancelAuthorize, error)
	//
	RetrieveEvent_FileBank_UploadDeclaration(blockhash types.Hash) (event.Event_UploadDeclaration, error)
	//
	RetrieveEvent_FileBank_CreateBucket(blockhash types.Hash) (event.Event_CreateBucket, error)
	//
	RetrieveEvent_FileBank_DeleteFile(blockhash types.Hash) (event.Event_DeleteFile, error)
	//
	RetrieveEvent_FileBank_TransferReport(blockhash types.Hash) (event.Event_TransferReport, error)
	//
	RetrieveEvent_FileBank_ClaimRestoralOrder(blockhash types.Hash) (event.Event_ClaimRestoralOrder, error)
	//
	RetrieveEvent_FileBank_RecoveryCompleted(blockhash types.Hash) (event.Event_RecoveryCompleted, error)
	//
	RetrieveEvent_FileBank_IdleSpaceCert(blockhash types.Hash) (event.Event_IdleSpaceCert, error)
	//
	RetrieveEvent_FileBank_ReplaceIdleSpace(blockhash types.Hash) (event.Event_ReplaceIdleSpace, error)
	//
	RetrieveEvent_FileBank_CalculateReport(blockhash types.Hash) (event.Event_CalculateReport, error)
	//
	RetrieveEvent_Sminer_Registered(blockhash types.Hash) (event.Event_Registered, error)
	//
	RetrieveEvent_Sminer_RegisterPoisKey(blockhash types.Hash) (event.Event_RegisterPoisKey, error)
	//
	RetrieveEvent_Sminer_UpdataIp(blockhash types.Hash) (event.Event_UpdataIp, error)
	//
	RetrieveEvent_Sminer_UpdataBeneficiary(blockhash types.Hash) (event.Event_UpdataBeneficiary, error)
	//
	RetrieveEvent_Sminer_MinerExitPrep(blockhash types.Hash) (event.Event_MinerExitPrep, error)
	//
	RetrieveEvent_Sminer_IncreaseCollateral(blockhash types.Hash) (event.Event_IncreaseCollateral, error)
	//
	RetrieveEvent_Sminer_Receive(blockhash types.Hash) (event.Event_Receive, error)
	//
	RetrieveEvent_Sminer_Withdraw(blockhash types.Hash) (event.Event_Withdraw, error)
	//
	RetrieveEvent_Sminer_IncreaseDeclarationSpace(blockhash types.Hash) (event.Event_IncreaseDeclarationSpace, error)
	//
	RetrieveEvent_StorageHandler_BuySpace(blockhash types.Hash) (event.Event_BuySpace, error)
	//
	RetrieveEvent_StorageHandler_ExpansionSpace(blockhash types.Hash) (event.Event_ExpansionSpace, error)
	//
	RetrieveEvent_StorageHandler_RenewalSpace(blockhash types.Hash) (event.Event_RenewalSpace, error)
	//
	RetrieveEvent_Balances_Transfer(blockhash types.Hash) (types.EventBalancesTransfer, error)
	//
	RetrieveEvent_FileBank_GenRestoralOrder(blockhash types.Hash) (event.Event_GenerateRestoralOrder, error)
	//
	RetrieveAllEvent_FileBank_UploadDeclaration(blockhash types.Hash) ([]event.AllUploadDeclarationEvent, error)
	//
	RetrieveAllEvent_FileBank_StorageCompleted(blockhash types.Hash) ([]string, error)
	//
	RetrieveAllEvent_FileBank_DeleteFile(blockhash types.Hash) ([]event.AllDeleteFileEvent, error)
}
