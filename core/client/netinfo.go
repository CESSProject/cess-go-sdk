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

func (c *Cli) QueryNetSnapShot() (chain.NetSnapShot, error) {
	return c.Chain.QueryNetSnapShot()
}

func (c *Cli) QueryChallenge(pubkey []byte) (ChallengeInfo, error) {
	var chal ChallengeInfo
	acc, err := types.NewAccountID(pubkey)
	if err != nil {
		return chal, err
	}
	netinfo, err := c.Chain.QueryNetSnapShot()
	if err != nil {
		return chal, err
	}
	chal.Random = make([]byte, len(netinfo.NetSnapShot.Random))
	for _, v := range netinfo.MinerSnapShot {
		if v.Miner == *acc {
			for i := 0; i < len(netinfo.NetSnapShot.Random); i++ {
				chal.Random[i] = byte(netinfo.NetSnapShot.Random[i])
			}
			chal.Start = uint32(netinfo.NetSnapShot.Start)
			break
		}
	}
	return chal, nil
}
