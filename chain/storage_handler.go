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

	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

// QueryUnitPrice query price per GiB space
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: price per GiB space
//   - error: error message
func (c *ChainClient) QueryUnitPrice(block int32) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	if !c.GetRpcState() {
		return "", ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, StorageHandler, UnitPrice)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), StorageHandler, UnitPrice, err)
		c.SetRpcState(false)
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), StorageHandler, UnitPrice, err)
			c.SetRpcState(false)
			return "", err
		}
		if !ok {
			return "", ERR_RPC_EMPTY_VALUE
		}

		return fmt.Sprintf("%v", data), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return "", err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), StorageHandler, UnitPrice, err)
		c.SetRpcState(false)
		return "", err
	}
	if !ok {
		return "", ERR_RPC_EMPTY_VALUE
	}

	return fmt.Sprintf("%v", data), nil
}

// QueryTotalIdleSpace query the size of all idle space
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint64: the size of all idle space
//   - error: error message
func (c *ChainClient) QueryTotalIdleSpace(block int32) (uint64, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	if !c.GetRpcState() {
		return 0, ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, StorageHandler, TotalIdleSpace)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), StorageHandler, TotalIdleSpace, err)
		c.SetRpcState(false)
		return 0, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), StorageHandler, TotalIdleSpace, err)
			c.SetRpcState(false)
			return 0, err
		}
		if !ok {
			return 0, ERR_RPC_EMPTY_VALUE
		}
		return data.Uint64(), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return 0, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), StorageHandler, TotalIdleSpace, err)
		c.SetRpcState(false)
		return 0, err
	}
	if !ok {
		return 0, ERR_RPC_EMPTY_VALUE
	}
	return data.Uint64(), nil
}

// QueryTotalServiceSpace query the size of all service space
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint64: the size of all service space
//   - error: error message
func (c *ChainClient) QueryTotalServiceSpace(block int32) (uint64, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	if !c.GetRpcState() {
		return 0, ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, StorageHandler, TotalServiceSpace)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), StorageHandler, TotalServiceSpace, err)
		return 0, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), StorageHandler, TotalServiceSpace, err)
			c.SetRpcState(false)
			return 0, err
		}
		if !ok {
			return 0, ERR_RPC_EMPTY_VALUE
		}
		return data.Uint64(), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return 0, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), StorageHandler, TotalServiceSpace, err)
		c.SetRpcState(false)
		return 0, err
	}
	if !ok {
		return 0, ERR_RPC_EMPTY_VALUE
	}
	return data.Uint64(), nil
}

// QueryPurchasedSpace query all purchased space size
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint64: all purchased space size
//   - error: error message
func (c *ChainClient) QueryPurchasedSpace(block int32) (uint64, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	if !c.GetRpcState() {
		return 0, ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, StorageHandler, PurchasedSpace)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), StorageHandler, PurchasedSpace, err)
		return 0, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), StorageHandler, PurchasedSpace, err)
			c.SetRpcState(false)
			return 0, err
		}
		if !ok {
			return 0, ERR_RPC_EMPTY_VALUE
		}
		return data.Uint64(), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return 0, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), StorageHandler, PurchasedSpace, err)
		c.SetRpcState(false)
		return 0, err
	}
	if !ok {
		return 0, ERR_RPC_EMPTY_VALUE
	}
	return data.Uint64(), nil
}

// QueryTerritory query territory info
//   - accountId: account id
//   - name: territory name
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - TerritoryInfo: territory info
//   - error: error message
func (c *ChainClient) QueryTerritory(accountId []byte, name string, block int32) (TerritoryInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data TerritoryInfo

	if !c.GetRpcState() {
		return data, ERR_RPC_CONNECTION
	}

	param2, err := codec.Encode(types.NewBytes([]byte(name)))
	if err != nil {
		return data, errors.New("invalid account id")
	}

	key, err := types.CreateStorageKey(c.metadata, StorageHandler, Territory, accountId, param2)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), StorageHandler, Territory, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), StorageHandler, Territory, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), StorageHandler, Territory, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// QueryConsignment query consignment info
