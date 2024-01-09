/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

func (c *chainClient) QuerySpacePricePerGib() (string, error) {
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.STORAGEHANDLER, pattern.UNITPRICE, err)
		c.SetChainState(false)
		return "", err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.STORAGEHANDLER, pattern.UNITPRICE, err)
		c.SetChainState(false)
		return "", err
	}
	if !ok {
		return "", pattern.ERR_RPC_EMPTY_VALUE
	}

	return fmt.Sprintf("%v", data), nil
}

func (c *chainClient) QueryUserSpaceInfo(puk []byte) (pattern.UserSpaceInfo, error) {
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.STORAGEHANDLER, pattern.USERSPACEINFO, err)
		c.SetChainState(false)
		return data, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.STORAGEHANDLER, pattern.USERSPACEINFO, err)
		c.SetChainState(false)
		return data, err
	}
	if !ok {
		return data, pattern.ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) QueryUserSpaceSt(puk []byte) (pattern.UserSpaceSt, error) {
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

func (c *chainClient) QueryTotalIdleSpace() (uint64, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	if !c.GetChainState() {
		return 0, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.STORAGEHANDLER, pattern.TOTALIDLESPACE)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.STORAGEHANDLER, pattern.TOTALIDLESPACE, err)
		c.SetChainState(false)
		return 0, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.STORAGEHANDLER, pattern.TOTALIDLESPACE, err)
		c.SetChainState(false)
		return 0, err
	}
	if !ok {
		return 0, pattern.ERR_RPC_EMPTY_VALUE
	}
	return data.Uint64(), nil
}

func (c *chainClient) QueryTotalServiceSpace() (uint64, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	if !c.GetChainState() {
		return 0, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.STORAGEHANDLER, pattern.TOTALSERVICESPACE)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.STORAGEHANDLER, pattern.TOTALSERVICESPACE, err)
		c.SetChainState(false)
		return 0, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.STORAGEHANDLER, pattern.TOTALSERVICESPACE, err)
		c.SetChainState(false)
		return 0, err
	}
	if !ok {
		return 0, pattern.ERR_RPC_EMPTY_VALUE
	}
	return data.Uint64(), nil
}

func (c *chainClient) QueryPurchasedSpace() (uint64, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	if !c.GetChainState() {
		return 0, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.STORAGEHANDLER, pattern.PURCHASEDSPACE)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.STORAGEHANDLER, pattern.PURCHASEDSPACE, err)
		c.SetChainState(false)
		return 0, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.STORAGEHANDLER, pattern.PURCHASEDSPACE, err)
		c.SetChainState(false)
		return 0, err
	}
	if !ok {
		return 0, pattern.ERR_RPC_EMPTY_VALUE
	}
	return data.Uint64(), nil
}

func (c *chainClient) BuySpace(count uint32) (string, error) {
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

	if count == 0 {
		return "", errors.New("[BuySpace] invalid count")
	}

	if !c.GetChainState() {
		return txhash, pattern.ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, pattern.TX_STORAGE_BUYSPACE, types.NewU32(count))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_BUYSPACE, err)
		c.SetChainState(false)
		return txhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_BUYSPACE, err)
		c.SetChainState(false)
		return txhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_BUYSPACE, err)
		c.SetChainState(false)
		return txhash, err
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_BUYSPACE, err)
		c.SetChainState(false)
		return txhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), pattern.ERR_RPC_PRIORITYTOOLOW) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return txhash, errors.Wrap(err, "[Sign]")
			}
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_BUYSPACE, err)
				c.SetChainState(false)
				return txhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_BUYSPACE, err)
			c.SetChainState(false)
			return txhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				txhash = status.AsInBlock.Hex()
				_, err = c.RetrieveEvent_StorageHandler_BuySpace(status.AsInBlock)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) ExpansionSpace(count uint32) (string, error) {
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

	if count == 0 {
		return "", errors.New("[ExpansionSpace] invalid count")
	}

	if !c.GetChainState() {
		return txhash, pattern.ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, pattern.TX_STORAGE_EXPANSIONSPACE, types.NewU32(count))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_EXPANSIONSPACE, err)
		c.SetChainState(false)
		return txhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_EXPANSIONSPACE, err)
		c.SetChainState(false)
		return txhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_EXPANSIONSPACE, err)
		c.SetChainState(false)
		return txhash, err
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_EXPANSIONSPACE, err)
		c.SetChainState(false)
		return txhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), pattern.ERR_RPC_PRIORITYTOOLOW) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return txhash, errors.Wrap(err, "[Sign]")
			}
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_EXPANSIONSPACE, err)
				c.SetChainState(false)
				return txhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_EXPANSIONSPACE, err)
			c.SetChainState(false)
			return txhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				txhash = status.AsInBlock.Hex()
				_, err = c.RetrieveEvent_StorageHandler_ExpansionSpace(status.AsInBlock)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) RenewalSpace(days uint32) (string, error) {
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

	if days == 0 {
		return "", errors.New("[RenewalSpace] invalid days")
	}

	if !c.GetChainState() {
		return txhash, pattern.ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, pattern.TX_STORAGE_RENEWALSPACE, types.NewU32(days))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_RENEWALSPACE, err)
		c.SetChainState(false)
		return txhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_RENEWALSPACE, err)
		c.SetChainState(false)
		return txhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_RENEWALSPACE, err)
		c.SetChainState(false)
		return txhash, err
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

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_RENEWALSPACE, err)
		c.SetChainState(false)
		return txhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), pattern.ERR_RPC_PRIORITYTOOLOW) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return txhash, errors.Wrap(err, "[Sign]")
			}
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_RENEWALSPACE, err)
				c.SetChainState(false)
				return txhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_STORAGE_RENEWALSPACE, err)
			c.SetChainState(false)
			return txhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				txhash = status.AsInBlock.Hex()
				_, err = c.RetrieveEvent_StorageHandler_RenewalSpace(status.AsInBlock)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}
