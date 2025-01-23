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
	"github.com/mr-tron/base58"
	"github.com/pkg/errors"
)

// QueryOss query oss info
//   - accountID: oss's account
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - OssInfo: oss info
//   - error: error message
func (c *ChainClient) QueryOss(accountID []byte, block int32) (OssInfo, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return OssInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Oss, Oss, ERR_RPC_CONNECTION.Error())
	}

	var data OssInfo

	key, err := types.CreateStorageKey(c.metadata, Oss, Oss, accountID)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Oss, Oss, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Oss, Oss, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Oss, Oss, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryAllOss query all oss info
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []OssInfo: all oss info
//   - error: error message
func (c *ChainClient) QueryAllOss(block int32) ([]OssInfo, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return []OssInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Oss, Oss, ERR_RPC_CONNECTION.Error())
	}

	var result []OssInfo

	key := CreatePrefixedKey(Oss, Oss)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), Oss, Oss, err)
		c.SetRpcState(false)
		return nil, err
	}

	var set []types.StorageChangeSet
	if block < 0 {
		set, err = c.api.RPC.State.QueryStorageAtLatest(keys)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), Oss, Oss, err)
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
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), Oss, Oss, err)
			c.SetRpcState(false)
			return nil, err
		}
	}
	for _, elem := range set {
		for _, change := range elem.Changes {
			var data OssInfo
			if err := codec.Decode(change.StorageData, &data); err != nil {
				continue
			}
			result = append(result, data)
		}
	}
	return result, nil
}

// QueryAllOssPeerId query all oss's peer id
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []string: all oss's peer id
//   - error: error message
func (c *ChainClient) QueryAllOssPeerId(block int32) ([]string, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return []string{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Oss, Oss, ERR_RPC_CONNECTION.Error())
	}

	var result []string

	key := CreatePrefixedKey(Oss, Oss)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), Oss, Oss, err)
		c.SetRpcState(false)
		return nil, err
	}

	var set []types.StorageChangeSet
	if block < 0 {
		set, err = c.api.RPC.State.QueryStorageAtLatest(keys)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), Oss, Oss, err)
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
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), Oss, Oss, err)
			c.SetRpcState(false)
			return nil, err
		}
	}

	for _, elem := range set {
		for _, change := range elem.Changes {
			var data OssInfo
			if err := codec.Decode(change.StorageData, &data); err != nil {
				continue
			}
			result = append(result, base58.Encode([]byte(string(data.Peerid[:]))))
		}
	}
	return result, nil
}

// QueryAuthorityList query authorised all accounts
//   - accountID: account to be queried
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []types.AccountID: authorised all accounts
//   - error: error message
func (c *ChainClient) QueryAuthorityList(accountID []byte, block int32) ([]types.AccountID, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return []types.AccountID{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Oss, AuthorityList, ERR_RPC_CONNECTION.Error())
	}

	var data []types.AccountID

	key, err := types.CreateStorageKey(c.metadata, Oss, AuthorityList, accountID)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Oss, AuthorityList, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Oss, AuthorityList, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Oss, AuthorityList, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// Authorize to authorise space usage to another account
//   - accountID: authorised account
//
// Return:
//   - string: block hash
//   - error: error message
//
// Node:
//   - accountID should be oss account
func (c *ChainClient) Authorize(accountID []byte) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	acc, err := types.NewAccountID(accountID)
	if err != nil {
		return "", errors.Wrap(err, "[NewAccountID]")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_Oss_authorize, *acc)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Oss_authorize, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Oss_authorize)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Oss_authorize, err)
	}

	return blockhash, nil
}

// CancelAuthorize cancels authorisation for an account
//   - accountID: account with cancelled authorisations
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) CancelAuthorize(accountID []byte) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_Oss_cancel_authorize, accountID)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Oss_cancel_authorize, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Oss_cancel_authorize)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Oss_cancel_authorize, err)
	}

	return blockhash, nil
}

// RegisterOss registered as oss role
//   - peerId: peer id
//   - domain: domain name, can be empty
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) RegisterOss(domain string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if len(domain) > int(MaxDomainNameLength) {
		return "", fmt.Errorf("register deoss: Domain name length cannot exceed %v characters", MaxDomainNameLength)
	}

	newcall, err := types.NewCall(c.metadata, ExtName_Oss_register, PeerId{}, types.NewBytes([]byte(domain)))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Oss_register, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Oss_register)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Oss_register, err)
	}

	return blockhash, nil
}

// UpdateOss update oss's peerId or domain
//   - peerId: peer id
//   - domain: domain name
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) UpdateOss(domain string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if len(domain) > int(MaxDomainNameLength) {
		return "", fmt.Errorf("update oss: domain name length cannot exceed %v", MaxDomainNameLength)
	}

	newcall, err := types.NewCall(c.metadata, ExtName_Oss_update, PeerId{}, types.NewBytes([]byte(domain)))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Oss_update, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Oss_update)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Oss_update, err)
	}

	return blockhash, nil
}

// DestroyOss destroys the oss role of the current account
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) DestroyOss() (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_Oss_destroy)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Oss_destroy, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Oss_destroy)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Oss_destroy, err)
	}

	return blockhash, nil
}
