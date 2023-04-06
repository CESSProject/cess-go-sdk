package client

import "github.com/CESSProject/sdk-go/core/chain"

func (c *Cli) CreateBucket(owner []byte, bucketname string) (string, error) {
	var err error
	var txhash string
	return txhash, err
}

func (c *Cli) QueryBuckets(owner []byte) ([]string, error) {
	bucketlist, err := c.Chain.GetBucketList(owner)
	if err != nil {
		if err.Error() != chain.ERR_Empty {
			return nil, err
		}
	}
	var buckets = make([]string, len(bucketlist))
	for i := 0; i < len(bucketlist); i++ {
		buckets[i] = string(bucketlist[i])
	}
	return buckets, nil
}

func (c *Cli) DeleteBucket(owner []byte, bucketName string) (string, error) {
	return c.Chain.DeleteBucket(owner, bucketName)
}
