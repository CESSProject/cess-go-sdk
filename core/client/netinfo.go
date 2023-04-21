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

type ChallengeInfo struct {
	Random []byte
	Start  uint32
}

func (c *Cli) QueryNetSnapShot() (chain.ChallengeSnapShot, error) {
	return c.Chain.QueryChallengeSnapshot()
}

func (c *Cli) QueryChallenge(pubkey []byte) (ChallengeInfo, error) {
	var chal ChallengeInfo
	acc, err := types.NewAccountID(pubkey)
	if err != nil {
		return chal, err
	}
	netinfo, err := c.Chain.QueryChallengeSnapshot()
	if err != nil {
		return chal, err
	}
	chal.Random = make([]byte, len(netinfo.NetSnapshot.Random))
	for _, v := range netinfo.MinerSnapShot {
		if v.Miner == *acc {
			for i := 0; i < len(netinfo.NetSnapshot.Random); i++ {
				chal.Random[i] = byte(netinfo.NetSnapshot.Random[i])
			}
			chal.Start = uint32(netinfo.NetSnapshot.Start)
			break
		}
	}
	return chal, nil
}
