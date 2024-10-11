/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// ******************************************************
// cess event type
// ******************************************************

// ------------------------Audit-------------------
type Event_VerifyProof struct {
	Phase     types.Phase
	TeeWorker WorkerPublicKey
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
	Tee    WorkerPublicKey
	Miner  types.AccountID
	Result types.Bool
	Topics []types.Hash
}

type Event_SubmitServiceVerifyResult struct {
	Phase  types.Phase
	Tee    WorkerPublicKey
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

type Event_Receive struct {
	Phase  types.Phase
	Acc    string
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
	Filehash []FileHash
	Topics   []types.Hash
}

type Event_FillerDelete struct {
	Phase      types.Phase
	Acc        types.AccountID
	FillerHash FileHash
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
	Deal_hash FileHash
	Topics    []types.Hash
}

type Event_CreateBucket struct {
	Phase      types.Phase
	Acc        types.AccountID
	Owner      types.AccountID
	BucketName types.Bytes
	Topics     []types.Hash
}

type Event_TerritorFileDelivery struct {
	Phase        types.Phase
	Filehash     FileHash
	NewTerritory types.Bytes
	Topics       []types.Hash
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
	DealHash FileHash
	Topics   []types.Hash
}

type Event_ReplaceFiller struct {
	Phase       types.Phase
	Acc         types.AccountID
	Filler_list []FileHash
	Topics      []types.Hash
}

type Event_CalculateEnd struct {
	Phase     types.Phase
	File_hash FileHash
	Topics    []types.Hash
}

type Event_GenerateRestoralOrder struct {
	Phase        types.Phase
	Miner        types.AccountID
	FragmentHash FileHash
	Topics       []types.Hash
}

type Event_ClaimRestoralOrder struct {
	Phase   types.Phase
	Miner   types.AccountID
	OrderId FileHash
	Topics  []types.Hash
}

type Event_RecoveryCompleted struct {
	Phase   types.Phase
	Miner   types.AccountID
	OrderId FileHash
	Topics  []types.Hash
}

type Event_StorageCompleted struct {
	Phase    types.Phase
	FileHash FileHash
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
	FileHash FileHash
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
	Pubkey              WorkerPublicKey
	AttestationProvider types.Option[types.U8]
	ConfidenceLevel     types.U8
	Topics              []types.Hash
}

type Event_KeyfairyAdded struct {
	Phase  types.Phase
	Pubkey WorkerPublicKey
	Topics []types.Hash
}

type Event_WorkerUpdated struct {
	Phase               types.Phase
	Pubkey              WorkerPublicKey
	AttestationProvider types.Option[types.U8]
	ConfidenceLevel     types.U8
	Topics              []types.Hash
}

type Event_MasterKeyRotated struct {
	Phase        types.Phase
	RotationId   types.U64
	MasterPubkey WorkerPublicKey
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
	Endpoint types.Bytes
	Topics   []types.Hash
}

type Event_OssUpdate struct {
	Phase       types.Phase
	Acc         types.AccountID
	NewEndpoint types.Bytes
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
