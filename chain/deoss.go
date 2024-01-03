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
	"github.com/mr-tron/base58"
	"github.com/pkg/errors"
)

func (c *chainClient) QueryDeOSSInfo(accountID []byte) (pattern.OssInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data pattern.OssInfo

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.OSS, pattern.OSS, accountID)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.OSS, pattern.OSS, err)
		c.SetChainState(false)
		return data, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.OSS, pattern.OSS, err)
		c.SetChainState(false)
		return data, err
	}
	if !ok {
		return data, pattern.ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) QueryAllDeOSSInfo() ([]pattern.OssInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var result []pattern.OssInfo

	if !c.GetChainState() {
		return nil, pattern.ERR_RPC_CONNECTION
	}

	key := createPrefixedKey(pattern.OSS, pattern.OSS)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), pattern.OSS, pattern.OSS, err)
		c.SetChainState(false)
		return nil, err
	}

	set, err := c.api.RPC.State.QueryStorageAtLatest(keys)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), pattern.OSS, pattern.OSS, err)
		c.SetChainState(false)
		return nil, err
	}

	for _, elem := range set {
		for _, change := range elem.Changes {
			var data pattern.OssInfo
			if err := codec.Decode(change.StorageData, &data); err != nil {
				continue
			}
			result = append(result, data)
		}
	}
	return result, nil
}

func (c *chainClient) QueryAllDeOSSPeerId() ([]string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var result []string

	if !c.GetChainState() {
		return nil, pattern.ERR_RPC_CONNECTION
	}

	key := createPrefixedKey(pattern.OSS, pattern.OSS)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetKeysLatest: %v", c.GetCurrentRpcAddr(), pattern.OSS, pattern.OSS, err)
		c.SetChainState(false)
		return nil, err
	}

	set, err := c.api.RPC.State.QueryStorageAtLatest(keys)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] QueryStorageAtLatest: %v", c.GetCurrentRpcAddr(), pattern.OSS, pattern.OSS, err)
		c.SetChainState(false)
		return nil, err
	}

	for _, elem := range set {
		for _, change := range elem.Changes {
			var data pattern.OssInfo
			if err := codec.Decode(change.StorageData, &data); err != nil {
				continue
			}
			result = append(result, base58.Encode([]byte(string(data.Peerid[:]))))
		}
	}
	return result, nil
}

func (c *chainClient) QueryAuthorizedAccountIDs(accountID []byte) ([]types.AccountID, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data []types.AccountID

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.OSS, pattern.AUTHORITYLIST, accountID)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.OSS, pattern.AUTHORITYLIST, err)
		c.SetChainState(false)
		return data, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.OSS, pattern.AUTHORITYLIST, err)
		c.SetChainState(false)
		return data, err
	}
	if !ok {
		return data, pattern.ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) QueryAuthorizedAccounts(accountID []byte) ([]string, error) {
	acc, err := c.QueryAuthorizedAccountIDs(accountID)
	if err != nil {
		return nil, err
	}
	var result = make([]string, len(acc))
	for k, v := range acc {
		result[k], _ = utils.EncodePublicKeyAsCessAccount(v[:])
	}
	return result, nil
}

