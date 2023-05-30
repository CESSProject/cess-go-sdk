package sdk

import (
	"math/big"

	"github.com/CESSProject/sdk-go/core/pattern"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type SDK interface {
	// QueryBlockHeight queries the block height corresponding to the block hash,
	// If the blockhash is empty, query the latest block height.
	QueryBlockHeight(blockhash string) (uint32, error)

	// QueryNodeSynchronizationSt returns the synchronization status of the current node.
	QueryNodeSynchronizationSt() (bool, error)

	// QueryNodeConnectionSt queries the connection status of the node.
	QueryNodeConnectionSt() bool

	// QueryDeoss queries deoss information.
	QueryDeoss(pubkey []byte) ([]byte, error)

	// QuaryAuthorizedAcc queries the account authorized by puk.
	QuaryAuthorizedAcc(puk []byte) (types.AccountID, error)
	QuaryAuthorizedAccount(puk []byte) (string, error)
	CheckSpaceUsageAuthorization(puk []byte) (bool, error)

	// QueryBucketInfo queries the information of the "bucketname" bucket of the puk.
	QueryBucketInfo(puk []byte, bucketname string) (pattern.BucketInfo, error)

	// QueryBucketList queries all buckets of the puk.
	// QueryAllBucketName
	QueryBucketList(puk []byte) ([]types.Bytes, error)
	QueryAllBucketName(puk []byte) ([]string, error)

	// QueryFileMetaInfo queries the metadata of the roothash file.
	QueryFileMetadata(roothash string) (pattern.FileMetadata, error)

	// QueryStorageMiner queries storage node information.
	QueryStorageMiner(puk []byte) (pattern.MinerInfo, error)

	// QuerySminerList queries the accounts of all storage miners.
	QuerySminerList() ([]types.AccountID, error)

	// QueryAccountInfo query account information.
	QueryAccountInfo(puk []byte) (types.AccountInfo, error)

	// QueryStorageOrder query storage order information.
	QueryStorageOrder(roothash string) (pattern.StorageOrder, error)

	// QueryPendingReplacements queries the amount of idle data that can be replaced.
	QueryPendingReplacements(puk []byte) (uint32, error)

	// QueryUserSpaceInfo queries the space information purchased by the user.
	QueryUserSpaceInfo(puk []byte) (pattern.UserSpaceInfo, error)
	QueryUserSpaceSt(puk []byte) (pattern.UserSpaceSt, error)

	// QuerySpacePricePerGib query space price per GiB.
	QuerySpacePricePerGib() (string, error)

	// QueryChallengeSnapshot query challenge information snapshot.
	QueryChallengeSnapshot() (pattern.ChallengeSnapShot, error)

	// QueryTeePodr2Puk queries the public key of the TEE.
	QueryTeePodr2Puk() ([]byte, error)

	// QueryTeePeerID queries the peerid of the Tee worker.
	QueryTeePeerID(puk []byte) ([]byte, error)

	// QueryTeeInfoList queries the information of all tee workers.
	QueryTeeInfoList() ([]pattern.TeeWorkerInfo, error)
	QueryTeeWorkerList() ([]pattern.TeeWorkerSt, error)

	//
	QueryAssignedProof() ([][]pattern.ProofAssignmentInfo, error)

	//
	QueryTeeAssignedProof(puk []byte) ([]pattern.ProofAssignmentInfo, error)

	//
	QueryMinerRewards(puk []byte) (pattern.MinerReward, error)
	QuaryRewards(puk []byte) (pattern.RewardsType, error)

	// Register is used to register OSS or BUCKET roles.
	Register(role string, puk []byte, earnings string, pledge uint64) (string, string, error)

	// UpdateAddress updates the address of oss or sminer.
	UpdateAddress(role, multiaddr string) (string, error)

	// UpdateIncomeAcc update income account.
	UpdateIncomeAcc(puk []byte) (string, error)
	UpdateIncomeAccount(income string) (string, error)

	// CreateBucket creates a bucket for puk.
	CreateBucket(puk []byte, bucketname string) (string, error)

	// DeleteBucket deletes buckets for puk.
	DeleteBucket(puk []byte, bucketname string) (string, error)

	// DeleteFile deletes files for puk.
	DeleteFile(puk []byte, roothash []string) (string, []pattern.FileHash, error)

	// UploadDeclaration creates a storage order.
	UploadDeclaration(roothash string, dealinfo []pattern.SegmentList, user pattern.UserBrief) (string, error)

	// SubmitIdleMetadata Submit idle file metadata.
	SubmitIdleMetadata(teeAcc []byte, idlefiles []pattern.IdleMetadata) (string, error)
	SubmitIdleFile(teeAcc []byte, idlefiles []pattern.IdleFileMeta) (string, error)

	// SubmitFileReport submits a stored file report.
	SubmitFileReport(roothash []pattern.FileHash) (string, []pattern.FileHash, error)
	ReportFiles(roothash []string) (string, []string, error)

	// ReplaceIdleFiles replaces idle files.
	ReplaceIdleFiles(roothash []pattern.FileHash) (string, []pattern.FileHash, error)
	ReplaceFile(roothash []string) (string, []string, error)

	// IncreaseStakes increase stakes.
	IncreaseStakes(tokens *big.Int) (string, error)
	IncreaseSminerStakes(token string) (string, error)

	// Exit exit the cess network.
	Exit(role string) (string, error)

	// ClaimRewards is used to claim rewards
	ClaimRewards() (string, error)

	// Withdraw is used to withdraw staking
	Withdraw() (string, error)

	//
	ReportProof(idlesigma, servicesigma string) (string, error)

	// ExtractAccountPuk extracts the public key of the account,
	// and returns its own public key if the account is empty.
	ExtractAccountPuk(account string) ([]byte, error)

	// GetSignatureAcc returns the signature account.
	GetSignatureAcc() string

	// GetSignatureURI to get the private key of the signing account
	GetSignatureURI() string

	// GetSubstrateAPI returns Substrate API
	GetSubstrateAPI() *gsrpc.SubstrateAPI

	// GetChainState returns chain node state
	GetChainState() bool

	//
	SetChainState(state bool)

	//
	Reconnect() error

	//
	GetMetadata() *types.Metadata

	//
	GetKeyEvents() types.StorageKey

	//
	SysProperties() (pattern.SysProperties, error)

	//
	SyncState() (pattern.SysSyncState, error)

	//
	SysVersion() (string, error)

	//
	NetListening() (bool, error)

	//
	Sign(msg []byte) ([]byte, error)
}