//   - token: territory key
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - ConsignmentInfo: consignment info
//   - error: error message
func (c *ChainClient) QueryConsignment(token string, block int32) (ConsignmentInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data ConsignmentInfo

	if !c.GetRpcState() {
		return data, ERR_RPC_CONNECTION
	}

	if len(token) != TerritoryKeyLen {
		return data, errors.New("invalid territory key")
	}
	var token_chain types.H256
	for i := 0; i < TerritoryKeyLen; i++ {
		token_chain[i] = byte(token[i])
	}

	param1, err := codec.Encode(token_chain)
	if err != nil {
		return data, errors.New("invalid territory key")
	}

	key, err := types.CreateStorageKey(c.metadata, StorageHandler, Consignment, param1)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), StorageHandler, Consignment, err)
		return data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), StorageHandler, Consignment, err)
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
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), StorageHandler, Consignment, err)
		c.SetRpcState(false)
		return data, err
	}
	if !ok {
		return data, ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

// MintTerritory purchase a territory
//   - gib_count: territory size
//   - territory_name: territory name
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) MintTerritory(gib_count uint32, territory_name string) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if gib_count == 0 {
		return "", errors.New("[MintTerritory] invalid gib_count")
	}

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, TX_StorageHandler_MintTerritory, types.NewU32(gib_count), types.NewBytes([]byte(territory_name)))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_MintTerritory, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_MintTerritory, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_MintTerritory, err)
		c.SetRpcState(false)
		return blockhash, err
	}
	if !ok {
		return blockhash, ERR_RPC_EMPTY_VALUE
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
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_MintTerritory, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_MintTerritory, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_MintTerritory, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_StorageHandler_mint_territory, StorageHandlerMintTerritory, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}

// ExpandingTerritory expanding the territory size
//   - territory_name: territory name
//   - gib_count: size to be expanded
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) ExpandingTerritory(territory_name string, gib_count uint32) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if gib_count == 0 {
		return "", errors.New("[ExpandingTerritory] invalid gib_count")
	}

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, TX_StorageHandler_ExpandingTerritory, types.NewBytes([]byte(territory_name)), types.NewU32(gib_count))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_ExpandingTerritory, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_ExpandingTerritory, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_ExpandingTerritory, err)
		c.SetRpcState(false)
		return blockhash, err
	}
	if !ok {
		return blockhash, ERR_RPC_EMPTY_VALUE
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
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_ExpandingTerritory, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_ExpandingTerritory, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_ExpandingTerritory, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_StorageHandler_expanding_territory, StorageHandlerExpansionTerritory, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}

// RenewalTerritory renewal of territory validity period
//   - territory_name: territory name
//   - days_count: renewal days
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) RenewalTerritory(territory_name string, days_count uint32) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if days_count == 0 {
		return "", errors.New("[RenewalTerritory] invalid days_count")
	}

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, TX_StorageHandler_RenewalTerritory, types.NewBytes([]byte(territory_name)), types.NewU32(days_count))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_RenewalTerritory, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_RenewalTerritory, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_RenewalTerritory, err)
		c.SetRpcState(false)
		return blockhash, err
	}
	if !ok {
		return blockhash, ERR_RPC_EMPTY_VALUE
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
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_RenewalTerritory, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_RenewalTerritory, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_RenewalTerritory, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_StorageHandler_renewal_territory, StorageHandlerRenewalTerritory, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}

// ReactivateTerritory reactivate expired territories
//   - territory_name: territory name
//   - days_count: number of days activated
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) ReactivateTerritory(territory_name string, days_count uint32) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if days_count == 0 {
		return "", errors.New("[ReactivateTerritory] invalid days_count")
	}

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, TX_StorageHandler_ReactivateTerritory, types.NewBytes([]byte(territory_name)), types.NewU32(days_count))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_ReactivateTerritory, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_ReactivateTerritory, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_ReactivateTerritory, err)
		c.SetRpcState(false)
		return blockhash, err
	}
	if !ok {
		return blockhash, ERR_RPC_EMPTY_VALUE
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
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_ReactivateTerritory, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_ReactivateTerritory, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_ReactivateTerritory, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_StorageHandler_reactivate_territory, StorageHandlerReactivateTerritory, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}

