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

// QueryMasterPubKey query master public key
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []byte: master public key
//   - error: error message
func (c *ChainClient) QueryMasterPubKey(block int32) ([]byte, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return []byte{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), TeeWorker, MasterPubkey, ERR_RPC_CONNECTION.Error())
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data MasterPublicKey

	key, err := types.CreateStorageKey(c.metadata, TeeWorker, MasterPubkey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TeeWorker, MasterPubkey, err)
		return nil, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			c.SetRpcState(false)
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TeeWorker, MasterPubkey, err)
			return nil, err
		}
		if !ok {
			return nil, ERR_RPC_EMPTY_VALUE
		}
		return []byte(string(data[:])), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return nil, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		c.SetRpcState(false)
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), TeeWorker, MasterPubkey, err)
		return nil, err
	}
	if !ok {
		return nil, ERR_RPC_EMPTY_VALUE
	}
	return []byte(string(data[:])), nil
}

// QueryWorkers query tee work info
//   - puk: tee's work public key
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - WorkerInfo: tee worker info
//   - error: error message
func (c *ChainClient) QueryWorkers(puk WorkerPublicKey, block int32) (WorkerInfo, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return WorkerInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), TeeWorker, Workers, ERR_RPC_CONNECTION.Error())
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data WorkerInfo

	publickey, err := codec.Encode(puk)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(c.metadata, TeeWorker, Workers, publickey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TeeWorker, Workers, err)
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			c.SetRpcState(false)
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TeeWorker, Workers, err)
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
		c.SetRpcState(false)
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TeeWorker, Workers, err)
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryAllWorkers query all tee work info
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []WorkerInfo: all tee worker info
//   - error: error message
func (c *ChainClient) QueryAllWorkers(block int32) ([]WorkerInfo, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return []WorkerInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), TeeWorker, Workers, ERR_RPC_CONNECTION.Error())
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var list []WorkerInfo

	key := CreatePrefixedKey(TeeWorker, Workers)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), TeeWorker, Workers, err)
		return list, err
	}
	var set []types.StorageChangeSet
	if block < 0 {
		set, err = c.api.RPC.State.QueryStorageAtLatest(keys)
		if err != nil {
			c.SetRpcState(false)
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), TeeWorker, Workers, err)
			return list, err
		}
	} else {
		blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
		if err != nil {
			return list, err
		}
		set, err = c.api.RPC.State.QueryStorageAt(keys, blockhash)
		if err != nil {
			c.SetRpcState(false)
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAt: %v", c.GetCurrentRpcAddr(), TeeWorker, Workers, err)
			return list, err
		}
	}
	for _, elem := range set {
		for _, change := range elem.Changes {
			var teeWorker WorkerInfo
			if err := codec.Decode(change.StorageData, &teeWorker); err != nil {
				continue
			}
			list = append(list, teeWorker)
		}
	}
	return list, nil
}

// QueryEndpoints query tee's endpoint
//   - puk: tee's work public key
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: tee's endpoint
//   - error: error message
func (c *ChainClient) QueryEndpoints(puk WorkerPublicKey, block int32) (string, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return "", fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), TeeWorker, Endpoints, ERR_RPC_CONNECTION.Error())
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.Text

	val, err := codec.Encode(puk)
	if err != nil {
		return "", errors.Wrap(err, "[Encode]")
	}
	key, err := types.CreateStorageKey(c.metadata, TeeWorker, Endpoints, val)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TeeWorker, Endpoints, err)
		return "", errors.Wrap(err, "[CreateStorageKey]")
	}
	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			c.SetRpcState(false)
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TeeWorker, Endpoints, err)
			return "", errors.Wrap(err, "[GetStorageLatest]")
		}
		if !ok {
			return "", ERR_RPC_EMPTY_VALUE
		}
		return string(data), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return string(data), err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		c.SetRpcState(false)
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), TeeWorker, Endpoints, err)
		return "", errors.Wrap(err, "[GetStorage]")
	}
	if !ok {
		return "", ERR_RPC_EMPTY_VALUE
	}
	return string(data), nil
}

// QueryWorkerAddedAt query tee work registered block
//   - puk: tee's work public key
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint32: tee work registered block
//   - error: error message
func (c *ChainClient) QueryWorkerAddedAt(puk WorkerPublicKey, block int32) (uint32, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), TeeWorker, WorkerAddedAt, ERR_RPC_CONNECTION.Error())
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U32

	val, err := codec.Encode(puk)
	if err != nil {
		return uint32(data), errors.Wrap(err, "[Encode]")
	}
	key, err := types.CreateStorageKey(c.metadata, TeeWorker, WorkerAddedAt, val)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TeeWorker, WorkerAddedAt, err)
		return uint32(data), err
	}
	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			c.SetRpcState(false)
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TeeWorker, WorkerAddedAt, err)
			return uint32(data), err
		}
		if !ok {
			return uint32(data), ERR_RPC_EMPTY_VALUE
		}
		return uint32(data), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return uint32(data), err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		c.SetRpcState(false)
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), TeeWorker, WorkerAddedAt, err)
		return uint32(data), err
	}
	if !ok {
		return uint32(data), ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}
