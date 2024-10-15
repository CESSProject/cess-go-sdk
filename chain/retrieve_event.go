/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
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

func (c *ChainClient) RetrieveEvent(blockhash types.Hash, extrinsic_name, event_name, signer string) error {
	if len(ExtrinsicsName) <= 0 {
		return errors.New("please call InitExtrinsicsName method first")
	}

	if len(extrinsic_name) <= 0 || len(event_name) <= 0 {
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
		if name == "" {
			if !e.Phase.IsApplyExtrinsic {
				continue
			}
			name, ok = ExtrinsicsName[block.Block.Extrinsics[e.Phase.AsApplyExtrinsic].Method.CallIndex]
			if !ok {
				continue
			}
			fmt.Println("name: ", name)
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
				return errors.New(ERR_Failed)
			}
		}
	}
	return errors.Errorf("transaction failed: no %s event found", event_name)
}

func (c *ChainClient) RetrieveExtrinsicsAndEvents(blockhash types.Hash) ([]string, map[string][]string, error) {
	var systemEvents = make([]string, 0)
	var extrinsicsEvents = make(map[string][]string, 0)
	block, err := c.GetSubstrateAPI().RPC.Chain.GetBlock(blockhash)
	if err != nil {
		return systemEvents, extrinsicsEvents, err
	}
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return systemEvents, extrinsicsEvents, err
	}
	for _, e := range events {
		if e.Phase.IsApplyExtrinsic {
			if name, ok := ExtrinsicsName[block.Block.Extrinsics[e.Phase.AsApplyExtrinsic].Method.CallIndex]; ok {
				if extrinsicsEvents[name] == nil {
					extrinsicsEvents[name] = make([]string, 0)
				}
				extrinsicsEvents[name] = append(extrinsicsEvents[name], e.Name)
			}
		} else {
			systemEvents = append(systemEvents, e.Name)
		}
	}
	return systemEvents, extrinsicsEvents, nil
}

func (c *ChainClient) RetrieveEvent_Sminer_Receive(blockhash types.Hash) (Event_Receive, error) {
	var result Event_Receive

	block, err := c.api.RPC.Chain.GetBlock(blockhash)
	if err != nil {
		return result, err
	}

	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}

	var signer string
	var earningsAcc string
	for _, e := range events {
		if e.Phase.IsApplyExtrinsic {
			if name, ok := ExtrinsicsName[block.Block.Extrinsics[e.Phase.AsApplyExtrinsic].Method.CallIndex]; ok {
				if name == ExtName_Sminer_receive_reward {
					switch e.Name {
					case SminerReceive:
						earningsAcc, _ = ParseAccountFromEvent(e)
						result.Acc = earningsAcc
					case TransactionPaymentTransactionFeePaid:
						signer, _, _ = parseSignerAndFeePaidFromEvent(e)
					case EvmAccountMappingTransactionFeePaid:
						signer, _, _ = parseSignerAndFeePaidFromEvent(e)
					case SystemExtrinsicSuccess:
						if signer == c.GetSignatureAcc() {
							return result, nil
						}
					case SystemExtrinsicFailed:
						if signer == c.GetSignatureAcc() {
							return result, errors.New(ERR_Failed)
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", SminerReceive)
}
