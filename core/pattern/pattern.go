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

// BlockInterval is the time interval for generating blocks, in seconds
const BlockInterval = time.Second * time.Duration(6)

const MaxSubmitedIdleFileMeta = 30

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
	// SMINER is a module about storage miners
	STORAGEHANDLER = "StorageHandler"
	// SYSTEM is a module about the system
	SYSTEM = "System"
)

// Chain state
const (
	//AUDIT
	UNVERIFYPROOF = "UnverifyProof"

	// OSS
	// OSS
	AUTHORITYLIST = "AuthorityList"

	// SMINER
	ALLMINER   = "AllMiner"
	MINERITEMS = "MinerItems"
	REWARDMAP  = "RewardMap"

	// TEEWORKER
	TEEWORKERMAP = "TeeWorkerMap"
	TEEPODR2Pk   = "TeePodr2Pk"

	// FILEBANK
	FILE               = "File"
	BUCKET             = "Bucket"
	BUCKETLIST         = "UserBucketList"
	DEALMAP            = "DealMap"
	FILLERMAP          = "FillerMap"
	PENDINGREPLACE     = "PendingReplacements"
	RESTORALORDER      = "RestoralOrder"
	RESTORALTARGETINFO = "RestoralTarget"

	// STORAGEHANDLER
	USERSPACEINFO = "UserOwnedSpace"
	UNITPRICE     = "UnitPrice"

	// NETSNAPSHOT
	CHALLENGESNAPSHOT = "ChallengeSnapShot"

	// SYSTEM
	ACCOUNT = "Account"
	EVENTS  = "Events"
)

// Extrinsics
const (
	//AUDIT
	TX_AUDIT_SUBMITPROOF = AUDIT + DOT + "submit_proof"

	// OSS
	TX_OSS_REGISTER = OSS + DOT + "register"
	TX_OSS_UPDATE   = OSS + DOT + "update"
	TX_OSS_DESTORY  = OSS + DOT + "destroy"

	// SMINER
	TX_SMINER_REGISTER       = SMINER + DOT + "regnstk"
	TX_SMINER_INCREASESTAKES = SMINER + DOT + "increase_collateral"
	TX_SMINER_UPDATEPEERID   = SMINER + DOT + "update_peer_id"
	TX_SMINER_UPDATEINCOME   = SMINER + DOT + "update_beneficiary"
	TX_SMINER_CLAIMREWARD    = SMINER + DOT + "receive_reward"

	// FILEBANK
	TX_FILEBANK_PUTBUCKET         = FILEBANK + DOT + "create_bucket"
	TX_FILEBANK_DELBUCKET         = FILEBANK + DOT + "delete_bucket"
	TX_FILEBANK_DELFILE           = FILEBANK + DOT + "delete_file"
	TX_FILEBANK_DELFILLER         = FILEBANK + DOT + "delete_filler"
	TX_FILEBANK_UPLOADDEC         = FILEBANK + DOT + "upload_declaration"
	TX_FILEBANK_UPLOADFILLER      = FILEBANK + DOT + "upload_filler"
	TX_FILEBANK_FILEREPORT        = FILEBANK + DOT + "transfer_report"
	TX_FILEBANK_REPLACEFILE       = FILEBANK + DOT + "replace_file_report"
	TX_FILEBANK_MINEREXITPREP     = FILEBANK + DOT + "miner_exit_prep"
	TX_FILEBANK_WITHDRAW          = FILEBANK + DOT + "miner_withdraw"
	TX_FILEBANK_GENRESTOREORDER   = FILEBANK + DOT + "generate_restoral_order"
	TX_FILEBANK_CLAIMRESTOREORDER = FILEBANK + DOT + "claim_restoral_order"
	TX_FILEBANK_CLAIMNOEXISTORDER = FILEBANK + DOT + "claim_restoral_noexist_order"
	TX_FILEBANK_RESTORALCOMPLETE  = FILEBANK + DOT + "restoral_order_complete"
)

// RPC Call
const (
	// System
	RPC_SYS_Properties = "system_properties"
	RPC_SYS_SyncState  = "system_syncState"
	RPC_SYS_Version    = "system_version"

	//Net
	RPC_NET_Listening = "net_listening"
)

const (
	Role_OSS    = "OSS"
	Role_DEOSS  = "DEOSS"
	Role_BUCKET = "BUCKET"
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
)

// byte size
const (
	SIZE_1KiB = 1024
	SIZE_1MiB = 1024 * SIZE_1KiB
	SIZE_1GiB = 1024 * SIZE_1MiB
)

const (
	SegmentSize  = 16 * SIZE_1MiB
	FragmentSize = 8 * SIZE_1MiB
	BlockNumber  = 1024
	DataShards   = 2
	ParShards    = 1
)

var (
	ERR_RPC_CONNECTION  = errors.New("rpc connection failed")
	ERR_RPC_IP_FORMAT   = errors.New("unsupported ip format")
	ERR_RPC_TIMEOUT     = errors.New("timeout")
	ERR_RPC_EMPTY_VALUE = errors.New("empty")
)

type FileHash [64]types.U8
type Random [20]types.U8
type TeePodr2Pk [270]types.U8
type PeerId [38]types.U8

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

type MinerInfo struct {
	BeneficiaryAcc types.AccountID
	PeerId         PeerId
	Collaterals    types.U128
	Debt           types.U128
	State          types.Bytes
	IdleSpace      types.U128
	ServiceSpace   types.U128
	LockSpace      types.U128
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
	Account types.AccountID
	Hash    []FileHash
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
	NeededList    []SegmentList
	User          UserBrief
	AssignedMiner []MinerTaskList
	ShareInfo     []SegmentInfo
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

type ChallengeSnapShot struct {
	NetSnapshot   NetSnapShot
	MinerSnapShot []MinerSnapShot
}

type NetSnapShot struct {
	Start             types.U32
	Life              types.U32
	TotalReward       types.U128
	TotalIdleSpace    types.U128
	TotalServiceSpace types.U128
	RandomIndexList   []types.U32
	Random            []Random
}

type MinerSnapShot struct {
	Miner        types.AccountID
	IdleSpace    types.U128
	ServiceSpace types.U128
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

type ChallengeInfo struct {
	Random          [][]byte
	RandomIndexList []uint32
	Start           uint32
}

type ChallengeSnapshot struct {
	NetSnapshot   NetSnapshot
	MinerSnapshot []MinerSnapshot
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
