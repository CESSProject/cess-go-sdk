/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/AstaFrode/go-substrate-rpc-client/v4/types"
	"github.com/AstaFrode/go-substrate-rpc-client/v4/types/codec"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/pkg/errors"
)

// QueryExpenders query expenders (idle data specification)
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - ExpendersInfo: idle data specification
//   - error: error message
func (c *ChainClient) QueryExpenders(block int32) (ExpendersInfo, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return ExpendersInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Sminer, Expenders, ERR_RPC_CONNECTION.Error())
		}
	}
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data ExpendersInfo

	key, err := types.CreateStorageKey(c.metadata, Sminer, Expenders)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Sminer, Expenders, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Sminer, Expenders, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Sminer, Expenders, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryMinerItems query storage miner info
//   - accountID: storage miner account
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - MinerInfo: storage miner info
//   - error: error message
func (c *ChainClient) QueryMinerItems(accountID []byte, block int32) (MinerInfo, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return MinerInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Sminer, MinerItems, ERR_RPC_CONNECTION.Error())
		}
	}
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data MinerInfo

	key, err := types.CreateStorageKey(c.metadata, Sminer, MinerItems, accountID)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Sminer, MinerItems, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Sminer, MinerItems, err)
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

	meta, err := c.api.RPC.State.GetMetadata(blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetMetadata: %v", c.GetCurrentRpcAddr(), Sminer, MinerItems, err)
		return data, err
	}
	key, err = types.CreateStorageKey(meta, Sminer, MinerItems, accountID)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Sminer, MinerItems, err)
		return data, err
	}
	time.Sleep(time.Second)
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Sminer, MinerItems, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryMinerItems query storage miner info
//   - accountID: storage miner account
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - MinerInfo: storage miner info
//   - error: error message
func (c *ChainClient) QueryMinerItemsV1(accountID []byte, block int32) (MinerInfoV1, error) {
	if !c.GetRpcState() {
		err := c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Sminer, MinerItems, ERR_RPC_CONNECTION.Error())
			return MinerInfoV1{}, err
		}
	}

	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data MinerInfoV1

	key, err := types.CreateStorageKey(c.metadata, Sminer, MinerItems, accountID)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Sminer, MinerItems, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Sminer, MinerItems, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Sminer, MinerItems, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryStakingStartBlock query storage miner's starting staking block
//   - accountID: storage miner account
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint32: starting staking block
//   - error: error message
func (c *ChainClient) QueryStakingStartBlock(accountID []byte, block int32) (uint32, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Sminer, StakingStartBlock, ERR_RPC_CONNECTION.Error())
		}
	}
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U32

	key, err := types.CreateStorageKey(c.metadata, Sminer, StakingStartBlock, accountID)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Sminer, StakingStartBlock, err)
		return 0, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Sminer, StakingStartBlock, err)
			c.SetRpcState(false)
			return 0, err
		}
		if !ok {
			return 0, ERR_RPC_EMPTY_VALUE
		}
		return uint32(data), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return 0, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Sminer, StakingStartBlock, err)
		c.SetRpcState(false)
		return 0, err
	}
	if !ok {
		return 0, ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

// QueryAllMiner query all storage miner accounts
//   - accountID: storage miner account
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []types.AccountID: all storage miner accounts
//   - error: error message
func (c *ChainClient) QueryAllMiner(block int32) ([]types.AccountID, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return []types.AccountID{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Sminer, AllMiner, ERR_RPC_CONNECTION.Error())
		}
	}
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data []types.AccountID

	key, err := types.CreateStorageKey(c.metadata, Sminer, AllMiner)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Sminer, AllMiner, err)
		return nil, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Sminer, AllMiner, err)
			c.SetRpcState(false)
			return nil, err
		}
		if !ok {
			return nil, ERR_RPC_EMPTY_VALUE
		}
		return data, nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return data, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Sminer, AllMiner, err)
		c.SetRpcState(false)
		return nil, err
	}
	if !ok {
		return nil, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryCounterForMinerItems query all storage miner count
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint32: all storage miner count
//   - error: error message
func (c *ChainClient) QueryCounterForMinerItems(block int32) (uint32, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Sminer, CounterForMinerItems, ERR_RPC_CONNECTION.Error())
		}
	}
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U32

	key, err := types.CreateStorageKey(c.metadata, Sminer, CounterForMinerItems)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Sminer, CounterForMinerItems, err)
		return 0, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Sminer, CounterForMinerItems, err)
			c.SetRpcState(false)
			return 0, err
		}
		if !ok {
			return 0, ERR_RPC_EMPTY_VALUE
		}
		return uint32(data), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return 0, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Sminer, CounterForMinerItems, err)
		c.SetRpcState(false)
		return 0, err
	}
	if !ok {
		return 0, ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

// QueryRewardMap query all reward information for storage miner
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - MinerReward: all reward information
//   - error: error message
func (c *ChainClient) QueryRewardMap(accountID []byte, block int32) (MinerReward, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return MinerReward{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Sminer, RewardMap, ERR_RPC_CONNECTION.Error())
		}
	}
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data MinerReward

	key, err := types.CreateStorageKey(c.metadata, Sminer, RewardMap, accountID)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Sminer, RewardMap, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Sminer, RewardMap, err)
			c.SetRpcState(false)
			return data, err
		}
		if !ok {
			return data, ERR_RPC_EMPTY_VALUE
		}
		if data.OrderList == nil {
			if data.RewardIssued.Int64() == 0 && data.TotalReward.Int64() == 0 {
				return data, ERR_RPC_EMPTY_VALUE
			}
		}
		return data, nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return data, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Sminer, RewardMap, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	if data.OrderList == nil {
		if data.RewardIssued.Int64() == 0 && data.TotalReward.Int64() == 0 {
			return data, ERR_RPC_EMPTY_VALUE
		}
	}
	return data, nil
}

