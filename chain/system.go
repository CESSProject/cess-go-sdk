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

// QueryNodeSynchronizationSt
func (c *ChainClient) QueryNodeSynchronizationSt() (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetChainState() {
		return false, pattern.ERR_RPC_CONNECTION
	}
	h, err := c.api.RPC.System.Health()
	if err != nil {
		return false, err
	}
	return h.IsSyncing, nil
}

// QueryBlockHeight
func (c *ChainClient) QueryBlockHeight(hash string) (uint32, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if hash != "" {
		var h types.Hash
		err := codec.DecodeFromHex(hash, &h)
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

// QueryAccountInfo
func (c *ChainClient) QueryAccountInfoByAccount(acc string) (types.AccountInfo, error) {
	puk, err := utils.ParsingPublickey(acc)
	if err != nil {
		return types.AccountInfo{}, err
	}
	return c.QueryAccountInfo(puk)
}

// QueryAccountInfo
func (c *ChainClient) QueryAccountInfo(puk []byte) (types.AccountInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.AccountInfo

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	b, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, b)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, pattern.ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryAllAccountInfoFromBlock
func (c *ChainClient) QueryAllAccountInfoFromBlock(block int) ([]types.AccountInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data []types.AccountInfo

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	key := createPrefixedKey(pattern.SYSTEM, pattern.ACCOUNT)

	if block < 0 {
		keys, err := c.api.RPC.State.GetKeysLatest(key)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), pattern.SYSTEM, pattern.ACCOUNT, err)
			c.SetChainState(false)
			return nil, err
		}
		set, err := c.api.RPC.State.QueryStorageAtLatest(keys)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), pattern.SYSTEM, pattern.ACCOUNT, err)
			c.SetChainState(false)
			return nil, err
		}
		for _, elem := range set {
			for _, change := range elem.Changes {
				var val types.AccountInfo
				if err := codec.Decode(change.StorageData, &val); err != nil {
					fmt.Println("Decode StorageData:", err)
					continue
				}
				var kkey types.AccountID
				if err := codec.Decode(change.StorageKey, &kkey); err != nil {
					fmt.Println("Decode StorageKey:", err)
					continue
				}
				fmt.Println(utils.EncodePublicKeyAsCessAccount(kkey[:]))
				data = append(data, val)
			}
		}
		return data, nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), pattern.SYSTEM, pattern.ACCOUNT, err)
		c.SetChainState(false)
		return data, err
	}

	fmt.Println(">>>>>")
	keys, err := c.api.RPC.State.GetKeys(key, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeys: %v", c.GetCurrentRpcAddr(), pattern.SYSTEM, pattern.ACCOUNT, err)
		c.SetChainState(false)
		return nil, err
	}
	set, err := c.api.RPC.State.QueryStorageAt(keys, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAt: %v", c.GetCurrentRpcAddr(), pattern.SYSTEM, pattern.ACCOUNT, err)
		c.SetChainState(false)
		return nil, err
	}
	for _, elem := range set {
		for _, change := range elem.Changes {
			if change.HasStorageData {
				var val types.AccountInfo

				var kkey types.AccountID
				if err := codec.Decode(change.StorageKey, &kkey); err != nil {
					fmt.Println("Decode StorageKey:", err)
					continue
				}
				if err := codec.Decode(change.StorageData, &val); err != nil {
					fmt.Println("Decode StorageData:", err)
					continue
				}
				fmt.Println(utils.EncodePublicKeyAsCessAccount(kkey[:]))
				data = append(data, val)
			}
		}
	}

	return data, nil
}

func (c *ChainClient) SysProperties() (pattern.SysProperties, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data pattern.SysProperties
	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, pattern.RPC_SYS_Properties)
	return data, err
}

func (c *ChainClient) SysChain() (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.Text
	if !c.GetChainState() {
		return "", pattern.ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, pattern.RPC_SYS_Chain)
	return string(data), err
}

func (c *ChainClient) SyncState() (pattern.SysSyncState, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data pattern.SysSyncState
	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, pattern.RPC_SYS_SyncState)
	return data, err
}

func (c *ChainClient) ChainVersion() (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.Text
	if !c.GetChainState() {
		return "", pattern.ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, pattern.RPC_SYS_Version)
	return string(data), err
}

func (c *ChainClient) NetListening() (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.Bool
	if !c.GetChainState() {
		return false, pattern.ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, pattern.RPC_NET_Listening)
	return bool(data), err
}
