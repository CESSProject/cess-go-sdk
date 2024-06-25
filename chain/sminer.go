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
	"strings"
	"time"

	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

// QueryExpenders query expenders (idle data specification)
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - ExpendersInfo: idle data specification
//   - error: error message
func (c *ChainClient) QueryExpenders(block int32) (ExpendersInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data ExpendersInfo

	if !c.GetRpcState() {
		return data, ERR_RPC_CONNECTION
	}

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
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data MinerInfo

	if !c.GetRpcState() {
		return data, ERR_RPC_CONNECTION
	}

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
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U32

	if !c.GetRpcState() {
		return 0, ERR_RPC_CONNECTION
	}

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
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data []types.AccountID

	if !c.GetRpcState() {
		return data, ERR_RPC_CONNECTION
	}

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
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U32

	if !c.GetRpcState() {
		return 0, ERR_RPC_CONNECTION
	}

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
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data MinerReward

	if !c.GetRpcState() {
		return data, ERR_RPC_CONNECTION
	}

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
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data RestoralTargetInfo

	if !c.GetRpcState() {
		return data, ERR_RPC_CONNECTION
	}

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
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var result []RestoralTargetInfo

	if !c.GetRpcState() {
		return nil, ERR_RPC_CONNECTION
	}

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
	defer func() {
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

	if !c.GetRpcState() {
		return data, ERR_RPC_CONNECTION
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
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data CompleteSnapShotType

	if !c.GetRpcState() {
		return 0, 0, ERR_RPC_CONNECTION
	}

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

// QueryCompleteMinerSnapShot query CompleteMinerSnapShot
//   - puk: account id
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - MinerCompleteInfo: CompleteMinerSnapShot
//   - error: error message
func (c *ChainClient) QueryCompleteMinerSnapShot(puk []byte, block int32) (MinerCompleteInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data MinerCompleteInfo

	if !c.GetRpcState() {
		return data, ERR_RPC_CONNECTION
	}

	param, err := codec.Encode(puk)
	if err != nil {
		return data, err
	}

	key, err := types.CreateStorageKey(c.metadata, Sminer, CompleteMinerSnapShot, param)
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
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	tokens, ok := new(big.Int).SetString(token, 10)
	if !ok {
		return "", fmt.Errorf("[IncreaseCollateral] invalid token: %s", token)
	}

	acc, err := types.NewAccountID(accountID)
	if err != nil {
		return blockhash, errors.Wrap(err, "[NewAccountID]")
	}

	call, err := types.NewCall(c.metadata, TX_Sminer_IncreaseCollateral, *acc, types.NewUCompact(tokens))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_Sminer_IncreaseCollateral, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_Sminer_IncreaseCollateral, err)
		return blockhash, err
	}

	ok, err = c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_Sminer_IncreaseCollateral, err)
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_Sminer_IncreaseCollateral, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_IncreaseCollateral, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_IncreaseCollateral, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_Sminer_increase_collateral, SminerIncreaseCollateral, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
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
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, TX_Sminer_IncreaseDeclarationSpace, types.NewU32(tibCount))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_Sminer_IncreaseDeclarationSpace, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_Sminer_IncreaseDeclarationSpace, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_Sminer_IncreaseDeclarationSpace, err)
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_Sminer_IncreaseDeclarationSpace, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_IncreaseDeclarationSpace, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_IncreaseDeclarationSpace, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_Sminer_increase_declaration_space, SminerIncreaseDeclarationSpace, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
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
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		err         error
		txhash      string
		call        types.Call
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return txhash, ERR_RPC_CONNECTION
	}

	acc, err := types.NewAccountID(c.GetSignatureAccPulickey())
	if err != nil {
		return txhash, errors.Wrap(err, "[NewAccountID]")
	}

	call, err = types.NewCall(c.metadata, TX_Sminer_MinerExitPrep, *acc)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_Sminer_MinerExitPrep, err)
		return txhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_Sminer_MinerExitPrep, err)
		return txhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_Sminer_MinerExitPrep, err)
		c.SetRpcState(false)
		return txhash, err
	}

	if !ok {
		return txhash, ERR_RPC_EMPTY_VALUE
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_Sminer_MinerExitPrep, err)
		return txhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return txhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_MinerExitPrep, err)
				c.SetRpcState(false)
				return txhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_MinerExitPrep, err)
			c.SetRpcState(false)
			return txhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				txhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_Sminer_miner_exit, SminerMinerExitPrep, c.signatureAcc)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, ERR_RPC_TIMEOUT
		}
	}
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
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return txhash, ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, TX_Sminer_MinerWithdraw)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_Sminer_MinerWithdraw, err)
		return txhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_Sminer_MinerWithdraw, err)
		return txhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_Sminer_MinerWithdraw, err)
		c.SetRpcState(false)
		return txhash, err
	}
	if !ok {
		return txhash, ERR_RPC_EMPTY_VALUE
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_Sminer_MinerWithdraw, err)
		return txhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return txhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_MinerWithdraw, err)
				c.SetRpcState(false)
				return txhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_MinerWithdraw, err)
			c.SetRpcState(false)
			return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				txhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_Sminer_miner_withdraw, SminerWithdraw, c.signatureAcc)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, ERR_RPC_TIMEOUT
		}
	}
}

