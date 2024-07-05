/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// chain client interface
type Chainer interface {
	// Audit
	QueryChallengeSnapShot(accountID []byte, block int32) (bool, ChallengeInfo, error)
	QueryCountedClear(accountID []byte, block int32) (uint8, error)
	QueryCountedServiceFailed(accountID []byte, block int32) (uint32, error)
	SubmitIdleProof(idleProof []types.U8) (string, error)
	SubmitServiceProof(serviceProof []types.U8) (string, error)

	// Babe
	QueryAuthorities(block int32) ([]ConsensusRrscAppPublic, error)

	// Balances
	QueryTotalIssuance(block int32) (string, error)
	QueryInactiveIssuance(block int32) (string, error)
	TransferToken(dest string, amount uint64) (string, error)

	// Oss
	QueryOss(accountID []byte, block int32) (OssInfo, error)
	QueryAllOss(block int32) ([]OssInfo, error)
	QueryAllOssPeerId(block int32) ([]string, error)
	QueryAuthorityList(accountID []byte, block int32) ([]types.AccountID, error)
	Authorize(accountID []byte) (string, error)
	CancelAuthorize(accountID []byte) (string, error)
	RegisterOss(peerId []byte, domain string) (string, error)
	UpdateOss(peerId string, domain string) (string, error)
	DestroyOss() (string, error)

	// EVM
	SendEvmCall(source types.H160, target types.H160, input types.Bytes, value types.U256, gasLimit types.U64, maxFeePerGas types.U256, accessList []AccessInfo) (string, error)

	// FileBank
	QueryBucket(accountID []byte, bucketName string, block int32) (BucketInfo, error)
	QueryDealMap(fid string, block int32) (StorageOrder, error)
	QueryFile(fid string, block int32) (FileMetadata, error)
	QueryRestoralOrder(fragmentHash string, block int32) (RestoralOrderInfo, error)
	QueryAllRestoralOrder(block int32) ([]RestoralOrderInfo, error)
	QueryAllBucketName(accountID []byte, block int32) ([]string, error)
	QueryUserHoldFileList(accountID []byte, block int32) ([]UserFileSliceInfo, error)
	QueryUserFidList(accountID []byte, block int32) ([]string, error)
	PlaceStorageOrder(fid, file_name, bucket_name, territory_name string, segment []SegmentDataInfo, owner []byte, file_size uint64) (string, error)
	UploadDeclaration(fid string, segment []SegmentList, user UserBrief, filesize uint64) (string, error)
	CreateBucket(owner []byte, bucketName string) (string, error)
	DeleteBucket(owner []byte, bucketName string) (string, error)
	DeleteFile(owner []byte, fid string) (string, error)
	TransferReport(index uint8, fid string) (string, error)
	GenerateRestoralOrder(fid, fragmentHash string) (string, error)
	ClaimRestoralOrder(fragmentHash string) (string, error)
	ClaimRestoralNoExistOrder(puk []byte, fid, fragmentHash string) (string, error)
	RestoralOrderComplete(fragmentHash string) (string, error)
	CertIdleSpace(spaceProofInfo SpaceProofInfo, teeSignWithAcc, teeSign types.Bytes, teePuk WorkerPublicKey) (string, error)
	ReplaceIdleSpace(spaceProofInfo SpaceProofInfo, teeSignWithAcc, teeSign types.Bytes, teePuk WorkerPublicKey) (string, error)
	CalculateReport(teeSig types.Bytes, tagSigInfo TagSigInfo) (string, error)
	TerritorFileDelivery(user []byte, fid string, target_territory string) (string, error)

	// SchedulerCredit
	QueryCurrentCounters(accountId []byte, block int32) (SchedulerCounterEntry, error)

	// Session
	QueryValidators(block int32) ([]types.AccountID, error)

	// Sminer
	QueryExpenders(block int32) (ExpendersInfo, error)
	QueryMinerItems(accountID []byte, block int32) (MinerInfo, error)
	QueryStakingStartBlock(accountID []byte, block int32) (uint32, error)
	QueryAllMiner(block int32) ([]types.AccountID, error)
	QueryCounterForMinerItems(block int32) (uint32, error)
	QueryRewardMap(accountID []byte, block int32) (MinerReward, error)
	QueryRestoralTarget(accountID []byte, block int32) (RestoralTargetInfo, error)
	QueryAllRestoralTarget(block int32) ([]RestoralTargetInfo, error)
	QueryPendingReplacements(accountID []byte, block int32) (types.U128, error)
	QueryCompleteSnapShot(era uint32, block int32) (uint32, uint64, error)
	QueryCompleteMinerSnapShot(puk []byte, block int32) (MinerCompleteInfo, error)
	IncreaseCollateral(accountID []byte, token string) (string, error)
	IncreaseDeclarationSpace(tibCount uint32) (string, error)
	MinerExitPrep() (string, error)
	MinerWithdraw() (string, error)
	ReceiveReward() (string, string, error)
	RegisterPoisKey(poisKey PoISKeyInfo, teeSignWithAcc, teeSign types.Bytes, teePuk WorkerPublicKey) (string, error)
	RegnstkSminer(earnings string, peerId []byte, staking uint64, tibCount uint32) (string, error)
	RegnstkAssignStaking(earnings string, peerId []byte, stakingAcc string, tibCount uint32) (string, error)
	UpdateBeneficiary(earnings string) (string, error)
	UpdateSminerPeerId(peerid PeerId) (string, error)

	// Staking
	QueryCounterForValidators(block int32) (uint32, error)
	QueryValidatorsCount(block int32) (uint32, error)
	QueryNominatorCount(block int32) (uint32, error)
	QueryErasTotalStake(era uint32, block int32) (string, error)
	QueryCurrentEra(block int32) (uint32, error)
	QueryErasRewardPoints(era uint32, block int32) (StakingEraRewardPoints, error)
	QueryAllNominators(block int32) ([]StakingNominations, error)
	QueryAllBonded(block int32) ([]types.AccountID, error)
	QueryValidatorCommission(accountID []byte, block int32) (uint8, error)
	QueryEraValidatorReward(era uint32, block int32) (string, error)
	QueryLedger(accountID []byte, block int32) (StakingLedger, error)
	QueryeErasStakers(era uint32, accountId []byte) (StakingExposure, error)
	QueryeAllErasStakersPaged(era uint32, accountId []byte) ([]StakingExposurePaged, error)
	QueryeErasStakersOverview(era uint32, accountId []byte) (PagedExposureMetadata, error)
	QueryeNominators(accountId []byte, block int32) (StakingNominations, error)

	// StorageHandler
	QueryUnitPrice(block int32) (string, error)
	QueryTotalIdleSpace(block int32) (uint64, error)
	QueryTotalServiceSpace(block int32) (uint64, error)
	QueryPurchasedSpace(block int32) (uint64, error)
	QueryTerritory(accountId []byte, name string, block int32) (TerritoryInfo, error)
	QueryConsignment(token types.H256, block int32) (ConsignmentInfo, error)
	MintTerritory(gib_count uint32, territory_name string) (string, error)
	ExpandingTerritory(territory_name string, gib_count uint32) (string, error)
	RenewalTerritory(territory_name string, days_count uint32) (string, error)
	ReactivateTerritory(territory_name string, days_count uint32) (string, error)
	TerritoryConsignment(territory_name string) (string, error)
	CancelConsignment(territory_name string) (string, error)
	BuyConsignment(token types.H256, territory_name string) (string, error)
	CancelPurchaseAction(token types.H256) (string, error)

	// System
	QueryBlockNumber(blockhash string) (uint32, error)
	QueryAccountInfo(account string, block int32) (types.AccountInfo, error)
	QueryAccountInfoByAccountID(accountID []byte, block int32) (types.AccountInfo, error)
	QueryAllAccountInfo(block int32) ([]types.AccountInfo, error)

	// TeeWorker
	QueryMasterPubKey(block int32) ([]byte, error)
	QueryWorkers(puk WorkerPublicKey, block int32) (WorkerInfo, error)
	QueryAllWorkers(block int32) ([]WorkerInfo, error)
	QueryEndpoints(puk WorkerPublicKey, block int32) (string, error)
	QueryWorkerAddedAt(puk WorkerPublicKey, block int32) (uint32, error)

	// CessTreasury
	QueryCurrencyReward(block int32) (string, error)
	QueryEraReward(block int32) (string, error)
	QueryReserveReward(block int32) (string, error)
	QueryRoundReward(era uint32, block int32) (string, error)

	// rpc_call
	ChainGetBlock(hash types.Hash) (types.SignedBlock, error)
	ChainGetBlockHash(block uint32) (types.Hash, error)
	ChainGetFinalizedHead() (types.Hash, error)
	NetListening() (bool, error)
	SystemProperties() (SysProperties, error)
	SystemChain() (string, error)
	SystemSyncState() (SysSyncState, error)
	SystemVersion() (string, error)

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
	GetBalances() uint64
	SetBalances(balance uint64)
	Sign(msg []byte) ([]byte, error)
	Verify(msg []byte, sig []byte) (bool, error)
	ReconnectRpc() error
	Close()

	// extrinsics
	InitExtrinsicsName() error
	ParseBlockData(blocknumber uint64) (BlockData, error)

	// event
	RetrieveAllEventName(blockhash types.Hash) ([]string, error)
	RetrieveEvent(blockhash types.Hash, extrinsic_name, event_name, signer string) error
	RetrieveExtrinsicsAndEvents(blockhash types.Hash) ([]string, map[string][]string, error)
}
