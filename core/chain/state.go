/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"
	"log"

	"github.com/CESSProject/sdk-go/core/utils"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/centrifuge/go-substrate-rpc-client/xxhash"
	"github.com/pkg/errors"
)

// QueryBlockHeight
func (c *chainClient) QueryBlockHeight(hash string) (uint32, error) {
	defer func() {
		recover()
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

// ExtractAccountPublickey
func (c *chainClient) ExtractAccountPuk(account string) ([]byte, error) {
	if account != "" {
		return utils.ParsingPublickey(account)
	}
	return c.keyring.PublicKey, nil
}

// QueryNodeSynchronizationSt
func (c *chainClient) QueryNodeSynchronizationSt() (bool, error) {
	if !c.IsChainClientOk() {
		return false, ERR_RPC_CONNECTION
	}
	h, err := c.api.RPC.System.Health()
	if err != nil {
		return false, err
	}
	return h.IsSyncing, nil
}

// QueryNodeConnectionSt
func (c *chainClient) QueryNodeConnectionSt() bool {
	return c.GetChainState()
}

// Get miner information on the chain
func (c *chainClient) QueryStorageMiner(pkey []byte) (MinerInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data MinerInfo

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		SMINER,
		MINERITEMS,
		pkey,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// Get all miner information on the cess chain
func (c *chainClient) GetAllStorageMiner() ([]types.AccountID, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data []types.AccountID

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		SMINER,
		ALLMINER,
	)
	if err != nil {
		return nil, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return nil, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return nil, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// Query file meta info
func (c *chainClient) GetFileMetaInfo(fid string) (FileMetaInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var (
		data FileMetaInfo
		hash FileHash
	)

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	if len(hash) != len(fid) {
		return data, errors.New("invalid filehash")
	}

	for i := 0; i < len(hash); i++ {
		hash[i] = types.U8(fid[i])
	}

	b, err := codec.Encode(hash)
	if err != nil {
		return data, errors.Wrap(err, "[Encode]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		FILEBANK,
		FILE,
		b,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) GetCessAccount() (string, error) {
	return utils.EncodePublicKeyAsCessAccount(c.keyring.PublicKey)
}

func (c *chainClient) GetAccountInfo(pkey []byte) (types.AccountInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data types.AccountInfo

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	acc, err := types.NewAccountID(pkey)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	b, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		SYSTEM,
		ACCOUNT,
		b,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) GetState(pubkey []byte) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data types.Bytes

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return "", ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	acc, err := types.NewAccountID(pubkey)
	if err != nil {
		return "", errors.Wrap(err, "[NewAccountID]")
	}

	b, err := codec.Encode(*acc)
	if err != nil {
		return "", errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		OSS,
		OSS,
		b,
	)
	if err != nil {
		return "", errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return "", errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return "", ERR_RPC_EMPTY_VALUE
	}

	return string(data), nil
}

func (c *chainClient) GetBucketInfo(owner_pkey []byte, name string) (BucketInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data BucketInfo

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	acc, err := types.NewAccountID(owner_pkey)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	owner, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	name_byte, err := codec.Encode(name)
	if err != nil {
		return data, errors.Wrap(err, "[Encode]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		FILEBANK,
		BUCKET,
		owner,
		name_byte,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) GetBucketList(owner_pkey []byte) ([]types.Bytes, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data []types.Bytes

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	acc, err := types.NewAccountID(owner_pkey)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	owner, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(
		c.metadata,
		FILEBANK,
		BUCKETLIST,
		owner,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) GetStorageOrder(roothash string) (StorageOrder, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data StorageOrder
	var hash FileHash

	if len(hash) != len(roothash) {
		return data, errors.New("invalid filehash")
	}

	for i := 0; i < len(hash); i++ {
		hash[i] = types.U8(roothash[i])
	}

	b, err := codec.Encode(hash)
	if err != nil {
		return data, errors.Wrap(err, "[Encode]")
	}

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		FILEBANK,
		DEALMAP,
		b,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) QueryPendingReplacements(owner_pkey []byte) (types.U32, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data types.U32

	acc, err := types.NewAccountID(owner_pkey)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	owner, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		FILEBANK,
		PENDINGREPLACE,
		owner,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) QueryUserSpaceInfo(pubkey []byte) (UserSpaceInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data UserSpaceInfo

	acc, err := types.NewAccountID(pubkey)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	owner, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		STORAGEHANDLER,
		USERSPACEINFO,
		owner,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) QuerySpacePricePerGib() (string, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data types.U128

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return "", ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		STORAGEHANDLER,
		UNITPRICE,
	)
	if err != nil {
		return "", errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return "", errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return "", ERR_RPC_EMPTY_VALUE
	}

	return fmt.Sprintf("%v", data), nil
}

func (c *chainClient) QueryNetSnapShot() (NetSnapShot, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data NetSnapShot

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return data, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		NETSNAPSHOT,
		CHALLENGESNAPSHOT,
	)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}

	return data, nil
}

func (c *chainClient) QueryTeePodr2Puk() ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data TeePodr2Pk

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return nil, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		TEEWORKER,
		TEEPODR2Pk,
	)
	if err != nil {
		return nil, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return nil, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return nil, ERR_RPC_EMPTY_VALUE
	}

	return []byte(string(data[:])), nil
}

func (c *chainClient) QueryTeeWorker(pubkey []byte) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			println(utils.RecoverError(err))
		}
	}()
	var data TeeWorkerInfo

	acc, err := types.NewAccountID(pubkey)
	if err != nil {
		return nil, errors.Wrap(err, "[NewAccountID]")
	}

	owner, err := codec.Encode(*acc)
	if err != nil {
		return nil, errors.Wrap(err, "[EncodeToBytes]")
	}

	if !c.IsChainClientOk() {
		c.SetChainState(false)
		return nil, ERR_RPC_CONNECTION
	}
	c.SetChainState(true)

	key, err := types.CreateStorageKey(
		c.metadata,
		TEEWORKER,
		TEEWORKERMAP,
		owner,
	)
	if err != nil {
		return nil, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return nil, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return nil, ERR_RPC_EMPTY_VALUE
	}

	return []byte(string(data.Peer_id[:])), nil
}

