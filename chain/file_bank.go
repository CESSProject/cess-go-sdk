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
	"path/filepath"

	"github.com/AstaFrode/go-substrate-rpc-client/v4/types"
	"github.com/AstaFrode/go-substrate-rpc-client/v4/types/codec"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/pkg/errors"
)

// QueryDealMap query file storage order
//   - fid: file identification
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - StorageOrder: file storage order
//   - error: error message
func (c *ChainClient) QueryDealMap(fid string, block int32) (StorageOrder, error) {
	if !c.GetRpcState() {
		err := c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), FileBank, DealMap, ERR_RPC_CONNECTION.Error())
			return StorageOrder{}, err
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		data StorageOrder
		hash FileHash
	)

	if len(fid) != FileHashLen {
		return data, errors.New("invalid filehash")
	}

	for i := 0; i < len(hash); i++ {
		hash[i] = types.U8(fid[i])
	}

	param_hash, err := codec.Encode(hash)
	if err != nil {
		return data, errors.Wrap(err, "[Encode]")
	}

	key, err := types.CreateStorageKey(c.metadata, FileBank, DealMap, param_hash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), FileBank, DealMap, err)
		return data, err
	}
	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), FileBank, DealMap, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), FileBank, DealMap, err)
		return data, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), FileBank, DealMap, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryDealMap query file storage order
//   - fid: file identification
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - StorageOrderV1: file storage order
//   - error: error message
func (c *ChainClient) QueryDealMapV1(fid string, block int32) (StorageOrderV1, error) {
	if !c.GetRpcState() {
		err := c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), FileBank, DealMap, ERR_RPC_CONNECTION.Error())
			return StorageOrderV1{}, err
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		data StorageOrderV1
		hash FileHash
	)

	if len(fid) != FileHashLen {
		return data, errors.New("invalid filehash")
	}

	for i := 0; i < len(hash); i++ {
		hash[i] = types.U8(fid[i])
	}

	param_hash, err := codec.Encode(hash)
	if err != nil {
		return data, errors.Wrap(err, "[Encode]")
	}

	key, err := types.CreateStorageKey(c.metadata, FileBank, DealMap, param_hash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), FileBank, DealMap, err)
		return data, err
	}
	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), FileBank, DealMap, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), FileBank, DealMap, err)
		return data, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), FileBank, DealMap, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryDealMapList query file storage order list
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []StorageOrder: file storage order list
//   - error: error message
func (c *ChainClient) QueryDealMapList(block int32) ([]StorageOrder, error) {
	if !c.GetRpcState() {
		err := c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), FileBank, DealMap, ERR_RPC_CONNECTION.Error())
			return []StorageOrder{}, err
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var result []StorageOrder

	key := CreatePrefixedKey(FileBank, DealMap)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), FileBank, DealMap, err)
		c.SetRpcState(false)
		return nil, err
	}

	var set []types.StorageChangeSet
	if block < 0 {
		set, err = c.api.RPC.State.QueryStorageAtLatest(keys)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), FileBank, DealMap, err)
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
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), FileBank, DealMap, err)
			c.SetRpcState(false)
			return nil, err
		}
	}
	for _, elem := range set {
		for _, change := range elem.Changes {
			var data StorageOrder
			if err := codec.Decode(change.StorageData, &data); err != nil {
				continue
			}
			result = append(result, data)
		}
	}
	return result, nil
}

// QueryFile query file metadata
//   - fid: file identification
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - FileMetadata: file metadata
//   - error: error message
func (c *ChainClient) QueryFile(fid string, block int32) (FileMetadata, error) {
	if !c.GetRpcState() {
		err := c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), FileBank, File, ERR_RPC_CONNECTION.Error())
			return FileMetadata{}, err
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		data FileMetadata
		hash FileHash
	)

	if len(fid) != FileHashLen {
		return data, errors.New("invalid filehash")
	}

	for i := 0; i < len(hash); i++ {
		hash[i] = types.U8(fid[i])
	}

	param_hash, err := codec.Encode(hash)
	if err != nil {
		return data, errors.Wrap(err, "[Encode]")
	}

	key, err := types.CreateStorageKey(c.metadata, FileBank, File, param_hash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), FileBank, File, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), FileBank, File, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), FileBank, File, err)
		return data, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), FileBank, File, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryFile query file metadata
