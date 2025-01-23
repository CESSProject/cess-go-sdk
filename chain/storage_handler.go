/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"
	"log"

	"github.com/AstaFrode/go-substrate-rpc-client/v4/types"
	"github.com/AstaFrode/go-substrate-rpc-client/v4/types/codec"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/pkg/errors"
)

// QueryUnitPrice query price per GiB space
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: price per GiB space
//   - error: error message
func (c *ChainClient) QueryUnitPrice(block int32) (string, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return "", fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), StorageHandler, UnitPrice, ERR_RPC_CONNECTION.Error())
	}

	var data types.U128

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
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), StorageHandler, TotalIdleSpace, ERR_RPC_CONNECTION.Error())
	}

	var data types.U128

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
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), StorageHandler, TotalServiceSpace, ERR_RPC_CONNECTION.Error())
	}

	var data types.U128

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
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), StorageHandler, PurchasedSpace, ERR_RPC_CONNECTION.Error())
	}

	var data types.U128

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
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return TerritoryInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), StorageHandler, Territory, ERR_RPC_CONNECTION.Error())
	}

	var data TerritoryInfo

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
func (c *ChainClient) QueryConsignment(token types.H256, block int32) (ConsignmentInfo, error) {
	c.rwlock.RLock()
	defer func() {
		c.rwlock.RUnlock()
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetRpcState() {
		return ConsignmentInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), StorageHandler, Consignment, ERR_RPC_CONNECTION.Error())
	}

	var data ConsignmentInfo

	param1, err := codec.Encode(token)
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
//   - days: the validity period of the territory, in days
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) MintTerritory(gib_count uint32, territory_name string, days uint32) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if gib_count == 0 {
		return "", errors.New("[MintTerritory] invalid gib_count")
	}

	if days == 0 {
		return "", errors.New("[MintTerritory] invalid days")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_StorageHandler_mint_territory, types.NewU32(gib_count), types.NewBytes([]byte(territory_name)), types.NewU32(days))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_mint_territory, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_StorageHandler_mint_territory)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_mint_territory, err)
	}

	return blockhash, nil
}

// ExpandingTerritory expanding the territory size
//   - territory_name: territory name
//   - gib_count: size to be expanded
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) ExpandingTerritory(territory_name string, gib_count uint32) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if gib_count == 0 {
		return "", errors.New("[ExpandingTerritory] invalid gib_count")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_StorageHandler_expanding_territory, types.NewBytes([]byte(territory_name)), types.NewU32(gib_count))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_expanding_territory, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_StorageHandler_expanding_territory)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_expanding_territory, err)
	}

	return blockhash, nil
}

// RenewalTerritory renewal of territory validity period
//   - territory_name: territory name
//   - days_count: renewal days
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) RenewalTerritory(territory_name string, days_count uint32) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if days_count == 0 {
		return "", errors.New("[RenewalTerritory] invalid days_count")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_StorageHandler_renewal_territory, types.NewBytes([]byte(territory_name)), types.NewU32(days_count))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_renewal_territory, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_StorageHandler_renewal_territory)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_renewal_territory, err)
	}

	return blockhash, nil
}

// ReactivateTerritory reactivate expired territories
//   - territory_name: territory name
//   - days_count: number of days activated
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) ReactivateTerritory(territory_name string, days_count uint32) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if days_count == 0 {
		return "", errors.New("[ReactivateTerritory] invalid days_count")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_StorageHandler_reactivate_territory, types.NewBytes([]byte(territory_name)), types.NewU32(days_count))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_reactivate_territory, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_StorageHandler_reactivate_territory)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_reactivate_territory, err)
	}

	return blockhash, nil
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
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_StorageHandler_territory_consignment, types.NewBytes([]byte(territory_name)))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_territory_consignment, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_StorageHandler_territory_consignment)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_territory_consignment, err)
	}

	return blockhash, nil
}

// CancelConsignment cancel consignment territory
//   - territory_name: territory name
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) CancelConsignment(territory_name string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_StorageHandler_cancel_consignment, types.NewBytes([]byte(territory_name)))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_cancel_consignment, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_StorageHandler_cancel_consignment)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_cancel_consignment, err)
	}

	return blockhash, nil
}

// BuyConsignment purchase territories for consignment
//   - token: territory key
//   - territory_name: renamed territory name
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) BuyConsignment(token types.H256, territory_name string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if len(territory_name) <= 0 {
		return "", errors.New("territory name is empty")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_StorageHandler_buy_consignment, token, types.NewBytes([]byte(territory_name)))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_buy_consignment, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_StorageHandler_buy_consignment)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_buy_consignment, err)
	}

	return blockhash, nil
}

// CancelPurchaseAction cancel purchase territories for consignment
//   - token: territory key
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) CancelPurchaseAction(token types.H256) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_StorageHandler_cancel_purchase_action, token)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_cancel_purchase_action, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_StorageHandler_cancel_purchase_action)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_StorageHandler_cancel_purchase_action, err)
	}

	return blockhash, nil
}
