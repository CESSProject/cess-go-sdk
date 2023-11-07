/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package pattern

import (
	"time"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

// DOT is "." character
const DOT = "."

// Unit precision of CESS token
const TokenPrecision_CESS = "000000000000"

const StakingStakePerTiB = 4000

// BlockInterval is the time interval for generating blocks, in seconds
const BlockInterval = time.Second * time.Duration(6)

const MaxSubmitedIdleFileMeta = 30

const PublicDeoss = "http://deoss-pub-gateway.cess.cloud/"
const PublicDeossAccount = "cXhwBytXqrZLr1qM5NHJhCzEMckSTzNKw17ci2aHft6ETSQm9"

// Pallets
const (
	//
	AUDIT = "Audit"
	// OSS is a module about DeOSS
	OSS = "Oss"
	// FILEBANK is a module about data metadata, bucket info, etc.
	FILEBANK = "FileBank"
	// TEEWOEKER is a module about TEE
	TEEWORKER = "TeeWorker"
	// SMINER is a module about storage miners
	SMINER = "Sminer"
	// STAKING is a module about staking
	STAKING = "Staking"
	// SMINER is a module about storage miners
	STORAGEHANDLER = "StorageHandler"
	// SYSTEM is a module about the system
	SYSTEM = "System"
)

// Chain state
const (
	//AUDIT
	UNVERIFYPROOF           = "UnverifyProof"
	CHALLENGEDURATION       = "ChallengeDuration"
	CHALLENGEVERIFYDURATION = "VerifyDuration"
	CHALLENGESNAPSHOT       = "ChallengeSnapShot"
	UNVERIFYIDLEPROOF       = "UnverifyIdleProof"
	UNVERIFYSERVICEPROOF    = "UnverifyServiceProof"

	// OSS
	// OSS
	AUTHORITYLIST = "AuthorityList"

	// SMINER
	ALLMINER           = "AllMiner"
	MINERITEMS         = "MinerItems"
	REWARDMAP          = "RewardMap"
	EXPENDERS          = "Expenders"
	RESTORALTARGETINFO = "RestoralTarget"

	// TEEWORKER
	TEEWORKERMAP = "TeeWorkerMap"
	TEEPODR2PK   = "TeePodr2Pk"

	// FILEBANK
	FILE           = "File"
	BUCKET         = "Bucket"
	BUCKETLIST     = "UserBucketList"
	DEALMAP        = "DealMap"
	FILLERMAP      = "FillerMap"
	PENDINGREPLACE = "PendingReplacements"
	RESTORALORDER  = "RestoralOrder"

	// STAKING
	COUNTERFORVALIDATORS = "CounterForValidators"

	// STORAGE_HANDLER
	USERSPACEINFO     = "UserOwnedSpace"
	UNITPRICE         = "UnitPrice"
	TOTALIDLESPACE    = "TotalIdleSpace"
	TOTALSERVICESPACE = "TotalServiceSpace"
	PURCHASEDSPACE    = "PurchasedSpace"

	// SYSTEM
	ACCOUNT = "Account"
	EVENTS  = "Events"
)

// Extrinsics
const (
	//AUDIT
	TX_AUDIT_SUBMITPROOF              = AUDIT + DOT + "submit_proof"
	TX_AUDIT_SUBMITIDLEPROOF          = AUDIT + DOT + "submit_idle_proof"
	TX_AUDIT_SUBMITSERVICEPROOF       = AUDIT + DOT + "submit_service_proof"
	TX_AUDIT_SUBMITIDLEPROOFRESULT    = AUDIT + DOT + "submit_verify_idle_result"
	TX_AUDIT_SUBMITSERVICEPROOFRESULT = AUDIT + DOT + "submit_verify_service_result"

	// OSS
	TX_OSS_REGISTER    = OSS + DOT + "register"
	TX_OSS_UPDATE      = OSS + DOT + "update"
	TX_OSS_DESTROY     = OSS + DOT + "destroy"
	TX_OSS_AUTHORIZE   = OSS + DOT + "authorize"
	TX_OSS_UNAUTHORIZE = OSS + DOT + "cancel_authorize"

	// SMINER
	TX_SMINER_REGISTER        = SMINER + DOT + "regnstk"
	TX_SMINER_INCREASESTAKES  = SMINER + DOT + "increase_collateral"
	TX_SMINER_UPDATEPEERID    = SMINER + DOT + "update_peer_id"
	TX_SMINER_UPDATEINCOME    = SMINER + DOT + "update_beneficiary"
	TX_SMINER_CLAIMREWARD     = SMINER + DOT + "receive_reward"
	TX_SMINER_MINEREXITPREP   = SMINER + DOT + "miner_exit_prep"
	TX_SMINER_WITHDRAW        = SMINER + DOT + "miner_withdraw"
	TX_SMINER_REGISTERPOISKEY = SMINER + DOT + "register_pois_key"

	// FILEBANK
	TX_FILEBANK_PUTBUCKET         = FILEBANK + DOT + "create_bucket"
	TX_FILEBANK_DELBUCKET         = FILEBANK + DOT + "delete_bucket"
	TX_FILEBANK_DELFILE           = FILEBANK + DOT + "delete_file"
	TX_FILEBANK_DELFILLER         = FILEBANK + DOT + "delete_filler"
	TX_FILEBANK_UPLOADDEC         = FILEBANK + DOT + "upload_declaration"
	TX_FILEBANK_UPLOADFILLER      = FILEBANK + DOT + "upload_filler"
	TX_FILEBANK_FILEREPORT        = FILEBANK + DOT + "transfer_report"
	TX_FILEBANK_REPLACEFILE       = FILEBANK + DOT + "replace_file_report"
	TX_FILEBANK_GENRESTOREORDER   = FILEBANK + DOT + "generate_restoral_order"
	TX_FILEBANK_CLAIMRESTOREORDER = FILEBANK + DOT + "claim_restoral_order"
	TX_FILEBANK_CLAIMNOEXISTORDER = FILEBANK + DOT + "claim_restoral_noexist_order"
	TX_FILEBANK_RESTORALCOMPLETE  = FILEBANK + DOT + "restoral_order_complete"
	TX_FILEBANK_CERTIDLESPACE     = FILEBANK + DOT + "cert_idle_space"
	TX_FILEBANK_REPLACEIDLESPACE  = FILEBANK + DOT + "replace_idle_space"

	// STORAGE_HANDLER
	TX_STORAGE_BUYSPACE       = STORAGEHANDLER + DOT + "buy_space"
	TX_STORAGE_EXPANSIONSPACE = STORAGEHANDLER + DOT + "expansion_space"
	TX_STORAGE_RENEWALSPACE   = STORAGEHANDLER + DOT + "renewal_space"

	// BALANCES
	TX_BALANCES_FORCETRANSFER = "Balances" + DOT + "transfer"
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
	Name_Deoss  = "deoss"
	Name_Sminer = "bucket"
	Name_SDK    = "client"
)

const DirMode = 0644

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
)