//   - fid: file identification
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - FileMetadataV1: file metadata
//   - error: error message
func (c *ChainClient) QueryFileV1(fid string, block int32) (FileMetadataV1, error) {
	if !c.GetRpcState() {
		err := c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), FileBank, File, ERR_RPC_CONNECTION.Error())
			return FileMetadataV1{}, err
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		data FileMetadataV1
		hash FileHash
	)

	if len(fid) != FileHashLen {
		return data, errors.New("invalid filehash")
	}

	for i := 0; i < len(hash); i++ {
		hash[i] = types.U8(fid[i])
	}

	param_hash, err := codec.Encode(hash)
	if err != nil {
		return data, errors.Wrap(err, "[Encode]")
	}

	key, err := types.CreateStorageKey(c.metadata, FileBank, File, param_hash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), FileBank, File, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), FileBank, File, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), FileBank, File, err)
		return data, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), FileBank, File, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryRestoralOrder query file restoral order
//   - fragmentHash: fragment hash
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - RestoralOrderInfo: restoral order info
//   - error: error message
func (c *ChainClient) QueryRestoralOrder(fragmentHash string, block int32) (RestoralOrderInfo, error) {
	if !c.GetRpcState() {
		err := c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), FileBank, RestoralOrder, ERR_RPC_CONNECTION.Error())
			return RestoralOrderInfo{}, err
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		data RestoralOrderInfo
		hash FileHash
	)

	if len(fragmentHash) != FileHashLen {
		return data, errors.New("invalid fragment hash")
	}

	for i := 0; i < len(hash); i++ {
		hash[i] = types.U8(fragmentHash[i])
	}

	param_hash, err := codec.Encode(hash)
	if err != nil {
		return data, errors.Wrap(err, "[Encode]")
	}

	key, err := types.CreateStorageKey(c.metadata, FileBank, RestoralOrder, param_hash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), FileBank, RestoralOrder, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), FileBank, RestoralOrder, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), FileBank, RestoralOrder, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryAllRestoralOrder query all file restoral order
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []RestoralOrderInfo: all restoral order info
//   - error: error message
func (c *ChainClient) QueryAllRestoralOrder(block int32) ([]RestoralOrderInfo, error) {
	if !c.GetRpcState() {
		err := c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), FileBank, RestoralOrder, ERR_RPC_CONNECTION.Error())
			return []RestoralOrderInfo{}, err
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var result []RestoralOrderInfo

	key := CreatePrefixedKey(FileBank, RestoralOrder)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), FileBank, RestoralOrder, err)
		c.SetRpcState(false)
		return nil, err
	}
	var set []types.StorageChangeSet
	if block < 0 {
		set, err = c.api.RPC.State.QueryStorageAtLatest(keys)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), FileBank, RestoralOrder, err)
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
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAt: %v", c.GetCurrentRpcAddr(), FileBank, RestoralOrder, err)
			c.SetRpcState(false)
			return nil, err
		}
	}

	for _, elem := range set {
		for _, change := range elem.Changes {
			var data RestoralOrderInfo
			if err := codec.Decode(change.StorageData, &data); err != nil {
				continue
			}
			result = append(result, data)
		}
	}
	return result, nil
}

