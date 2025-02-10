/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"
	"log"

	"github.com/AstaFrode/go-substrate-rpc-client/v4/types"
	"github.com/CESSProject/cess-go-sdk/utils"
)

// QueryAuthorities query consensus rrsc public
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []ConsensusRrscAppPublic: all consensus rrsc public
//   - error: error message
func (c *ChainClient) QueryAuthorities(block int32) ([]ConsensusRrscAppPublic, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return []ConsensusRrscAppPublic{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Babe, Authorities, ERR_RPC_CONNECTION.Error())
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data []ConsensusRrscAppPublic

	key, err := types.CreateStorageKey(c.metadata, Babe, Authorities)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Babe, Authorities, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Babe, Authorities, err)
			c.SetRpcState(false)
			return data, err
		}
		if !ok {
			return data, ERR_RPC_EMPTY_VALUE
		}
		return data, nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return data, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Babe, Authorities, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}
