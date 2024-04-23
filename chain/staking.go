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
	"github.com/pkg/errors"
)

func (c *ChainClient) QueryAllValidatorCount(block int) (uint32, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U32

	if !c.GetChainState() {
		return uint32(data), pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.STAKING, pattern.COUNTERFORVALIDATORS)
	if err != nil {
		return uint32(data), errors.Wrap(err, "[CreateStorageKey]")
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			return uint32(data), errors.Wrap(err, "[GetStorageLatest]")
		}
		if !ok {
			return uint32(data), pattern.ERR_RPC_EMPTY_VALUE
		}
		return uint32(data), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.COUNTERFORVALIDATORS, err)
		c.SetChainState(false)
		return uint32(data), err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		return uint32(data), errors.Wrap(err, "[GetStorage]")
	}
	if !ok {
		return uint32(data), pattern.ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

func (c *ChainClient) QueryValidatorsCount(block int) (uint32, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U32

	if !c.GetChainState() {
		return uint32(data), pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.STAKING, pattern.ValidatorCount)
	if err != nil {
		return uint32(data), errors.Wrap(err, "[CreateStorageKey]")
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			return uint32(data), errors.Wrap(err, "[GetStorageLatest]")
		}
		if !ok {
			return uint32(data), pattern.ERR_RPC_EMPTY_VALUE
		}
		return uint32(data), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.ValidatorCount, err)
		c.SetChainState(false)
		return uint32(data), err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		return uint32(data), errors.Wrap(err, "[GetStorage]")
	}
	if !ok {
		return uint32(data), pattern.ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

func (c *ChainClient) QueryNominatorCount(block int) (uint32, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U32

	if !c.GetChainState() {
		return uint32(data), pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.STAKING, pattern.CounterForNominators)
	if err != nil {
		return uint32(data), errors.Wrap(err, "[CreateStorageKey]")
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			return uint32(data), errors.Wrap(err, "[GetStorageLatest]")
		}
		if !ok {
			return uint32(data), pattern.ERR_RPC_EMPTY_VALUE
		}
		return uint32(data), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.CounterForNominators, err)
		c.SetChainState(false)
		return uint32(data), err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		return uint32(data), errors.Wrap(err, "[GetStorage]")
	}
	if !ok {
		return uint32(data), pattern.ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

func (c *ChainClient) QueryErasTotalStake(era uint32, block int) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	if !c.GetChainState() {
		return "", pattern.ERR_RPC_CONNECTION
	}

	param, err := codec.Encode(era)
	if err != nil {
		return "", err
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.STAKING, pattern.ErasTotalStake, param)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.ErasTotalStake, err)
		c.SetChainState(false)
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.ErasTotalStake, err)
			c.SetChainState(false)
			return "", err
		}
		if !ok {
			return "", pattern.ERR_RPC_EMPTY_VALUE
		}
		return data.String(), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.ErasTotalStake, err)
		c.SetChainState(false)
		return data.String(), err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.ErasTotalStake, err)
		c.SetChainState(false)
		return "", err
	}
	if !ok {
		return "", pattern.ERR_RPC_EMPTY_VALUE
	}
	return data.String(), nil
}

func (c *ChainClient) QueryCurrentEra(block int) (uint32, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U32

	if !c.GetChainState() {
		return 0, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.STAKING, pattern.CurrentEra)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.CurrentEra, err)
		c.SetChainState(false)
		return 0, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.CurrentEra, err)
			c.SetChainState(false)
			return 0, err
		}
		if !ok {
			return 0, pattern.ERR_RPC_EMPTY_VALUE
		}
		return uint32(data), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.CurrentEra, err)
		c.SetChainState(false)
		return uint32(data), err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.CurrentEra, err)
		c.SetChainState(false)
		return 0, err
	}
	if !ok {
		return 0, pattern.ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

// func (c *ChainClient) QueryeErasStakers(era uint32) ([]pattern.StakingExposure, error) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.Println(utils.RecoverError(err))
// 		}
// 	}()
// 	var result []pattern.StakingExposure

// 	if !c.GetChainState() {
// 		return result, pattern.ERR_RPC_CONNECTION
// 	}

// 	// key := createPrefixedKey(pattern.STAKING, pattern.ErasStakers)
// 	param1, err := codec.Encode(types.NewU32(era))
// 	if err != nil {
// 		return result, err
// 	}

// 	var p2 types.OptionAccountID

// 	pk, err := utils.ParsingPublickey("cXjGzUnWJcNXBzKBEJKvs3ZBJ5f1aEEca38abpuNxvwGNZ5Gy")
// 	if err != nil {
// 		return result, err
// 	}

// 	accid, err := types.NewAccountID(pk)
// 	if err != nil {
// 		return result, err
// 	}
// 	p2.SetSome(*accid)
// 	b := bytes.NewBuffer(make([]byte, 0))

// 	err = p2.Encode(*scale.NewEncoder(b))
// 	if err != nil {
// 		return result, err
// 	}
// 	param2, err := codec.Encode(p2)
// 	if err != nil {
// 		return result, err
// 	}
// 	_ = param2
// 	//fmt.Println(p2.HasValue())
// 	key, err := types.CreateStorageKey(c.metadata, pattern.STAKING, "ErasStakersClipped", param1, param2)
// 	if err != nil {
// 		return result, err
// 	}
// 	_ = key
// 	kkey := createPrefixedKey(pattern.STAKING, pattern.ErasStakers)
// 	//kkey = append(kkey, []byte(" ")...)
// 	//kkey = append(kkey, param1...) //xxhash.New128(param1).Sum(nil)...)

// 	entryMeta, err := c.GetMetadata().FindStorageEntryMetadata(pattern.STAKING, pattern.ErasStakers)
// 	if err != nil {
// 		return nil, err
// 	}
// 	hashers, err := entryMeta.Hashers()
// 	if err != nil {
// 		return nil, err
// 	}
// 	_, err = hashers[0].Write(param1)
// 	if err != nil {
// 		return nil, err
// 	}
// 	kkey = append(kkey, hashers[0].Sum(nil)...)
// 	_, err = hashers[1].Write(param2)
// 	if err != nil {
// 		return nil, err
// 	}
// 	kkey = append(kkey, hashers[1].Sum(nil)...)
// 	var result1 pattern.StakingExposure
// 	ok, err := c.api.RPC.State.GetStorageLatest(kkey, &result1)
// 	if err != nil {
// 		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.ErasStakers, err)
// 		c.SetChainState(false)
// 		return result, err
// 	}
// 	fmt.Println(result1)
// 	if !ok {
// 		return result, pattern.ERR_RPC_EMPTY_VALUE
// 	}
// 	return result, nil
// }

func (c *ChainClient) QueryStakingEraRewardPoints(era uint32) (pattern.StakingEraRewardPoints, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var result pattern.StakingEraRewardPoints

	if !c.GetChainState() {
		return result, pattern.ERR_RPC_CONNECTION
	}

	param1, err := codec.Encode(types.NewU32(era))
	if err != nil {
		return result, err
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.STAKING, pattern.ErasRewardPoints, param1)
	if err != nil {
		return result, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &result)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.ErasRewardPoints, err)
		c.SetChainState(false)
		return result, err
	}
	if !ok {
		return result, pattern.ERR_RPC_EMPTY_VALUE
	}
	return result, nil
}

