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

// the minimum account balance required for transaction
const MinTransactionBalance = "1000000000000000000"

const StakingStakePerTiB = 4000

const BlockIntervalSec = 6

const CESSWalletLen = 49

// BlockInterval is the time interval for generating blocks, in seconds
const BlockInterval = time.Second * time.Duration(BlockIntervalSec)

const TreasuryAccount = "cXhT9Xh3DhrBMDmXcGeMPDmTzDm1J8vDxBtKvogV33pPshnWS"

const (
	SIZE_1KiB = 1024
	SIZE_1MiB = 1024 * SIZE_1KiB
	SIZE_1GiB = 1024 * SIZE_1MiB
	SIZE_1TiB = 1024 * SIZE_1GiB

	NumberOfDataCopies = 3

	SegmentSize  = 32 * SIZE_1MiB
	FragmentSize = 8 * SIZE_1MiB
	DataShards   = 4
	ParShards    = 8
	TotalShards  = DataShards + ParShards
)

const (
	MinBucketNameLength uint8 = 3
	MaxBucketNameLength uint8 = 63
	MaxDomainNameLength uint8 = 100
)

const ZeroFileHash_8M string = "2daeb1f36095b44b318410b3f4e8b5d989dcc7bb023d1426c492dab0a3053e74"

// pallet names
const (
	// Audit
	Audit = "Audit"
	// Babe
	Babe = "Babe"
	// Balances
	Balances = "Balances"
	// CessTreasury
	CessTreasury = "CessTreasury"
	// EVM
	EVM = "EVM"
	// FileBank
	FileBank = "FileBank"
	// Oss
	Oss = "Oss"

	// SchedulerCredit
	SchedulerCredit = "SchedulerCredit"
	// Session
	Session = "Session"
	// Sminer
	Sminer = "Sminer"
	// Staking
	Staking = "Staking"
	// StorageHandler
	StorageHandler = "StorageHandler"

	// System
	System = "System"
	// TeeWorker
	TeeWorker = "TeeWorker"
)

// chain state
const (
	// Audit
	ChallengeSlip        = "ChallengeSlip"
	ChallengeSnapShot    = "ChallengeSnapShot"
	CountedClear         = "CountedClear"
	CountedServiceFailed = "CountedServiceFailed"
	VerifySlip           = "VerifySlip"

	// Babe
	Authorities = "Authorities"

	// Balances
	TotalIssuance    = "TotalIssuance"
	InactiveIssuance = "InactiveIssuance"

	// CessTreasury
	CurrencyReward = "CurrencyReward"
	EraReward      = "EraReward"
	ReserveReward  = "ReserveReward"
	RoundReward    = "RoundReward"

	// FileBank
	File                = "File"
	Bucket              = "Bucket"
	DealMap             = "DealMap"
	FillerMap           = "FillerMap"
	PendingReplacements = "PendingReplacements"
	RestoralOrder       = "RestoralOrder"
	UserBucketList      = "UserBucketList"
	UserHoldFileList    = "UserHoldFileList"

	// Oss
	// Oss
	AuthorityList = "AuthorityList"

	// SchedulerCredit
	CurrentCounters = "CurrentCounters"

	// Session
	// Validators = "Validators"
	KeyOwner = "keyOwner"

	// Sminer
	AllMiner              = "AllMiner"
	CounterForMinerItems  = "CounterForMinerItems"
	MinerItems            = "MinerItems"
	RewardMap             = "RewardMap"
	Expenders             = "Expenders"
	RestoralTarget        = "RestoralTarget"
	StakingStartBlock     = "StakingStartBlock"
	CompleteSnapShot      = "CompleteSnapShot"
	CompleteMinerSnapShot = "CompleteMinerSnapShot"

	// Staking
	CounterForValidators = "CounterForValidators"
	CounterForNominators = "CounterForNominators"
	CurrentEra           = "CurrentEra"
	ErasTotalStake       = "ErasTotalStake"
	ErasStakers          = "ErasStakers"
	ErasStakersPaged     = "ErasStakersPaged"
	ErasStakersOverview  = "ErasStakersOverview"
	ErasRewardPoints     = "ErasRewardPoints"
	Ledger               = "Ledger"
	Nominators           = "Nominators"
	Bonded               = "Bonded"
	Validators           = "Validators"
	ErasValidatorReward  = "ErasValidatorReward"
	ValidatorCount       = "ValidatorCount"

	// StorageHandler
	UserOwnedSpace    = "UserOwnedSpace"
	UnitPrice         = "UnitPrice"
	TotalIdleSpace    = "TotalIdleSpace"
	TotalServiceSpace = "TotalServiceSpace"
	PurchasedSpace    = "PurchasedSpace"
	Territory         = "Territory"
	PayOrder          = "PayOrder"
	Consignment       = "Consignment"

	// System
	Account = "Account"
	Events  = "Events"

	// TeeWorker
	Workers       = "Workers"
	MasterPubkey  = "MasterPubkey"
	Endpoints     = "Endpoints"
	WorkerAddedAt = "WorkerAddedAt"
)

