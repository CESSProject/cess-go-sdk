/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/CESSProject/cess-go-sdk/core/event"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

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
	for k, e := range events {
		result[k] = e.Name
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

func (c *chainClient) RetrieveEvent_FileBank_ClaimRestoralOrder(blockhash types.Hash) (event.Event_ClaimRestoralOrder, error) {
	var result event.Event_ClaimRestoralOrder
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankClaimRestoralOrder {
			for _, v := range e.Fields {
				vf := reflect.ValueOf(v.Value)
				if vf.Len() > 0 {
					allValue := fmt.Sprintf("%v", vf.Index(0))
					if strings.Contains(v.Name, "AccountId32") {
						temp := strings.Split(allValue, "] ")
						puk := make([]byte, types.AccountIDLen)
						for _, v := range temp {
							if strings.Count(v, " ") == (types.AccountIDLen - 1) {
								subValue := strings.TrimPrefix(v, "[")
								ids := strings.Split(subValue, " ")
								if len(ids) != types.AccountIDLen {
									continue
								}
								for kk, vv := range ids {
									intv, _ := strconv.Atoi(vv)
									puk[kk] = byte(intv)
								}
							}
						}

						if !utils.CompareSlice(puk, c.GetSignatureAccPulickey()) {
							continue
						}
						accid, err := types.NewAccountID(puk)
						if err != nil {
							continue
						}
						result.Miner = *accid
						return result, nil
					}
					// if strings.Contains(v.Name, "cp_cess_common.Hash") {
					// 	temp := strings.Split(allValue, "] ")
					// 	for _, v := range temp {
					// 		if strings.Count(v, " ") == (len(pattern.FileHash{}) - 1) {
					// 			subValue := strings.TrimPrefix(v, "[")
					// 			ids := strings.Split(subValue, " ")
					// 			if len(ids) != len(pattern.FileHash{}) {
					// 				continue
					// 			}
					// 			for kk, vv := range ids {
					// 				intv, _ := strconv.Atoi(vv)
					// 				result.OrderId[kk] = types.U8(intv)
					// 			}
					// 			if suc {
					// 				return result, nil
					// 			}
					// 		}
					// 	}
					// }
				}
			}
		}
	}
	return result, errors.New("failed: no FileBank_ClaimRestoralOrder event found")
}

func (c *chainClient) RetrieveEvent_Audit_SubmitIdleProof(blockhash types.Hash) (event.Event_SubmitIdleProof, error) {
	var result event.Event_SubmitIdleProof
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.AuditSubmitIdleProof {
			for _, v := range e.Fields {
				vf := reflect.ValueOf(v.Value)
				if vf.Len() > 0 {
					allValue := fmt.Sprintf("%v", vf.Index(0))
					if strings.Contains(v.Name, "AccountId32") {
						temp := strings.Split(allValue, "] ")
						puk := make([]byte, types.AccountIDLen)
						for _, v := range temp {
							if strings.Count(v, " ") == (types.AccountIDLen - 1) {
								subValue := strings.TrimPrefix(v, "[")
								ids := strings.Split(subValue, " ")
								if len(ids) != types.AccountIDLen {
									continue
								}
								for kk, vv := range ids {
									intv, _ := strconv.Atoi(vv)
									puk[kk] = byte(intv)
								}
							}
						}

						if !utils.CompareSlice(puk, c.GetSignatureAccPulickey()) {
							continue
						}
						accid, err := types.NewAccountID(puk)
						if err != nil {
							continue
						}
						result.Miner = *accid
						return result, nil
					}
				}
			}
		}
	}
	return result, errors.New("failed: no Audit_SubmitIdleProof event found")
}

func (c *chainClient) RetrieveEvent_Audit_SubmitServiceProof(blockhash types.Hash) (event.Event_SubmitServiceProof, error) {
	var result event.Event_SubmitServiceProof
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.AuditSubmitServiceProof {
			for _, v := range e.Fields {
				vf := reflect.ValueOf(v.Value)
				if vf.Len() > 0 {
					allValue := fmt.Sprintf("%v", vf.Index(0))
					if strings.Contains(v.Name, "AccountId32") {
						temp := strings.Split(allValue, "] ")
						puk := make([]byte, types.AccountIDLen)
						for _, v := range temp {
							if strings.Count(v, " ") == (types.AccountIDLen - 1) {
								subValue := strings.TrimPrefix(v, "[")
								ids := strings.Split(subValue, " ")
								if len(ids) != types.AccountIDLen {
									continue
								}
								for kk, vv := range ids {
									intv, _ := strconv.Atoi(vv)
									puk[kk] = byte(intv)
								}
							}
						}

						if !utils.CompareSlice(puk, c.GetSignatureAccPulickey()) {
							continue
						}
						accid, err := types.NewAccountID(puk)
						if err != nil {
							continue
						}
						result.Miner = *accid
						return result, nil
					}
				}
			}
		}
	}
	return result, errors.New("failed: no Audit_SubmitServiceProof event found")
}

