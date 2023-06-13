/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"
	"log"
	"time"

	"github.com/CESSProject/sdk-go/core/event"
	"github.com/CESSProject/sdk-go/core/pattern"
	"github.com/CESSProject/sdk-go/core/utils"
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

func (c *ChainSDK) GenerateRestoralOrder(rootHash, fragmentHash string) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	if !c.GetChainState() {
		return txhash, pattern.ERR_RPC_CONNECTION
	}

	var rooth pattern.FileHash
	var fragh pattern.FileHash

	if len(rootHash) != len(rooth) {
		return txhash, errors.New("invalid root hash")
	}

	if len(fragmentHash) != len(fragh) {
		return txhash, errors.New("invalid fragment hash")
	}

	for i := 0; i < len(rootHash); i++ {
		rooth[i] = types.U8(rootHash[i])
	}

	for i := 0; i < len(fragmentHash); i++ {
		fragh[i] = types.U8(fragmentHash[i])
	}

	call, err := types.NewCall(c.metadata, pattern.TX_FILEBANK_GENRESTOREORDER, rooth, fragh)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewCall]")
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		return txhash, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return txhash, pattern.ERR_RPC_EMPTY_VALUE
	}

	o := types.SignatureOptions{
		BlockHash:          c.genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        c.genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(accountInfo.Nonce)),
		SpecVersion:        c.runtimeVersion.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: c.runtimeVersion.TransactionVersion,
	}

	ext := types.NewExtrinsic(call)

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		return txhash, errors.Wrap(err, "[Sign]")
	}

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.timeForBlockOut)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				events := event.EventRecords{}
				txhash, _ = codec.EncodeToHex(status.AsInBlock)
				h, err := c.api.RPC.State.GetStorageRaw(c.keyEvents, status.AsInBlock)
				if err != nil {
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil || len(events.FileBank_RestoralOrderComplete) > 0 {
					return txhash, nil
				}
				return txhash, errors.New(pattern.ERR_Failed)
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}