const (
	ERR_Failed  = "failed"
	ERR_Timeout = "timeout"
	ERR_Empty   = "empty"
)

const (
	MinBucketNameLength = 3
	MaxBucketNameLength = 63
	MaxDomainNameLength = 50
)

// byte size
const (
	SIZE_1KiB = 1024
	SIZE_1MiB = 1024 * SIZE_1KiB
	SIZE_1GiB = 1024 * SIZE_1MiB
)

const (
	SegmentSize  = 64 * SIZE_1MiB
	FragmentSize = 16 * SIZE_1MiB
	DataShards   = 4
	ParShards    = 2
)

var (
	ERR_RPC_CONNECTION     = errors.New("rpc connection failed")
	ERR_RPC_IP_FORMAT      = errors.New("unsupported ip format")
	ERR_RPC_TIMEOUT        = errors.New("timeout")
	ERR_RPC_EMPTY_VALUE    = errors.New("empty")
	ERR_RPC_PRIORITYTOOLOW = "Priority is too low"
)

const (
	FileHashLen            = 64
	RandomLen              = 20
	TeePodr2PkLen          = 270
	PeerIdPublicKeyLen     = 38
	PoISKeyLen             = 256
	TeeSignatureLen        = 256
	AccumulatorLen         = 256
	SpaceChallengeParamLen = 8
	BloomFilterLen         = 256
	MaxSegmentNum          = 1000
)