// TerritoryConsignment consignment territory
//   - territory_name: territory name
//
// Return:
//   - string: block hash
//   - error: error message
//
// Tip:
//   - The territory must be in an active state
//   - Remaining lease term greater than 1 day
func (c *ChainClient) TerritoryConsignment(territory_name string) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, TX_StorageHandler_TerritoryConsignment, types.NewBytes([]byte(territory_name)))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_TerritoryConsignment, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_TerritoryConsignment, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_TerritoryConsignment, err)
		c.SetRpcState(false)
		return blockhash, err
	}
	if !ok {
		return blockhash, ERR_RPC_EMPTY_VALUE
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
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_TerritoryConsignment, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_TerritoryConsignment, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_TerritoryConsignment, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_StorageHandler_territory_consignment, StorageHandlerConsignment, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}

// CancelConsignment cancel consignment territory
//   - territory_name: territory name
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) CancelConsignment(territory_name string) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	call, err := types.NewCall(c.metadata, TX_StorageHandler_CancelConsignment, types.NewBytes([]byte(territory_name)))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_CancelConsignment, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_CancelConsignment, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_CancelConsignment, err)
		c.SetRpcState(false)
		return blockhash, err
	}
	if !ok {
		return blockhash, ERR_RPC_EMPTY_VALUE
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
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_CancelConsignment, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_CancelConsignment, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_CancelConsignment, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_StorageHandler_cancel_consignment, StorageHandlerCancleConsignment, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}

// BuyConsignment purchase territories for consignment
//   - token: territory key
//   - territory_name: territory name
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) BuyConsignment(token string, territory_name string) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	if len(territory_name) <= 0 {
		return blockhash, errors.New("territory name is empty")
	}

	var token_chain types.H256
	for i := 0; i < TerritoryKeyLen; i++ {
		token_chain[i] = byte(token[i])
	}

	call, err := types.NewCall(c.metadata, TX_StorageHandler_BuyConsignment, token_chain, types.NewBytes([]byte(territory_name)))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_BuyConsignment, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_BuyConsignment, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_BuyConsignment, err)
		c.SetRpcState(false)
		return blockhash, err
	}
	if !ok {
		return blockhash, ERR_RPC_EMPTY_VALUE
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
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_BuyConsignment, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_BuyConsignment, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_BuyConsignment, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_StorageHandler_buy_consignment, StorageHandlerBuyConsignment, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}

// CancelPurchaseAction cancel purchase territories for consignment
//   - token: territory key
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) CancelPurchaseAction(token string) (string, error) {
	c.lock.Lock()
	defer func() {
		c.lock.Unlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var (
		blockhash   string
		accountInfo types.AccountInfo
	)

	if !c.GetRpcState() {
		return blockhash, ERR_RPC_CONNECTION
	}

	var token_chain types.H256
	for i := 0; i < TerritoryKeyLen; i++ {
		token_chain[i] = byte(token[i])
	}

	call, err := types.NewCall(c.metadata, TX_StorageHandler_CancelPurchaseAction, token_chain)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_CancelPurchaseAction, err)
		return blockhash, err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_CancelPurchaseAction, err)
		return blockhash, err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_CancelPurchaseAction, err)
		c.SetRpcState(false)
		return blockhash, err
	}
	if !ok {
		return blockhash, ERR_RPC_EMPTY_VALUE
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
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_CancelPurchaseAction, err)
		return blockhash, err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		if strings.Contains(err.Error(), ERR_PriorityIsTooLow) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return blockhash, errors.Wrap(err, "[Sign]")
			}
			<-c.txTicker.C
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_CancelPurchaseAction, err)
				c.SetRpcState(false)
				return blockhash, err
			}
		} else {
			err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_StorageHandler_CancelPurchaseAction, err)
			c.SetRpcState(false)
			return blockhash, err
		}
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				err = c.RetrieveEvent(status.AsInBlock, ExtName_StorageHandler_cancel_purchase_action, StorageHandlerCancelPurchaseAction, c.signatureAcc)
				return blockhash, err
			}
		case err = <-sub.Err():
			return blockhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, ERR_RPC_TIMEOUT
		}
	}
}
