/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"time"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

// DOT is "." character
const DOT = "."

// Unit precision of CESS token
const TokenPrecision_CESS = "000000000000000000"

const StakingStakePerTiB = 4000

const BlockIntervalSec = 6

// BlockInterval is the time interval for generating blocks, in seconds
const BlockInterval = time.Second * time.Duration(BlockIntervalSec)

const MaxSubmitedIdleFileMeta = 30

const PublicDeoss = "http://deoss-pub-gateway.cess.cloud/"
const PublicDeossAccount = "cXhwBytXqrZLr1qM5NHJhCzEMckSTzNKw17ci2aHft6ETSQm9"

// pallet names
const (
	// Audit
	Audit = "Audit"
	// OSS is a module about DeOSS
	Oss = "Oss"
	// FILEBANK is a module about data metadata, bucket info, etc.
	FileBank = "FileBank"
	// TEEWOEKER is a module about TEE
	TeeWorker = "TeeWorker"
	// SMINER is a module about storage miners
	Sminer = "Sminer"
	// STAKING is a module about staking
	Staking = "Staking"
	// SMINER is a module about storage miners
	StorageHandler = "StorageHandler"
	// BALANCES is a module about the balances
	Balances = "Balances"
	// SYSTEM is a module about the system
	System = "System"
	// EVM is a module about the evm contract
	EVM = "EVM"
	//
	CessTreasury = "CessTreasury"
)

// chain state
const (
	// Audit
	ChallengeSlip        = "ChallengeSlip"
	ChallengeSnapShot    = "ChallengeSnapShot"
	CountedClear         = "CountedClear"
	CountedServiceFailed = "CountedServiceFailed"
	VerifySlip           = "VerifySlip"

	// OSS
	// OSS
	AuthorityList = "AuthorityList"

	// SMINER
	AllMiner          = "AllMiner"
	MinerItems        = "MinerItems"
	RewardMap         = "RewardMap"
	Expenders         = "Expenders"
	RestoralTarget    = "RestoralTarget"
	StakingStartBlock = "StakingStartBlock"
	CompleteSnapShot  = "CompleteSnapShot"

	// TEEWORKER
	TEEWorkers       = "Workers"
	TEEMasterPubkey  = "MasterPubkey"
	TEEEndpoints     = "Endpoints"
	TEEWorkerAddedAt = "WorkerAddedAt"

	// FILEBANK
	File                = "File"
	Bucket              = "Bucket"
	DealMap             = "DealMap"
	FillerMap           = "FillerMap"
	PendingReplacements = "PendingReplacements"
	RestoralOrder       = "RestoralOrder"
	UserBucketList      = "UserBucketList"
	UserHoldFileList    = "UserHoldFileList"

	// STAKING
	COUNTERFORVALIDATORS = "CounterForValidators"
	CounterForNominators = "CounterForNominators"
	ErasTotalStake       = "ErasTotalStake"
	CurrentEra           = "CurrentEra"
	ErasStakers          = "ErasStakers"
	ErasRewardPoints     = "ErasRewardPoints"
	Nominators           = "Nominators"
	Bonded               = "Bonded"
	Validators           = "Validators"
	ErasValidatorReward  = "ErasValidatorReward"
	ValidatorCount       = "ValidatorCount"

	// STORAGE_HANDLER
	UserOwnedSpace    = "UserOwnedSpace"
	UnitPrice         = "UnitPrice"
	TotalIdleSpace    = "TotalIdleSpace"
	TotalServiceSpace = "TotalServiceSpace"
	PurchasedSpace    = "PurchasedSpace"

	// Balances
	TotalIssuance    = "TotalIssuance"
	InactiveIssuance = "InactiveIssuance"

	// SYSTEM
	Account = "Account"
	Events  = "Events"

	// CessTreasury
	CurrencyReward = "CurrencyReward"
	EraReward      = "EraReward"
	ReserveReward  = "ReserveReward"
	RoundReward    = "RoundReward"
)