// QueryUserHoldFileList query user's all files
//   - accountID: user account
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []UserFileSliceInfo: file list
//   - error: error message
func (c *ChainClient) QueryUserHoldFileList(accountID []byte, block int32) ([]UserFileSliceInfo, error) {
	if !c.GetRpcState() {
		err := c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), FileBank, UserHoldFileList, ERR_RPC_CONNECTION.Error())
			return []UserFileSliceInfo{}, err
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data []UserFileSliceInfo

	acc, err := types.NewAccountID(accountID)
	if err != nil {
		return nil, errors.Wrap(err, "[NewAccountID]")
	}

	owner, err := codec.Encode(*acc)
	if err != nil {
		return nil, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(c.metadata, FileBank, UserHoldFileList, owner)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), FileBank, UserHoldFileList, err)
		return nil, err
	}
	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), FileBank, UserHoldFileList, err)
			c.SetRpcState(false)
			return nil, err
		}
		if !ok {
			return data, ERR_RPC_EMPTY_VALUE
		}
		return data, nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return nil, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), FileBank, UserHoldFileList, err)
		c.SetRpcState(false)
		return nil, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryUserFidList query user's all fid
//   - accountID: user account
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - []string: all fid
//   - error: error message
func (c *ChainClient) QueryUserFidList(accountID []byte, block int32) ([]string, error) {
	if !c.GetRpcState() {
		err := c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), FileBank, UserHoldFileList, ERR_RPC_CONNECTION.Error())
			return []string{}, err
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data []UserFileSliceInfo

	acc, err := types.NewAccountID(accountID)
	if err != nil {
		return nil, errors.Wrap(err, "[NewAccountID]")
	}

	owner, err := codec.Encode(*acc)
	if err != nil {
		return nil, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(c.metadata, FileBank, UserHoldFileList, owner)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), FileBank, UserHoldFileList, err)
		return nil, err
	}
	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), FileBank, UserHoldFileList, err)
			c.SetRpcState(false)
			return nil, err
		}
		if !ok {
			return []string{}, ERR_RPC_EMPTY_VALUE
		}
		var value = make([]string, len(data))
		for i := 0; i < len(data); i++ {
			value[i] = string(data[i].Filehash[:])
		}
		return value, nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return nil, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), FileBank, UserHoldFileList, err)
		c.SetRpcState(false)
		return nil, err
	}
	if !ok {
		return []string{}, ERR_RPC_EMPTY_VALUE
	}
	var value = make([]string, len(data))
	for i := 0; i < len(data); i++ {
		value[i] = string(data[i].Filehash[:])
	}
	return value, nil
}

// PlaceStorageOrder place an order for storage file
//   - fid: file identification
//   - file_name: file name
//   - territory_name: territory name
//   - segment: segment info
//   - owner: account of the file owner
//   - filesize: file size
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) PlaceStorageOrder(fid, file_name, territory_name string, segment []SegmentDataInfo, owner []byte, file_size uint64) (string, error) {
	var err error
	var segmentList = make([]SegmentList, len(segment))
	var user UserBrief

	for i := 0; i < len(segment); i++ {
		hash := filepath.Base(segment[i].SegmentHash)
		for k := 0; k < len(hash); k++ {
			segmentList[i].SegmentHash[k] = types.U8(hash[k])
		}
		segmentList[i].FragmentHash = make([]FileHash, len(segment[i].FragmentHash))
		for j := 0; j < len(segment[i].FragmentHash); j++ {
			hash := filepath.Base(segment[i].FragmentHash[j])
			for k := 0; k < len(hash); k++ {
				segmentList[i].FragmentHash[j][k] = types.U8(hash[k])
			}
		}
	}

	acc, err := types.NewAccountID(owner)
	if err != nil {
		return "", err
	}
	user.User = *acc
	user.FileName = types.NewBytes([]byte(file_name))
	user.TerriortyName = types.NewBytes([]byte(territory_name))
	return c.UploadDeclaration(fid, segmentList, user, file_size)
}

