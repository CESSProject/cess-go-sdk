/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

// DOT is "." character
const DOT = "."

// Unit precision of CESS token
const TokenPrecision_CESS = "000000000000"

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
	FILE           = "File"
	BUCKET         = "Bucket"
	BUCKETLIST     = "UserBucketList"
	DEALMAP        = "DealMap"
	PENDINGREPLACE = "PendingReplacements"

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
	TX_FILEBANK_PUTBUCKET     = FILEBANK + DOT + "create_bucket"
	TX_FILEBANK_DELBUCKET     = FILEBANK + DOT + "delete_bucket"
	TX_FILEBANK_DELFILE       = FILEBANK + DOT + "delete_file"
	TX_FILEBANK_UPLOADDEC     = FILEBANK + DOT + "upload_declaration"
	TX_FILEBANK_UPLOADFILLER  = FILEBANK + DOT + "upload_filler"
	TX_FILEBANK_FILEREPORT    = FILEBANK + DOT + "transfer_report"
	TX_FILEBANK_REPLACEFILE   = FILEBANK + DOT + "replace_file_report"
	TX_FILEBANK_MINEREXITPREP = FILEBANK + DOT + "miner_exit_prep"
	TX_FILEBANK_WITHDRAW      = FILEBANK + DOT + "withdraw"
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
	Ss58Format    types.U8
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
	Completion  types.U32
	State       types.U8
	SegmentList []SegmentInfo
	Owner       []UserBrief
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
	SegmentList   []SegmentList
	NeededList    []SegmentList
	User          UserBrief
	AssignedMiner []MinerTaskList
	ShareInfo     []SegmentInfo
	CompleteList  []types.AccountID
}

type IdleMetadata struct {
	Size      types.U64
	BlockNum  types.U32
	BlockSize types.U32
	ScanSize  types.U32
	Acc       types.AccountID
	Hash      FileHash
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

func CompareSlice(s1, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
