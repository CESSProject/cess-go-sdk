/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package event

const (
	// AUDIT
	AuditVerifyProof               = "Audit.VerifyProof"
	AuditSubmitProof               = "Audit.SubmitProof"
	AuditGenerateChallenge         = "Audit.GenerateChallenge"
	AuditSubmitIdleProof           = "Audit.SubmitIdleProof"
	AuditSubmitServiceProof        = "Audit.SubmitServiceProof"
	AuditSubmitIdleVerifyResult    = "Audit.SubmitIdleVerifyResult"
	AuditSubmitServiceVerifyResult = "Audit.SubmitServiceVerifyResult"

	// BALANCE
	BalanceTransfer = "Balances.Transfer"

	// FILEBANK
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

	// OSS
	OssAuthorize       = "Oss.Authorize"
	OssCancelAuthorize = "Oss.CancelAuthorize"
	OssOssRegister     = "Oss.OssRegister"
	OssOssUpdate       = "Oss.OssUpdate"
	OssOssDestroy      = "Oss.OssDestroy"

	// SMINER
	SminerRegistered               = "Sminer.Registered"
	SminerRegisterPoisKey          = "Sminer.RegisterPoisKey"
	SminerDrawFaucetMoney          = "Sminer.DrawFaucetMoney"
	SminerFaucetTopUpMoney         = "Sminer.FaucetTopUpMoney"
	SminerIncreaseCollateral       = "Sminer.IncreaseCollateral"
	SminerDeposit                  = "Sminer.Deposit"
	SminerUpdateBeneficiary        = "Sminer.UpdateBeneficiary"
	SminerUpdatePeerId             = "Sminer.UpdatePeerId"
	SminerReceive                  = "Sminer.Receive"
	SminerMinerExitPrep            = "Sminer.MinerExitPrep"
	SminerWithdraw                 = "Sminer.Withdraw"
	SminerIncreaseDeclarationSpace = "Sminer.IncreaseDeclarationSpace"

	// StorageHandler
	StorageHandlerBuySpace       = "StorageHandler.BuySpace"
	StorageHandlerExpansionSpace = "StorageHandler.ExpansionSpace"
	StorageHandlerRenewalSpace   = "StorageHandler.RenewalSpace"

	// TeeWorker
	TeeWorkerRegistrationTeeWorker         = "TeeWorker.RegistrationTeeWorker"
	TeeWorkerUpdatePeerId                  = "TeeWorker.UpdatePeerId"
	TeeWorkerExit                          = "TeeWorker.Exit"
	TeeWorkerMasterKeyLaunched             = "TeeWorker.MasterKeyLaunched"
	TeeWorkerKeyfairyAdded                 = "TeeWorker.KeyfairyAdded"
	TeeWorkerWorkerAdded                   = "TeeWorker.WorkerAdded"
	TeeWorkerWorkerUpdated                 = "TeeWorker.WorkerUpdated"
	TeeWorkerMasterKeyRotated              = "TeeWorker.MasterKeyRotated"
	TeeWorkerMasterKeyRotationFailed       = "TeeWorker.MasterKeyRotationFailed"
	TeeWorkerMinimumCesealVersionChangedTo = "TeeWorker.MinimumCesealVersionChangedTo"
)