// GenerateStorageOrder generate a file storage order
//   - fid: file identification
//   - segment: segment info
//   - user: UserBrief
//   - filename: file name
//   - filesize: file size
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) UploadDeclaration(fid string, segment []SegmentList, user UserBrief, filesize uint64) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var hash FileHash

	if len(fid) != FileHashLen {
		return "", errors.New("invalid filehash")
	}
	if filesize <= 0 {
		return "", errors.New("invalid filesize")
	}
	for i := 0; i < len(hash); i++ {
		hash[i] = types.U8(fid[i])
	}

	newcall, err := types.NewCall(c.metadata, ExtName_FileBank_upload_declaration, hash, segment, user, types.NewU128(*new(big.Int).SetUint64(filesize)))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_upload_declaration, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_FileBank_upload_declaration)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_upload_declaration, err)
	}

	return blockhash, nil
}

// DeleteFile delete a bucket for owner
//   - owner: file owner account
//   - fid: file identification
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - if you are not the owner, the owner account must be authorised to you
func (c *ChainClient) DeleteFile(owner []byte, fid string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if len(fid) != FileHashLen {
		return "", errors.New("invalid fid")
	}

	acc, err := types.NewAccountID(owner)
	if err != nil {
		return "", errors.Wrap(err, "[NewAccountID]")
	}

	var fhash FileHash
	for i := 0; i < len(fid); i++ {
		fhash[i] = types.U8(fid[i])
	}

	newcall, err := types.NewCall(c.metadata, ExtName_FileBank_delete_file, *acc, fhash)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_delete_file, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_FileBank_delete_file)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_delete_file, err)
	}

	return blockhash, nil
}

// TransferReport is used by miners to report that a file has been transferred
//   - index: index of the file fragment
//   - fid: file identification
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - for storage miner use only
func (c *ChainClient) TransferReport(index uint8, fid string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if index <= 0 || int(index) > (DataShards+ParShards) {
		return "", errors.New("invalid index")
	}

	var fhash FileHash

	for j := 0; j < len(fid); j++ {
		fhash[j] = types.U8(fid[j])
	}

	newcall, err := types.NewCall(c.metadata, ExtName_FileBank_transfer_report, types.NewU8(index), fhash)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_transfer_report, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_FileBank_transfer_report)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_transfer_report, err)
	}

	return blockhash, nil
}

// GenerateRestoralOrder generate restoral orders for file fragment
//   - fid: file identification
//   - fragmentHash: fragment hash
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - for storage miner use only
func (c *ChainClient) GenerateRestoralOrder(fid, fragmentHash string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var rooth FileHash
	var fragh FileHash

	if len(fid) != FileHashLen {
		return "", errors.New("invalid file hash")
	}

	if len(fragmentHash) != FileHashLen {
		return "", errors.New("invalid fragment hash")
	}

	for i := 0; i < len(fid); i++ {
		rooth[i] = types.U8(fid[i])
	}

	for i := 0; i < len(fragmentHash); i++ {
		fragh[i] = types.U8(fragmentHash[i])
	}

	newcall, err := types.NewCall(c.metadata, ExtName_FileBank_generate_restoral_order, rooth, fragh)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_generate_restoral_order, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_FileBank_generate_restoral_order)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_generate_restoral_order, err)
	}

	return blockhash, nil
}

// ClaimRestoralOrder claim a restoral order
//   - fragmentHash: fragment hash
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - for storage miner use only
func (c *ChainClient) ClaimRestoralOrder(fragmentHash string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if len(fragmentHash) != FileHashLen {
		return "", errors.New("invalid fragment hash")
	}

	var fragh FileHash
	for i := 0; i < len(fragmentHash); i++ {
		fragh[i] = types.U8(fragmentHash[i])
	}

	newcall, err := types.NewCall(c.metadata, ExtName_FileBank_claim_restoral_order, fragh)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_claim_restoral_order, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_FileBank_claim_restoral_order)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_claim_restoral_order, err)
	}

	return blockhash, nil
}

