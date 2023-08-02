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
	"strconv"
	"time"

	"github.com/CESSProject/cess-go-sdk/core/event"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

// QueryStorageMiner
func (c *chainClient) QueryStorageMiner(puk []byte) (pattern.MinerInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.MinerInfo

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SMINER, pattern.MINERITEMS, puk)
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

func (c *chainClient) QueryStorageMiner_V2(puk []byte) (pattern.MinerInfo_V2, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.MinerInfo_V2

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SMINER, pattern.MINERITEMS, puk)
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

// QuerySminerList
func (c *chainClient) QuerySminerList() ([]types.AccountID, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data []types.AccountID

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SMINER, pattern.ALLMINER)
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
	return data, nil
}

// QueryMinerRewards
func (c *chainClient) QueryStorageNodeReward(puk []byte) (pattern.MinerReward, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.MinerReward

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SMINER, pattern.REWARDMAP, puk)
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

func (c *chainClient) QuaryStorageNodeRewardInfo(puk []byte) (pattern.RewardsType, error) {
	var reward pattern.RewardsType
	rewards, err := c.QueryStorageNodeReward(puk)
	if err != nil {
		return reward, err
	}

	reward.Total = rewards.TotalReward.String()
	reward.Claimed = rewards.RewardIssued.String()
	reward.Available = rewards.CurrentlyAvailableReward.String()

	return reward, nil
}

