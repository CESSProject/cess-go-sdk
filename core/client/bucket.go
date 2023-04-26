/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"github.com/CESSProject/sdk-go/core/chain"
	"github.com/CESSProject/sdk-go/core/utils"
)

func (c *Cli) QueryAllBucketName(owner []byte) ([]string, error) {
	bucketlist, err := c.Chain.QueryBucketList(owner)
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

func (c *Cli) CheckBucketName(bucketname string) bool {
	return utils.CheckBucketName(bucketname)
}