func (c *chainClient) RegisterDeOSS(peerId []byte, domain string) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		err         error
		txhash      string
		call        types.Call
		accountInfo types.AccountInfo
	)

	if !c.GetChainState() {
		return txhash, pattern.ERR_RPC_CONNECTION
	}

	var peerid pattern.PeerId
	if len(peerid) != len(peerId) {
		return txhash, errors.New("register deoss: invalid peerid")
	}
	for i := 0; i < len(peerid); i++ {
		peerid[i] = types.U8(peerId[i])
	}

	if len(domain) > pattern.MaxDomainNameLength {
		return txhash, fmt.Errorf("register deoss: Domain name length cannot exceed %v characters", pattern.MaxDomainNameLength)
	}

	err = utils.CheckDomain(domain)
	if err != nil {
		return txhash, errors.New("register deoss: invalid domain")
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_REGISTER, err)
		c.SetChainState(false)
		return txhash, err
	}

	call, err = types.NewCall(c.metadata, pattern.TX_OSS_REGISTER, peerid, types.NewBytes([]byte(domain)))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_REGISTER, err)
		c.SetChainState(false)
		return txhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_REGISTER, err)
		c.SetChainState(false)
		return txhash, err
	}

	if !ok {
		keyStr, _ := utils.NumsToByteStr(key, map[string]bool{})
		return txhash, fmt.Errorf(
			"chain rpc.state.GetStorageLatest[%v]: %v",
			keyStr,
			pattern.ERR_RPC_EMPTY_VALUE,
		)
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
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_REGISTER, err)
		c.SetChainState(false)
		return txhash, err
	}

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
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_REGISTER, err)
				c.SetChainState(false)
				return txhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_REGISTER, err)
			c.SetChainState(false)
			return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
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
				_, err = c.RetrieveEvent_Oss_OssRegister(status.AsInBlock)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) UpdateDeOSS(peerId string, domain string) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		err         error
		txhash      string
		call        types.Call
		accountInfo types.AccountInfo
	)

	if !c.GetChainState() {
		return txhash, pattern.ERR_RPC_CONNECTION
	}

	var peerid pattern.PeerId
	if len(peerid) != len(peerId) {
		return txhash, errors.New("update deoss: invalid peerid")
	}
	for i := 0; i < len(peerid); i++ {
		peerid[i] = types.U8(peerId[i])
	}

	if len(domain) > pattern.MaxSubmitedIdleFileMeta {
		return txhash, fmt.Errorf("register deoss: domain name length cannot exceed %v", pattern.MaxSubmitedIdleFileMeta)
	}

	err = utils.CheckDomain(domain)
	if err != nil {
		return txhash, errors.New("register deoss: invalid domain name")
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_UPDATE, err)
		c.SetChainState(false)
		return txhash, err
	}

	call, err = types.NewCall(c.metadata, pattern.TX_OSS_UPDATE, peerid, types.NewBytes([]byte(domain)))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_UPDATE, err)
		c.SetChainState(false)
		return txhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_UPDATE, err)
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

	ext := types.NewExtrinsic(call)

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_UPDATE, err)
		c.SetChainState(false)
		return txhash, err
	}

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
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_UPDATE, err)
				c.SetChainState(false)
				return txhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_UPDATE, err)
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
				_, err = c.RetrieveEvent_Oss_OssUpdate(status.AsInBlock)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) ExitDeOSS() (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		err         error
		txhash      string
		call        types.Call
		accountInfo types.AccountInfo
	)

	if !c.GetChainState() {
		return txhash, pattern.ERR_RPC_CONNECTION
	}

	call, err = types.NewCall(c.metadata, pattern.TX_OSS_DESTROY)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_DESTROY, err)
		c.SetChainState(false)
		return txhash, err
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_DESTROY, err)
		c.SetChainState(false)
		return txhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_DESTROY, err)
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

	ext := types.NewExtrinsic(call)

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_DESTROY, err)
		c.SetChainState(false)
		return txhash, err
	}

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
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_DESTROY, err)
				c.SetChainState(false)
				return txhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_DESTROY, err)
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
				_, err = c.RetrieveEvent_Oss_OssDestroy(status.AsInBlock)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) AuthorizeSpace(ossAccount string) (string, error) {
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

	puk, err := utils.ParsingPublickey(ossAccount)
	if err != nil {
		return txhash, errors.Wrap(err, "[ParsingPublickey]")
	}

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewAccountID]")
	}

	list, err := c.QueryAuthorizedAccounts(c.GetSignatureAccPulickey())
	if err != nil {
		if err.Error() != pattern.ERR_Empty {
			return txhash, errors.Wrap(err, "[QueryAuthorizedAccounts]")
		}
	} else {
		for _, v := range list {
			if v == ossAccount {
				return "", nil
			}
		}
	}

	call, err := types.NewCall(c.metadata, pattern.TX_OSS_AUTHORIZE, *acc)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_AUTHORIZE, err)
		c.SetChainState(false)
		return txhash, err
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_AUTHORIZE, err)
		c.SetChainState(false)
		return txhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_AUTHORIZE, err)
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

	ext := types.NewExtrinsic(call)

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_AUTHORIZE, err)
		c.SetChainState(false)
		return txhash, err
	}

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
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_AUTHORIZE, err)
				c.SetChainState(false)
				return txhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_AUTHORIZE, err)
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
				_, err = c.RetrieveEvent_Oss_Authorize(status.AsInBlock)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) UnAuthorizeSpace(oss_acc string) (string, error) {
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

	pubkey, err := utils.ParsingPublickey(oss_acc)
	if err != nil {
		return txhash, errors.Wrap(err, "[ParsingPublickey]")
	}

	call, err := types.NewCall(c.metadata, pattern.TX_OSS_UNAUTHORIZE)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_AUTHORIZE, err)
		c.SetChainState(false)
		return txhash, err
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, pubkey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_AUTHORIZE, err)
		c.SetChainState(false)
		return txhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_AUTHORIZE, err)
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

	ext := types.NewExtrinsic(call)

	// Sign the transaction
	err = ext.Sign(c.keyring, o)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_AUTHORIZE, err)
		c.SetChainState(false)
		return txhash, err
	}

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
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_AUTHORIZE, err)
				c.SetChainState(false)
				return txhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), pattern.TX_OSS_AUTHORIZE, err)
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
				_, err = c.RetrieveEvent_Oss_CancelAuthorize(status.AsInBlock)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}
