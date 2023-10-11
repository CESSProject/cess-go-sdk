/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"log"
	"strings"
	"time"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/pkg/errors"
)

func (c *chainClient) QueryAssignedProof() ([][]pattern.ProofAssignmentInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	if !c.GetChainState() {
		return nil, pattern.ERR_RPC_CONNECTION
	}
	var list [][]pattern.ProofAssignmentInfo
	key := createPrefixedKey(pattern.AUDIT, pattern.UNVERIFYPROOF)
	keys, err := c.api.RPC.State.GetKeysLatest(key)
	if err != nil {
		return list, errors.Wrap(err, "[GetKeysLatest]")
	}
	set, err := c.api.RPC.State.QueryStorageAtLatest(keys)

	if err != nil {
		return list, errors.Wrap(err, "[QueryStorageAtLatest]")
	}
	for _, elem := range set {
		for _, change := range elem.Changes {
			var data []pattern.ProofAssignmentInfo
			if err := codec.Decode(change.StorageData, &data); err != nil {
				log.Println(err)
				continue
			}
			list = append(list, data)
		}
	}
	return list, nil
}

func (c *chainClient) QueryTeeAssignedProof(puk []byte) ([]pattern.ProofAssignmentInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data []pattern.ProofAssignmentInfo

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.AUDIT, pattern.UNVERIFYPROOF, puk)
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

func (c *chainClient) QueryChallengeExpiration() (uint32, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.U32

	if !c.GetChainState() {
		return 0, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.AUDIT, pattern.CHALLENGEDURATION)
	if err != nil {
		return 0, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return 0, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return 0, pattern.ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

func (c *chainClient) QueryChallengeVerifyExpiration() (uint32, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data types.U32

	if !c.GetChainState() {
		return 0, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.AUDIT, pattern.CHALLENGEVERIFYDURATION)
	if err != nil {
		return 0, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return 0, errors.Wrap(err, "[GetStorageLatest]")
	}
	if !ok {
		return 0, pattern.ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

func (c *chainClient) QueryChallengeInfo(puk []byte) (bool, pattern.ChallengeInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data pattern.ChallengeInfo

	if !c.GetChainState() {
		return false, data, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.AUDIT, pattern.CHALLENGESNAPSHOT, puk)
	if err != nil {
		return false, data, errors.Wrap(err, "[CreateStorageKey]")
	}

	ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
	if err != nil {
		return false, data, errors.Wrap(err, "[GetStorageLatest]")
	}

	return ok, data, nil
}

func (c *chainClient) QueryUnverifiedIdleProof(puk []byte) ([]pattern.IdleProofInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data []pattern.IdleProofInfo

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.AUDIT, pattern.UNVERIFYIDLEPROOF, puk)
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

func (c *chainClient) QueryUnverifiedServiceProof(puk []byte) ([]pattern.ServiceProofInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()
	var data []pattern.ServiceProofInfo

	if !c.GetChainState() {
		return data, pattern.ERR_RPC_CONNECTION
	}

	key, err := types.CreateStorageKey(c.metadata, pattern.AUDIT, pattern.UNVERIFYSERVICEPROOF, puk)
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

func (c *chainClient) SubmitIdleProof(idleProve []types.U8) (string, error) {
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

	call, err := types.NewCall(c.metadata, pattern.TX_AUDIT_SUBMITIDLEPROOF, idleProve)
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
		if strings.Contains(err.Error(), pattern.ERR_RPC_PRIORITYTOOLOW) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return txhash, errors.Wrap(err, "[Sign]")
			}
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				c.SetChainState(false)
				return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
			}
		} else {
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
				_, err = c.RetrieveEvent_Audit_SubmitIdleProof(status.AsInBlock)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) SubmitServiceProof(serviceProof []types.U8) (string, error) {
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

	call, err := types.NewCall(c.metadata, pattern.TX_AUDIT_SUBMITSERVICEPROOF, serviceProof)
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
		if strings.Contains(err.Error(), pattern.ERR_RPC_PRIORITYTOOLOW) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return txhash, errors.Wrap(err, "[Sign]")
			}
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				c.SetChainState(false)
				return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
			}
		} else {
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
				_, err = c.RetrieveEvent_Audit_SubmitServiceProof(status.AsInBlock)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) SubmitIdleProofResult(idleproveHash []types.U8, front, rear types.U64, accumulator pattern.Accumulator, result types.Bool, signature pattern.TeeSignature, tee_acc []byte) (string, error) {
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

	teeacc, err := types.NewAccountID(tee_acc)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewAccountID]")
	}

	call, err := types.NewCall(c.metadata, pattern.TX_AUDIT_SUBMITIDLEPROOFRESULT, idleproveHash, front, rear, accumulator, result, signature, *teeacc)
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
		if strings.Contains(err.Error(), pattern.ERR_RPC_PRIORITYTOOLOW) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return txhash, errors.Wrap(err, "[Sign]")
			}
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				c.SetChainState(false)
				return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
			}
		} else {
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
				_, err = c.RetrieveEvent_Audit_SubmitIdleVerifyResult(status.AsInBlock)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}

func (c *chainClient) SubmitServiceProofResult(result types.Bool, signature pattern.TeeSignature, bloomFilter pattern.BloomFilter, tee_acc []byte) (string, error) {
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

	teeacc, err := types.NewAccountID(tee_acc)
	if err != nil {
		return txhash, errors.Wrap(err, "[NewAccountID]")
	}

	call, err := types.NewCall(c.metadata, pattern.TX_AUDIT_SUBMITSERVICEPROOFRESULT, result, signature, bloomFilter, *teeacc)
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
		if strings.Contains(err.Error(), pattern.ERR_RPC_PRIORITYTOOLOW) {
			o.Nonce = types.NewUCompactFromUInt(uint64(accountInfo.Nonce + 1))
			err = ext.Sign(c.keyring, o)
			if err != nil {
				return txhash, errors.Wrap(err, "[Sign]")
			}
			sub, err = c.api.RPC.Author.SubmitAndWatchExtrinsic(ext)
			if err != nil {
				c.SetChainState(false)
				return txhash, errors.Wrap(err, "[SubmitAndWatchExtrinsic]")
			}
		} else {
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
				_, err = c.RetrieveEvent_Audit_SubmitServiceVerifyResult(status.AsInBlock)
				return txhash, err
			}
		case err = <-sub.Err():
			return txhash, errors.Wrap(err, "[sub]")
		case <-timeout.C:
			return txhash, pattern.ERR_RPC_TIMEOUT
		}
	}
}
