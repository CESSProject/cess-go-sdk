/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"
	"log"
	"math/big"

	"github.com/AstaFrode/go-substrate-rpc-client/v4/types"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/pkg/errors"
)

// QueryTotalIssuance query the total amount of token issuance
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: the total amount of token issuance
//   - error: error message
func (c *ChainClient) QueryTotalIssuance(block int32) (string, error) {
	if !c.GetRpcState() {
		err := c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Balances, TotalIssuance, ERR_RPC_CONNECTION.Error())
			return "", err
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	key, err := types.CreateStorageKey(c.metadata, Balances, TotalIssuance)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Balances, TotalIssuance, err)
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Balances, TotalIssuance, err)
			c.SetRpcState(false)
			return "", err
		}
		if !ok {
			return "", ERR_RPC_EMPTY_VALUE
		}
		return data.String(), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Balances, TotalIssuance, err)
		return "", err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Balances, TotalIssuance, err)
		c.SetRpcState(false)
		return "", err
	}
	if !ok {
		return "0", nil
	}
	if data.String() == "" {
		return "0", nil
	}
	return data.String(), nil
}

// QueryInactiveIssuance query the amount of inactive token issuance
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: the amount of inactive token issuance
//   - error: error message
func (c *ChainClient) QueryInactiveIssuance(block int32) (string, error) {
	if !c.GetRpcState() {
		err := c.ReconnectRpc()
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Balances, InactiveIssuance, ERR_RPC_CONNECTION.Error())
			return "", err
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U128

	key, err := types.CreateStorageKey(c.metadata, Balances, InactiveIssuance)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Balances, InactiveIssuance, err)
		return "", err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Balances, InactiveIssuance, err)
			c.SetRpcState(false)
			return "", err
		}
		if !ok {
			return "", ERR_RPC_EMPTY_VALUE
		}
		return data.String(), nil
	}

	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetBlockHash: %v", c.GetCurrentRpcAddr(), Balances, InactiveIssuance, err)
		return "", err
	}

	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Balances, InactiveIssuance, err)
		c.SetRpcState(false)
		return "", err
	}
	if !ok {
		return "0", nil
	}
	if data.String() == "" {
		return "0", nil
	}
	return data.String(), nil
}

// TransferToken transfers to other accounts
//   - dest: target account
//   - amount: transfer amount, It is the smallest unit. If you need to use CESS as the unit, you need to add 18 zeros.
//     For example, if you transfer 1 CESS, you need to fill in "1000000000000000000"
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) TransferToken(dest string, amount string) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	pubkey, err := utils.ParsingPublickey(dest)
	if err != nil {
		return "", errors.Wrapf(err, "[ParsingPublickey]")
	}

	address, err := types.NewMultiAddressFromAccountID(pubkey)
	if err != nil {
		return "", errors.Wrapf(err, "[NewMultiAddressFromAccountID]")
	}

	amount_bg, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return "", errors.New("[TransferToken] invalid amount")
	}

	newcall, err := types.NewCall(c.metadata, ExtName_Balances_transferKeepAlive, address, types.NewUCompact(amount_bg))
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Balances_transferKeepAlive, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Balances_transferKeepAlive)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Balances_transferKeepAlive, err)
	}
	return blockhash, nil
}
