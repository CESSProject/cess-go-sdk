/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"fmt"

	"github.com/CESSProject/sdk-go/core/chain"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

func (c *Cli) SubmitIdleFile(size uint64, blockNum, blocksize, scansize uint32, pubkey []byte, hash string) (string, error) {
	acc, err := types.NewAccountID(pubkey)
	if err != nil {
		return "", err
	}
	if len(hash) != len(chain.FileHash{}) {
		return "", fmt.Errorf("Invalid file hash: %s", hash)
	}
	var filehash chain.FileHash

	for i := 0; i < len(hash); i++ {
		filehash[i] = types.U8(hash[i])
	}
	var idlefiles = []chain.IdleMetadata{
		{
			Size:      types.NewU64(size),
			BlockNum:  types.NewU32(blockNum),
			BlockSize: types.NewU32(blocksize),
			ScanSize:  types.NewU32(scansize),
			Acc:       *acc,
			Hash:      filehash,
		},
	}
	return c.Chain.SubmitIdleMetadata(idlefiles)
}
