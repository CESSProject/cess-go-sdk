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

	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

// QueryTotalIssuance query the total amount of token issuance
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - string: the total amount of token issuance
//   - error: error message
func (c *ChainClient) QueryTotalIssuance(block int) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.U128

	if !c.GetRpcState() {
		return "", ERR_RPC_CONNECTION
	}

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
func (c *ChainClient) QueryInactiveIssuance(block int) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.U128

	if !c.GetRpcState() {
		return "", ERR_RPC_CONNECTION
	}

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
//   - amount: transfer amount
//
// Return:
//   - string: block hash
//   - string: target account
//   - error: error message
func (c *ChainClient) TransferToken(dest string, amount uint64) (string, string, error) {
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
		return blockhash, "", ERR_RPC_CONNECTION
	}

	pubkey, err := utils.ParsingPublickey(dest)
	if err != nil {
		return blockhash, "", errors.Wrapf(err, "[ParsingPublickey]")
	}

	address, err := types.NewMultiAddressFromAccountID(pubkey)
	if err != nil {
		return blockhash, "", errors.Wrapf(err, "[NewMultiAddressFromAccountID]")
	}

	call, err := types.NewCall(c.metadata, TX_Balances_Transfer, address, types.NewUCompactFromUInt(amount))
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), TX_Balances_Transfer, err)
		return blockhash, "", err
	}

	ext := types.NewExtrinsic(call)

	key, err := types.CreateStorageKey(c.metadata, System, Account, c.keyring.PublicKey)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), TX_Balances_Transfer, err)
		return blockhash, "", err
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), TX_Balances_Transfer, err)
		c.SetRpcState(false)
		return blockhash, "", err
	}
	if !ok {
		return blockhash, "", ERR_RPC_EMPTY_VALUE
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
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] Sign: %v", c.GetCurrentRpcAddr(), TX_Balances_Transfer, err)
		return blockhash, "", err
	}

	<-c.txTicker.C

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitAndWatchExtrinsic: %v", c.GetCurrentRpcAddr(), TX_Balances_Transfer, err)
		c.SetRpcState(false)
		return blockhash, "", err
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				blockhash = status.AsInBlock.Hex()
				_, err = c.RetrieveEvent_Balances_Transfer(status.AsInBlock)
				return blockhash, dest, err
			}
		case err = <-sub.Err():
			return blockhash, "", errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return blockhash, "", ERR_RPC_TIMEOUT
		}
	}
}
