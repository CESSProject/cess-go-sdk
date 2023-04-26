/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"github.com/CESSProject/sdk-go/core/chain"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

func (c *Cli) ReportFiles(roothash []string) (string, []string, error) {
	var hashs = make([]chain.FileHash, len(roothash))
	for i := 0; i < len(roothash); i++ {
		for j := 0; j < len(roothash[i]); j++ {
			hashs[i][j] = types.U8(roothash[i][j])
		}
	}
	txhash, failed, err := c.Chain.SubmitFileReport(hashs)
	var failedfiles = make([]string, len(failed))
	for k, v := range failed {
		failedfiles[k] = string(v[:])
	}
	return txhash, failedfiles, err
}
