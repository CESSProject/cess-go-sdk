/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

const (
	// Audit
	AuditVerifyProof               = "Audit.VerifyProof"
	AuditSubmitProof               = "Audit.SubmitProof"
	AuditGenerateChallenge         = "Audit.GenerateChallenge"
	AuditSubmitIdleProof           = "Audit.SubmitIdleProof"
	AuditSubmitServiceProof        = "Audit.SubmitServiceProof"
	AuditSubmitIdleVerifyResult    = "Audit.SubmitIdleVerifyResult"
	AuditSubmitServiceVerifyResult = "Audit.SubmitServiceVerifyResult"

	// Balances
	BalancesWithdraw = "Balances.Withdraw"
	BalancesTransfer = "Balances.Transfer"

	// FileBank
	FileBankDeleteFile            = "FileBank.DeleteFile"
	FileBankFillerDelete          = "FileBank.FillerDelete"
	FileBankFillerUpload          = "FileBank.FillerUpload"
	FileBankUploadDeclaration     = "FileBank.UploadDeclaration"
	FileBankCreateBucket          = "FileBank.CreateBucket"
	FileBankDeleteBucket          = "FileBank.DeleteBucket"
	FileBankTransferReport        = "FileBank.TransferReport"
	FileBankReplaceFiller         = "FileBank.ReplaceFiller"
	FileBankGenerateRestoralOrder = "FileBank.GenerateRestoralOrder"
	FileBankClaimRestoralOrder    = "FileBank.ClaimRestoralOrder"
	FileBankRecoveryCompleted     = "FileBank.RecoveryCompleted"
	FileBankStorageCompleted      = "FileBank.StorageCompleted"
	FileBankIdleSpaceCert         = "FileBank.IdleSpaceCert"
	FileBankReplaceIdleSpace      = "FileBank.ReplaceIdleSpace"
	FileBankCalculateReport       = "FileBank.CalculateReport"

	FileBankTerritorFileDelivery = "FileBank.TerritorFileDelivery"

	// Oss
	OssAuthorize       = "Oss.Authorize"
	OssCancelAuthorize = "Oss.CancelAuthorize"
	OssOssRegister     = "Oss.OssRegister"
	OssOssUpdate       = "Oss.OssUpdate"
	OssOssDestroy      = "Oss.OssDestroy"

	// Sminer
	SminerRegistered               = "Sminer.Registered"
	SminerRegisterPoisKey          = "Sminer.RegisterPoisKey"
	SminerDrawFaucetMoney          = "Sminer.DrawFaucetMoney"
	SminerFaucetTopUpMoney         = "Sminer.FaucetTopUpMoney"
	SminerIncreaseCollateral       = "Sminer.IncreaseCollateral"
	SminerDeposit                  = "Sminer.Deposit"
	SminerUpdateBeneficiary        = "Sminer.UpdateBeneficiary"
	SminerUpdateEndpoint           = "Sminer.UpdateEndPoint"
	SminerReceive                  = "Sminer.Receive"
	SminerMinerExitPrep            = "Sminer.MinerExitPrep"
	SminerWithdraw                 = "Sminer.Withdraw"
	SminerIncreaseDeclarationSpace = "Sminer.IncreaseDeclarationSpace"

	// Staking
	StakingStakersElected = "Staking.StakersElected"
	StakingEraPaid        = "Staking.EraPaid"
	StakingPayoutStarted  = "Staking.PayoutStarted"
	StakingRewarded       = "Staking.Rewarded"
	StakingUnbonded       = "Staking.Unbonded"

	// StorageHandler
	StorageHandlerMintTerritory        = "StorageHandler.MintTerritory"
	StorageHandlerExpansionTerritory   = "StorageHandler.ExpansionTerritory"
	StorageHandlerRenewalTerritory     = "StorageHandler.RenewalTerritory"
	StorageHandlerReactivateTerritory  = "StorageHandler.ReactivateTerritory"
	StorageHandlerConsignment          = "StorageHandler.Consignment"
	StorageHandlerCancleConsignment    = "StorageHandler.CancleConsignment"
	StorageHandlerBuyConsignment       = "StorageHandler.BuyConsignment"
	StorageHandlerCancelPurchaseAction = "StorageHandler.CancelPurchaseAction"

	// TeeWorker
	TeeWorkerExit                          = "TeeWorker.Exit"
	TeeWorkerMasterKeyLaunched             = "TeeWorker.MasterKeyLaunched"
	TeeWorkerKeyfairyAdded                 = "TeeWorker.KeyfairyAdded"
	TeeWorkerWorkerAdded                   = "TeeWorker.WorkerAdded"
	TeeWorkerWorkerUpdated                 = "TeeWorker.WorkerUpdated"
	TeeWorkerMasterKeyRotated              = "TeeWorker.MasterKeyRotated"
	TeeWorkerMasterKeyRotationFailed       = "TeeWorker.MasterKeyRotationFailed"
	TeeWorkerMinimumCesealVersionChangedTo = "TeeWorker.MinimumCesealVersionChangedTo"

	//
	TransactionPaymentTransactionFeePaid = "TransactionPayment.TransactionFeePaid"
	//
	EvmAccountMappingTransactionFeePaid = "EvmAccountMapping.TransactionFeePaid"

	// System
	SystemExtrinsicSuccess = "System.ExtrinsicSuccess"
	SystemExtrinsicFailed  = "System.ExtrinsicFailed"
	SystemNewAccount       = "System.NewAccount"
)
