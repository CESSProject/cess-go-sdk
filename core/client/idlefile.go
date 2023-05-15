/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"github.com/CESSProject/sdk-go/core/chain"
	"github.com/CESSProject/sdk-go/core/rule"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type IdleFileMeta struct {
	Size      uint64
	BlockNum  uint32
	BlockSize uint32
	ScanSize  uint32
	minerAcc  []byte
	Hash      string
}

func (c *Cli) SubmitIdleFile(teeAcc []byte, idlefiles []IdleFileMeta) (string, error) {
	var submit = make([]chain.IdleMetadata, 0)
	for i := 0; i < len(idlefiles); i++ {
		var filehash chain.FileHash
		acc, err := types.NewAccountID(idlefiles[i].minerAcc)
		if err != nil {
			continue
		}

		if len(idlefiles[i].Hash) != len(chain.FileHash{}) {
			continue
		}

		for j := 0; j < len(idlefiles[i].Hash); j++ {
			filehash[j] = types.U8(idlefiles[i].Hash[j])
		}

		var ele = chain.IdleMetadata{
			Size:      types.NewU64(idlefiles[i].Size),
			BlockNum:  types.NewU32(idlefiles[i].BlockNum),
			BlockSize: types.NewU32(idlefiles[i].BlockSize),
			ScanSize:  types.NewU32(idlefiles[i].ScanSize),
			Acc:       *acc,
			Hash:      filehash,
		}
		submit = append(submit, ele)
		if len(submit) >= rule.MaxSubmitedIdleFileMeta {
			break
		}
	}
	return c.Chain.SubmitIdleMetadata(teeAcc, submit)
}
