/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"log"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

func (c *Sdk) QueryChallengeSnapshot() (pattern.ChallengeSnapShot, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.ChallengeSnapShot

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.AUDIT, pattern.CHALLENGESNAPSHOT)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, pattern.ERR_RPC_EMPTY_VALUE
	}

	return data, nil
}

func (c *Sdk) QueryChallengeSt() (pattern.ChallengeSnapshot, error) {
	var challengeSnapshot pattern.ChallengeSnapshot
	chall, err := c.QueryChallengeSnapshot()
	if err != nil {
		return challengeSnapshot, err
	}
	challengeSnapshot.NetSnapshot.Start = uint32(chall.NetSnapshot.Start)
	challengeSnapshot.NetSnapshot.Life = uint32(chall.NetSnapshot.Life)
	challengeSnapshot.NetSnapshot.Total_idle_space = chall.NetSnapshot.TotalIdleSpace.String()
	challengeSnapshot.NetSnapshot.Total_reward = chall.NetSnapshot.TotalReward.String()
	challengeSnapshot.NetSnapshot.Random_index_list = make([]uint32, len(chall.NetSnapshot.RandomIndexList))
	for k, v := range chall.NetSnapshot.RandomIndexList {
		challengeSnapshot.NetSnapshot.Random_index_list[k] = uint32(v)
	}
	challengeSnapshot.NetSnapshot.Random = make([][]byte, len(chall.NetSnapshot.Random))
	for k, v := range chall.NetSnapshot.Random {
		challengeSnapshot.NetSnapshot.Random[k] = []byte(string(v[:]))
	}
	challengeSnapshot.MinerSnapshot = make([]pattern.MinerSnapshot, len(chall.MinerSnapShot))
	for k, v := range chall.MinerSnapShot {
		challengeSnapshot.MinerSnapshot[k].Idle_space = v.IdleSpace.String()
		challengeSnapshot.MinerSnapshot[k].Service_space = v.ServiceSpace.String()
		challengeSnapshot.MinerSnapshot[k].Miner, err = utils.EncodePublicKeyAsCessAccount(v.Miner[:])
		if err != nil {
			return challengeSnapshot, err
		}
	}
	return challengeSnapshot, nil
}

func (c *Sdk) QueryChallenge(pubkey []byte) (pattern.ChallengeInfo, error) {
	var chal pattern.ChallengeInfo
	acc, err := types.NewAccountID(pubkey)
	if err != nil {
		return chal, err
	}
	netinfo, err := c.QueryChallengeSnapshot()
	if err != nil {
		return chal, err
	}
	chal.RandomIndexList = make([]uint32, len(netinfo.NetSnapshot.RandomIndexList))
	chal.Random = make([][]byte, len(netinfo.NetSnapshot.Random))
	for _, v := range netinfo.MinerSnapShot {
		if v.Miner == *acc {
			for k, value := range netinfo.NetSnapshot.Random {
				chal.Random[k] = []byte(string(value[:]))
				chal.RandomIndexList[k] = uint32(netinfo.NetSnapshot.RandomIndexList[k])
			}
			chal.Start = uint32(netinfo.NetSnapshot.Start)
			break
		}
	}
	return chal, nil
}