func (c *chainClient) RetrieveEvent_Audit_SubmitIdleVerifyResult(blockhash types.Hash) (event.Event_SubmitIdleVerifyResult, error) {
	var result event.Event_SubmitIdleVerifyResult
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.AuditSubmitIdleVerifyResult {
			for _, v := range e.Fields {
				vf := reflect.ValueOf(v.Value)
				if vf.Len() > 0 {
					allValue := fmt.Sprintf("%v", vf.Index(0))
					// if strings.Contains(v.Name, "AccountId32.tee") {
					// 	temp := strings.Split(allValue, "] ")
					// 	puk := make([]byte, types.AccountIDLen)
					// 	for _, v := range temp {
					// 		if strings.Count(v, " ") == (types.AccountIDLen - 1) {
					// 			subValue := strings.TrimPrefix(v, "[")
					// 			ids := strings.Split(subValue, " ")
					// 			if len(ids) != types.AccountIDLen {
					// 				continue
					// 			}
					// 			for kk, vv := range ids {
					// 				intv, _ := strconv.Atoi(vv)
					// 				puk[kk] = byte(intv)
					// 			}
					// 		}
					// 	}

					// 	accid, err := types.NewAccountID(puk)
					// 	if err != nil {
					// 		continue
					// 	}
					// 	suc = true
					// 	result.Tee = *accid
					// }
					if strings.Contains(v.Name, "AccountId32.miner") {
						temp := strings.Split(allValue, "] ")
						puk := make([]byte, types.AccountIDLen)
						for _, v := range temp {
							if strings.Count(v, " ") == (types.AccountIDLen - 1) {
								subValue := strings.TrimPrefix(v, "[")
								ids := strings.Split(subValue, " ")
								if len(ids) != types.AccountIDLen {
									continue
								}
								for kk, vv := range ids {
									intv, _ := strconv.Atoi(vv)
									puk[kk] = byte(intv)
								}
							}
						}
						if !utils.CompareSlice(puk, c.GetSignatureAccPulickey()) {
							continue
						}
						accid, err := types.NewAccountID(puk)
						if err != nil {
							continue
						}
						result.Tee = *accid
						return result, nil
					}
				}
			}
		}
	}
	return result, errors.New("failed: no Audit_SubmitIdleVerifyResult event found")
}

func (c *chainClient) RetrieveEvent_Audit_SubmitServiceVerifyResult(blockhash types.Hash) (event.Event_SubmitServiceVerifyResult, error) {
	var result event.Event_SubmitServiceVerifyResult
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.AuditSubmitServiceVerifyResult {
			for _, v := range e.Fields {
				vf := reflect.ValueOf(v.Value)
				if vf.Len() > 0 {
					allValue := fmt.Sprintf("%v", vf.Index(0))
					if strings.Contains(v.Name, "AccountId32.miner") {
						temp := strings.Split(allValue, "] ")
						puk := make([]byte, types.AccountIDLen)
						for _, v := range temp {
							if strings.Count(v, " ") == (types.AccountIDLen - 1) {
								subValue := strings.TrimPrefix(v, "[")
								ids := strings.Split(subValue, " ")
								if len(ids) != types.AccountIDLen {
									continue
								}
								for kk, vv := range ids {
									intv, _ := strconv.Atoi(vv)
									puk[kk] = byte(intv)
								}
							}
						}
						if !utils.CompareSlice(puk, c.GetSignatureAccPulickey()) {
							continue
						}
						accid, err := types.NewAccountID(puk)
						if err != nil {
							continue
						}
						result.Tee = *accid
						return result, nil
					}
				}
			}
		}
	}
	return result, errors.New("failed: no Audit_SubmitServiceVerifyResult event found")
}
