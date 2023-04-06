package client

import "github.com/CESSProject/sdk-go/core/chain"

func (c *Cli) QueryStorageMiner(pubkey []byte) (chain.MinerInfo, error) {
	return c.Chain.QueryStorageMiner(pubkey)
}

func (c *Cli) QueryDeoss(pubkey []byte) (string, error) {
	return c.Chain.QueryDeoss(pubkey)
}

func (c *Cli) QueryFile(roothash string) (chain.FileMetaInfo, error) {
	return c.Chain.GetFileMetaInfo(roothash)
}

func (c *Cli) QueryBucket(owner []byte, bucketname string) (chain.BucketInfo, error) {
	return c.Chain.GetBucketInfo(owner, bucketname)
}
