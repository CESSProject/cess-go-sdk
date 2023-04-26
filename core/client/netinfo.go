/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

func (c *Cli) QueryChallengeSt() (ChallengeSnapshot, error) {
	var challengeSnapshot ChallengeSnapshot
	chall, err := c.Chain.QueryChallengeSnapshot()
	if err != nil {
		return challengeSnapshot, err
	}
	challengeSnapshot.NetSnapshot.Start = uint32(chall.NetSnapshot.Start)
	challengeSnapshot.NetSnapshot.Total_idle_space = chall.NetSnapshot.Total_idle_space.String()
	challengeSnapshot.NetSnapshot.Total_reward = chall.NetSnapshot.Total_reward.String()
	challengeSnapshot.MinerSnapshot = make([]MinerSnapshot, len(chall.MinerSnapShot))
	for k, v := range chall.MinerSnapShot {
		challengeSnapshot.MinerSnapshot[k].Idle_space = v.Idle_space.String()
		challengeSnapshot.MinerSnapshot[k].Service_space = v.Service_space.String()
		challengeSnapshot.MinerSnapshot[k].Miner, err = utils.EncodePublicKeyAsCessAccount(v.Miner[:])
		if err != nil {
			return challengeSnapshot, err
		}
	}
	return challengeSnapshot, nil
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
