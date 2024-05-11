/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdk

import (
	"github.com/CESSProject/cess-go-sdk/chain"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// CESS Go SDK Interface Description
type SDK interface {
	// Audit
	QueryChallengeSnapShot(accountID []byte, block int32) (bool, chain.ChallengeInfo, error)
	QueryCountedClear(accountID []byte, block int32) (uint8, error)
	QueryCountedServiceFailed(accountID []byte, block int32) (uint32, error)
	SubmitIdleProof(idleProof []types.U8) (string, error)
	SubmitServiceProof(serviceProof []types.U8) (string, error)

	// Balances
	QueryTotalIssuance(block int) (string, error)
	QueryInactiveIssuance(block int) (string, error)
	TransferToken(dest string, amount uint64) (string, string, error)

	// Oss
	QueryOss(accountID []byte, block int32) (chain.OssInfo, error)
	QueryAllOss(block int32) ([]chain.OssInfo, error)
	QueryAllOssPeerId(block int32) ([]string, error)
	QueryAuthorityList(accountID []byte, block int32) ([]types.AccountID, error)
	Authorize(accountID []byte) (string, error)
	CancelAuthorize(accountID []byte) (string, error)
	RegisterOss(peerId []byte, domain string) (string, error)
	UpdateOss(peerId string, domain string) (string, error)
	DestroyOss() (string, error)

	// EVM
	SendEvmCall(source types.H160, target types.H160, input types.Bytes, value types.U256, gasLimit types.U64, maxFeePerGas types.U256, accessList []chain.AccessInfo) (string, error)

	// FileBank
	QueryBucket(accountID []byte, bucketName string, block int32) (chain.BucketInfo, error)
	QueryDealMap(fid string, block int32) (chain.StorageOrder, error)
	QueryFile(fid string, block int32) (chain.FileMetadata, error)
	QueryRestoralOrder(fragmentHash string, block int32) (chain.RestoralOrderInfo, error)
	QueryAllRestoralOrder(block int32) ([]chain.RestoralOrderInfo, error)
	QueryAllBucketName(accountID []byte, block int32) ([]string, error)
	QueryAllUserFiles(accountID []byte, block int32) ([]string, error)
	GenerateStorageOrder(fid string, segment []chain.SegmentDataInfo, owner []byte, filename string, buckname string, filesize uint64) (string, error)
	UploadDeclaration(fid string, segment []chain.SegmentList, user chain.UserBrief, filesize uint64) (string, error)
	CreateBucket(owner []byte, bucketName string) (string, error)
	DeleteBucket(owner []byte, bucketName string) (string, error)
	DeleteFile(owner []byte, fid string) (string, error)
	TransferReport(index uint8, fid string) (string, error)
	GenerateRestoralOrder(fid, fragmentHash string) (string, error)
	ClaimRestoralOrder(fragmentHash string) (string, error)
	ClaimRestoralNoExistOrder(puk []byte, fid, fragmentHash string) (string, error)
	RestoralOrderComplete(fragmentHash string) (string, error)
	CertIdleSpace(spaceProofInfo chain.SpaceProofInfo, teeSignWithAcc, teeSign types.Bytes, teePuk chain.WorkerPublicKey) (string, error)
	ReplaceIdleSpace(spaceProofInfo chain.SpaceProofInfo, teeSignWithAcc, teeSign types.Bytes, teePuk chain.WorkerPublicKey) (string, error)
	CalculateReport(teeSig types.Bytes, tagSigInfo chain.TagSigInfo) (string, error)

	// Sminer
	QueryExpenders(block int32) (chain.ExpendersInfo, error)
	QueryMinerItems(accountID []byte, block int32) (chain.MinerInfo, error)
	QueryStakingStartBlock(accountID []byte, block int32) (uint32, error)
	QueryAllMiner(block int32) ([]types.AccountID, error)
	QueryCounterForMinerItems(block int32) (uint32, error)
	QueryRewardMap(accountID []byte, block int32) (chain.MinerReward, error)
	QueryRestoralTarget(accountID []byte, block int32) (chain.RestoralTargetInfo, error)
	QueryAllRestoralTarget(block int32) ([]chain.RestoralTargetInfo, error)
	QueryPendingReplacements(accountID []byte, block int32) (types.U128, error)
	QueryCompleteSnapShot(era uint32, block int32) (uint32, uint64, error)
	IncreaseCollateral(accountID []byte, token string) (string, error)
	IncreaseDeclarationSpace(tibCount uint32) (string, error)
	MinerExitPrep() (string, error)
	MinerWithdraw() (string, error)
	ReceiveReward() (string, error)
	RegisterPoisKey(poisKey chain.PoISKeyInfo, teeSignWithAcc, teeSign types.Bytes, teePuk chain.WorkerPublicKey) (string, error)
	RegnstkSminer(earnings string, peerId []byte, staking uint64, tibCount uint32) (string, error)
	RegnstkAssignStaking(earnings string, peerId []byte, stakingAcc string, tibCount uint32) (string, error)
	UpdateBeneficiary(earnings string) (string, error)
	UpdateSminerPeerId(peerid chain.PeerId) (string, error)

	// Staking
	QueryCounterForValidators(block int) (uint32, error)
	QueryValidatorsCount(block int) (uint32, error)
	QueryNominatorCount(block int) (uint32, error)
	QueryErasTotalStake(era uint32, block int) (string, error)
	QueryCurrentEra(block int) (uint32, error)
	QueryErasRewardPoints(era uint32, block int32) (chain.StakingEraRewardPoints, error)
	QueryAllNominators(block int32) ([]chain.StakingNominations, error)
	QueryAllBonded(block int32) ([]types.AccountID, error)
	QueryValidatorCommission(accountID []byte, block int) (uint8, error)
	QueryEraValidatorReward(era uint32, block int) (string, error)

	// StorageHandler
	QueryUnitPrice(block int32) (string, error)
	QueryUserOwnedSpace(accountID []byte, block int32) (chain.UserSpaceInfo, error)
	QueryTotalIdleSpace(block int32) (uint64, error)
	QueryTotalServiceSpace(block int32) (uint64, error)
	QueryPurchasedSpace(block int32) (uint64, error)
	BuySpace(count uint32) (string, error)
	ExpansionSpace(count uint32) (string, error)
	RenewalSpace(days uint32) (string, error)

	// System
	QueryBlockNumber(blockhash string) (uint32, error)
	QueryAccountInfo(account string, block int32) (types.AccountInfo, error)
	QueryAccountInfoByAccountID(accountID []byte, block int32) (types.AccountInfo, error)

	// TeeWorker
	QueryMasterPubKey(block int32) ([]byte, error)
	QueryWorkers(puk chain.WorkerPublicKey, block int32) (chain.WorkerInfo, error)
	QueryAllWorkers(block int32) ([]chain.WorkerInfo, error)
	QueryEndpoints(puk chain.WorkerPublicKey, block int32) (string, error)
	QueryWorkerAddedAt(puk chain.WorkerPublicKey, block int32) (uint32, error)

	// CessTreasury
	QueryCurrencyReward(block int32) (string, error)
	QueryEraReward(block int32) (string, error)
	QueryReserveReward(block int32) (string, error)
	QueryRoundReward(era uint32, block int32) (string, error)

	// rpc_call
	SystemProperties() (chain.SysProperties, error)
	SystemChain() (string, error)
	SystemSyncState() (chain.SysSyncState, error)
	SystemVersion() (string, error)
	NetListening() (bool, error)

	// chain_client
	GetSDKName() string
	GetCurrentRpcAddr() string
	GetRpcState() bool
	SetRpcState(state bool)
	GetSignatureAcc() string
	GetSignatureAccPulickey() []byte
	GetSubstrateAPI() *gsrpc.SubstrateAPI
	GetMetadata() *types.Metadata
	GetTokenSymbol() string
	GetNetworkEnv() string
	GetURI() string
	Sign(msg []byte) ([]byte, error)
	Verify(msg []byte, sig []byte) (bool, error)
	ReconnectRpc() error
	Close()

	// Extrinsics
	InitExtrinsicsName() error

	// event
	DecodeEventNameFromBlock(block uint64) ([]string, error)
	DecodeEventNameFromBlockhash(blockhash types.Hash) ([]string, error)
	QueryAllAccountInfoFromBlock(block int) ([]types.AccountInfo, error)
	// retrieve event
	RetrieveEvent_Audit_SubmitIdleProof(blockhash types.Hash) (chain.Event_SubmitIdleProof, error)
	RetrieveEvent_Audit_SubmitServiceProof(blockhash types.Hash) (chain.Event_SubmitServiceProof, error)
	RetrieveEvent_Audit_SubmitIdleVerifyResult(blockhash types.Hash) (chain.Event_SubmitIdleVerifyResult, error)
	RetrieveEvent_Audit_SubmitServiceVerifyResult(blockhash types.Hash) (chain.Event_SubmitServiceVerifyResult, error)
	RetrieveEvent_Oss_OssUpdate(blockhash types.Hash) (chain.Event_OssUpdate, error)
	RetrieveEvent_Oss_OssRegister(blockhash types.Hash) (chain.Event_OssRegister, error)
	RetrieveEvent_Oss_OssDestroy(blockhash types.Hash) (chain.Event_OssDestroy, error)
	RetrieveEvent_Oss_Authorize(blockhash types.Hash) (chain.Event_Authorize, error)
	RetrieveEvent_Oss_CancelAuthorize(blockhash types.Hash) (chain.Event_CancelAuthorize, error)
	RetrieveEvent_FileBank_UploadDeclaration(blockhash types.Hash) (chain.Event_UploadDeclaration, error)
	RetrieveEvent_FileBank_CreateBucket(blockhash types.Hash) (chain.Event_CreateBucket, error)
	RetrieveEvent_FileBank_DeleteFile(blockhash types.Hash) (chain.Event_DeleteFile, error)
	RetrieveEvent_FileBank_TransferReport(blockhash types.Hash) (chain.Event_TransferReport, error)
	RetrieveEvent_FileBank_ClaimRestoralOrder(blockhash types.Hash) (chain.Event_ClaimRestoralOrder, error)
	RetrieveEvent_FileBank_RecoveryCompleted(blockhash types.Hash) (chain.Event_RecoveryCompleted, error)
	RetrieveEvent_FileBank_IdleSpaceCert(blockhash types.Hash) (chain.Event_IdleSpaceCert, error)
	RetrieveEvent_FileBank_ReplaceIdleSpace(blockhash types.Hash) (chain.Event_ReplaceIdleSpace, error)
	RetrieveEvent_FileBank_CalculateReport(blockhash types.Hash) (chain.Event_CalculateReport, error)
	RetrieveEvent_Sminer_Registered(blockhash types.Hash) (chain.Event_Registered, error)
	RetrieveEvent_Sminer_RegisterPoisKey(blockhash types.Hash) (chain.Event_RegisterPoisKey, error)
	RetrieveEvent_Sminer_UpdataIp(blockhash types.Hash) (chain.Event_UpdatePeerId, error)
	RetrieveEvent_Sminer_UpdataBeneficiary(blockhash types.Hash) (chain.Event_UpdateBeneficiary, error)
	RetrieveEvent_Sminer_MinerExitPrep(blockhash types.Hash) (chain.Event_MinerExitPrep, error)
	RetrieveEvent_Sminer_IncreaseCollateral(blockhash types.Hash) (chain.Event_IncreaseCollateral, error)
	RetrieveEvent_Sminer_Receive(blockhash types.Hash) (chain.Event_Receive, error)
	RetrieveEvent_Sminer_Withdraw(blockhash types.Hash) (chain.Event_Withdraw, error)
	RetrieveEvent_Sminer_IncreaseDeclarationSpace(blockhash types.Hash) (chain.Event_IncreaseDeclarationSpace, error)
	RetrieveEvent_StorageHandler_BuySpace(blockhash types.Hash) (chain.Event_BuySpace, error)
	RetrieveEvent_StorageHandler_ExpansionSpace(blockhash types.Hash) (chain.Event_ExpansionSpace, error)
	RetrieveEvent_StorageHandler_RenewalSpace(blockhash types.Hash) (chain.Event_RenewalSpace, error)
	RetrieveEvent_Balances_Transfer(blockhash types.Hash) (types.EventBalancesTransfer, error)
	RetrieveEvent_FileBank_GenRestoralOrder(blockhash types.Hash) (chain.Event_GenerateRestoralOrder, error)
	RetrieveAllEvent_FileBank_UploadDeclaration(blockhash types.Hash) ([]chain.AllUploadDeclarationEvent, error)
	RetrieveAllEvent_FileBank_StorageCompleted(blockhash types.Hash) ([]string, error)
	RetrieveAllEvent_FileBank_DeleteFile(blockhash types.Hash) ([]chain.AllDeleteFileEvent, error)
	RetrieveAllEventFromBlock(blockhash types.Hash) ([]string, map[string][]string, error)
	RetrieveBlock(blocknumber uint64) ([]string, []chain.ExtrinsicsInfo, []chain.TransferInfo, string, string, string, string, int64, error)
	RetrieveBlockAndAll(blocknumber uint64) ([]string, []chain.ExtrinsicsInfo, []chain.TransferInfo, []string, []string, string, string, string, string, string, int64, error)
	ParseBlockData(blocknumber uint64) (chain.BlockData, error)
}
