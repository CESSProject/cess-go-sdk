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

// QueryCounterForValidators query validator number (waiting nodes included)
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint32: validator number
//   - error: error message
func (c *ChainClient) QueryCounterForValidators(block int32) (uint32, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, CounterForValidators, ERR_RPC_CONNECTION.Error())
	}

	var data types.U32

	key, err := types.CreateStorageKey(c.metadata, Staking, CounterForValidators)
	if err != nil {
		return uint32(data), errors.Wrap(err, "[CreateStorageKey]")
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			c.SetRpcState(false)
			return uint32(data), errors.Wrap(err, "[GetStorageLatest]")
		}
		if !ok {
			return uint32(data), ERR_RPC_EMPTY_VALUE
		}
		return uint32(data), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Staking, CounterForValidators, err)
		return uint32(data), err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		c.SetRpcState(false)
		return uint32(data), errors.Wrap(err, "[GetStorage]")
	}
	if !ok {
		return uint32(data), ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

// QueryValidatorsCount query validator number (waiting nodes not included)
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint32: validator number
//   - error: error message
func (c *ChainClient) QueryValidatorsCount(block int32) (uint32, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, ValidatorCount, ERR_RPC_CONNECTION.Error())
	}

	var data types.U32

	key, err := types.CreateStorageKey(c.metadata, Staking, ValidatorCount)
	if err != nil {
		return uint32(data), errors.Wrap(err, "[CreateStorageKey]")
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			c.SetRpcState(false)
			return uint32(data), errors.Wrap(err, "[GetStorageLatest]")
		}
		if !ok {
			return uint32(data), ERR_RPC_EMPTY_VALUE
		}
		return uint32(data), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Staking, ValidatorCount, err)
		return uint32(data), err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		c.SetRpcState(false)
		return uint32(data), errors.Wrap(err, "[GetStorage]")
	}
	if !ok {
		return uint32(data), ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

// QueryNominatorCount query nominator number
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint32: nominator number
//   - error: error message
func (c *ChainClient) QueryNominatorCount(block int32) (uint32, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, CounterForNominators, ERR_RPC_CONNECTION.Error())
	}

	var data types.U32

	key, err := types.CreateStorageKey(c.metadata, Staking, CounterForNominators)
	if err != nil {
		return uint32(data), errors.Wrap(err, "[CreateStorageKey]")
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			c.SetRpcState(false)
			return uint32(data), errors.Wrap(err, "[GetStorageLatest]")
		}
		if !ok {
			return uint32(data), ERR_RPC_EMPTY_VALUE
		}
		return uint32(data), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Staking, CounterForNominators, err)
		return uint32(data), err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		c.SetRpcState(false)
		return uint32(data), errors.Wrap(err, "[GetStorage]")
	}
	if !ok {
		return uint32(data), ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

// QueryErasTotalStake query the total number of staking for each era
//   - era: era id
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: the total number of staking
//   - error: error message
func (c *ChainClient) QueryErasTotalStake(era uint32, block int32) (string, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return "", fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, ErasTotalStake, ERR_RPC_CONNECTION.Error())
	}

	var data types.U128

	param, err := codec.Encode(era)
	if err != nil {
		return "", err
	}

	key, err := types.CreateStorageKey(c.metadata, Staking, ErasTotalStake, param)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Staking, ErasTotalStake, err)
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Staking, ErasTotalStake, err)
			c.SetRpcState(false)
			return "", err
		}
		if !ok {
			return "", ERR_RPC_EMPTY_VALUE
		}
		return data.String(), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Staking, ErasTotalStake, err)
		return data.String(), err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Staking, ErasTotalStake, err)
		c.SetRpcState(false)
		return "", err
	}
	if !ok {
		return "", ERR_RPC_EMPTY_VALUE
	}
	return data.String(), nil
}

// QueryCurrentEra query the current era id
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint32: era id
//   - error: error message
func (c *ChainClient) QueryCurrentEra(block int32) (uint32, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, CurrentEra, ERR_RPC_CONNECTION.Error())
	}

	var data types.U32

	key, err := types.CreateStorageKey(c.metadata, Staking, CurrentEra)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Staking, CurrentEra, err)
		return 0, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Staking, CurrentEra, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Staking, CurrentEra, err)
		return uint32(data), err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Staking, CurrentEra, err)
		c.SetRpcState(false)
		return 0, err
	}
	if !ok {
		return 0, ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

// QueryErasRewardPoints query the rewards of consensus nodes in each era
//   - era: era id
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - StakingEraRewardPoints: the rewards of consensus nodes
//   - error: error message
func (c *ChainClient) QueryErasRewardPoints(era uint32, block int32) (StakingEraRewardPoints, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return StakingEraRewardPoints{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, ErasRewardPoints, ERR_RPC_CONNECTION.Error())
	}

	var result StakingEraRewardPoints

	param1, err := codec.Encode(types.NewU32(era))
	if err != nil {
		return result, err
	}

	key, err := types.CreateStorageKey(c.metadata, Staking, ErasRewardPoints, param1)
	if err != nil {
		return result, err
	}
	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &result)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Staking, ErasRewardPoints, err)
			c.SetRpcState(false)
			return result, err
		}
		if !ok {
			return result, ERR_RPC_EMPTY_VALUE
		}
		return result, nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return result, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &result, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Staking, ErasRewardPoints, err)
		c.SetRpcState(false)
		return result, err
	}
	if !ok {
		return result, ERR_RPC_EMPTY_VALUE
	}
	return result, nil
}