// QueryRestoralTarget query the data recovery information of exited storage miner
//   - accountID: storage miner account
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - RestoralTargetInfo: the data recovery information
//   - error: error message
func (c *ChainClient) QueryRestoralTarget(accountID []byte, block int32) (RestoralTargetInfo, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return RestoralTargetInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Sminer, RestoralTarget, ERR_RPC_CONNECTION.Error())
		}
	}
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data RestoralTargetInfo

	acc, err := types.NewAccountID(accountID)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	account, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(c.metadata, Sminer, RestoralTarget, account)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Sminer, RestoralTarget, err)
		return data, err
	}
	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Sminer, RestoralTarget, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Sminer, RestoralTarget, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryAllRestoralTarget query the data recovery information of all exited storage miner
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []RestoralTargetInfo: all the data recovery information
//   - error: error message
func (c *ChainClient) QueryAllRestoralTarget(block int32) ([]RestoralTargetInfo, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return []RestoralTargetInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Sminer, RestoralTarget, ERR_RPC_CONNECTION.Error())
		}
	}
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var result []RestoralTargetInfo

	key := CreatePrefixedKey(Sminer, RestoralTarget)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), Sminer, RestoralTarget, err)
		c.SetRpcState(false)
		return nil, err
	}
	var set []types.StorageChangeSet
	if block < 0 {
		set, err = c.api.RPC.State.QueryStorageAtLatest(keys)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), Sminer, RestoralTarget, err)
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
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAt: %v", c.GetCurrentRpcAddr(), Sminer, RestoralTarget, err)
			c.SetRpcState(false)
			return nil, err
		}
	}
	for _, elem := range set {
		for _, change := range elem.Changes {
			var data RestoralTargetInfo
			if err := codec.Decode(change.StorageData, &data); err != nil {
				continue
			}
			result = append(result, data)
		}
	}
	return result, nil
}

// QueryPendingReplacements query the size of the storage miner's replaceable idle data
//   - accountID: storage miner account
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - types.U128: the size of replaceable idle data
//   - error: error message
func (c *ChainClient) QueryPendingReplacements(accountID []byte, block int32) (types.U128, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return types.U128{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Sminer, PendingReplacements, ERR_RPC_CONNECTION.Error())
		}
	}
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	acc, err := types.NewAccountID(accountID)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	account, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(c.metadata, Sminer, PendingReplacements, account)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Sminer, PendingReplacements, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Sminer, PendingReplacements, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Sminer, PendingReplacements, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryCompleteSnapShot query the number of storage miners and storage miner power in each era
//   - era: era id
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint32: the number of storage miners in current era
//   - uint64: all storage miners power in current era
//   - error: error message
func (c *ChainClient) QueryCompleteSnapShot(era uint32, block int32) (uint32, uint64, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return 0, 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Sminer, CompleteSnapShot, ERR_RPC_CONNECTION.Error())
		}
	}
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data CompleteSnapShotType

	param, err := codec.Encode(era)
	if err != nil {
		return 0, 0, err
	}

	key, err := types.CreateStorageKey(c.metadata, Sminer, CompleteSnapShot, param)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Sminer, CompleteSnapShot, err)
		return 0, 0, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Sminer, CompleteSnapShot, err)
			c.SetRpcState(false)
			return 0, 0, err
		}
		if !ok {
			return 0, 0, ERR_RPC_EMPTY_VALUE
		}
		return uint32(data.MinerCount), data.TotalPower.Uint64(), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Sminer, CompleteSnapShot, err)
		return 0, 0, err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Sminer, CompleteSnapShot, err)
		c.SetRpcState(false)
		return 0, 0, err
	}
	if !ok {
		return 0, 0, ERR_RPC_EMPTY_VALUE
	}

	return uint32(data.MinerCount), data.TotalPower.Uint64(), nil
}

