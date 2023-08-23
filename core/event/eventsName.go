/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package event

const (
	// AUDIT
	AuditVerifyProof               = "Audit.VerifyProof "
	AuditSubmitProof               = "Audit.SubmitProof"
	AuditGenerateChallenge         = "Audit.GenerateChallenge"
	AuditSubmitIdleProof           = "Audit.SubmitIdleProof"
	AuditSubmitServiceProof        = "Audit.SubmitServiceProof"
	AuditSubmitIdleVerifyResult    = "Audit.SubmitIdleVerifyResult"
	AuditSubmitServiceVerifyResult = "Audit.SubmitServiceVerifyResult"

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

	// OSS
	OssAuthorize       = "Oss.Authorize"
	OssCancelAuthorize = "Oss.CancelAuthorize"
	OssOssRegister     = "Oss.OssRegister"
	OssOssUpdate       = "Oss.OssUpdate"
	OssOssDestroy      = "Oss.OssDestroy"

	// SMINER
	SminerRegistered         = "Sminer.Registered"
	SminerDrawFaucetMoney    = "Sminer.DrawFaucetMoney"
	SminerFaucetTopUpMoney   = "Sminer.FaucetTopUpMoney"
	SminerIncreaseCollateral = "Sminer.IncreaseCollateral"
	SminerDeposit            = "Sminer.Deposit"
	SminerUpdataBeneficiary  = "Sminer.UpdataBeneficiary"
	SminerUpdataIp           = "Sminer.UpdataIp"
	SminerReceive            = "Sminer.Receive"
	SminerMinerExitPrep      = "Sminer.MinerExitPrep"
	SminerWithdraw           = "Sminer.Withdraw"

	// StorageHandler
	StorageHandlerBuySpace       = "StorageHandler.BuySpace"
	StorageHandlerExpansionSpace = "StorageHandler.ExpansionSpace"
	StorageHandlerRenewalSpace   = "StorageHandler.RenewalSpace"

	// TeeWorker
	TeeWorkerRegistrationTeeWorker = "TeeWorker.RegistrationTeeWorker"
	TeeWorkerUpdatePeerId          = "TeeWorker.UpdatePeerId"
	TeeWorkerExit                  = "TeeWorker.Exit"
)
