/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package event

import (
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// ******************************************************
// cess event type
// ******************************************************

// ------------------------Audit-------------------
type Event_VerifyProof struct {
	Phase     types.Phase
	TeeWorker types.AccountID
	Miner     types.AccountID
	Topics    []types.Hash
}

type Event_SubmitProof struct {
	Phase  types.Phase
	Miner  types.AccountID
	Topics []types.Hash
}

type Event_GenerateChallenge struct {
	Phase  types.Phase
	Topics []types.Hash
}

type Event_SubmitIdleProof struct {
	Phase  types.Phase
	Miner  types.AccountID
	Topics []types.Hash
}

type Event_SubmitServiceProof struct {
	Phase  types.Phase
	Miner  types.AccountID
	Topics []types.Hash
}

type Event_SubmitIdleVerifyResult struct {
	Phase  types.Phase
	Tee    types.AccountID
	Miner  types.AccountID
	Result types.Bool
	Topics []types.Hash
}

type Event_SubmitServiceVerifyResult struct {
	Phase  types.Phase
	Tee    types.AccountID
	Miner  types.AccountID
	Result types.Bool
	Topics []types.Hash
}

// ------------------------Sminer------------------------
type Event_Registered struct {
	Phase      types.Phase
	Acc        types.AccountID
	StakingVal types.U128
	Topics     []types.Hash
}

type Event_DrawFaucetMoney struct {
	Phase  types.Phase
	Topics []types.Hash
}

type Event_FaucetTopUpMoney struct {
	Phase  types.Phase
	Acc    types.AccountID
	Topics []types.Hash
}

type Event_LessThan24Hours struct {
	Phase  types.Phase
	Last   types.U32
	Now    types.U32
	Topics []types.Hash
}
type Event_AlreadyFrozen struct {
	Phase  types.Phase
	Acc    types.AccountID
	Topics []types.Hash
}

type Event_MinerExit struct {
	Phase  types.Phase
	Acc    types.AccountID
	Topics []types.Hash
}

type Event_MinerClaim struct {
	Phase  types.Phase
	Acc    types.AccountID
	Topics []types.Hash
}

type Event_IncreaseCollateral struct {
	Phase   types.Phase
	Acc     types.AccountID
	Balance types.U128
	Topics  []types.Hash
}

type Event_Deposit struct {
	Phase   types.Phase
	Balance types.U128
	Topics  []types.Hash
}

type Event_UpdataBeneficiary struct {
	Phase  types.Phase
	Acc    types.AccountID
	New    types.AccountID
	Topics []types.Hash
}

type Event_UpdataIp struct {
	Phase  types.Phase
	Acc    types.AccountID
	Old    pattern.PeerId
	New    pattern.PeerId
	Topics []types.Hash
}

type Event_Receive struct {
	Phase  types.Phase
	Acc    types.AccountID
	Reward types.U128
	Topics []types.Hash
}

type Event_MinerExitPrep struct {
	Phase  types.Phase
	Miner  types.AccountID
	Topics []types.Hash
}

// ------------------------FileBank----------------------
type Event_DeleteFile struct {
	Phase    types.Phase
	Operator types.AccountID
	Owner    types.AccountID
	Filehash []pattern.FileHash
	Topics   []types.Hash
}

type Event_FillerDelete struct {
	Phase      types.Phase
	Acc        types.AccountID
	FillerHash pattern.FileHash
	Topics     []types.Hash
}

type Event_FillerUpload struct {
	Phase    types.Phase
	Acc      types.AccountID
	Filesize types.U64
	Topics   []types.Hash
}

type Event_UploadDeclaration struct {
	Phase     types.Phase
	Operator  types.AccountID
	Owner     types.AccountID
	Deal_hash pattern.FileHash
	Topics    []types.Hash
}

type Event_CreateBucket struct {
	Phase      types.Phase
	Acc        types.AccountID
	Owner      types.AccountID
	BucketName types.Bytes
	Topics     []types.Hash
}

type Event_DeleteBucket struct {
	Phase      types.Phase
	Acc        types.AccountID
	Owner      types.AccountID
	BucketName types.Bytes
	Topics     []types.Hash
}

type Event_TransferReport struct {
	Phase       types.Phase
	Acc         types.AccountID
	Failed_list []pattern.FileHash
	Topics      []types.Hash
}

type Event_ReplaceFiller struct {
	Phase       types.Phase
	Acc         types.AccountID
	Filler_list []pattern.FileHash
	Topics      []types.Hash
}

type Event_CalculateEnd struct {
	Phase     types.Phase
	File_hash pattern.FileHash
	Topics    []types.Hash
}

type Event_GenerateRestoralOrder struct {
	Phase        types.Phase
	Miner        types.AccountID
	FragmentHash pattern.FileHash
	Topics       []types.Hash
}

type Event_ClaimRestoralOrder struct {
	Phase   types.Phase
	Miner   types.AccountID
	OrderId pattern.FileHash
	Topics  []types.Hash
}

type Event_RecoveryCompleted struct {
	Phase   types.Phase
	Miner   types.AccountID
	OrderId pattern.FileHash
	Topics  []types.Hash
}

type Event_StorageCompleted struct {
	Phase    types.Phase
	FileHash pattern.FileHash
	Topics   []types.Hash
}

type Event_Withdraw struct {
	Phase  types.Phase
	Acc    types.AccountID
	Topics []types.Hash
}

type Event_IdleSpaceCert struct {
	Phase  types.Phase
	Acc    types.AccountID
	Space  types.U128
	Topics []types.Hash
}

type Event_ReplaceIdleSpace struct {
	Phase  types.Phase
	Acc    types.AccountID
	Space  types.U128
	Topics []types.Hash
}

// ------------------------StorageHandler--------------------------------
type Event_BuySpace struct {
	Phase            types.Phase
	Acc              types.AccountID
	Storage_capacity types.U128
	Spend            types.U128
	Topics           []types.Hash
}

type Event_ExpansionSpace struct {
	Phase           types.Phase
	Acc             types.AccountID
	Expansion_space types.U128
	Fee             types.U128
	Topics          []types.Hash
}

type Event_RenewalSpace struct {
	Phase        types.Phase
	Acc          types.AccountID
	Renewal_days types.U32
	Fee          types.U128
	Topics       []types.Hash
}

type Event_LeaseExpired struct {
	Phase  types.Phase
	Acc    types.AccountID
	Size   types.U128
	Topics []types.Hash
}

type Event_LeaseExpireIn24Hours struct {
	Phase  types.Phase
	Acc    types.AccountID
	Size   types.U128
	Topics []types.Hash
}

// ------------------------TEE Worker--------------------
type Event_RegistrationTeeWorker struct {
	Phase  types.Phase
	Acc    types.AccountID
	PeerId pattern.PeerId
	Topics []types.Hash
}

type Event_UpdatePeerId struct {
	Phase  types.Phase
	Acc    types.AccountID
	Topics []types.Hash
}

type Event_Exit struct {
	Phase  types.Phase
	Acc    types.AccountID
	Topics []types.Hash
}

// ------------------------Oss---------------------------
type Event_OssRegister struct {
	Phase    types.Phase
	Acc      types.AccountID
	Endpoint pattern.PeerId
	Topics   []types.Hash
}

type Event_OssUpdate struct {
	Phase       types.Phase
	Acc         types.AccountID
	NewEndpoint pattern.PeerId
	Topics      []types.Hash
}

type Event_OssDestroy struct {
	Phase  types.Phase
	Acc    types.AccountID
	Topics []types.Hash
}

type Event_Authorize struct {
	Phase    types.Phase
	Acc      types.AccountID
	Operator types.AccountID
	Topics   []types.Hash
}

type Event_CancelAuthorize struct {
	Phase  types.Phase
	Acc    types.AccountID
	Topics []types.Hash
}

// ------------------------system------------------------
type Event_ElectionFinalized struct {
	Phase   types.Phase
	Compute types.U8
	Score   ElectionScore
	Topics  []types.Hash
}

// *******************************************************
type ElectionScore struct {
	/// The minimal winner, in terms of total backing stake.
	///
	/// This parameter should be maximized.
	Minimal_stake types.U128
	/// The sum of the total backing of all winners.
	///
	/// This parameter should maximized
	Sum_stake types.U128
	/// The sum squared of the total backing of all winners, aka. the variance.
	///
	/// Ths parameter should be minimized.
	Sum_stake_squared types.U128
}

//*******************************************************

// Events
type EventRecords struct {
	// AUDIT
	Audit_VerifyProof               []Event_VerifyProof
	Audit_SubmitProof               []Event_SubmitProof
	Audit_GenerateChallenge         []Event_GenerateChallenge
	Audit_SubmitIdleProof           []Event_SubmitIdleProof
	Audit_SubmitServiceProof        []Event_SubmitServiceProof
	Audit_SubmitIdleVerifyResult    []Event_SubmitIdleVerifyResult
	Audit_SubmitServiceVerifyResult []Event_SubmitServiceVerifyResult

	// Cacher

	// FILEBANK
	FileBank_DeleteFile            []Event_DeleteFile
	FileBank_FillerDelete          []Event_FillerDelete
	FileBank_FillerUpload          []Event_FillerUpload
	FileBank_UploadDeclaration     []Event_UploadDeclaration
	FileBank_CreateBucket          []Event_CreateBucket
	FileBank_DeleteBucket          []Event_DeleteBucket
	FileBank_TransferReport        []Event_TransferReport
	FileBank_ReplaceFiller         []Event_ReplaceFiller
	FileBank_CalculateEnd          []Event_CalculateEnd
	FileBank_GenerateRestoralOrder []Event_GenerateRestoralOrder
	FileBank_ClaimRestoralOrder    []Event_ClaimRestoralOrder
	FileBank_RecoveryCompleted     []Event_RecoveryCompleted
	FileBank_StorageCompleted      []Event_StorageCompleted
	FileBank_Withdraw              []Event_Withdraw
	FileBank_IdleSpaceCert         []Event_IdleSpaceCert
	FileBank_ReplaceIdleSpace      []Event_ReplaceIdleSpace

	// OSS
	Oss_Authorize       []Event_Authorize
	Oss_CancelAuthorize []Event_CancelAuthorize
	Oss_OssRegister     []Event_OssRegister
	Oss_OssUpdate       []Event_OssUpdate
	Oss_OssDestroy      []Event_OssDestroy

	// SMINER
	Sminer_Registered         []Event_Registered
	Sminer_DrawFaucetMoney    []Event_DrawFaucetMoney
	Sminer_FaucetTopUpMoney   []Event_FaucetTopUpMoney
	Sminer_LessThan24Hours    []Event_LessThan24Hours
	Sminer_AlreadyFrozen      []Event_AlreadyFrozen
	Sminer_IncreaseCollateral []Event_IncreaseCollateral
	Sminer_Deposit            []Event_Deposit
	Sminer_UpdataBeneficiary  []Event_UpdataBeneficiary
	Sminer_UpdataIp           []Event_UpdataIp
	Sminer_Receive            []Event_Receive
	Sminer_MinerExitPrep      []Event_MinerExitPrep

	// StorageHandler
	StorageHandler_BuySpace             []Event_BuySpace
	StorageHandler_ExpansionSpace       []Event_ExpansionSpace
	StorageHandler_RenewalSpace         []Event_RenewalSpace
	StorageHandler_LeaseExpired         []Event_LeaseExpired
	StorageHandler_LeaseExpireIn24Hours []Event_LeaseExpireIn24Hours

	// TeeWorker
	TeeWorker_RegistrationTeeWorker []Event_RegistrationTeeWorker
	TeeWorker_UpdatePeerId          []Event_UpdatePeerId
	TeeWorker_Exit                  []Event_Exit

	// system
	ElectionProviderMultiPhase_ElectionFinalized []Event_ElectionFinalized

	// system-gsrpc
	types.EventRecords
}