// QueryAllNominators query all nominators info
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []StakingNominations: all nominators info
//   - error: error message
func (c *ChainClient) QueryAllNominators(block int32) ([]StakingNominations, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return []StakingNominations{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, Nominators, ERR_RPC_CONNECTION.Error())
	}

	var result []StakingNominations

	key := CreatePrefixedKey(Staking, Nominators)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), Staking, Nominators, err)
		c.SetRpcState(false)
		return nil, err
	}
	var set []types.StorageChangeSet
	if block < 0 {
		set, err = c.api.RPC.State.QueryStorageAtLatest(keys)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), Staking, Nominators, err)
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
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAt: %v", c.GetCurrentRpcAddr(), Staking, Nominators, err)
			c.SetRpcState(false)
			return nil, err
		}
	}

	for _, elem := range set {
		for _, change := range elem.Changes {
			if change.HasStorageData {
				var data StakingNominations
				if err := codec.Decode(change.StorageData, &data); err != nil {
					continue
				}
				result = append(result, data)
			}

		}
	}
	return result, nil
}

// QueryAllBonded query all consensus and nominators accounts
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []types.AccountID: all consensus and nominators accounts
//   - error: error message
func (c *ChainClient) QueryAllBonded(block int32) ([]types.AccountID, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return []types.AccountID{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, Bonded, ERR_RPC_CONNECTION.Error())
	}

	var result []types.AccountID

	key := CreatePrefixedKey(Staking, Bonded)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), Staking, Bonded, err)
		c.SetRpcState(false)
		return nil, err
	}

	var set []types.StorageChangeSet
	if block < 0 {
		set, err = c.api.RPC.State.QueryStorageAtLatest(keys)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), Staking, Bonded, err)
			c.SetRpcState(false)
			return nil, err
		}
	} else {
		blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Staking, Bonded, err)
			return result, err
		}

		set, err = c.api.RPC.State.QueryStorageAt(keys, blockhash)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAt: %v", c.GetCurrentRpcAddr(), Staking, Bonded, err)
			c.SetRpcState(false)
			return nil, err
		}
	}

	for _, elem := range set {
		for _, change := range elem.Changes {
			if change.HasStorageData {
				var data types.AccountID
				if err := codec.Decode(change.StorageData, &data); err != nil {
					continue
				}
				result = append(result, data)
			}

		}
	}
	return result, nil
}

// QueryValidatorCommission query validator commission
//   - accountID: validator account
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint8: validator commission
//   - error: error message
func (c *ChainClient) QueryValidatorCommission(accountID []byte, block int32) (uint8, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, Validators, ERR_RPC_CONNECTION.Error())
	}

	var result StakingValidatorPrefs

	key, err := types.CreateStorageKey(c.metadata, Staking, Validators, accountID)
	if err != nil {
		return 0, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &result)
		if err != nil {
			return 0, nil
		}
		if !ok {
			return 0, ERR_RPC_EMPTY_VALUE
		}
		return uint8(uint32(result.Commission-2) / uint32(40000000)), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Staking, Validators, err)
		return 0, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &result, blockhash)
	if err != nil {
		return 0, nil
	}
	if !ok {
		return 0, ERR_RPC_EMPTY_VALUE
	}
	return uint8(uint32(result.Commission-2) / uint32(40000000)), nil
}

// QueryEraValidatorReward query the total rewards for each era
//   - era: era id
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: total rewards
//   - error: error message
func (c *ChainClient) QueryEraValidatorReward(era uint32, block int32) (string, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return "", fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, ErasValidatorReward, ERR_RPC_CONNECTION.Error())
	}

	var result types.U128

	param, err := codec.Encode(types.NewU32(era))
	if err != nil {
		return "", err
	}

	key, err := types.CreateStorageKey(c.metadata, Staking, ErasValidatorReward, param)
	if err != nil {
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &result)
		if err != nil {
			return "", nil
		}
		if !ok {
			return "", ERR_RPC_EMPTY_VALUE
		}
		return result.String(), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return "", err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &result, blockhash)
	if err != nil {
		return "", nil
	}
	if !ok {
		return "", ERR_RPC_EMPTY_VALUE
	}
	return result.String(), nil
}