func (c *chainClient) QueryTeeWorkerList() ([]TeeWorkerInfo, error) {
	var list []TeeWorkerInfo
	key := createPrefixedKey(TEEWORKER, TEEWORKERMAP)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		return list, errors.Wrap(err, "[GetKeysLatest]")
	}
	set, err := c.api.RPC.State.QueryStorageAtLatest(keys)
	if err != nil {
		return list, errors.Wrap(err, "[QueryStorageAtLatest]")
	}
	for _, elem := range set {
		for _, change := range elem.Changes {
			var teeWorker TeeWorkerInfo
			if err := codec.Decode(change.StorageData, &teeWorker); err != nil {
				log.Println(err)
				continue
			}
			list = append(list, teeWorker)
		}
	}
	return list, nil
}

// Pallert
const (
	_FILEBANK = "FileBank"
	_SYSTEM   = "System"
	_CACHER   = "Cacher"
)

// Chain state
const (
	// System
	_SYSTEM_ACCOUNT = "Account"
	_SYSTEM_EVENTS  = "Events"
	// FileMap
	_FILEMAP_FILEMETA = "File"
	// Miner
	_CACHER_CACHER = "Cachers"
)

type CacherInfo struct {
	Acc       types.AccountID
	Ip        types.Bytes
	BytePrice types.U128
}

func (c *chainClient) GetCachers() ([]CacherInfo, error) {
	var list []CacherInfo
	key := createPrefixedKey(_CACHER_CACHER, _CACHER)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		return list, errors.Wrap(err, "get cachers info error")
	}
	set, err := c.api.RPC.State.QueryStorageAtLatest(keys)
	if err != nil {
		return list, errors.Wrap(err, "get cachers info error")
	}
	for _, elem := range set {
		for _, change := range elem.Changes {
			var cacher CacherInfo
			if err := codec.Decode(change.StorageData, &cacher); err != nil {
				//logger.Uld.Sugar().Error("get cachers info error,hash:", err)
				log.Println(err)
				continue
			}
			list = append(list, cacher)
		}
	}
	return list, nil
}

func createPrefixedKey(method, prefix string) []byte {
	return append(xxhash.New128([]byte(prefix)).Sum(nil), xxhash.New128([]byte(method)).Sum(nil)...)
}