// QueryCompleteMinerSnapShot query the completed challenge snapshots of miners
//   - puk: account id
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []MinerCompleteInfo: list of completed challenge snapshots
//   - error: error message
func (c *ChainClient) QueryCompleteMinerSnapShot(puk []byte, block int32) ([]MinerCompleteInfo, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return nil, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Sminer, CompleteMinerSnapShot, ERR_RPC_CONNECTION.Error())
		}
	}
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data []MinerCompleteInfo

	key, err := types.CreateStorageKey(c.metadata, Sminer, CompleteMinerSnapShot, puk)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Sminer, CompleteMinerSnapShot, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Sminer, CompleteMinerSnapShot, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Sminer, CompleteMinerSnapShot, err)
		return data, err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Sminer, CompleteMinerSnapShot, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}

	return data, nil
}

// IncreaseCollateral increases the number of staking for storage miner
//   - accountID: storage miner account
//   - token: number of staking
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - The number of staking to be added is calculated in the smallest unit,
//     if you want to add 1CESS staking, you need to fill in "1000000000000000000"
func (c *ChainClient) IncreaseCollateral(accountID []byte, token string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	tokens, ok := new(big.Int).SetString(token, 10)
	if !ok {
		return "", fmt.Errorf("[IncreaseCollateral] invalid token: %s", token)
	}

	acc, err := types.NewAccountID(accountID)
	if err != nil {
		return "", errors.Wrap(err, "[NewAccountID]")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_Sminer_increase_collateral, *acc, types.NewUCompact(tokens))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_increase_collateral, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Sminer_increase_collateral)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_increase_collateral, err)
	}

	return blockhash, nil
}

// IncreaseDeclarationSpace increases the size of space declared on the chain
//   - tibCount: the size of the declaration space increased, in TiB
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - the size of the declared space cannot be reduced
//   - when the staking does not meet the declared space size, you will be frozen
func (c *ChainClient) IncreaseDeclarationSpace(tibCount uint32) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_Sminer_miner_exit, types.NewU32(tibCount))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_miner_exit, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Sminer_miner_exit)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_miner_exit, err)
	}

	return blockhash, nil
}

// MinerExitPrep pre-exit storage miner
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - after pre-exit, you need to wait for one day before it will automatically exit
//   - cannot register as a storage miner again after pre-exit
func (c *ChainClient) MinerExitPrep() (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	acc, err := types.NewAccountID(c.GetSignatureAccPulickey())
	if err != nil {
		return "", errors.Wrap(err, "[NewAccountID]")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_Sminer_miner_exit, *acc)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_miner_exit, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Sminer_miner_exit)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_miner_exit, err)
	}

	return blockhash, nil
}

// MinerWithdraw withdraws all staking
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - must be an exited miner to withdraw
//   - wait a day to withdraw after pre-exit
func (c *ChainClient) MinerWithdraw() (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_Sminer_miner_withdraw)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_miner_withdraw, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Sminer_miner_withdraw)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_miner_withdraw, err)
	}

	return blockhash, nil
}

// ReceiveReward to receive rewards
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - for storage miner only
//   - pass at least one idle and service challenge at the same time to get the reward
func (c *ChainClient) ReceiveReward() (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	call, err := types.NewCall(c.metadata, ExtName_Sminer_receive_reward)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_receive_reward, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_receive_reward, err)
		return blockhash, err
	}

	if !c.GetRpcState() {
		err = c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] %s", c.GetCurrentRpcAddr(), ExtName_Sminer_receive_reward, ERR_RPC_CONNECTION.Error())
			return blockhash, err
		}
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_receive_reward, err)
		c.SetRpcState(false)
		return blockhash, err
	}
	if !ok {
		return blockhash, ERR_RPC_EMPTY_VALUE
	}

	o := types.SignatureOptions{
		BlockHash:          c.genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        c.genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(accountInfo.Nonce)),
		SpecVersion:        c.runtimeVersion.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: c.runtimeVersion.TransactionVersion,
	}

	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_receive_reward, err)
		return blockhash, err
	}

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_receive_reward, err)
		c.SetRpcState(false)
		return blockhash, err
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_Sminer_receive_reward, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}