// QueryLedger query the staking ledger
//   - accountID: account id
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - StakingLedger: staking ledger
//   - error: error message
func (c *ChainClient) QueryLedger(accountID []byte, block int32) (StakingLedger, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return StakingLedger{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, Ledger, ERR_RPC_CONNECTION.Error())
	}

	var result StakingLedger

	key, err := types.CreateStorageKey(c.metadata, Staking, Ledger, accountID)
	if err != nil {
		return result, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &result)
		if err != nil {
			return result, nil
		}
		if !ok {
			return result, ERR_RPC_EMPTY_VALUE
		}
		return result, nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return result, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &result, blockhash)
	if err != nil {
		return result, nil
	}
	if !ok {
		return result, ERR_RPC_EMPTY_VALUE
	}
	return result, nil
}

// QueryeErasStakers query the staking exposure
//   - era: era id
//   - accountId: account id
//
// Return:
//   - StakingExposure: staking exposure
//   - error: error message
func (c *ChainClient) QueryeErasStakers(era uint32, accountId []byte) (StakingExposure, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return StakingExposure{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, ErasStakers, ERR_RPC_CONNECTION.Error())
	}

	var result StakingExposure

	param1, err := codec.Encode(types.NewU32(era))
	if err != nil {
		return result, err
	}

	key, err := types.CreateStorageKey(c.metadata, Staking, ErasStakers, param1, accountId)
	if err != nil {
		return result, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &result)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Staking, ErasStakers, err)
		c.SetRpcState(false)
		return result, err
	}
	if !ok {
		return result, ERR_RPC_EMPTY_VALUE
	}
	return result, nil
}

// QueryeNominators query the nominator info
//   - accountId: account id
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - StakingNominations: nominator info
//   - error: error message
func (c *ChainClient) QueryeNominators(accountId []byte, block int32) (StakingNominations, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return StakingNominations{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, Nominators, ERR_RPC_CONNECTION.Error())
	}

	var result StakingNominations

	key, err := types.CreateStorageKey(c.metadata, Staking, Nominators, accountId)
	if err != nil {
		return result, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &result)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Staking, Nominators, err)
			c.SetRpcState(false)
			return result, err
		}
		if !ok {
			return result, ERR_RPC_EMPTY_VALUE
		}
		return result, nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return result, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &result, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Staking, Nominators, err)
		c.SetRpcState(false)
		return result, err
	}
	if !ok {
		return result, ERR_RPC_EMPTY_VALUE
	}
	return result, nil
}

// QueryeAllErasStakersPaged query all the staking exposure
//   - era: era id
//   - accountId: account id
//
// Return:
//   - []QueryeErasStakersPaged: all staking exposure
//   - error: error message
func (c *ChainClient) QueryeAllErasStakersPaged(era uint32, accountId []byte) ([]StakingExposurePaged, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return []StakingExposurePaged{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, ErasStakersPaged, ERR_RPC_CONNECTION.Error())
	}

	var result []StakingExposurePaged

	param1, err := codec.Encode(types.NewU32(era))
	if err != nil {
		return result, err
	}

	for i := 0; i < 256; i++ {
		param3, err := codec.Encode(types.U32(i))
		if err != nil {
			return result, err
		}
		key, err := types.CreateStorageKey(c.metadata, Staking, ErasStakersPaged, param1, accountId, param3)
		if err != nil {
			return result, err
		}
		var data StakingExposurePaged
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Staking, ErasStakers, err)
			c.SetRpcState(false)
			return result, err
		}
		if !ok {
			break
		}
		result = append(result, data)
	}
	return result, nil
}

// QueryeErasStakersOverview query the PagedExposureMetadata
//   - era: era id
//   - accountId: account id
//
// Return:
//   - PagedExposureMetadata: PagedExposureMetadata
//   - error: error message
func (c *ChainClient) QueryeErasStakersOverview(era uint32, accountId []byte) (PagedExposureMetadata, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return PagedExposureMetadata{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Staking, ErasStakersOverview, ERR_RPC_CONNECTION.Error())
	}

	var result PagedExposureMetadata

	param1, err := codec.Encode(types.NewU32(era))
	if err != nil {
		return result, err
	}

	key, err := types.CreateStorageKey(c.metadata, Staking, ErasStakersOverview, param1, accountId)
	if err != nil {
		return result, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &result)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Staking, ErasStakers, err)
		c.SetRpcState(false)
		return result, err
	}
	if !ok {
		return result, ERR_RPC_EMPTY_VALUE
	}

	return result, nil
}
