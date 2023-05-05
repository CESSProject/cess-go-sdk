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
	//
	NETSNAPSHOT = "NetSnapShotStorage"
	// SYSTEM is a module about the system
	SYSTEM = "System"
)

// Chain state
const (
	// OSS
	// OSS
	AUTHORITYLIST = "AuthorityList"

	// SMINER
	ALLMINER   = "AllMiner"
	MINERITEMS = "MinerItems"

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
	// OSS
	TX_OSS_REGISTER = OSS + DOT + "register"
	TX_OSS_UPDATE   = OSS + DOT + "update"
	TX_OSS_DESTORY  = OSS + DOT + "destroy"

	// SMINER
	TX_SMINER_REGISTER       = SMINER + DOT + "regnstk"
	TX_SMINER_EXIT           = SMINER + DOT + "exit_miner"
	TX_SMINER_INCREASESTAKES = SMINER + DOT + "increase_collateral"
	TX_SMINER_UPDATEADDR     = SMINER + DOT + "update_ip"
	TX_SMINER_UPDATEINCOME   = SMINER + DOT + "update_beneficiary"

	// FILEBANK
	TX_FILEBANK_PUTBUCKET    = FILEBANK + DOT + "create_bucket"
	TX_FILEBANK_DELBUCKET    = FILEBANK + DOT + "delete_bucket"
	TX_FILEBANK_DELFILE      = FILEBANK + DOT + "delete_file"
	TX_FILEBANK_UPLOADDEC    = FILEBANK + DOT + "upload_declaration"
	TX_FILEBANK_ADDIDLESPACE = FILEBANK + DOT + "test_add_idle_space"
	TX_FILEBANK_FILEREPORT   = FILEBANK + DOT + "transfer_report"
	TX_FILEBANK_REPLACEFILE  = FILEBANK + DOT + "replace_file_report"
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
type TeePodr2Pk [294]types.U8
type PeerID [53]types.U8

type MinerInfo struct {
	BeneficiaryAcc types.AccountID
	PeerId         types.Bytes
	Collaterals    types.U128
	Debt           types.U128
	State          types.Bytes
	IdleSpace      types.U128
	ServiceSpace   types.U128
	LockSpace      types.U128
}

type RewardInfo struct {
	TotalReward              types.U128
	RewardIssued             types.U128
	CurrentlyAvailableReward types.U128
	OrderList                []RewardOrder
}

type RewardOrder struct {
	OrderReward types.U128
	EachShare   types.U128
	AwardCount  types.U128
	HasIssued   types.U128
}

type FileMetadata struct {
	Completion  types.U32
	State       types.U8
	SegmentList []SegmentInfo
	Owner       []UserBrief
}

type BucketInfo struct {
	Objects_list []FileHash
	Authority    []types.AccountID
}

type UserBrief struct {
	User        types.AccountID
	File_name   types.Bytes
	Bucket_name types.Bytes
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
	Start               types.U32
	Total_reward        types.U128
	Total_idle_space    types.U128
	Total_service_space types.U128
	Random              Random
}

type MinerSnapShot struct {
	Miner         types.AccountID
	Idle_space    types.U128
	Service_space types.U128
}

type NodePublickey struct {
	Node_publickey [32]types.U8
}

type TeeWorkerInfo struct {
	Controller_account types.AccountID
	Peer_id            PeerID
	Node_key           NodePublickey
	Stash_account      types.AccountID
}