// RPC Call
const (
	// Chain
	RPC_Chain_getBlock         = "chain_getBlock"
	RPC_Chain_getBlockHash     = "chain_getBlockHash"
	RPC_Chain_getFinalizedHead = "chain_getFinalizedHead"

	//Net
	RPC_NET_Listening = "net_listening"

	// System
	RPC_SYS_Properties = "system_properties"
	RPC_SYS_SyncState  = "system_syncState"
	RPC_SYS_Version    = "system_version"
	RPC_SYS_Chain      = "system_chain"
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
	ERR_Failed           = "failed"
	ERR_Timeout          = "timeout"
	ERR_Empty            = "empty"
	ERR_PriorityIsTooLow = "Priority is too low"
)

var (
	ERR_RPC_CONNECTION   = errors.New("rpc err: connection failed")
	ERR_RPC_IP_FORMAT    = errors.New("unsupported ip format")
	ERR_RPC_TIMEOUT      = errors.New("timeout")
	ERR_RPC_EMPTY_VALUE  = errors.New("empty")
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
	RrscAppPublicLen       = 32
	TerritoryKeyLen        = 32
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
type RrscAppPublic [RrscAppPublicLen]types.U8
type AppPublicType [4]types.U8

var RrscAppPublicType = AppPublicType{'r', 'r', 's', 'c'}
var AudiAppPublicType = AppPublicType{'a', 'u', 'd', 'i'}
var GranAppPublicType = AppPublicType{'g', 'r', 'a', 'n'}
var ImonAppPublicType = AppPublicType{'i', 'm', 'o', 'n'}

// Audit
type ChallengeInfo struct {
	MinerSnapshot    MinerSnapShot
	ChallengeElement ChallengeElement
	ProveInfo        ProveInfo
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

type SpaceProofInfo struct {
	Miner       types.AccountID
	Front       types.U64
	Rear        types.U64
	PoisKey     PoISKeyInfo
	Accumulator Accumulator
}

type PoISKeyInfo struct {
	G PoISKey_G
	N PoISKey_N
}

type ProveInfo struct {
	Assign       types.U8
	IdleProve    types.Option[IdleProveInfo]
	ServiceProve types.Option[ServiceProveInfo]
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

// babe
type ConsensusRrscAppPublic struct {
	Public  RrscAppPublic
	Unknown types.U64
}

// Oss
type BucketInfo struct {
	FileList  []FileHash
	Authority []types.AccountID
}

type OssInfo struct {
	Peerid PeerId
	Domain types.Bytes
}

// FileBank
type StorageOrder struct {
	FileSize     types.U128
	SegmentList  []SegmentList
	User         UserBrief
	CompleteList []CompleteInfo
}

type SegmentList struct {
	SegmentHash  FileHash
	FragmentHash []FileHash
}

type CompleteInfo struct {
	Index types.U8
	Miner types.AccountID
}

type FileMetadata struct {
	SegmentList []SegmentInfo
	Owner       []UserBrief
	FileSize    types.U128
	Completion  types.U32
	State       types.U8
}

type SegmentInfo struct {
	Hash         FileHash
	FragmentList []FragmentInfo
}

type FragmentInfo struct {
	Hash  FileHash
	Avail types.Bool
	Tag   types.Option[types.U32]
	Miner types.AccountID
}

type UserBrief struct {
	User          types.AccountID
	FileName      types.Bytes
	BucketName    types.Bytes
	TerriortyName types.Bytes
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

// SchedulerCredit
type SchedulerCounterEntry struct {
	ProceedBlockSize uint64
	PunishmentCount  uint32
}

// Session
type KeyOwnerParam struct {
	PublicType AppPublicType
	Public     types.Bytes
}

// Sminer
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

// StorageHandler
// type UserSpaceInfo struct {
// 	TotalSpace     types.U128
// 	UsedSpace      types.U128
// 	LockedSpace    types.U128
// 	RemainingSpace types.U128
// 	Start          types.U32
// 	Deadline       types.U32
// 	State          types.Bytes
// }

type TerritoryInfo struct {
	Token          types.H256
	TotalSpace     types.U128
	UsedSpace      types.U128
	LockedSpace    types.U128
	RemainingSpace types.U128
	Start          types.U32
	Deadline       types.U32
	State          types.U8 //0: Active 1: Frozen 2: Expired 3: OnConsignment
}

// type OrderInfo struct {
// 	TerritoryName [TerritoryKeyLen]types.U8
// 	Pay           types.U128
// 	GibCount      types.U32
// 	Days          types.U32
// 	Expired       types.U32
// 	TargetAcc     types.AccountID
// 	OrderType     types.U8
// }

type ConsignmentInfo struct {
	User   types.AccountID
	Price  types.U128
	Buyers types.OptionAccountID
	Exec   types.OptionU32
	Locked types.Bool
}

type UserFileSliceInfo struct {
	TerritoryName types.Bytes
	Filehash      FileHash
	FileSize      types.U128
}

// Staking
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

type StakingLedger struct {
	Stash          types.AccountID
	Total          types.UCompact
	Active         types.UCompact
	Unlocking      []UnlockChunk
	ClaimedRewards []types.U32
}

type UnlockChunk struct {
	Value types.UCompact
	Era   types.BlockNumber
}

// System
type SysProperties struct {
	IsEthereum    types.Bool
	Ss58Format    types.U32
	TokenDecimals types.U8
	TokenSymbol   types.Text
}

type SysSyncState struct {
	StartingBlock types.U32
	CurrentBlock  types.U32
	HighestBlock  types.U32
}

// TeeWorker
type WorkerInfo struct {
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
	Total  types.UCompact
	Own    types.UCompact
	Others []OtherStakingExposure
}

type StakingExposurePaged struct {
	PageTotal types.UCompact
	Others    []OtherStakingExposure
}

type OtherStakingExposure struct {
	Who   types.AccountID
	Value types.UCompact
}

type PagedExposureMetadata struct {
	Total          types.UCompact
	Own            types.UCompact
	NominatorCount types.U32
	PageCount      types.U32
}

type StakingValidatorPrefs struct {
	Commission types.U32
	Blocked    types.Bool
}

type CompleteSnapShotType struct {
	MinerCount types.U32
	TotalPower types.U128
}

type MinerCompleteInfo struct {
	EraIndex    types.U32
	Issued      types.Bool
	FinishBlock types.U32
	Power       types.U128
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
	R types.H160
	C []types.H160
}

type BlockData struct {
	BlockHash           string
	PreHash             string
	ExtHash             string
	StHash              string
	AllGasFee           string
	Timestamp           int64
	BlockId             uint32
	IsNewEra            bool
	EraPaid             EraPaid
	SysEvents           []string
	NewAccounts         []string
	GenChallenge        []string
	StorageCompleted    []string
	MinerReg            []MinerRegInfo
	Extrinsics          []ExtrinsicsInfo
	TransferInfo        []TransferInfo
	UploadDecInfo       []UploadDecInfo
	DeleteFileInfo      []DeleteFileInfo
	CreateBucketInfo    []CreateBucketInfo
	DeleteBucketInfo    []DeleteBucketInfo
	SubmitIdleProve     []SubmitIdleProve
	SubmitServiceProve  []SubmitServiceProve
	SubmitIdleResult    []SubmitIdleResult
	SubmitServiceResult []SubmitServiceResult
	Punishment          []Punishment
	MinerRegPoiskeys    []MinerRegPoiskey
	GatewayReg          []GatewayReg
	StakingPayouts      []StakingPayout
	Unbonded            []Unbonded
	MintTerritory       []MintTerritory
}

type ExtrinsicsInfo struct {
	Name    string
	Signer  string
	Hash    string
	FeePaid string
	Result  bool
	Events  []string
}

type TransferInfo struct {
	ExtrinsicName string
	ExtrinsicHash string
	From          string
	To            string
	Amount        string
	Result        bool
}

type UploadDecInfo struct {
	ExtrinsicHash string
	Owner         string
	Fid           string
}

type DeleteFileInfo struct {
	ExtrinsicHash string
	Owner         string
	Fid           string
}

type MinerRegInfo struct {
	ExtrinsicHash string
	Account       string
}

type CreateBucketInfo struct {
	ExtrinsicHash string
	Owner         string
	BucketName    string
}

type DeleteBucketInfo struct {
	ExtrinsicHash string
	Owner         string
	BucketName    string
}

type SubmitIdleProve struct {
	ExtrinsicHash string
	Miner         string
}

type SubmitServiceProve struct {
	ExtrinsicHash string
	Miner         string
}

type SubmitIdleResult struct {
	ExtrinsicHash string
	Miner         string
	Result        bool
}

type SubmitServiceResult struct {
	ExtrinsicHash string
	Miner         string
	Result        bool
}

type Punishment struct {
	ExtrinsicName string
	ExtrinsicHash string
	From          string
	To            string
	Amount        string
}

type MinerRegPoiskey struct {
	ExtrinsicHash string
	Miner         string
}

type GatewayReg struct {
	ExtrinsicHash string
	Account       string
}

type EraPaid struct {
	HaveValue       bool
	EraIndex        uint32
	ValidatorPayout string
	Remainder       string
}

type StakingPayout struct {
	EraIndex      uint32
	ExtrinsicHash string
	ClaimedAcc    string
	Amount        string
}

type Unbonded struct {
	ExtrinsicHash string
	Account       string
	Amount        string
}

type MintTerritory struct {
	ExtrinsicHash  string
	Account        string
	TerritoryToken string
	TerritoryName  string
	TerritorySize  uint64
}
