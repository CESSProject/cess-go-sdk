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

func (c *Cli) Workspace() string {
	return c.Node.Workspace()
}

func (c *Cli) QueryStorageMiner(pubkey []byte) (chain.MinerInfo, error) {
	return c.Chain.QueryStorageMiner(pubkey)
}

func (c *Cli) QueryDeoss(pubkey []byte) (string, error) {
	return c.Chain.QueryDeoss(pubkey)
}

func (c *Cli) QueryFile(roothash string) (chain.FileMetadata, error) {
	return c.Chain.QueryFileMetadata(roothash)
}

func (c *Cli) QueryBucket(owner []byte, bucketname string) (chain.BucketInfo, error) {
	return c.Chain.QueryBucketInfo(owner, bucketname)
}

func (c *Cli) QueryGrantor(puk []byte) (bool, error) {
	grantor, err := c.Chain.QuaryAuthorizedAcc(puk)
	if err != nil {
		if err.Error() == chain.ERR_Empty {
			return false, nil
		}
		return false, err
	}
	account_chain, _ := utils.EncodePublicKeyAsCessAccount(grantor[:])
	account_local, _ := c.GetCessAccount()

	return account_chain == account_local, nil
}

func (c *Cli) QueryStorageOrder(roothash string) (chain.StorageOrder, error) {
	return c.Chain.GetStorageOrder(roothash)
}

func (c *Cli) QueryReplacements(pubkey []byte) (uint32, error) {
	num, err := c.Chain.QueryPendingReplacements(pubkey)
	if err != nil {
		return 0, err
	}
	return uint32(num), nil
}

func (c *Cli) QueryUserSpaceInfo(pubkey []byte) (chain.UserSpaceInfo, error) {
	return c.Chain.QueryUserSpaceInfo(pubkey)
}

func (c *Cli) QuerySpacePricePerGib() (string, error) {
	return c.Chain.QuerySpacePricePerGib()
}

func (c *Cli) QueryTeePodr2Puk() ([]byte, error) {
	return c.Chain.QueryTeePodr2Puk()
}

func (c *Cli) QueryTeeWorkerList() ([]chain.TeeWorkerInfo, error) {
	return c.Chain.QueryTeeWorkerList()
}

func (c *Cli) QueryTeeWorkerPeerID(pubkey []byte) ([]byte, error) {
	return c.Chain.QueryTeeWorker(pubkey)
}
