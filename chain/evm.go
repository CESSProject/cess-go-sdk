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

func (c *ChainClient) SendEvmCall(source types.H160, target types.H160, input types.Bytes, value types.U256, gasLimit types.U64, maxFeePerGas types.U256, accessList []AccessInfo) (string, error) {
	<-c.tradeCh
	defer func() {
		c.tradeCh <- true
		if err := recover(); err != nil {
			log.Println(utils.RecoverError(err))
		}
	}()

	var nonce types.Option[types.U256]
	nonce.SetNone()

	var maxPriorityFeePerGas types.Option[types.U256]
	maxPriorityFeePerGas.SetNone()

	newcall, err := types.NewCall(c.metadata, ExtName_Evm_call, source, target, input, value, gasLimit, maxFeePerGas, maxPriorityFeePerGas, nonce, accessList)
	if err != nil {
		return "", fmt.Errorf("rpc err: [%s] [tx] [%s] NewCall: %v", c.GetCurrentRpcAddr(), ExtName_Evm_call, err)
	}

	blockhash, err := c.SubmitExtrinsic(newcall, ExtName_Evm_call)
	if err != nil {
		return blockhash, fmt.Errorf("rpc err: [%s] [tx] [%s] SubmitExtrinsic: %v", c.GetCurrentRpcAddr(), ExtName_Evm_call, err)
	}

	return blockhash, nil
}