type FileHash [FileHashLen]types.U8
type Random [RandomLen]types.U8
type TeePodr2Pk [TeePodr2PkLen]types.U8
type PeerId [PeerIdPublicKeyLen]types.U8
type PoISKey_G [PoISKeyLen]types.U8
type PoISKey_N [PoISKeyLen]types.U8
type TeeSignature [TeeSignatureLen]types.U8
type Accumulator [AccumulatorLen]types.U8
type SpaceChallengeParam [SpaceChallengeParamLen]types.U64
type BloomFilter [BloomFilterLen]types.U64

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
	BeneficiaryAcc     types.AccountID
	PeerId             PeerId
	Collaterals        types.U128
	Debt               types.U128
	State              []types.U8
	IdleSpace          types.U128
	ServiceSpace       types.U128
	LockSpace          types.U128
	SpaceProofInfo     types.Option[SpaceProofInfo]
	ServiceBloomFilter BloomFilter
	TeeSignature       TeeSignature
}

type SpaceProofInfo struct {
	Miner       types.AccountID
	Front       types.U64
	Rear        types.U64
	PoisKey     PoISKeyInfo
	Accumulator Accumulator
}

type RewardOrder struct {
	OrderReward types.U128
	EachShare   types.U128
	AwardCount  types.U8
	HasIssued   types.Bool
}

type MinerReward struct {
	TotalReward              types.U128
	RewardIssued             types.U128
	CurrentlyAvailableReward types.U128
	OrderList                []RewardOrder
}

type FileMetadata struct {
	SegmentList []SegmentInfo
	Owner       []UserBrief
	FileSize    types.U128
	Completion  types.U32
	State       types.U8
}

type BucketInfo struct {
	ObjectsList []FileHash
	Authority   []types.AccountID
}

type UserBrief struct {
	User       types.AccountID
	FileName   types.Bytes
	BucketName types.Bytes
}

type SegmentList struct {
	SegmentHash  FileHash
	FragmentHash []FileHash
}

type MinerTaskList struct {
	Index        types.U8
	Miner        types.AccountID
	FragmentList []FileHash
}

type SegmentInfo struct {
	Hash         FileHash
	FragmentList []FragmentList
}

type FragmentList struct {
	Hash  FileHash
	Avail types.Bool
	Miner types.AccountID
}

type StorageOrder struct {
	Stage         types.U8
	Count         types.U8
	FileSize      types.U128
	SegmentList   []SegmentList
	User          UserBrief
	MinerTaskList []MinerTaskList
	CompleteList  []types.AccountID
}

type IdleMetadata struct {
	BlockNum types.U32
	Acc      types.AccountID
	Hash     FileHash
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

type ChallengeInfo struct {
	MinerSnapshot    MinerSnapShot
	ChallengeElement ChallengeElement
	ProveInfo        ProveInfo
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
	TeeSignature       TeeSignature
}

type IdleProveInfo struct {
	TeeAcc       types.AccountID
	IdleProve    types.Bytes
	VerifyResult types.Option[bool]
}

type ServiceProveInfo struct {
	TeeAcc       types.AccountID
	ServiceProve types.Bytes
	VerifyResult types.Option[bool]
}

type MinerSnapShot_V2 struct {
	Miner              types.AccountID
	IdleLife           types.U32
	ServiceLife        types.U32
	IdleSpace          types.U128
	ServiceSpace       types.U128
	IdleSubmitted      types.Bool
	ServiceSubmitted   types.Bool
	ServiceBloomFilter BloomFilter
	SpaceProofInfo     SpaceProofInfo
	TeeSignature       TeeSignature
}

type NodePublickey struct {
	NodePublickey [32]types.U8
}

type TeeWorkerInfo struct {
	ControllerAccount types.AccountID
	PeerId            PeerId
	NodeKey           NodePublickey
	StashAccount      types.AccountID
}

type ProofAssignmentInfo struct {
	SnapShot     MinerSnapShot
	IdleProve    types.Bytes
	ServiceProve types.Bytes
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

type IdleProofInfo struct {
	MinerSnapShot MinerSnapShot_V2
	IdleProofs    []types.U8
}

type ServiceProofInfo struct {
	MinerSnapShot MinerSnapShot_V2
	ServiceProofs []types.U8
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

// type ChallengeSnapshot struct {
// 	NetSnapshot   NetSnapshot
// 	MinerSnapshot []MinerSnapshot
// }

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

type TeeWorkerSt struct {
	Controller_account string
	Peer_id            []byte
	Node_key           []byte
	Stash_account      string
}

type RewardsType struct {
	Total     string
	Claimed   string
	Available string
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