func (c *ChainClient) QueryNominatorsLatest() ([]pattern.StakingNominations, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	//var keyList []string
	var result []pattern.StakingNominations

	if !c.GetChainState() {
		return nil, pattern.ERR_RPC_CONNECTION
	}

	key := createPrefixedKey(pattern.STAKING, pattern.Nominators)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.Nominators, err)
		c.SetChainState(false)
		return nil, err
	}

	set, err := c.api.RPC.State.QueryStorageAtLatest(keys)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.Nominators, err)
		c.SetChainState(false)
		return nil, err
	}

	for _, elem := range set {
		for _, change := range elem.Changes {
			if change.HasStorageData {
				var data pattern.StakingNominations
				if err := codec.Decode(change.StorageData, &data); err != nil {
					continue
				}
				result = append(result, data)
			}

		}
	}
	return result, nil
}

func (c *ChainClient) QueryBondedList(block int32) ([]types.AccountID, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	//var keyList []string
	var result []types.AccountID

	if !c.GetChainState() {
		return nil, pattern.ERR_RPC_CONNECTION
	}

	key := createPrefixedKey(pattern.STAKING, pattern.Bonded)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.Bonded, err)
		c.SetChainState(false)
		return nil, err
	}

	if block < 0 {
		set, err := c.api.RPC.State.QueryStorageAtLatest(keys)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.Bonded, err)
			c.SetChainState(false)
			return nil, err
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

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.Bonded, err)
		c.SetChainState(false)
		return result, err
	}

	set, err := c.api.RPC.State.QueryStorageAt(keys, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.Bonded, err)
		c.SetChainState(false)
		return nil, err
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

func (c *ChainClient) QueryValidatorCommission(account string, block int) (uint8, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	//var keyList []string
	var result pattern.StakingValidatorPrefs

	if !c.GetChainState() {
		return 0, pattern.ERR_RPC_CONNECTION
	}

	pk, _ := utils.ParsingPublickey(account)
	acc, _ := types.NewAccountID(pk)
	// key := createPrefixedKey(pattern.STAKING, pattern.ErasStakers)
	param1, err := codec.Encode(*acc)
	if err != nil {
		return 0, err
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.STAKING, pattern.Validators, param1)
	if err != nil {
		return 0, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &result)
		if err != nil {
			return 0, nil
		}
		if !ok {
			return 0, pattern.ERR_RPC_EMPTY_VALUE
		}
		return uint8(uint32(result.Commission-2) / uint32(40000000)), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), pattern.STAKING, pattern.Validators, err)
		c.SetChainState(false)
		return 0, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &result, blockhash)
	if err != nil {
		return 0, nil
	}
	if !ok {
		return 0, pattern.ERR_RPC_EMPTY_VALUE
	}
	return uint8(uint32(result.Commission-2) / uint32(40000000)), nil
}

func (c *ChainClient) QueryEraValidatorReward(era uint32) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	//var keyList []string
	var result types.U128

	if !c.GetChainState() {
		return "", pattern.ERR_RPC_CONNECTION
	}

	param, err := codec.Encode(types.NewU32(era))
	if err != nil {
		return "", err
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.STAKING, pattern.ErasValidatorReward, param)
	if err != nil {
		return "", err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &result)
	if err != nil {
		return "", nil
	}
	if !ok {
		return "", pattern.ERR_RPC_EMPTY_VALUE
	}
	return result.String(), nil
}
