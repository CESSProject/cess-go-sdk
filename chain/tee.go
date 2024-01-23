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

func (c *chainClient) QueryTeeWorkerMap(puk []byte) (pattern.TeeWorkerInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.TeeWorkerInfo

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	owner, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.TEEWORKER, pattern.TEEWORKERMAP, owner)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.TEEWORKER, pattern.TEEWORKERMAP, err)
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.TEEWORKER, pattern.TEEWORKERMAP, err)
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, pattern.ERR_RPC_EMPTY_VALUE
	}

	return data, nil
}

func (c *chainClient) QueryTeeInfo(puk []byte) (pattern.TeeInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.TeeInfo

	teeWorkerInfo, err := c.QueryTeeWorkerMap(puk)
	if err != nil {
		return data, err
	}
	data.Pubkey = string(teeWorkerInfo.Pubkey[:])
	data.EcdhPubkey = string(teeWorkerInfo.EcdhPubkey[:])
	data.Version = uint32(teeWorkerInfo.Version)
	data.LastUpdated = uint64(teeWorkerInfo.LastUpdated)
	if teeWorkerInfo.StashAccount.HasValue() {
		if ok, puk := teeWorkerInfo.StashAccount.Unwrap(); ok {
			data.StashAccount, _ = utils.EncodePublicKeyAsCessAccount(puk[:])
		}
	}
	if teeWorkerInfo.AttestationProvider.HasValue() {
		if ok, val := teeWorkerInfo.AttestationProvider.Unwrap(); ok {
			data.AttestationProvider = uint8(val)
		}
	}
	data.ConfidenceLevel = uint8(teeWorkerInfo.ConfidenceLevel)
	data.Features = make([]uint32, len(teeWorkerInfo.Features))
	for i := 0; i < len(teeWorkerInfo.Features); i++ {
		data.Features[i] = uint32(teeWorkerInfo.Features[i])
	}
	data.WorkerRole = uint8(teeWorkerInfo.Role)
	return data, nil
}

func (c *chainClient) QueryTeePodr2Puk() ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.TeePodr2Pk

	if !c.GetChainState() {
		return nil, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.TEEWORKER, pattern.TEEPODR2PK)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.TEEWORKER, pattern.TEEPODR2PK, err)
		return nil, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.TEEWORKER, pattern.TEEPODR2PK, err)
		return nil, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return nil, pattern.ERR_RPC_EMPTY_VALUE
	}

	return []byte(string(data[:])), nil
}

func (c *chainClient) QueryAllTeeWorkerMap() ([]pattern.TeeWorkerInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var list []pattern.TeeWorkerInfo

	if !c.GetChainState() {
		return list, pattern.ERR_RPC_CONNECTION
	}

	key := createPrefixedKey(pattern.TEEWORKER, pattern.TEEWORKERMAP)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), pattern.TEEWORKER, pattern.TEEWORKERMAP, err)
		return list, errors.Wrap(err, "[GetKeysLatest]")
	}

	set, err := c.api.RPC.State.QueryStorageAtLatest(keys)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), pattern.TEEWORKER, pattern.TEEWORKERMAP, err)
		return list, errors.Wrap(err, "[QueryStorageAtLatest]")
	}

	for _, elem := range set {
		for _, change := range elem.Changes {
			var teeWorker pattern.TeeWorkerInfo
			if err := codec.Decode(change.StorageData, &teeWorker); err != nil {
				log.Println(err)
				continue
			}
			list = append(list, teeWorker)
		}
	}
	return list, nil
}

func (c *chainClient) QueryAllTeeInfo() ([]pattern.TeeInfo, error) {
	teelist, err := c.QueryAllTeeWorkerMap()
	if err != nil {
		return nil, err
	}
	var results = make([]pattern.TeeInfo, len(teelist))
	for k, v := range teelist {
		results[k].Pubkey = string(v.Pubkey[:])
		results[k].EcdhPubkey = string(v.EcdhPubkey[:])
		results[k].Version = uint32(v.Version)
		results[k].LastUpdated = uint64(v.LastUpdated)
		if v.StashAccount.HasValue() {
			if ok, puk := v.StashAccount.Unwrap(); ok {
				results[k].StashAccount, _ = utils.EncodePublicKeyAsCessAccount(puk[:])
			}
		}
		if v.AttestationProvider.HasValue() {
			if ok, val := v.AttestationProvider.Unwrap(); ok {
				results[k].AttestationProvider = uint8(val)
			}
		}
		results[k].ConfidenceLevel = uint8(v.ConfidenceLevel)
		results[k].Features = make([]uint32, len(v.Features))
		for i := 0; i < len(v.Features); i++ {
			results[k].Features[i] = uint32(v.Features[i])
		}
		results[k].WorkerRole = uint8(v.Role)
	}
	return results, nil
}

func (c *chainClient) QueryTeeWorkEndpoint(workPuk pattern.WorkerPublicKey) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.Text

	if !c.GetChainState() {
		return "", pattern.ERR_RPC_CONNECTION
	}

	val, err := codec.Encode(workPuk)
	if err != nil {
		return "", errors.Wrap(err, "[Encode]")
	}
	key, err := types.CreateStorageKey(c.metadata, pattern.TEEWORKER, pattern.TEEEndpoints, val)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.TEEWORKER, pattern.TEEEndpoints, err)
		return "", errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.TEEWORKER, pattern.TEEEndpoints, err)
		return "", errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return "", pattern.ERR_RPC_EMPTY_VALUE
	}

	return string(data), nil
}
