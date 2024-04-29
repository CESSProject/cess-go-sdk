/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"
	"log"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
)

func (c *ChainClient) QueryRoundReward(era uint32, block int32) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.RoundRewardType

	if !c.GetChainState() {
		return "", pattern.ERR_RPC_CONNECTION
	}

	param, err := codec.Encode(era)
	if err != nil {
		return "", err
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.CessTreasury, pattern.RoundReward, param)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.CessTreasury, pattern.RoundReward, err)
		c.SetChainState(false)
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.CessTreasury, pattern.RoundReward, err)
			c.SetChainState(false)
			return "", err
		}
		if !ok {
			return "", pattern.ERR_RPC_EMPTY_VALUE
		}
		return data.TotalReward.String(), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), pattern.CessTreasury, pattern.RoundReward, err)
		c.SetChainState(false)
		return "", err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), pattern.CessTreasury, pattern.RoundReward, err)
		c.SetChainState(false)
		return "", err
	}
	if !ok {
		return "", pattern.ERR_RPC_EMPTY_VALUE
	}

	return data.TotalReward.String(), nil
}
