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
	"github.com/CESSProject/cess-go-sdk/utils"
)

// QueryChallengeSnapShot query challenge snapshot data
//   - accountID: signature account of the storage miner
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - bool: is there a value
//   - ChallengeInfo: challenge snapshot data
//   - error: error message
func (c *ChainClient) QueryChallengeSnapShot(accountID []byte, block int32) (bool, ChallengeInfo, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return false, ChallengeInfo{}, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Audit, ChallengeSnapShot, ERR_RPC_CONNECTION.Error())
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data ChallengeInfo

	key, err := types.CreateStorageKey(c.metadata, Audit, ChallengeSnapShot, accountID)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Audit, ChallengeSnapShot, err)
		return false, data, err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Audit, ChallengeSnapShot, err)
			c.SetRpcState(false)
			return false, data, err
		}
		return ok, data, nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return false, data, err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Audit, ChallengeSnapShot, err)
		c.SetRpcState(false)
		return false, data, err
	}
	return ok, data, nil
}

// QueryCounterdClear query the number of times to clear the challenge failure count
//   - accountID: signature account of the storage miner
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint8: cleanup count
//   - error: error message
func (c *ChainClient) QueryCountedClear(accountID []byte, block int32) (uint8, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Audit, CountedClear, ERR_RPC_CONNECTION.Error())
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U8

	key, err := types.CreateStorageKey(c.metadata, Audit, CountedClear, accountID)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Audit, CountedClear, err)
		return uint8(data), err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Audit, CountedClear, err)
			c.SetRpcState(false)
			return uint8(data), err
		}
		if !ok {
			return uint8(data), ERR_RPC_EMPTY_VALUE
		}
		return uint8(data), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return uint8(data), err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Audit, CountedClear, err)
		c.SetRpcState(false)
		return uint8(data), err
	}
	if !ok {
		return uint8(data), ERR_RPC_EMPTY_VALUE
	}
	return uint8(data), nil
}

// QueryCountedServiceFailed query the number of failed service data challenge
//   - accountID: signature account of the storage miner
//   - block: block number, less than 0 indicates the latest block
//
// Return:
//   - uint32: Is there a value
//   - error: error message
func (c *ChainClient) QueryCountedServiceFailed(accountID []byte, block int32) (uint32, error) {
	if !c.GetRpcState() {
		if err := c.ReconnectRpc(); err != nil {
			return 0, fmt.Errorf("rpc err: [%s] [st] [%s.%s] %s", c.GetCurrentRpcAddr(), Audit, CountedServiceFailed, ERR_RPC_CONNECTION.Error())
		}
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var data types.U32

	key, err := types.CreateStorageKey(c.metadata, Audit, CountedServiceFailed, accountID)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] CreateStorageKey: %v", c.GetCurrentRpcAddr(), Audit, CountedServiceFailed, err)
		return uint32(data), err
	}

	if block < 0 {
		ok, err := c.api.RPC.State.GetStorageLatest(key, &data)
		if err != nil {
			err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorageLatest: %v", c.GetCurrentRpcAddr(), Audit, CountedServiceFailed, err)
			c.SetRpcState(false)
			return 0, err
		}
		if !ok {
			return 0, ERR_RPC_EMPTY_VALUE
		}
		return uint32(data), nil
	}
	blockhash, err := c.api.RPC.Chain.GetBlockHash(uint64(block))
	if err != nil {
		return uint32(data), err
	}
	ok, err := c.api.RPC.State.GetStorage(key, &data, blockhash)
	if err != nil {
		err = fmt.Errorf("rpc err: [%s] [st] [%s.%s] GetStorage: %v", c.GetCurrentRpcAddr(), Audit, CountedServiceFailed, err)
		c.SetRpcState(false)
		return uint32(data), err
	}
	if !ok {
		return uint32(data), ERR_RPC_EMPTY_VALUE
	}
	return uint32(data), nil
}

// SubmitIdleProof submit idle data proof to the chain
//   - idleProof: idle data proof
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) SubmitIdleProof(idleProof []types.U8) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	if len(idleProof) == 0 {
		return "", ERR_IdleProofIsEmpty
	}

	newcall, err := types.NewCall(c.metadata, ExtName_Audit_submit_idle_proof, idleProof)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Audit_submit_idle_proof, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Audit_submit_idle_proof)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Audit_submit_idle_proof, err)
	}

	return blockhash, nil
}

// SubmitServiceProof submit service data proof to the chain
//   - serviceProof: service data proof
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) SubmitServiceProof(serviceProof []types.U8) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_Audit_submit_service_proof, serviceProof)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Audit_submit_service_proof, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Audit_submit_service_proof)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Audit_submit_service_proof, err)
	}

	return blockhash, nil
}

// SubmitVerifyIdleResult submit validation result of idle data proof to the chain
//   - totalProofHash: total idle data proof hash value
//   - front: idle data pre-offset
//   - rear: back offset of idle data
//   - accumulator: accumulator value
//   - result: validation result of idle data proof
//   - sig: signature from tee
//   - teePuk: tee's work public key
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) SubmitVerifyIdleResult(totalProofHash []types.U8, front, rear types.U64, accumulator Accumulator, result types.Bool, sig types.Bytes, teePuk WorkerPublicKey) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_Audit_submit_verify_idle_result, totalProofHash, front, rear, accumulator, result, sig, teePuk)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Audit_submit_verify_idle_result, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Audit_submit_verify_idle_result)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Audit_submit_verify_idle_result, err)
	}

	return blockhash, nil
}

// SubmitVerifyServiceResult submit validation result of service data proof to the chain
//   - result: validation result of idle data proof
//   - sig: signature from tee
//   - bloomFilter: bloom filter value
//   - teePuk: tee's work public key
//
// Return:
//   - string: block hash
//   - error: error message
func (c *ChainClient) SubmitVerifyServiceResult(result types.Bool, sign types.Bytes, bloomFilter BloomFilter, teePuk WorkerPublicKey) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	newcall, err := types.NewCall(c.metadata, ExtName_Audit_submit_verify_service_result, result, sign, bloomFilter, teePuk)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Audit_submit_verify_service_result, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Audit_submit_verify_service_result)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Audit_submit_verify_service_result, err)
	}

	return blockhash, nil
}