// ReceiveReward to receive rewards
//
// Return:
//   - string: block hash
//   - string: earnings account for receiving payments
//   - error: error message
//
// Note:
//   - for storage miner only
//   - pass at least one idle and service challenge at the same time to get the reward
func (c *ChainClient) ReceiveReward() (string, string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		earningsAcc string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return blockhash, earningsAcc, ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, TX_Sminer_ReceiveReward)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_Sminer_ReceiveReward, err)
		return blockhash, earningsAcc, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_Sminer_ReceiveReward, err)
		return blockhash, earningsAcc, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_Sminer_ReceiveReward, err)
		c.SetRpcState(false)
		return blockhash, earningsAcc, err
	}
	if !ok {
		return blockhash, earningsAcc, ERR_RPC_EMPTY_VALUE
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_Sminer_ReceiveReward, err)
		return blockhash, earningsAcc, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, earningsAcc, errors.Wrap(err, "[Sign]")
			}
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_ReceiveReward, err)
				c.SetRpcState(false)
				return blockhash, earningsAcc, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_ReceiveReward, err)
			c.SetRpcState(false)
			return blockhash, earningsAcc, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				receive_event, err := c.RetrieveEvent_Sminer_Receive(status.AsInBlock)
				return blockhash, receive_event.Acc, err
			}
		case err = <-sub.Err():
			return blockhash, earningsAcc, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, earningsAcc, ERR_RPC_TIMEOUT
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
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, TX_Sminer_RegisterPoisKey, poisKey, teeSignWithAcc, teeSign, teePuk)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_Sminer_RegisterPoisKey, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_Sminer_RegisterPoisKey, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_Sminer_RegisterPoisKey, err)
		c.SetRpcState(false)
		return blockhash, err
	}

	if !ok {
		keyStr, _ := utils.NumsToByteStr(key, map[string]bool{})
		return blockhash, fmt.Errorf(
			"chain rpc.state.GetStorageLatest[%v]: %v",
			keyStr,
			ERR_RPC_EMPTY_VALUE,
		)
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_Sminer_RegisterPoisKey, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_RegisterPoisKey, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_RegisterPoisKey, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_Sminer_register_pois_key, SminerRegisterPoisKey, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}

// RegnstkSminer registers as a storage miner,
// which is the first stage of storage miner registration.
//
//   - earnings: earnings account
//   - peerId: peer id
//   - staking: number of staking, the unit is CESS
//   - tibCount: the size of declaration space, in TiB
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) RegnstkSminer(earnings string, peerId []byte, staking uint64, tibCount uint32) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	var peerid PeerId
	if len(peerId) != PeerIdPublicKeyLen {
		return blockhash, fmt.Errorf("[RegnstkSminer] invalid peerid: %v", peerId)
	}

	for i := 0; i < len(peerid); i++ {
		peerid[i] = types.U8(peerId[i])
	}

	pubkey, err := utils.ParsingPublickey(earnings)
	if err != nil {
		return blockhash, errors.Wrap(err, "[DecodeToPub]")
	}
	acc, err := types.NewAccountID(pubkey)
	if err != nil {
		return blockhash, errors.Wrap(err, "[NewAccountID]")
	}
	realTokens, ok := new(big.Int).SetString(strconv.FormatUint(staking, 10)+TokenPrecision_CESS, 10)
	if !ok {
		return blockhash, errors.New("[big.Int.SetString]")
	}
	call, err := types.NewCall(c.metadata, TX_Sminer_Regnstk, *acc, peerid, types.NewU128(*realTokens), types.U32(tibCount))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_Sminer_Regnstk, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_Sminer_Regnstk, err)
		return blockhash, err
	}

	ok, err = c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_Sminer_Regnstk, err)
		c.SetRpcState(false)
		return blockhash, err
	}

	if !ok {
		keyStr, _ := utils.NumsToByteStr(key, map[string]bool{})
		return blockhash, fmt.Errorf(
			"chain rpc.state.GetStorageLatest[%v]: %v",
			keyStr,
			ERR_RPC_EMPTY_VALUE,
		)
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_Sminer_Regnstk, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_Regnstk, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_Regnstk, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_Sminer_regnstk, SminerRegistered, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}

