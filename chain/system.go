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

	"github.com/CESSProject/cess-go-sdk/core/event"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

// QueryNodeSynchronizationSt
func (c *chainClient) QueryNodeSynchronizationSt() (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if !c.GetChainState() {
		return false, pattern.ERR_RPC_CONNECTION
	}
	h, err := c.api.RPC.System.Health()
	if err != nil {
		return false, err
	}
	return h.IsSyncing, nil
}

// QueryBlockHeight
func (c *chainClient) QueryBlockHeight(hash string) (uint32, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if hash != "" {
		var h types.Hash
		err := codec.DecodeFromHex(hash, &h)
		if err != nil {
			return 0, err
		}
		block, err := c.api.RPC.Chain.GetBlock(h)
		if err != nil {
			return 0, errors.Wrap(err, "[GetBlock]")
		}
		return uint32(block.Block.Header.Number), nil
	}

	block, err := c.api.RPC.Chain.GetBlockLatest()
	if err != nil {
		return 0, errors.Wrap(err, "[GetBlockLatest]")
	}
	return uint32(block.Block.Header.Number), nil
}

// QueryAccountInfo
func (c *chainClient) QueryAccountInfo(puk []byte) (types.AccountInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.AccountInfo

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return data, errors.Wrap(err, "[NewAccountID]")
	}

	b, err := codec.Encode(*acc)
	if err != nil {
		return data, errors.Wrap(err, "[EncodeToBytes]")
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, b)
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

func (c *chainClient) SysProperties() (pattern.SysProperties, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data pattern.SysProperties
	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, pattern.RPC_SYS_Properties)
	return data, err
}

func (c *chainClient) SysChain() (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.Text
	if !c.GetChainState() {
		return "", pattern.ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, pattern.RPC_SYS_Chain)
	return string(data), err
}

func (c *chainClient) SyncState() (pattern.SysSyncState, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data pattern.SysSyncState
	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, pattern.RPC_SYS_SyncState)
	return data, err
}

func (c *chainClient) SysVersion() (string, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.Text
	if !c.GetChainState() {
		return "", pattern.ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, pattern.RPC_SYS_Version)
	return string(data), err
}

func (c *chainClient) NetListening() (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.Bool
	if !c.GetChainState() {
		return false, pattern.ERR_RPC_CONNECTION
	}
	err := c.api.Client.Call(&data, pattern.RPC_NET_Listening)
	return bool(data), err
}

func (c *chainClient) TransferToken(dest string, amount uint64) (string, string, error) {
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
		return txhash, "", pattern.ERR_RPC_CONNECTION
	}

	pubkey, err := utils.ParsingPublickey(dest)
	if err != nil {
		return "", "", errors.Wrapf(err, "[ParsingPublickey]")
	}

	address, err := types.NewMultiAddressFromAccountID(pubkey)
	if err != nil {
		return "", "", errors.Wrapf(err, "[NewAddressFromAccountID]")
	}

	call, err := types.NewCall(c.metadata, pattern.TX_BALANCES_FORCETRANSFER, address, types.NewUCompactFromUInt(amount))
	if err != nil {
		return txhash, "", errors.Wrap(err, "[NewCall]")
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		return txhash, "", errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, "", errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return txhash, "", pattern.ERR_RPC_EMPTY_VALUE
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
		return txhash, "", errors.Wrap(err, "[Sign]")
	}

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		c.SetChainState(false)
		return txhash, "", errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
	}
	defer sub.Unsubscribe()

	timeout := time.NewTimer(c.packingTime)
	defer timeout.Stop()

	for {
		select {
		case status := <-sub.Chan():
			if status.IsInBlock {
				events := event.EventRecords{}
				txhash, _ = codec.EncodeToHex(status.AsInBlock)
				h, err := c.api.RPC.State.GetStorageRaw(c.keyEvents, status.AsInBlock)
				if err != nil {
					return txhash, "", errors.Wrap(err, "[GetStorageRaw]")
				}
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil {
					return txhash, "", nil
				}
				if len(events.Balances_Transfer) > 0 {
					for _, v := range events.Balances_Transfer {
						if utils.CompareSlice(v.From[:], c.GetSignatureAccPulickey()) {
							to, _ := utils.EncodePublicKeyAsCessAccount(v.To[:])
							return txhash, to, nil
						}
					}
				}
				return txhash, "", errors.New(pattern.ERR_Failed)
			}
		case err = <-sub.Err():
			return txhash, "", errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, "", pattern.ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) DecodeEventNameFromBlock(block uint64) ([]string, error) {
	blockHash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return nil, err
	}
	events, err := c.eventRetriever.GetEvents(blockHash)
	if err != nil {
		return nil, err
	}
	var result = make([]string, len(events))
	for k, v := range events {
		result[k] = v.Name
		if v.Name == "fileBank.ClaimRestoralOrder" {
			for _, vv := range v.Fields {
				fmt.Printf("%s", vv.Name)
				fmt.Printf("  %v\n", vv.Value)
			}
		}
	}
	return result, nil
}

func (c *chainClient) DecodeEventNameFromBlockhash(blockhash types.Hash) ([]string, error) {
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return nil, err
	}
	var result = make([]string, len(events))
	for k, v := range events {
		result[k] = v.Name
	}
	return result, nil
}

func (c *chainClient) FindEventFromBlock(block uint64, eventname string) ([]string, error) {
	blockHash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return nil, err
	}
	events, err := c.eventRetriever.GetEvents(blockHash)
	if err != nil {
		return nil, err
	}
	var result = make([]string, len(events))
	for k, v := range events {
		result[k] = v.Name
	}
	return result, nil
}