func (c *chainClient) RegisterOrUpdateSminer(peerId []byte, earnings string, pledge uint64) (string, string, error) {
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
		pubkey      []byte
		minerinfo   pattern.MinerInfo
		acc         *types.AccountID
		call        types.Call
		accountInfo types.AccountInfo
	)

	if !c.GetChainState() {
		return txhash, earnings, pattern.ERR_RPC_CONNECTION
	}

	c.SetSdkName(pattern.Name_Sminer)

	var peerid pattern.PeerId
	if len(peerid) != len(peerId) {
		return txhash, earnings, fmt.Errorf("invalid peerid: %v", peerId)
	}

	for i := 0; i < len(peerid); i++ {
		peerid[i] = types.U8(peerId[i])
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		return txhash, earnings, errors.Wrap(err, "[CreateStorageKey]")
	}

	minerinfo, err = c.QueryStorageMiner(c.keyring.PublicKey)
	if err != nil {
		if err.Error() != pattern.ERR_Empty {
			return txhash, earnings, err
		}
	} else {
		if !utils.CompareSlice([]byte(string(minerinfo.PeerId[:])), peerId) {
			txhash, err = c.updateSminerPeerId(key, peerid)
			if err != nil {
				return txhash, earnings, err
			}
		}
		acc, _ := utils.EncodePublicKeyAsCessAccount(minerinfo.BeneficiaryAcc[:])
		if earnings != "" {
			if acc != earnings {
				puk, err := utils.ParsingPublickey(earnings)
				if err != nil {
					return txhash, acc, err
				}
				txhash, err = c.updateEarningsAcc(key, puk)
				return txhash, earnings, err
			}
		}
		return "", acc, nil
	}

	pubkey, err = utils.ParsingPublickey(earnings)
	if err != nil {
		return txhash, earnings, errors.Wrap(err, "[DecodeToPub]")
	}
	acc, err = types.NewAccountID(pubkey)
	if err != nil {
		return txhash, earnings, errors.Wrap(err, "[NewAccountID]")
	}
	realTokens, ok := new(big.Int).SetString(strconv.FormatUint(pledge, 10)+pattern.TokenPrecision_CESS, 10)
	if !ok {
		return txhash, earnings, errors.New("[big.Int.SetString]")
	}
	call, err = types.NewCall(c.metadata, pattern.TX_SMINER_REGISTER, *acc, peerid, types.NewU128(*realTokens))
	if err != nil {
		return txhash, earnings, errors.Wrap(err, "[NewCall]")
	}

	ok, err = c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, earnings, errors.Wrap(err, "[GetStorageLatest]")
	}

	if !ok {
		keyStr, _ := utils.NumsToByteStr(key, map[string]bool{})
		return txhash, earnings, fmt.Errorf(
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
		return txhash, earnings, errors.Wrap(err, "[Sign]")
	}

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		c.SetChainState(false)
		return txhash, earnings, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
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
					return txhash, earnings, errors.Wrap(err, "[GetStorageRaw]")
				}
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil {
					return txhash, earnings, nil
				}

				if len(events.Sminer_Registered) > 0 {
					return txhash, earnings, nil
				}
			}
		case err = <-sub.Err():
			return txhash, earnings, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, earnings, pattern.ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) updateSminerPeerId(key types.StorageKey, peerid pattern.PeerId) (string, error) {
	var (
		err         error
		txhash      string
		call        types.Call
		accountInfo types.AccountInfo
	)

	call, err = types.NewCall(c.metadata, pattern.TX_SMINER_UPDATEPEERID, peerid)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewCall]")
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
		c.SetChainState(false)
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
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
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil || len(events.Sminer_UpdataIp) > 0 {
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

func (c *chainClient) UpdateSminerPeerId(peerid pattern.PeerId) (string, error) {
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

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		return txhash, errors.Wrap(err, "[CreateStorageKey]")
	}

	call, err := types.NewCall(c.metadata, pattern.TX_SMINER_UPDATEPEERID, peerid)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewCall]")
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
		c.SetChainState(false)
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
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
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil || len(events.Sminer_UpdataIp) > 0 {
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

func (c *chainClient) ExitSminer() (string, error) {
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

	call, err = types.NewCall(c.metadata, pattern.TX_FILEBANK_MINEREXITPREP)
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
		c.SetChainState(false)
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
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
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil || len(events.Sminer_MinerExitPrep) > 0 {
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

func (c *chainClient) UpdateEarningsAcc(puk []byte) (string, error) {
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

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewAccountID]")
	}

	call, err := types.NewCall(c.metadata, pattern.TX_SMINER_UPDATEINCOME, *acc)
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
		c.SetChainState(false)
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
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
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil || len(events.Sminer_UpdataBeneficiary) > 0 {
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

func (c *chainClient) updateEarningsAcc(key types.StorageKey, puk []byte) (string, error) {
	var (
		txhash      string
		accountInfo types.AccountInfo
	)

	acc, err := types.NewAccountID(puk)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewAccountID]")
	}

	call, err := types.NewCall(c.metadata, pattern.TX_SMINER_UPDATEINCOME, *acc)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewCall]")
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
		c.SetChainState(false)
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
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
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil || len(events.Sminer_UpdataBeneficiary) > 0 {
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

func (c *chainClient) UpdateEarningsAccount(earnings string) (string, error) {
	puk, err := utils.ParsingPublickey(earnings)
	if err != nil {
		return "", err
	}
	return c.UpdateEarningsAcc(puk)
}

// Storage miners increase deposit function
func (c *chainClient) IncreaseStakingAmount(tokens *big.Int) (string, error) {
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

	call, err := types.NewCall(c.metadata, pattern.TX_SMINER_INCREASESTAKES, types.NewUCompact(tokens))
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
		c.SetChainState(false)
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
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
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil || len(events.Sminer_IncreaseCollateral) > 0 {
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

func (c *chainClient) IncreaseStorageNodeStakingAmount(token string) (string, error) {
	tokens, ok := new(big.Int).SetString(token+pattern.TokenPrecision_CESS, 10)
	if !ok {
		return "", fmt.Errorf("Invalid tokens: %s", token)
	}
	return c.IncreaseStakingAmount(tokens)
}

// ClaimRewards
func (c *chainClient) ClaimRewards() (string, error) {
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

	call, err := types.NewCall(c.metadata, pattern.TX_SMINER_CLAIMREWARD)
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
		c.SetChainState(false)
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
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
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil || len(events.Sminer_Receive) > 0 {
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

// Withdraw
func (c *chainClient) Withdraw() (string, error) {
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

	call, err := types.NewCall(c.metadata, pattern.TX_FILEBANK_WITHDRAW)
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
		c.SetChainState(false)
		return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
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
					return txhash, errors.Wrap(err, "[GetStorageRaw]")
				}
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil || len(events.FileBank_Withdraw) > 0 {
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

func (s *chainClient) Expenders() (pattern.ExpendersInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.ExpendersInfo

	if !s.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(s.metadata, pattern.SMINER, pattern.EXPENDERS)
	if err != nil {
		return data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := s.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return data, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return data, pattern.ERR_RPC_EMPTY_VALUE
	}
	return data, nil
}

func (c *chainClient) RegisterOrUpdateSminer_V2(peerId []byte, earnings string, pledge uint64, poisKey pattern.PoISKeyInfo, sign pattern.TeeSignature) (string, string, error) {
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
		pubkey      []byte
		minerinfo   pattern.MinerInfo
		acc         *types.AccountID
		call        types.Call
		accountInfo types.AccountInfo
	)

	if !c.GetChainState() {
		return txhash, earnings, pattern.ERR_RPC_CONNECTION
	}

	c.SetSdkName(pattern.Name_Sminer)

	var peerid pattern.PeerId
	if len(peerid) != len(peerId) {
		return txhash, earnings, fmt.Errorf("invalid peerid: %v", peerId)
	}

	for i := 0; i < len(peerid); i++ {
		peerid[i] = types.U8(peerId[i])
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.SYSTEM, pattern.ACCOUNT, c.keyring.PublicKey)
	if err != nil {
		return txhash, earnings, errors.Wrap(err, "[CreateStorageKey]")
	}

	minerinfo, err = c.QueryStorageMiner(c.keyring.PublicKey)
	if err != nil {
		if err.Error() != pattern.ERR_Empty {
			return txhash, earnings, err
		}
	} else {
		if !utils.CompareSlice([]byte(string(minerinfo.PeerId[:])), peerId) {
			txhash, err = c.updateSminerPeerId(key, peerid)
			if err != nil {
				return txhash, earnings, err
			}
		}
		acc, _ := utils.EncodePublicKeyAsCessAccount(minerinfo.BeneficiaryAcc[:])
		if earnings != "" {
			if acc != earnings {
				puk, err := utils.ParsingPublickey(earnings)
				if err != nil {
					return txhash, acc, err
				}
				txhash, err = c.updateEarningsAcc(key, puk)
				return txhash, earnings, err
			}
		}
		return "", acc, nil
	}

	pubkey, err = utils.ParsingPublickey(earnings)
	if err != nil {
		return txhash, earnings, errors.Wrap(err, "[DecodeToPub]")
	}
	acc, err = types.NewAccountID(pubkey)
	if err != nil {
		return txhash, earnings, errors.Wrap(err, "[NewAccountID]")
	}
	realTokens, ok := new(big.Int).SetString(strconv.FormatUint(pledge, 10)+pattern.TokenPrecision_CESS, 10)
	if !ok {
		return txhash, earnings, errors.New("[big.Int.SetString]")
	}
	call, err = types.NewCall(c.metadata, pattern.TX_SMINER_REGISTER, *acc, peerid, types.NewU128(*realTokens), poisKey, sign)
	if err != nil {
		return txhash, earnings, errors.Wrap(err, "[NewCall]")
	}

	ok, err = c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil {
		return txhash, earnings, errors.Wrap(err, "[GetStorageLatest]")
	}

	if !ok {
		keyStr, _ := utils.NumsToByteStr(key, map[string]bool{})
		return txhash, earnings, fmt.Errorf(
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
		return txhash, earnings, errors.Wrap(err, "[Sign]")
	}

	// Do the transfer and track the actual status
	sub, err := c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
	if err != nil {
		c.SetChainState(false)
		return txhash, earnings, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
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
					return txhash, earnings, errors.Wrap(err, "[GetStorageRaw]")
				}
				err = types.EventRecordsRaw(*h).DecodeEventRecords(c.metadata, &events)
				if err != nil {
					return txhash, earnings, nil
				}

				if len(events.Sminer_Registered) > 0 {
					return txhash, earnings, nil
				}
			}
		case err = <-sub.Err():
			return txhash, earnings, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, earnings, pattern.ERR_RPC_TIMEOUT
		}
	}
}