// Extrinsics
const (
	//AUDIT
	TX_Audit_SubmitIdleProof           = Audit + DOT + "submit_idle_proof"
	TX_Audit_SubmitServiceProof        = Audit + DOT + "submit_service_proof"
	TX_Audit_SubmitVerifyIdleResult    = Audit + DOT + "submit_verify_idle_result"
	TX_Audit_SubmitVerifyServiceResult = Audit + DOT + "submit_verify_service_result"

	// OSS
	TX_OSS_REGISTER    = OSS + DOT + "register"
	TX_OSS_UPDATE      = OSS + DOT + "update"
	TX_OSS_DESTROY     = OSS + DOT + "destroy"
	TX_OSS_AUTHORIZE   = OSS + DOT + "authorize"
	TX_OSS_UNAUTHORIZE = OSS + DOT + "cancel_authorize"

	// SMINER
	TX_SMINER_REGISTER              = SMINER + DOT + "regnstk"
	TX_SMINER_REGISTERASSIGNSTAKING = SMINER + DOT + "regnstk_assign_staking"
	TX_SMINER_INCREASESTAKES        = SMINER + DOT + "increase_collateral"
	TX_SMINER_UPDATEPEERID          = SMINER + DOT + "update_peer_id"
	TX_SMINER_UPDATEINCOME          = SMINER + DOT + "update_beneficiary"
	TX_SMINER_CLAIMREWARD           = SMINER + DOT + "receive_reward"
	TX_SMINER_MINEREXITPREP         = SMINER + DOT + "miner_exit_prep"
	TX_SMINER_WITHDRAW              = SMINER + DOT + "miner_withdraw"
	TX_SMINER_REGISTERPOISKEY       = SMINER + DOT + "register_pois_key"
	TX_SMINER_INCREASEDECSPACE      = SMINER + DOT + "increase_declaration_space"

	// FileBank
	TX_FileBank_CreateBucket              = FileBank + DOT + "create_bucket"
	TX_FileBank_DeleteBucket              = FileBank + DOT + "delete_bucket"
	TX_FileBank_DeleteFile                = FileBank + DOT + "delete_file"
	TX_FileBank_UploadDeclaration         = FileBank + DOT + "upload_declaration"
	TX_FileBank_TransferReport            = FileBank + DOT + "transfer_report"
	TX_FileBank_GenerateRestoralOrder     = FileBank + DOT + "generate_restoral_order"
	TX_FileBank_ClaimRestoralOrder        = FileBank + DOT + "claim_restoral_order"
	TX_FileBank_ClaimRestoralNoexistOrder = FileBank + DOT + "claim_restoral_noexist_order"
	TX_FileBank_RestoralOrderComplete     = FileBank + DOT + "restoral_order_complete"
	TX_FileBank_CertIdleSpace             = FileBank + DOT + "cert_idle_space"
	TX_FileBank_ReplaceIdleSpace          = FileBank + DOT + "replace_idle_space"
	TX_FileBank_CalculateReport           = FileBank + DOT + "calculate_report"

	// STORAGE_HANDLER
	TX_STORAGE_BUYSPACE       = STORAGEHANDLER + DOT + "buy_space"
	TX_STORAGE_EXPANSIONSPACE = STORAGEHANDLER + DOT + "expansion_space"
	TX_STORAGE_RENEWALSPACE   = STORAGEHANDLER + DOT + "renewal_space"

	// BALANCES
	TX_Balances_Transfer = "Balances" + DOT + "transfer"

	// EVM
	TX_EVM_Call = EVM + DOT + "call"
)

// RPC Call
const (
	// System
	RPC_SYS_Properties = "system_properties"
	RPC_SYS_SyncState  = "system_syncState"
	RPC_SYS_Version    = "system_version"
	RPC_SYS_Chain      = "system_chain"

	//Net
	RPC_NET_Listening = "net_listening"
)

const (
	Active = iota
	Calculate
	Missing
	Recovery
)

const (
	MINER_STATE_POSITIVE = "positive"
	MINER_STATE_FROZEN   = "frozen"
	MINER_STATE_EXIT     = "exit"
	MINER_STATE_LOCK     = "lock"
	MINER_STATE_OFFLINE  = "offline"
)

// 0:Full 1:Verifier 2:Marker
const (
	TeeType_Full     uint8 = 0
	TeeType_Verifier uint8 = 1
	TeeType_Marker   uint8 = 2
)

const (
	ERR_Failed  = "failed"
	ERR_Timeout = "timeout"
	ERR_Empty   = "empty"
)

var (
	ERR_RPC_CONNECTION     = errors.New("rpc err: connection failed")
	ERR_RPC_IP_FORMAT      = errors.New("unsupported ip format")
	ERR_RPC_TIMEOUT        = errors.New("timeout")
	ERR_RPC_EMPTY_VALUE    = errors.New("empty")
	ERR_RPC_PRIORITYTOOLOW = "Priority is too low"

	ERR_IdleProofIsEmpty = errors.New("idle data proof is empty")
)

