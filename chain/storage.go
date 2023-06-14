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

func (c *ChainSDK) QuerySpacePricePerGib() (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	if !c.GetChainState() {
		return "", pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.STORAGEHANDLER, pattern.UNITPRICE)
	if err != nil {
		return "", errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return "", errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return "", pattern.ERR_RPC_EMPTY_VALUE
	}

	return fmt.Sprintf("%v", data), nil
}

func (c *ChainSDK) QueryUserSpaceInfo(puk []byte) (pattern.UserSpaceInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.UserSpaceInfo

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

	key, err := types.CreateStorageKey(c.metadata, pattern.STORAGEHANDLER, pattern.USERSPACEINFO, owner)
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

func (c *ChainSDK) QueryUserSpaceSt(puk []byte) (pattern.UserSpaceSt, error) {
	var userSpaceSt pattern.UserSpaceSt
	spaceinfo, err := c.QueryUserSpaceInfo(puk)
	if err != nil {
		return userSpaceSt, err
	}
	userSpaceSt.Start = uint32(spaceinfo.Start)
	userSpaceSt.Deadline = uint32(spaceinfo.Deadline)
	userSpaceSt.TotalSpace = spaceinfo.TotalSpace.String()
	userSpaceSt.UsedSpace = spaceinfo.UsedSpace.String()
	userSpaceSt.RemainingSpace = spaceinfo.RemainingSpace.String()
	userSpaceSt.LockedSpace = spaceinfo.LockedSpace.String()
	userSpaceSt.State = string(spaceinfo.State)
	return userSpaceSt, nil
}
