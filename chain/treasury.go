/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"
	"log"

	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
)

// QueryCurrencyReward query the currency rewards
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: currency rewards
//   - error: error message
func (c *ChainClient) QueryCurrencyReward(block int32) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	if !c.GetChainState() {
		return "", ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, CessTreasury, CurrencyReward)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), CessTreasury, CurrencyReward, err)
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), CessTreasury, CurrencyReward, err)
			c.SetChainState(false)
			return "", err
		}
		if !ok {
			return "0", ERR_RPC_EMPTY_VALUE
		}
		if data.String() == "" {
			return "0", nil
		}
		return data.String(), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), CessTreasury, CurrencyReward, err)
		c.SetChainState(false)
		return "", err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), CessTreasury, CurrencyReward, err)
		c.SetChainState(false)
		return "", err
	}
	if !ok {
		return "0", ERR_RPC_EMPTY_VALUE
	}

	if data.String() == "" {
		return "0", nil
	}
	return data.String(), nil
}

// QueryEraReward query the rewards in era
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: rewards in era
//   - error: error message
func (c *ChainClient) QueryEraReward(block int32) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	if !c.GetChainState() {
		return "", ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, CessTreasury, EraReward)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), CessTreasury, EraReward, err)
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), CessTreasury, EraReward, err)
			c.SetChainState(false)
			return "", err
		}
		if !ok {
			return "0", ERR_RPC_EMPTY_VALUE
		}
		if data.String() == "" {
			return "0", nil
		}
		return data.String(), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), CessTreasury, EraReward, err)
		c.SetChainState(false)
		return "", err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), CessTreasury, EraReward, err)
		c.SetChainState(false)
		return "", err
	}
	if !ok {
		return "0", ERR_RPC_EMPTY_VALUE
	}

	if data.String() == "" {
		return "0", nil
	}
	return data.String(), nil
}

// QueryReserveReward query the reserve rewards
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: reserve rewards
//   - error: error message
func (c *ChainClient) QueryReserveReward(block int32) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	if !c.GetChainState() {
		return "", ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, CessTreasury, ReserveReward)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), CessTreasury, ReserveReward, err)
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), CessTreasury, ReserveReward, err)
			c.SetChainState(false)
			return "", err
		}
		if !ok {
			return "0", ERR_RPC_EMPTY_VALUE
		}
		if data.String() == "" {
			return "0", nil
		}
		return data.String(), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), CessTreasury, ReserveReward, err)
		c.SetChainState(false)
		return "", err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), CessTreasury, ReserveReward, err)
		c.SetChainState(false)
		return "", err
	}
	if !ok {
		return "0", ERR_RPC_EMPTY_VALUE
	}

	if data.String() == "" {
		return "0", nil
	}
	return data.String(), nil
}

// QueryRoundReward querie the rewards in each era
//   - era: era id
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: rewards in an era
//   - error: error message
func (c *ChainClient) QueryRoundReward(era uint32, block int32) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data RoundRewardType

	if !c.GetChainState() {
		return "", ERR_RPC_CONNECTION
	}

	param, err := codec.Encode(era)
	if err != nil {
		return "", err
	}

	key, err := types.CreateStorageKey(c.metadata, CessTreasury, RoundReward, param)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), CessTreasury, RoundReward, err)
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), CessTreasury, RoundReward, err)
			c.SetChainState(false)
			return "", err
		}
		if !ok {
			return "0", ERR_RPC_EMPTY_VALUE
		}
		if data.TotalReward.String() == "" {
			return "0", nil
		}
		return data.TotalReward.String(), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), CessTreasury, RoundReward, err)
		c.SetChainState(false)
		return "", err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), CessTreasury, RoundReward, err)
		c.SetChainState(false)
		return "", err
	}
	if !ok {
		return "0", ERR_RPC_EMPTY_VALUE
	}

	if data.TotalReward.String() == "" {
		return "0", nil
	}
	return data.TotalReward.String(), nil
}