const (
	FileHashLen            = 64
	RandomLen              = 20
	PeerIdPublicKeyLen     = 38
	PoISKeyLen             = 256
	TeeSignatureLen        = 256
	AccumulatorLen         = 256
	SpaceChallengeParamLen = 8
	BloomFilterLen         = 256
	MaxSegmentNum          = 1000
	WorkerPublicKeyLen     = 32
	MasterPublicKeyLen     = 32
	EcdhPublicKeyLen       = 32
	TeeSigLen              = 64
)

type FileHash [FileHashLen]types.U8
type Random [RandomLen]types.U8
type PeerId [PeerIdPublicKeyLen]types.U8
type PoISKey_G [PoISKeyLen]types.U8
type PoISKey_N [PoISKeyLen]types.U8
type TeeSignature [TeeSignatureLen]types.U8
type Accumulator [AccumulatorLen]types.U8
type SpaceChallengeParam [SpaceChallengeParamLen]types.U64
type BloomFilter [BloomFilterLen]types.U64
type WorkerPublicKey [WorkerPublicKeyLen]types.U8
type MasterPublicKey [MasterPublicKeyLen]types.U8
type EcdhPublicKey [EcdhPublicKeyLen]types.U8
type TeeSig [TeeSigLen]types.U8

// Audit
type ChallengeInfo struct {
	MinerSnapshot    MinerSnapShot
	ChallengeElement ChallengeElement
	ProveInfo        ProveInfo
}

// Oss
type BucketInfo struct {
	FileList  []FileHash
	Authority []types.AccountID
}

// FileBank
type StorageOrder struct {
	FileSize     types.U128
	SegmentList  []SegmentList
	User         UserBrief
	CompleteList []CompleteInfo
}

type FileMetadata struct {
	SegmentList []SegmentInfo
	Owner       []UserBrief
	FileSize    types.U128
	Completion  types.U32
	State       types.U8
}

type RestoralOrderInfo struct {
	Count        types.U32
	Miner        types.AccountID
	OriginMiner  types.AccountID
	FragmentHash FileHash
	FileHash     FileHash
	GenBlock     types.U32
	Deadline     types.U32
}

type RestoralTargetInfo struct {
	Miner         types.AccountID
	ServiceSpace  types.U128
	RestoredSpace types.U128
	CoolingBlock  types.U32
}

type UserFileSliceInfo struct {
	Filehash FileHash
	Filesize types.U128
}

type SpaceProofInfo struct {
	Miner       types.AccountID
	Front       types.U64
	Rear        types.U64
	PoisKey     PoISKeyInfo
	Accumulator Accumulator
}

type SysProperties struct {
	Ss58Format    types.Bytes
	TokenDecimals types.U8
	TokenSymbol   types.Text
	SS58Prefix    types.U32
}

type SysSyncState struct {
	StartingBlock types.U32
	CurrentBlock  types.U32
	HighestBlock  types.U32
}

type OssInfo struct {
	Peerid PeerId
	Domain types.Bytes
}

type MinerInfo struct {
	BeneficiaryAccount types.AccountID
	StakingAccount     types.AccountID
	PeerId             PeerId
	Collaterals        types.U128
	Debt               types.U128
	State              types.Bytes // positive, exit, frozen, lock
	DeclarationSpace   types.U128
	IdleSpace          types.U128
	ServiceSpace       types.U128
	LockSpace          types.U128
	SpaceProofInfo     types.Option[SpaceProofInfo]
	ServiceBloomFilter BloomFilter
	TeeSig             TeeSig
}

type MinerReward struct {
	TotalReward  types.U128
	RewardIssued types.U128
	OrderList    []RewardOrder
}

type RewardOrder struct {
	ReceiveCount     types.U8
	MaxCount         types.U8
	Atonce           types.Bool
	OrderReward      types.U128
	EachAmount       types.U128
	LastReceiveBlock types.U32
}

type SegmentInfo struct {
	Hash         FileHash
	FragmentList []FragmentInfo
}

type UserBrief struct {
	User       types.AccountID
	FileName   types.Bytes
	BucketName types.Bytes
}

type FragmentInfo struct {
	Hash  FileHash
	Avail types.Bool
	Tag   types.Option[types.U32]
	Miner types.AccountID
}

type SegmentList struct {
	SegmentHash  FileHash
	FragmentHash []FileHash
}

type CompleteInfo struct {
	Index types.U8
	Miner types.AccountID
}

