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
	"github.com/AstaFrode/go-substrate-rpc-client/v4/types/codec"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/pkg/errors"
)

// QueryBlockNumber query the block number based on the block hash
//   - blockhash: hex-encoded block hash, if empty query the latest block number
//
// Return:
//   - uint32: block number
//   - error: error message
func (c *ChainClient) QueryBlockNumber(blockhash string) (uint32, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return 0, fmt.Errorf("rpc err: [%s] [st] [GetBlockLatest] %s", c.GetCurrentRpcAddr(), ERR_RPC_CONNECTION.Error())
	}

	if blockhash != "" {
		var h types.Hash
		err := codec.DecodeFromHex(blockhash, &h)
		if err != nil {
			return 0, err
		}
		block, err := c.api.RPC.Chain.GetBlock(h)
		if err != nil {
			return 0, errors.Wrap(err, "[GetBlock]")
		}
		return uint32(block.Block.Header.Number), nil
	}

	block, err := c.api.RPC.Chain.GetBlockLatest()
	if err != nil {
		return 0, errors.Wrap(err, "[GetBlockLatest]")
	}
	return uint32(block.Block.Header.Number), nil
}

// QueryAccountInfo query account info
//   - account: account
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - types.AccountInfo: account info
//   - error: error message
func (c *ChainClient) QueryAccountInfo(account string, block int32) (types.AccountInfo, error) {
	puk, err := utils.ParsingPublickey(account)
	if err != nil {
		return types.AccountInfo{}, err
	}
	return c.QueryAccountInfoByAccountID(puk, block)
}

// QueryAccountInfoByAccountID query account info
//   - accountID: account id
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - types.AccountInfo: account info
//   - error: error message
func (c *ChainClient) QueryAccountInfoByAccountID(accountID []byte, block int32) (types.AccountInfo, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return types.AccountInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), System, Account, ERR_RPC_CONNECTION.Error())
	}

	var data types.AccountInfo
	acc, err := types.NewAccountID(accountID)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	b, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(c.metadata, System, Account, b)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			return data, errors.Wrap(err, "[GetStorageLatest]")
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
		return data, errors.Wrap(err, "[GetStorage]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryAllAccountInfo query all account info
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []types.AccountInfo: all account info
//   - error: error message
func (c *ChainClient) QueryAllAccountInfo(block int32) ([]types.AccountInfo, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return []types.AccountInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), System, Account, ERR_RPC_CONNECTION.Error())
	}

	var result []types.AccountInfo
	key := CreatePrefixedKey(System, Account)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), System, Account, err)
		c.SetRpcState(false)
		return nil, err
	}

	var set []types.StorageChangeSet
	if block < 0 {
		set, err = c.api.RPC.State.QueryStorageAtLatest(keys)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), System, Account, err)
			c.SetRpcState(false)
			return nil, err
		}
	} else {
		blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
		if err != nil {
			return nil, err
		}
		set, err = c.api.RPC.State.QueryStorageAt(keys, blockhash)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), System, Account, err)
			c.SetRpcState(false)
			return nil, err
		}
	}
	for _, elem := range set {
		for _, change := range elem.Changes {
			var data types.AccountInfo
			if err := codec.Decode(change.StorageData, &data); err != nil {
				continue
			}
			result = append(result, data)
		}
	}
	return result, nil
}