// RegnstkAssignStaking is registered as a storage miner, unlike RegnstkSminer,
// needs to be actively staking by the staking account, which is the first stage
// of storage miner registration.
//
//   - earnings: earnings account
//   - peerId: peer id
//   - stakingAcc: staking account
//   - tibCount: the size of declaration space, in TiB
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) RegnstkAssignStaking(earnings string, peerId []byte, stakingAcc string, tibCount uint32) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	var peerid PeerId
	if len(peerId) != PeerIdPublicKeyLen {
		return blockhash, fmt.Errorf("invalid peerid: %v", peerId)
	}

	for i := 0; i < len(peerid); i++ {
		peerid[i] = types.U8(peerId[i])
	}

	pubkey, err := utils.ParsingPublickey(earnings)
	if err != nil {
		return blockhash, errors.Wrap(err, "[DecodeToPub]")
	}
	beneficiaryacc, err := types.NewAccountID(pubkey)
	if err != nil {
		return blockhash, errors.Wrap(err, "[NewAccountID]")
	}
	pubkey, err = utils.ParsingPublickey(stakingAcc)
	if err != nil {
		return blockhash, errors.Wrap(err, "[DecodeToPub]")
	}
	stakingacc, err := types.NewAccountID(pubkey)
	if err != nil {
		return blockhash, errors.Wrap(err, "[NewAccountID]")
	}
	call, err := types.NewCall(
		c.metadata,
		TX_Sminer_RegnstkAssignStaking,
		*beneficiaryacc,
		peerid,
		*stakingacc,
		types.U32(tibCount),
	)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_Sminer_RegnstkAssignStaking, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_Sminer_RegnstkAssignStaking, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_Sminer_RegnstkAssignStaking, err)
		c.SetRpcState(false)
		return blockhash, err
	}

	if !ok {
		keyStr, _ := utils.NumsToByteStr(key, map[string]bool{})
		return blockhash, fmt.Errorf(
			"chain rpc.state.GetStorageLatest[%v]: %v",
			keyStr,
			ERR_RPC_EMPTY_VALUE,
		)
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_Sminer_RegnstkAssignStaking, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_RegnstkAssignStaking, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_RegnstkAssignStaking, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_Sminer_regnstk_assign_staking, SminerRegistered, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}

// UpdateBeneficiary updates earnings account for storage miner
//
//   - earnings: earnings account
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) UpdateBeneficiary(earnings string) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	puk, err := utils.ParsingPublickey(earnings)
	if err != nil {
		return "", err
	}

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return blockhash, errors.Wrap(err, "[NewAccountID]")
	}

	call, err := types.NewCall(c.metadata, TX_Sminer_UpdateBeneficiary, *acc)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_Sminer_UpdateBeneficiary, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_Sminer_UpdateBeneficiary, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_Sminer_UpdateBeneficiary, err)
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_Sminer_UpdateBeneficiary, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_UpdateBeneficiary, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_UpdateBeneficiary, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_Sminer_update_beneficiary, SminerUpdateBeneficiary, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}

// UpdateSminerPeerId update peer id for storage miner
//
//   - peerid: peer id
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) UpdateSminerPeerId(peerid PeerId) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, TX_Sminer_UpdatePeerId, peerid)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_Sminer_UpdatePeerId, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_Sminer_UpdatePeerId, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_Sminer_UpdatePeerId, err)
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_Sminer_UpdatePeerId, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_UpdatePeerId, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Sminer_UpdatePeerId, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()
	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()
	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_Sminer_update_peer_id, SminerUpdatePeerId, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}
