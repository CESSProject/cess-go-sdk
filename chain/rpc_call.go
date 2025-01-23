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

// ChainGetBlock get SignedBlock info by block hash
//
// Return:
//   - types.SignedBlock: SignedBlock info
//   - error: error message
func (c *ChainClient) ChainGetBlock(hash types.Hash) (types.SignedBlock, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return types.SignedBlock{}, fmt.Errorf("rpc err: [%s] [rpc_call] [%s] %s", c.GetCurrentRpcAddr(), RPC_Chain_getBlock, ERR_RPC_CONNECTION.Error())
	}

	var data types.SignedBlock
	err := c.api.Client.Call(&data, RPC_Chain_getBlock, hash)
	return data, err
}

// ChainGetBlockHash get block hash by block number
//
// Return:
//   - types.Hash: block hash
//   - error: error message
func (c *ChainClient) ChainGetBlockHash(block uint32) (types.Hash, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return types.Hash{}, fmt.Errorf("rpc err: [%s] [rpc_call] [%s] %s", c.GetCurrentRpcAddr(), RPC_Chain_getBlockHash, ERR_RPC_CONNECTION.Error())
	}

	var data types.Hash
	err := c.api.Client.Call(&data, RPC_Chain_getBlockHash, types.NewU32(block))
	return data, err
}

// ChainGetFinalizedHead get finalized block hash
//
// Return:
//   - types.Hash: block hash
//   - error: error message
func (c *ChainClient) ChainGetFinalizedHead() (types.Hash, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return types.Hash{}, fmt.Errorf("rpc err: [%s] [rpc_call] [%s] %s", c.GetCurrentRpcAddr(), RPC_Chain_getFinalizedHead, ERR_RPC_CONNECTION.Error())
	}

	var data types.Hash
	err := c.api.Client.Call(&data, RPC_Chain_getFinalizedHead)
	return data, err
}

// SystemProperties query system properties
//
// Return:
//   - SysProperties: system properties
//   - error: error message
func (c *ChainClient) SystemProperties() (SysProperties, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return SysProperties{}, fmt.Errorf("rpc err: [%s] [rpc_call] [%s] %s", c.GetCurrentRpcAddr(), RPC_SYS_Properties, ERR_RPC_CONNECTION.Error())
	}

	var data SysProperties
	err := c.api.Client.Call(&data, RPC_SYS_Properties)
	return data, err
}

// SystemProperties query system properties
//
// Return:
//   - string: system chain
//   - error: error message
func (c *ChainClient) SystemChain() (string, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return "", fmt.Errorf("rpc err: [%s] [rpc_call] [%s] %s", c.GetCurrentRpcAddr(), RPC_SYS_Chain, ERR_RPC_CONNECTION.Error())
	}

	var data types.Text
	err := c.api.Client.Call(&data, RPC_SYS_Chain)
	return string(data), err
}

// SystemSyncState query system sync state
//
// Return:
//   - SysSyncState: system sync state
//   - error: error message
func (c *ChainClient) SystemSyncState() (SysSyncState, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return SysSyncState{}, fmt.Errorf("rpc err: [%s] [rpc_call] [%s] %s", c.GetCurrentRpcAddr(), RPC_SYS_SyncState, ERR_RPC_CONNECTION.Error())
	}

	var data SysSyncState
	err := c.api.Client.Call(&data, RPC_SYS_SyncState)
	return data, err
}

// SystemVersion query system version
//
// Return:
//   - string: system version
//   - error: error message
func (c *ChainClient) SystemVersion() (string, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return "", fmt.Errorf("rpc err: [%s] [rpc_call] [%s] %s", c.GetCurrentRpcAddr(), RPC_SYS_Version, ERR_RPC_CONNECTION.Error())
	}

	var data types.Text
	err := c.api.Client.Call(&data, RPC_SYS_Version)
	return string(data), err
}

// NetListening query net listenning
//
// Return:
//   - bool: net listenning
//   - error: error message
func (c *ChainClient) NetListening() (bool, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return false, fmt.Errorf("rpc err: [%s] [rpc_call] [%s] %s", c.GetCurrentRpcAddr(), RPC_NET_Listening, ERR_RPC_CONNECTION.Error())
	}

	var data types.Bool
	err := c.api.Client.Call(&data, RPC_NET_Listening)
	return bool(data), err
}
