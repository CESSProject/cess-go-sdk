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

func (c *chainClient) QueryTeeInfoList() ([]pattern.TeeWorkerInfo, error) {
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
		return list, errors.Wrap(err, "[GetKeysLatest]")
	}

	set, err := c.api.RPC.State.QueryStorageAtLatest(keys)
	if err != nil {
		return list, errors.Wrap(err, "[QueryStorageAtLatest]")
	}

	for _, elem := range set {
		for _, change := range elem.Changes {
			var teeWorker pattern.TeeWorkerInfo
			if err := codec.Decode(change.StorageData, &teeWorker); err != nil {
				fmt.Println(err)
				continue
			}
			list = append(list, teeWorker)
		}
	}
	return list, nil
}

func (c *chainClient) QueryTeePeerID(puk []byte) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.TeeWorkerInfo

	if !c.GetChainState() {
		return nil, pattern.ERR_RPC_CONNECTION
	}

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return nil, errors.Wrap(err, "[NewAccountID]")
	}

	owner, err := codec.Encode(*acc)
	if err != nil {
		return nil, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.TEEWORKER, pattern.TEEWORKERMAP, owner)
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

	return []byte(string(data.PeerId[:])), nil
}

func (c *chainClient) QueryTeeWorkerList() ([]pattern.TeeWorkerSt, error) {
	teelist, err := c.QueryTeeInfoList()
	if err != nil {
		return nil, err
	}
	var results = make([]pattern.TeeWorkerSt, len(teelist))
	for k, v := range teelist {
		results[k].Node_key = []byte(string(v.NodeKey.NodePublickey[:]))
		results[k].Peer_id = []byte(string(v.PeerId[:]))
		results[k].Controller_account, err = utils.EncodePublicKeyAsCessAccount(v.ControllerAccount[:])
		if err != nil {
			return results, err
		}
		results[k].Stash_account, err = utils.EncodePublicKeyAsCessAccount(v.StashAccount[:])
		if err != nil {
			return results, err
		}
	}
	return results, nil
}
