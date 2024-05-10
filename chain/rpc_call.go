/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"log"

	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// SystemProperties query system properties
//
// Return:
//   - SysProperties: system properties
//   - error: error message
func (c *ChainClient) SystemProperties() (SysProperties, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data SysProperties
	if !c.GetRpcState() {
		return data, ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, RPC_SYS_Properties)
	return data, err
}

// SystemProperties query system properties
//
// Return:
//   - string: system chain
//   - error: error message
func (c *ChainClient) SystemChain() (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.Text
	if !c.GetRpcState() {
		return "", ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, RPC_SYS_Chain)
	return string(data), err
}

// SystemSyncState query system sync state
//
// Return:
//   - SysSyncState: system sync state
//   - error: error message
func (c *ChainClient) SystemSyncState() (SysSyncState, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data SysSyncState
	if !c.GetRpcState() {
		return data, ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, RPC_SYS_SyncState)
	return data, err
}

// SystemVersion query system version
//
// Return:
//   - string: system version
//   - error: error message
func (c *ChainClient) SystemVersion() (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.Text
	if !c.GetRpcState() {
		return "", ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, RPC_SYS_Version)
	return string(data), err
}

// NetListening query net listenning
//
// Return:
//   - bool: net listenning
//   - error: error message
func (c *ChainClient) NetListening() (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.Bool
	if !c.GetRpcState() {
		return false, ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, RPC_NET_Listening)
	return bool(data), err
}
