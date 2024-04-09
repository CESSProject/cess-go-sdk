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
	TeeWorker pattern.WorkerPublicKey
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
	Tee    pattern.WorkerPublicKey
	Miner  types.AccountID
	Result types.Bool
	Topics []types.Hash
}

type Event_SubmitServiceVerifyResult struct {
	Phase  types.Phase
	Tee    pattern.WorkerPublicKey
	Miner  types.AccountID
	Result types.Bool
	Topics []types.Hash
}

// ------------------------Sminer------------------------
type Event_Registered struct {
	Phase  types.Phase
	Acc    types.AccountID
	Topics []types.Hash
}

type Event_RegisterPoisKey struct {
	Phase  types.Phase
	Miner  types.AccountID
	Topics []types.Hash
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

type Event_UpdateBeneficiary struct {
	Phase  types.Phase
	Acc    types.AccountID
	New    types.AccountID
	Topics []types.Hash
}

type Event_UpdatePeerId struct {
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
	Phase    types.Phase
	Acc      types.AccountID
	DealHash pattern.FileHash
	Topics   []types.Hash
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

type Event_IncreaseDeclarationSpace struct {
	Phase  types.Phase
	Miner  types.AccountID
	Space  types.U128
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

type Event_CalculateReport struct {
	Phase    types.Phase
	Miner    types.AccountID
	FileHash pattern.FileHash
	Topics   []types.Hash
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
type Event_Exit struct {
	Phase  types.Phase
	Acc    types.AccountID
	Topics []types.Hash
}

type Event_MasterKeyLaunched struct {
	Phase  types.Phase
	Topics []types.Hash
}

type Event_WorkerAdded struct {
	Phase               types.Phase
	Pubkey              pattern.WorkerPublicKey
	AttestationProvider types.Option[types.U8]
	ConfidenceLevel     types.U8
	Topics              []types.Hash
}

type Event_KeyfairyAdded struct {
	Phase  types.Phase
	Pubkey pattern.WorkerPublicKey
	Topics []types.Hash
}

type Event_WorkerUpdated struct {
	Phase               types.Phase
	Pubkey              pattern.WorkerPublicKey
	AttestationProvider types.Option[types.U8]
	ConfidenceLevel     types.U8
	Topics              []types.Hash
}

type Event_MasterKeyRotated struct {
	Phase        types.Phase
	RotationId   types.U64
	MasterPubkey pattern.WorkerPublicKey
	Topics       []types.Hash
}

type Event_MasterKeyRotationFailed struct {
	Phase              types.Phase
	RotationLock       types.Option[types.U64]
	KeyfairyRotationId types.U64
	Topics             []types.Hash
}

type Event_MinimumCesealVersionChangedTo struct {
	Phase  types.Phase
	Elem1  types.U32
	Elem2  types.U32
	Elem3  types.U32
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

type Event_PhaseTransitioned struct {
	Phase  types.Phase
	From   Signed
	To     Unsigneds
	Round  types.U32
	Topics []types.Hash
}

type Signed struct {
	Index types.U8
	Value types.U32
}

type Unsigneds struct {
	Index         types.U8
	UnsignedValue []UnsignedValue
}

type UnsignedValue struct {
	Bool types.Bool
	Bn   types.U32
}

type Event_SolutionStored struct {
	Phase       types.Phase
	Compute     ElectionCompute
	Origin      types.Option[types.AccountID]
	PrevEjected types.Bool
	Topics      []types.Hash
}

type ElectionCompute struct {
	Index types.U8
	Value types.U8
}

type Event_Locked struct {
	Phase  types.Phase
	Who    types.AccountID
	Amount types.U128
	Topics []types.Hash
}

type Event_ServiceFeePaid struct {
	Phase       types.Phase
	Who         types.AccountID
	ActualFee   types.U128
	ExpectedFee types.U128
	Topics      []types.Hash
}

type Event_CallDone struct {
	Phase      types.Phase
	Who        types.AccountID
	CallResult Result
	Topics     []types.Hash
}

type Result struct {
	Index    types.U8
	ResultOk ResultOk
}

type ResultOk struct {
	ActualWeight types.Option[ActualWeightType]
	PaysFee      types.U8
}

type ActualWeightType struct {
	RefTime   types.U64
	ProofSize types.U64
}

type Event_TransactionFeePaid struct {
	Phase     types.Phase
	Who       types.AccountID
	ActualFee types.U128
	Tip       types.U128
	Topics    []types.Hash
}

type Event_ValidatorPrefsSet struct {
	Phase  types.Phase
	Stash  types.AccountID
	Prefs  ValidatorPrefs
	Topics []types.Hash
}

type ValidatorPrefs struct {
	Commission types.U32
	Blocked    types.Bool
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

type AllUploadDeclarationEvent struct {
	Operator string
	Owner    string
	Filehash string
}

type AllDeleteFileEvent struct {
	Operator string
	Owner    string
	Filehash string
}

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
	FileBank_IdleSpaceCert         []Event_IdleSpaceCert
	FileBank_ReplaceIdleSpace      []Event_ReplaceIdleSpace
	FileBank_CalculateReport       []Event_CalculateReport

	// OSS
	Oss_Authorize       []Event_Authorize
	Oss_CancelAuthorize []Event_CancelAuthorize
	Oss_OssRegister     []Event_OssRegister
	Oss_OssUpdate       []Event_OssUpdate
	Oss_OssDestroy      []Event_OssDestroy

	// SMINER
	Sminer_Registered               []Event_Registered
	Sminer_RegisterPoisKey          []Event_RegisterPoisKey
	Sminer_DrawFaucetMoney          []Event_DrawFaucetMoney
	Sminer_FaucetTopUpMoney         []Event_FaucetTopUpMoney
	Sminer_LessThan24Hours          []Event_LessThan24Hours
	Sminer_AlreadyFrozen            []Event_AlreadyFrozen
	Sminer_IncreaseCollateral       []Event_IncreaseCollateral
	Sminer_Deposit                  []Event_Deposit
	Sminer_UpdateBeneficiary        []Event_UpdateBeneficiary
	Sminer_UpdatePeerId             []Event_UpdatePeerId
	Sminer_Receive                  []Event_Receive
	Sminer_MinerExitPrep            []Event_MinerExitPrep
	Sminer_Withdraw                 []Event_Withdraw
	Sminer_IncreaseDeclarationSpace []Event_IncreaseDeclarationSpace

	// StorageHandler
	StorageHandler_BuySpace             []Event_BuySpace
	StorageHandler_ExpansionSpace       []Event_ExpansionSpace
	StorageHandler_RenewalSpace         []Event_RenewalSpace
	StorageHandler_LeaseExpired         []Event_LeaseExpired
	StorageHandler_LeaseExpireIn24Hours []Event_LeaseExpireIn24Hours

	// TeeWorker
	TeeWorker_Exit                          []Event_Exit
	TeeWorker_MasterKeyLaunched             []Event_MasterKeyLaunched
	TeeWorker_WorkerAdded                   []Event_WorkerAdded
	TeeWorker_KeyfairyAdded                 []Event_KeyfairyAdded
	TeeWorker_WorkerUpdated                 []Event_WorkerUpdated
	TeeWorker_MasterKeyRotated              []Event_MasterKeyRotated
	TeeWorker_MasterKeyRotationFailed       []Event_MasterKeyRotationFailed
	TeeWorker_MinimumCesealVersionChangedTo []Event_MinimumCesealVersionChangedTo

	// system - Staking
	Balances_Locked []Event_Locked

	// system - EvmAccountMapping
	EvmAccountMapping_ServiceFeePaid     []Event_ServiceFeePaid
	EvmAccountMapping_CallDone           []Event_CallDone
	EvmAccountMapping_TransactionFeePaid []Event_TransactionFeePaid

	// system - Staking
	Staking_ValidatorPrefsSet []Event_ValidatorPrefsSet

	// system - ElectionProviderMultiPhase
	ElectionProviderMultiPhase_ElectionFinalized []Event_ElectionFinalized
	ElectionProviderMultiPhase_PhaseTransitioned []Event_PhaseTransitioned
	ElectionProviderMultiPhase_SolutionStored    []Event_SolutionStored

	// system-gsrpc
	types.EventRecords
}

type BlockData struct {
	BlockHash        string
	PreHash          string
	ExtHash          string
	StHash           string
	AllGasFee        string
	Timestamp        int64
	BlockId          uint32
	IsNewEra         bool
	SysEvents        []string
	NewAccounts      []string
	MinerReg         []MinerRegInfo
	Extrinsics       []ExtrinsicsInfo
	TransferInfo     []TransferInfo
	UploadDecInfo    []UploadDecInfo
	DeleteFileInfo   []DeleteFileInfo
	CreateBucketInfo []CreateBucketInfo
	DeleteBucketInfo []DeleteBucketInfo
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
