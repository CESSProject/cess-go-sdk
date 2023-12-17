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
	"github.com/CESSProject/cess-go-sdk/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

func (c *chainClient) QueryTeeWorkerMap(puk []byte) (pattern.TeeWorkerMap, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.TeeWorkerMap

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
	data.EndPoint = string(teeWorkerInfo.EndPoint)
	data.WorkAccount, _ = utils.EncodePublicKeyAsCessAccount(teeWorkerInfo.WorkAccount[:])
	data.TeeType = uint8(teeWorkerInfo.TeeType)
	data.PeerId = []byte(string(teeWorkerInfo.PeerId[:]))
	if teeWorkerInfo.BondStash.HasValue() {
		ok, val := teeWorkerInfo.BondStash.Unwrap()
		if ok {
			data.StashAccount, _ = utils.EncodePublicKeyAsCessAccount(val[:])
		}
	}
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
		return nil, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return nil, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return nil, pattern.ERR_RPC_EMPTY_VALUE
	}

	return []byte(string(data[:])), nil
}

func (c *chainClient) QueryAllTeeWorkerMap() ([]pattern.TeeWorkerMap, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var list []pattern.TeeWorkerMap

	if !c.GetChainState() {
		return list, pattern.ERR_RPC_CONNECTION
	}

	key := createPrefixedKey(pattern.TEEWORKER, pattern.TEEWORKERMAP)
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
			var teeWorker pattern.TeeWorkerMap
			if err := codec.Decode(change.StorageData, &teeWorker); err != nil {
				fmt.Println(err)
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
		results[k].EndPoint = string(v.EndPoint[:])
		results[k].PeerId = []byte(string(v.PeerId[:]))
		results[k].TeeType = uint8(v.TeeType)
		results[k].WorkAccount, _ = utils.EncodePublicKeyAsCessAccount(v.WorkAccount[:])
		if v.BondStash.HasValue() {
			ok, acc := v.BondStash.Unwrap()
			if ok {
				results[k].StashAccount, err = utils.EncodePublicKeyAsCessAccount(acc[:])
				if err != nil {
					return results, err
				}
			}
		}
	}
	return results, nil
}
