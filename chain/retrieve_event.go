/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"

	"github.com/AstaFrode/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

func (c *ChainClient) RetrieveAllEventName(blockhash types.Hash) ([]string, error) {
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

func (c *ChainClient) RetrieveEvent(blockhash types.Hash, extrinsic_name, signer string) error {
	if len(ExtrinsicsName) <= 0 {
		return errors.New("please call InitExtrinsicsName method first")
	}

	if len(extrinsic_name) <= 0 {
		return errors.New("extrinsic_name or event_name is empty")
	}

	if len(signer) != CESSWalletLen {
		return errors.New("invalid wallet account")
	}

	block, err := c.api.RPC.Chain.GetBlock(blockhash)
	if err != nil {
		return err
	}

	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return err
	}

	var (
		ok               bool
		name             string
		extrinsic_signer string
	)
	for _, e := range events {
		if !e.Phase.IsApplyExtrinsic {
			continue
		}
		if name == "" {
			name, ok = ExtrinsicsName[block.Block.Extrinsics[e.Phase.AsApplyExtrinsic].Method.CallIndex]
			if !ok {
				continue
			}
		}
		if name != extrinsic_name {
			if e.Name == SystemExtrinsicSuccess || e.Name == SystemExtrinsicFailed {
				name = ""
			}
			continue
		}
		switch e.Name {
		case TransactionPaymentTransactionFeePaid:
			extrinsic_signer, _, _ = parseSignerAndFeePaidFromEvent(e)
		case EvmAccountMappingTransactionFeePaid:
			extrinsic_signer, _, _ = parseSignerAndFeePaidFromEvent(e)
		case SystemExtrinsicSuccess:
			name = ""
			if extrinsic_signer == signer {
				return nil
			}
		case SystemExtrinsicFailed:
			name = ""
			if extrinsic_signer == signer {
				return errors.New(SystemExtrinsicFailed)
			}
		}
	}
	return fmt.Errorf("not found extrinsic: %s", extrinsic_name)
}