// ClaimRestoralNoExistOrder claim the restoral order of an exited storage miner
//   - puk: storage miner account
//   - fid: file identification
//   - fragmentHash: fragment hash
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - for storage miner use only
func (c *ChainClient) ClaimRestoralNoExistOrder(puk []byte, fid, fragmentHash string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return "", errors.Wrap(err, "[NewAccountID]")
	}

	var rooth FileHash
	var fragh FileHash

	if len(fid) != FileHashLen {
		return "", errors.New("invalid file hash")
	}

	if len(fragmentHash) != FileHashLen {
		return "", errors.New("invalid fragment hash")
	}

	for i := 0; i < len(fid); i++ {
		rooth[i] = types.U8(fid[i])
	}

	for i := 0; i < len(fragmentHash); i++ {
		fragh[i] = types.U8(fragmentHash[i])
	}

	newcall, err := types.NewCall(c.metadata, ExtName_FileBank_claim_restoral_noexist_order, *acc, rooth, fragh)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_claim_restoral_noexist_order, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_FileBank_claim_restoral_noexist_order)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_claim_restoral_noexist_order, err)
	}

	return blockhash, nil
}

// RestoralOrderComplete submits the restored completed order
//   - fragmentHash: fragment hash
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - for storage miner use only
func (c *ChainClient) RestoralOrderComplete(fragmentHash string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var fragh FileHash

	if len(fragmentHash) != FileHashLen {
		return "", errors.New("invalid fragment hash")
	}

	for i := 0; i < len(fragmentHash); i++ {
		fragh[i] = types.U8(fragmentHash[i])
	}

	newcall, err := types.NewCall(c.metadata, ExtName_FileBank_restoral_order_complete, fragh)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_restoral_order_complete, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_FileBank_restoral_order_complete)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_restoral_order_complete, err)
	}

	return blockhash, nil
}

// CertIdleSpace authenticates idle file to the chain
//   - spaceProofInfo: space proof info
//   - teeSignWithAcc: tee sign with account
//   - teeSign: tee sign
//   - teePuk: tee work public key
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - for storage miner use only
func (c *ChainClient) CertIdleSpace(spaceProofInfo SpaceProofInfo, teeSignWithAcc, teeSign types.Bytes, teePuk WorkerPublicKey) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_FileBank_cert_idle_space, spaceProofInfo, teeSignWithAcc, teeSign, teePuk)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_cert_idle_space, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_FileBank_cert_idle_space)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_cert_idle_space, err)
	}

	return blockhash, nil
}

// ReplaceIdleSpace replaces idle files with service files
//   - spaceProofInfo: space proof info
//   - teeSignWithAcc: tee sign with account
//   - teeSign: tee sign
//   - teePuk: tee work public key
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - for storage miner use only
func (c *ChainClient) ReplaceIdleSpace(spaceProofInfo SpaceProofInfo, teeSignWithAcc, teeSign types.Bytes, teePuk WorkerPublicKey) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_FileBank_replace_idle_space, spaceProofInfo, teeSignWithAcc, teeSign, teePuk)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_replace_idle_space, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_FileBank_replace_idle_space)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_replace_idle_space, err)
	}

	return blockhash, nil
}

// CalculateReport report file tag calculation completed
//   - teeSig: tee sign
//   - tagSigInfo: tag sig info
//
// Return:
//   - string: block hash
//   - error: error message
//
// Note:
//   - for storage miner use only
func (c *ChainClient) CalculateReport(teeSig types.Bytes, tagSigInfo TagSigInfo) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_FileBank_calculate_report, teeSig, tagSigInfo)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_calculate_report, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_FileBank_calculate_report)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_calculate_report, err)
	}

	return blockhash, nil
}

// TerritoryFileDelivery transfer files to another territory
//   - user: file owner account
//   - fid: file id
//   - target_territory: transfer to the target territory
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) TerritoryFileDelivery(user []byte, fid string, target_territory string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	acc, err := types.NewAccountID(user)
	if err != nil {
		return "", errors.Wrap(err, "[NewAccountID]")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_FileBank_territory_file_delivery, *acc, types.NewBytes([]byte(fid)), types.NewBytes([]byte(target_territory)))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_territory_file_delivery, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_FileBank_territory_file_delivery)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_FileBank_territory_file_delivery, err)
	}

	return blockhash, nil
}