type UserSpaceInfo struct {
	TotalSpace     types.U128
	UsedSpace      types.U128
	LockedSpace    types.U128
	RemainingSpace types.U128
	Start          types.U32
	Deadline       types.U32
	State          types.Bytes
}

type ProveInfo struct {
	Assign       types.U8
	IdleProve    types.Option[IdleProveInfo]
	ServiceProve types.Option[ServiceProveInfo]
}

type ChallengeElement struct {
	Start        types.U32
	IdleSlip     types.U32
	ServiceSlip  types.U32
	VerifySlip   types.U32
	SpaceParam   SpaceChallengeParam
	ServiceParam QElement
}

type QElement struct {
	Index []types.U32
	Value []Random
}

type MinerSnapShot struct {
	IdleSpace          types.U128
	ServiceSpace       types.U128
	ServiceBloomFilter BloomFilter
	SpaceProofInfo     SpaceProofInfo
	TeeSig             TeeSig
}

type IdleProveInfo struct {
	TeePubkey    WorkerPublicKey
	IdleProve    types.Bytes
	VerifyResult types.Option[bool]
}

type ServiceProveInfo struct {
	TeePubkey    WorkerPublicKey
	ServiceProve types.Bytes
	VerifyResult types.Option[bool]
}

type TeeWorkerInfo struct {
	Pubkey              WorkerPublicKey
	EcdhPubkey          EcdhPublicKey
	Version             types.U32
	LastUpdated         types.U64
	StashAccount        types.Option[types.AccountID]
	AttestationProvider types.Option[types.U8]
	ConfidenceLevel     types.U8
	Features            []types.U32
	Role                types.U8 // 0:Full 1:Verifier 2:Marker
}

type ExpendersInfo struct {
	K types.U64
	N types.U64
	D types.U64
}

type PoISKeyInfo struct {
	G PoISKey_G
	N PoISKey_N
}

type IdleSignInfo struct {
	Miner              types.AccountID
	Rear               types.U64
	Front              types.U64
	Accumulator        Accumulator
	LastOperationBlock types.U32
	PoisKey            PoISKeyInfo
}
type TagSigInfo struct {
	Miner    types.AccountID
	Digest   []DigestInfo
	Filehash FileHash
}

type DigestInfo struct {
	Fragment  FileHash
	TeePubkey WorkerPublicKey
}

type StakingExposure struct {
	Total  types.U128
	Own    types.U128
	Others []OtherStakingExposure
}

type OtherStakingExposure struct {
	Who   types.AccountID
	Value types.U128
}

type StakingEraRewardPoints struct {
	Total      types.U32
	Individual []Individual
}

type Individual struct {
	Acc    types.AccountID
	Reward types.U32
}

type StakingNominations struct {
	Targets     []types.AccountID
	SubmittedIn types.U32
	Suppressed  types.Bool
}

type StakingValidatorPrefs struct {
	Commission types.U32
	Blocked    types.Bool
}

type CompleteSnapShotType struct {
	MinerCount types.U32
	TotalPower types.U128
}

type RoundRewardType struct {
	TotalReward types.U128
	OtherReward types.U128
}

// --------------------customer-----------------
type IdleFileMeta struct {
	BlockNum uint32
	MinerAcc []byte
	Hash     string
}

type UserSpaceSt struct {
	TotalSpace     string
	UsedSpace      string
	LockedSpace    string
	RemainingSpace string
	State          string
	Start          uint32
	Deadline       uint32
}

type NetSnapshot struct {
	Start               uint32
	Life                uint32
	Total_reward        string
	Total_idle_space    string
	Total_service_space string
	Random_index_list   []uint32
	Random              [][]byte
}

type MinerSnapshot struct {
	Miner         string
	Idle_space    string
	Service_space string
}

type TeeInfo struct {
	Pubkey              string
	EcdhPubkey          string
	Version             uint32
	LastUpdated         uint64
	StashAccount        string
	AttestationProvider uint8
	ConfidenceLevel     uint8
	Features            []uint32
	WorkerRole          uint8 // 0:Full 1:Verifier 2:Marker
}

type RewardsType struct {
	Total   string
	Claimed string
}

type SegmentDataInfo struct {
	SegmentHash  string
	FragmentHash []string
}

type UserInfo struct {
	UserAccount string
	FileName    string
	BucketName  string
	FileSize    uint64
}

type AccessInfo struct {
	r types.H160
	c []types.H160
}