// RegisterPoisKey register pois key, storage miner registration
// requires two stages, this is the second one.
//
//   - poisKey: pois key
//   - teeSignWithAcc: tee's sign with account
//   - teeSign: tee's sign
//   - teePuk: tee's work public key
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - storage miners must complete the first stage to register for the second stage
func (c *ChainClient) RegisterPoisKey(poisKey PoISKeyInfo, teeSignWithAcc, teeSign types.Bytes, teePuk WorkerPublicKey) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_Sminer_register_pois_key, teeSignWithAcc, teeSign, teePuk)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_register_pois_key, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Sminer_register_pois_key)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_register_pois_key, err)
	}

	return blockhash, nil
}

// RegnstkSminer registers as a storage miner,
// which is the first stage of storage miner registration.
//
//   - earnings: earnings account
//   - endpoint: communications endpoint
//   - staking: number of staking, the unit is CESS
//   - tibCount: the size of declaration space, in TiB
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) RegnstkSminer(earnings string, endpoint []byte, staking uint64, tibCount uint32) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if len(endpoint) < 0 {
		return "", errors.New("empty endpoint")
	}

	pubkey, err := utils.ParsingPublickey(earnings)
	if err != nil {
		return "", errors.Wrap(err, "[DecodeToPub]")
	}
	acc, err := types.NewAccountID(pubkey)
	if err != nil {
		return "", errors.Wrap(err, "[NewAccountID]")
	}
	realTokens, ok := new(big.Int).SetString(strconv.FormatUint(staking, 10)+TokenPrecision_CESS, 10)
	if !ok {
		return "", errors.New("[big.Int.SetString]")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_Sminer_regnstk, *acc, types.NewBytes(endpoint), types.NewU128(*realTokens), types.U32(tibCount))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_regnstk, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Sminer_regnstk)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_regnstk, err)
	}

	return blockhash, nil
}

// RegnstkAssignStaking is registered as a storage miner, unlike RegnstkSminer,
// needs to be actively staking by the staking account, which is the first stage
// of storage miner registration.
//
//   - earnings: earnings account
//   - endpoint: communications endpoint
//   - stakingAcc: staking account
//   - tibCount: the size of declaration space, in TiB
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) RegnstkAssignStaking(earnings string, endpoint []byte, stakingAcc string, tibCount uint32) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if len(endpoint) <= 0 {
		return "", errors.New("empty endpoint")
	}

	pubkey, err := utils.ParsingPublickey(earnings)
	if err != nil {
		return "", errors.Wrap(err, "[DecodeToPub]")
	}
	beneficiaryacc, err := types.NewAccountID(pubkey)
	if err != nil {
		return "", errors.Wrap(err, "[NewAccountID]")
	}
	pubkey, err = utils.ParsingPublickey(stakingAcc)
	if err != nil {
		return "", errors.Wrap(err, "[DecodeToPub]")
	}
	stakingacc, err := types.NewAccountID(pubkey)
	if err != nil {
		return "", errors.Wrap(err, "[NewAccountID]")
	}
	newcall, err := types.NewCall(c.metadata, ExtName_Sminer_regnstk_assign_staking, *beneficiaryacc, types.NewBytes(endpoint), *stakingacc, types.U32(tibCount))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_regnstk_assign_staking, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Sminer_regnstk_assign_staking)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_regnstk_assign_staking, err)
	}

	return blockhash, nil
}

// UpdateBeneficiary updates earnings account for storage miner
//
//   - earnings: earnings account
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) UpdateBeneficiary(earnings string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	puk, err := utils.ParsingPublickey(earnings)
	if err != nil {
		return "", err
	}

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return "", errors.Wrap(err, "[NewAccountID]")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_Sminer_update_beneficiary, *acc)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_update_beneficiary, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Sminer_update_beneficiary)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_update_beneficiary, err)
	}

	return blockhash, nil
}

// UpdateSminerEndpoint update address for storage miner
//   - endpoint: address
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) UpdateSminerEndpoint(endpoint []byte) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if len(endpoint) <= 0 {
		return "", errors.New("empty endpoint")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_Sminer_update_endpoint, types.NewBytes(endpoint))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_update_endpoint, err)
		return "", err
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Sminer_update_endpoint)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Sminer_update_endpoint, err)
	}

	return blockhash, nil
}
