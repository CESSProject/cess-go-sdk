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
)

// QueryTotalIssuance query the total amount of token issuance
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: the total amount of token issuance
//   - error: error message
func (c *ChainClient) QueryTotalIssuance(block int) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.U128

	if !c.GetChainState() {
		return "", ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, Balances, TotalIssuance)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Balances, TotalIssuance, err)
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Balances, TotalIssuance, err)
			c.SetChainState(false)
			return "", err
		}
		if !ok {
			return "", ERR_RPC_EMPTY_VALUE
		}
		return data.String(), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Balances, TotalIssuance, err)
		c.SetChainState(false)
		return "", err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Balances, TotalIssuance, err)
		c.SetChainState(false)
		return "", err
	}
	if !ok {
		return "0", nil
	}
	if data.String() == "" {
		return "0", nil
	}
	return data.String(), nil
}

// QueryInactiveIssuance query the amount of inactive token issuance
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: the amount of inactive token issuance
//   - error: error message
func (c *ChainClient) QueryInactiveIssuance(block int) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.U128

	if !c.GetChainState() {
		return "", ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, Balances, InactiveIssuance)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Balances, InactiveIssuance, err)
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Balances, InactiveIssuance, err)
			c.SetChainState(false)
			return "", err
		}
		if !ok {
			return "", ERR_RPC_EMPTY_VALUE
		}
		return data.String(), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Balances, InactiveIssuance, err)
		c.SetChainState(false)
		return "", err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Balances, InactiveIssuance, err)
		c.SetChainState(false)
		return "", err
	}
	if !ok {
		return "0", nil
	}
	if data.String() == "" {
		return "0", nil
	}
	return data.String(), nil
}
