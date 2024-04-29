/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"bytes"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/CESSProject/cess-go-sdk/core/event"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/parser"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/vedhavyas/go-subkey/scale"
	"golang.org/x/crypto/blake2b"
)

func (c *ChainClient) DecodeEventNameFromBlock(block uint64) ([]string, error) {
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

func (c *ChainClient) DecodeEventNameFromBlockhash(blockhash types.Hash) ([]string, error) {
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

func (c *ChainClient) RetrieveEvent_FileBank_ClaimRestoralOrder(blockhash types.Hash) (event.Event_ClaimRestoralOrder, error) {
	var result event.Event_ClaimRestoralOrder
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankClaimRestoralOrder {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
	}
	return result, errors.Errorf("failed: no %s event found", event.FileBankClaimRestoralOrder)
}

func (c *ChainClient) RetrieveEvent_Audit_SubmitIdleProof(blockhash types.Hash) (event.Event_SubmitIdleProof, error) {
	var result event.Event_SubmitIdleProof
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.AuditSubmitIdleProof {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
	}
	return result, errors.Errorf("failed: no %s event found", event.AuditSubmitIdleProof)
}

func (c *ChainClient) RetrieveEvent_Audit_SubmitServiceProof(blockhash types.Hash) (event.Event_SubmitServiceProof, error) {
	var result event.Event_SubmitServiceProof
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.AuditSubmitServiceProof {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
	}
	return result, errors.Errorf("failed: no %s event found", event.AuditSubmitServiceProof)
}

func (c *ChainClient) RetrieveEvent_Audit_SubmitIdleVerifyResult(blockhash types.Hash) (event.Event_SubmitIdleVerifyResult, error) {
	var result event.Event_SubmitIdleVerifyResult
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.AuditSubmitIdleVerifyResult {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Miner = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.AuditSubmitIdleVerifyResult)
}

func (c *ChainClient) RetrieveEvent_Audit_SubmitServiceVerifyResult(blockhash types.Hash) (event.Event_SubmitServiceVerifyResult, error) {
	var result event.Event_SubmitServiceVerifyResult
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.AuditSubmitServiceVerifyResult {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Miner = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.AuditSubmitServiceVerifyResult)
}

func (c *ChainClient) RetrieveEvent_Oss_OssUpdate(blockhash types.Hash) (event.Event_OssUpdate, error) {
	var result event.Event_OssUpdate
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.OssOssUpdate {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.OssOssUpdate)
}

func (c *ChainClient) RetrieveEvent_Oss_OssRegister(blockhash types.Hash) (event.Event_OssRegister, error) {
	var result event.Event_OssRegister
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.OssOssRegister {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.OssOssRegister)
}

func (c *ChainClient) RetrieveEvent_Oss_OssDestroy(blockhash types.Hash) (event.Event_OssDestroy, error) {
	var result event.Event_OssDestroy
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.OssOssDestroy {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.OssOssDestroy)
}

func (c *ChainClient) RetrieveEvent_Oss_Authorize(blockhash types.Hash) (event.Event_Authorize, error) {
	var result event.Event_Authorize
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.OssAuthorize {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.OssAuthorize)
}

func (c *ChainClient) RetrieveEvent_Oss_CancelAuthorize(blockhash types.Hash) (event.Event_CancelAuthorize, error) {
	var result event.Event_CancelAuthorize
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.OssCancelAuthorize {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.OssCancelAuthorize)
}

func (c *ChainClient) RetrieveEvent_FileBank_UploadDeclaration(blockhash types.Hash) (event.Event_UploadDeclaration, error) {
	var result event.Event_UploadDeclaration
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankUploadDeclaration {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
					vf := reflect.ValueOf(v.Value)
					if vf.Len() > 0 {
						allValue := fmt.Sprintf("%v", vf.Index(0))
						if strings.Contains(v.Name, "AccountId32.operator") {
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
							result.Operator = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.FileBankUploadDeclaration)
}

func (c *ChainClient) RetrieveEvent_FileBank_CreateBucket(blockhash types.Hash) (event.Event_CreateBucket, error) {
	var result event.Event_CreateBucket
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankCreateBucket {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
					vf := reflect.ValueOf(v.Value)
					if vf.Len() > 0 {
						allValue := fmt.Sprintf("%v", vf.Index(0))
						if strings.Contains(v.Name, "AccountId32.operator") {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.FileBankCreateBucket)
}

func (c *ChainClient) RetrieveEvent_FileBank_DeleteBucket(blockhash types.Hash) (event.Event_DeleteBucket, error) {
	var result event.Event_DeleteBucket
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankDeleteBucket {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
					vf := reflect.ValueOf(v.Value)
					if vf.Len() > 0 {
						allValue := fmt.Sprintf("%v", vf.Index(0))
						if strings.Contains(v.Name, "AccountId32.operator") {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.FileBankDeleteBucket)
}

func (c *ChainClient) RetrieveEvent_FileBank_DeleteFile(blockhash types.Hash) (event.Event_DeleteFile, error) {
	var result event.Event_DeleteFile
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankDeleteFile {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
					vf := reflect.ValueOf(v.Value)
					if vf.Len() > 0 {
						allValue := fmt.Sprintf("%v", vf.Index(0))
						if strings.Contains(v.Name, "AccountId32.operator") {
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
							result.Operator = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.FileBankDeleteFile)
}

func (c *ChainClient) RetrieveEvent_FileBank_TransferReport(blockhash types.Hash) (event.Event_TransferReport, error) {
	var result event.Event_TransferReport
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankTransferReport {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.FileBankTransferReport)
}

func (c *ChainClient) RetrieveEvent_FileBank_RecoveryCompleted(blockhash types.Hash) (event.Event_RecoveryCompleted, error) {
	var result event.Event_RecoveryCompleted
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankRecoveryCompleted {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
	}
	return result, errors.Errorf("failed: no %s event found", event.FileBankRecoveryCompleted)
}

func (c *ChainClient) RetrieveEvent_FileBank_IdleSpaceCert(blockhash types.Hash) (event.Event_IdleSpaceCert, error) {
	var result event.Event_IdleSpaceCert
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankIdleSpaceCert {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.FileBankIdleSpaceCert)
}

func (c *ChainClient) RetrieveEvent_FileBank_ReplaceIdleSpace(blockhash types.Hash) (event.Event_ReplaceIdleSpace, error) {
	var result event.Event_ReplaceIdleSpace
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankReplaceIdleSpace {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.FileBankReplaceIdleSpace)
}

func (c *ChainClient) RetrieveEvent_FileBank_CalculateReport(blockhash types.Hash) (event.Event_CalculateReport, error) {
	var result event.Event_CalculateReport
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankCalculateReport {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
	}
	return result, errors.Errorf("failed: no %s event found", event.FileBankCalculateReport)
}

func (c *ChainClient) RetrieveEvent_Sminer_UpdataIp(blockhash types.Hash) (event.Event_UpdatePeerId, error) {
	var result event.Event_UpdatePeerId
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.SminerUpdatePeerId {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.SminerUpdatePeerId)
}

func (c *ChainClient) RetrieveEvent_Sminer_UpdataBeneficiary(blockhash types.Hash) (event.Event_UpdateBeneficiary, error) {
	var result event.Event_UpdateBeneficiary
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.SminerUpdateBeneficiary {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.SminerUpdateBeneficiary)
}

func (c *ChainClient) RetrieveEvent_Sminer_Registered(blockhash types.Hash) (event.Event_Registered, error) {
	var result event.Event_Registered
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.SminerRegistered {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.SminerRegistered)
}

func (c *ChainClient) RetrieveEvent_Sminer_RegisterPoisKey(blockhash types.Hash) (event.Event_RegisterPoisKey, error) {
	var result event.Event_RegisterPoisKey
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.SminerRegisterPoisKey {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
	}
	return result, errors.Errorf("failed: no %s event found", event.SminerRegisterPoisKey)
}

func (c *ChainClient) RetrieveEvent_Sminer_MinerExitPrep(blockhash types.Hash) (event.Event_MinerExitPrep, error) {
	var result event.Event_MinerExitPrep
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.SminerMinerExitPrep {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
	}
	return result, errors.Errorf("failed: no %s event found", event.SminerMinerExitPrep)
}

func (c *ChainClient) RetrieveEvent_Sminer_IncreaseCollateral(blockhash types.Hash) (event.Event_IncreaseCollateral, error) {
	var result event.Event_IncreaseCollateral
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.SminerIncreaseCollateral {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.SminerIncreaseCollateral)
}

func (c *ChainClient) RetrieveEvent_Sminer_IncreaseDeclarationSpace(blockhash types.Hash) (event.Event_IncreaseDeclarationSpace, error) {
	var result event.Event_IncreaseDeclarationSpace
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.SminerIncreaseDeclarationSpace {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
	}
	return result, errors.Errorf("failed: no %s event found", event.SminerIncreaseDeclarationSpace)
}

func (c *ChainClient) RetrieveEvent_Sminer_Receive(blockhash types.Hash) (event.Event_Receive, error) {
	var result event.Event_Receive
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.SminerReceive {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.SminerReceive)
}

func (c *ChainClient) RetrieveEvent_Sminer_Withdraw(blockhash types.Hash) (event.Event_Withdraw, error) {
	var result event.Event_Withdraw
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.SminerWithdraw {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.SminerWithdraw)
}

func (c *ChainClient) RetrieveEvent_StorageHandler_BuySpace(blockhash types.Hash) (event.Event_BuySpace, error) {
	var result event.Event_BuySpace
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.StorageHandlerBuySpace {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.StorageHandlerBuySpace)
}

func (c *ChainClient) RetrieveEvent_StorageHandler_ExpansionSpace(blockhash types.Hash) (event.Event_ExpansionSpace, error) {
	var result event.Event_ExpansionSpace
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.StorageHandlerExpansionSpace {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.StorageHandlerExpansionSpace)
}

func (c *ChainClient) RetrieveEvent_StorageHandler_RenewalSpace(blockhash types.Hash) (event.Event_RenewalSpace, error) {
	var result event.Event_RenewalSpace
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.StorageHandlerRenewalSpace {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
							result.Acc = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.StorageHandlerRenewalSpace)
}

func (c *ChainClient) RetrieveEvent_Balances_Transfer(blockhash types.Hash) (types.EventBalancesTransfer, error) {
	var result types.EventBalancesTransfer
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.BalanceTransfer {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
					vf := reflect.ValueOf(v.Value)
					if vf.Len() > 0 {
						allValue := fmt.Sprintf("%v", vf.Index(0))
						if strings.Contains(v.Name, "AccountId32.from") {
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
							result.From = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.BalanceTransfer)
}

func (c *ChainClient) RetrieveEvent_FileBank_GenRestoralOrder(blockhash types.Hash) (event.Event_GenerateRestoralOrder, error) {
	var result event.Event_GenerateRestoralOrder
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankGenerateRestoralOrder {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
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
	}
	return result, errors.Errorf("failed: no %s event found", event.FileBankGenerateRestoralOrder)
}

func (c *ChainClient) RetrieveAllEvent_FileBank_UploadDeclaration(blockhash types.Hash) ([]event.AllUploadDeclarationEvent, error) {
	var result = make([]event.AllUploadDeclarationEvent, 0)
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}

	for _, e := range events {
		if e.Name == event.FileBankUploadDeclaration {
			var ele event.AllUploadDeclarationEvent
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
					vf := reflect.ValueOf(v.Value)
					if vf.Len() > 0 {
						allValue := fmt.Sprintf("%v", vf.Index(0))
						if strings.Contains(v.Name, "AccountId32.operator") {
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
							ele.Operator, _ = utils.EncodePublicKeyAsCessAccount(puk)
						} else if strings.Contains(v.Name, "AccountId32.owner") {
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
							ele.Owner, _ = utils.EncodePublicKeyAsCessAccount(puk)
						} else if strings.Contains(v.Name, "deal_hash") {
							temp := strings.Split(allValue, "] ")
							for _, v := range temp {
								if strings.Count(v, " ") == (pattern.FileHashLen - 1) {
									subValue := strings.TrimPrefix(v, "[")
									ids := strings.Split(subValue, " ")
									if len(ids) != pattern.FileHashLen {
										continue
									}
									var fhash pattern.FileHash
									for kk, vv := range ids {
										intv, _ := strconv.Atoi(vv)
										fhash[kk] = types.U8(intv)
									}
									ele.Filehash = string(fhash[:])
								}
							}
						}
					}
				}
			}
			result = append(result, ele)
		}
	}
	return result, nil
}

func (c *ChainClient) RetrieveAllEvent_FileBank_StorageCompleted(blockhash types.Hash) ([]string, error) {
	var result = make([]string, 0)
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankStorageCompleted {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
					vf := reflect.ValueOf(v.Value)
					if vf.Len() > 0 {
						allValue := fmt.Sprintf("%v", vf.Index(0))
						if strings.Contains(v.Name, "file_hash") {
							temp := strings.Split(allValue, "] ")
							for _, v := range temp {
								if strings.Count(v, " ") == (pattern.FileHashLen - 1) {
									subValue := strings.TrimPrefix(v, "[")
									ids := strings.Split(subValue, " ")
									if len(ids) != pattern.FileHashLen {
										continue
									}
									var fhash pattern.FileHash
									for kk, vv := range ids {
										intv, _ := strconv.Atoi(vv)
										fhash[kk] = types.U8(intv)
									}
									result = append(result, string(fhash[:]))
								}
							}
						}
					}
				}
			}
		}
	}
	return result, nil
}

func (c *ChainClient) RetrieveAllEvent_FileBank_DeleteFile(blockhash types.Hash) ([]event.AllDeleteFileEvent, error) {
	var result = make([]event.AllDeleteFileEvent, 0)
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == event.FileBankDeleteFile {
			var ele event.AllDeleteFileEvent
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
					vf := reflect.ValueOf(v.Value)
					if vf.Len() > 0 {
						if strings.Contains(v.Name, "AccountId32.operator") {
							allValue := fmt.Sprintf("%v", vf.Index(0))
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
							ele.Operator, _ = utils.EncodePublicKeyAsCessAccount(puk)
						} else if strings.Contains(v.Name, "AccountId32.owner") {
							allValue := fmt.Sprintf("%v", vf.Index(0))
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
							ele.Owner, _ = utils.EncodePublicKeyAsCessAccount(puk)
						} else if strings.Contains(v.Name, "file_hash") {
							allValue := fmt.Sprintf("%v", vf.Index(0))
							temp := strings.Split(allValue, "] ")
							for _, v := range temp {
								if strings.Count(v, " ") == (pattern.FileHashLen - 1) {
									subValue := strings.TrimPrefix(v, "[")
									ids := strings.Split(subValue, " ")
									if len(ids) != pattern.FileHashLen {
										continue
									}
									var fhash pattern.FileHash
									for kk, vv := range ids {
										intv, _ := strconv.Atoi(vv)
										fhash[kk] = types.U8(intv)
									}
									ele.Filehash = string(fhash[:])
								}
							}
						}
					}
				}
			}
			result = append(result, ele)
		}
	}
	return result, nil
}

// func (c *ChainClient) RetrieveAllEvent(blockhash types.Hash) ([]string, []string, error) {
// 	var flag bool
// 	var systemEvents = make([]string, 0)
// 	var extrinsicsEvents = make([]string, 0)
// 	events, err := c.eventRetriever.GetEvents(blockhash)
// 	if err != nil {
// 		return systemEvents, extrinsicsEvents, err
// 	}
// 	for _, e := range events {
// 		fmt.Println("event name: ", e.Name)
// 		if e.Name == "System.ExtrinsicSuccess" {
// 			flag = true
// 		}
// 		if flag {
// 			extrinsicsEvents = append(extrinsicsEvents, e.Name)
// 		} else {
// 			systemEvents = append(systemEvents, e.Name)
// 		}
// 	}
// 	return systemEvents, extrinsicsEvents, nil
// }

func (c *ChainClient) RetrieveAllEventFromBlock(blockhash types.Hash) ([]string, map[string][]string, error) {
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

func (c *ChainClient) RetrieveBlock(blocknumber uint64) ([]string, []event.ExtrinsicsInfo, []event.TransferInfo, string, string, string, string, int64, error) {
	var timeUnixMilli int64
	var systemEvents = make([]string, 0)
	var extrinsicsInfo = make([]event.ExtrinsicsInfo, 0)
	var transferInfo = make([]event.TransferInfo, 0)
	blockhash, err := c.GetSubstrateAPI().RPC.Chain.GetBlockHash(blocknumber)
	if err != nil {
		return systemEvents, extrinsicsInfo, transferInfo, "", "", "", "", 0, err
	}
	block, err := c.GetSubstrateAPI().RPC.Chain.GetBlock(blockhash)
	if err != nil {
		return systemEvents, extrinsicsInfo, transferInfo, "", "", "", "", 0, err
	}
	if blocknumber == 0 {
		return systemEvents, extrinsicsInfo, transferInfo, blockhash.Hex(), block.Block.Header.ParentHash.Hex(), block.Block.Header.ExtrinsicsRoot.Hex(), block.Block.Header.StateRoot.Hex(), 0, nil
	}
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return systemEvents, extrinsicsInfo, transferInfo, "", "", "", "", 0, err
	}
	var eventsBuf = make([]string, 0)
	var signer string
	var fee string
	var ok bool
	var name string
	//var parsedBalancesTransfer = true
	for _, e := range events {
		if e.Phase.IsApplyExtrinsic {
			if name, ok = ExtrinsicsName[block.Block.Extrinsics[e.Phase.AsApplyExtrinsic].Method.CallIndex]; ok {
				if name == ExtName_Timestamp_set {
					timeDecoder := scale.NewDecoder(bytes.NewReader(block.Block.Extrinsics[e.Phase.AsApplyExtrinsic].Method.Args))
					timestamp, err := timeDecoder.DecodeUintCompact()
					if err != nil {
						return systemEvents, extrinsicsInfo, transferInfo, "", "", "", "", 0, err
					}
					timeUnixMilli = timestamp.Int64()
					extrinsicsInfo = append(extrinsicsInfo, event.ExtrinsicsInfo{
						Name:   name,
						Events: []string{e.Name},
						Result: true,
					})
					continue
				}
				eventsBuf = append(eventsBuf, e.Name)
				if e.Name == event.TransactionPaymentTransactionFeePaid ||
					e.Name == event.EvmAccountMappingTransactionFeePaid {
					signer, fee, _ = parseSignerAndFeePaidFromEvent(e)
				} else if e.Name == event.BalancesTransfer {
					//parsedBalancesTransfer = false
					from, to, amount, _ := ParseTransferInfoFromEvent(e)
					transferInfo = append(transferInfo, event.TransferInfo{
						From:   from,
						To:     to,
						Amount: amount,
						Result: true,
					})
					// transfers, err := c.parseTransferInfoFromBlock(blockhash)
					// if err != nil {
					// 	return systemEvents, extrinsicsInfo, transferInfo, "", "", "", "", 0, err
					// }
					// if len(transfers) > 0 {
					// 	transferInfo = append(transferInfo, transfers...)
					// }
				} else if e.Name == event.SystemExtrinsicSuccess {
					if len(eventsBuf) > 0 {
						extrinsicsInfo = append(extrinsicsInfo, event.ExtrinsicsInfo{
							Name:    name,
							Signer:  signer,
							FeePaid: fee,
							Result:  true,
							Events:  append(make([]string, 0), eventsBuf...),
						})
						eventsBuf = make([]string, 0)
					}
				} else if e.Name == event.SystemExtrinsicFailed {
					if len(eventsBuf) > 0 {
						extrinsicsInfo = append(extrinsicsInfo, event.ExtrinsicsInfo{
							Name:    name,
							Signer:  signer,
							FeePaid: fee,
							Result:  false,
							Events:  append(make([]string, 0), eventsBuf...),
						})
						eventsBuf = make([]string, 0)
					}
				}
			}
		} else {
			systemEvents = append(systemEvents, e.Name)
		}
	}
	return systemEvents, extrinsicsInfo, transferInfo, blockhash.Hex(), block.Block.Header.ParentHash.Hex(), block.Block.Header.ExtrinsicsRoot.Hex(), block.Block.Header.StateRoot.Hex(), timeUnixMilli, nil
}

func (c *ChainClient) RetrieveBlockAndAll(blocknumber uint64) ([]string, []event.ExtrinsicsInfo, []event.TransferInfo, []string, []string, string, string, string, string, string, int64, error) {
	var timeUnixMilli int64
	var systemEvents = make([]string, 0)
	var extrinsicsInfo = make([]event.ExtrinsicsInfo, 0)
	var transferInfo = make([]event.TransferInfo, 0)
	var sminerRegInfo = make([]string, 0)
	var newAccounts = make([]string, 0)
	var allGasFee = new(big.Int)
	var allExtrinsicsHash = make([]string, 0)
	blockhash, err := c.GetSubstrateAPI().RPC.Chain.GetBlockHash(blocknumber)
	if err != nil {
		return systemEvents, extrinsicsInfo, transferInfo, sminerRegInfo, newAccounts, "", "", "", "", "", 0, err
	}
	block, err := c.GetSubstrateAPI().RPC.Chain.GetBlock(blockhash)
	if err != nil {
		return systemEvents, extrinsicsInfo, transferInfo, sminerRegInfo, newAccounts, "", "", "", "", "", 0, err
	}
	for _, v := range block.Block.Extrinsics {
		extBytes, err := codec.Encode(v)
		if err != nil {
			continue
		}
		h := blake2b.Sum256(extBytes)
		allExtrinsicsHash = append(allExtrinsicsHash, hexutil.Encode(h[:]))
	}
	if blocknumber == 0 {
		return systemEvents, extrinsicsInfo, transferInfo, sminerRegInfo, newAccounts, blockhash.Hex(), block.Block.Header.ParentHash.Hex(), block.Block.Header.ExtrinsicsRoot.Hex(), block.Block.Header.StateRoot.Hex(), "", 0, nil
	}
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return systemEvents, extrinsicsInfo, transferInfo, sminerRegInfo, newAccounts, "", "", "", "", "", 0, err
	}
	var eventsBuf = make([]string, 0)
	var signer string
	var fee string
	var ok bool
	var name string
	var extInfo event.ExtrinsicsInfo
	for _, e := range events {
		if e.Phase.IsApplyExtrinsic {
			if name, ok = ExtrinsicsName[block.Block.Extrinsics[e.Phase.AsApplyExtrinsic].Method.CallIndex]; ok {
				if name == ExtName_Timestamp_set {
					extInfo = event.ExtrinsicsInfo{}
					timeDecoder := scale.NewDecoder(bytes.NewReader(block.Block.Extrinsics[e.Phase.AsApplyExtrinsic].Method.Args))
					timestamp, err := timeDecoder.DecodeUintCompact()
					if err != nil {
						return systemEvents, extrinsicsInfo, transferInfo, sminerRegInfo, newAccounts, "", "", "", "", "", 0, err
					}
					timeUnixMilli = timestamp.Int64()
					extrinsicsInfo = append(extrinsicsInfo, event.ExtrinsicsInfo{
						Name:   name,
						Events: []string{e.Name},
						Result: true,
					})
					continue
				}
				eventsBuf = append(eventsBuf, e.Name)
				if e.Name == event.TransactionPaymentTransactionFeePaid ||
					e.Name == event.EvmAccountMappingTransactionFeePaid {
					signer, fee, _ = parseSignerAndFeePaidFromEvent(e)
					tmp, ok := new(big.Int).SetString(fee, 10)
					if ok {
						allGasFee = allGasFee.Add(allGasFee, tmp)
					}
				} else if e.Name == event.BalancesTransfer {
					from, to, amount, _ := ParseTransferInfoFromEvent(e)
					transferInfo = append(transferInfo, event.TransferInfo{
						ExtrinsicName: name,
						From:          from,
						To:            to,
						Amount:        amount,
						Result:        true,
					})
					// extInfo.From = from
					// extInfo.To = to
				} else if e.Name == event.SminerRegistered {
					acc, err := ParseAccountFromEvent(e)
					if err == nil {
						sminerRegInfo = append(sminerRegInfo, acc)
					}
				} else if e.Name == event.SystemNewAccount {
					acc, err := ParseAccountFromEvent(e)
					if err == nil {
						newAccounts = append(newAccounts, acc)
					}
				} else if e.Name == event.SystemExtrinsicSuccess {
					if len(eventsBuf) > 0 {
						extInfo.Name = name
						extInfo.Signer = signer
						extInfo.FeePaid = fee
						extInfo.Result = true
						extInfo.Events = append(make([]string, 0), eventsBuf...)
						extrinsicsInfo = append(extrinsicsInfo, extInfo)
						eventsBuf = make([]string, 0)
					}
					extInfo = event.ExtrinsicsInfo{}
				} else if e.Name == event.SystemExtrinsicFailed {
					if len(eventsBuf) > 0 {
						extInfo.Name = name
						extInfo.Signer = signer
						extInfo.FeePaid = fee
						extInfo.Result = false
						extInfo.Events = append(make([]string, 0), eventsBuf...)
						extrinsicsInfo = append(extrinsicsInfo, extInfo)
						eventsBuf = make([]string, 0)
					}
					extInfo = event.ExtrinsicsInfo{}
				}
			}
		} else {
			systemEvents = append(systemEvents, e.Name)
		}
	}
	if len(allExtrinsicsHash) != len(extrinsicsInfo) {
		return systemEvents, extrinsicsInfo, transferInfo, sminerRegInfo, newAccounts, "", "", "", "", "", 0, errors.New("The number of transaction hashes does not equal the number of transactions")
	}
	for i := 0; i < len(allExtrinsicsHash); i++ {
		extrinsicsInfo[i].Hash = allExtrinsicsHash[i]
	}

	// for i := 0; i < len(extrinsicsInfo); i++ {
	// 	for j := 0; j < len(transferInfo); j++ {
	// 		if extrinsicsInfo[i].Name == transferInfo[j].ExtrinsicName {
	// 			if transferInfo[j].ExtrinsicName == ExtName_Sminer_faucet {
	// 				if extrinsicsInfo[i].From == transferInfo[j].From &&
	// 					extrinsicsInfo[i].To == transferInfo[j].To {
	// 					transferInfo[j].ExtrinsicHash = extrinsicsInfo[i].Hash
	// 				}
	// 			}
	// 			if extrinsicsInfo[i].Signer == transferInfo[j].From {
	// 				transferInfo[j].ExtrinsicHash = extrinsicsInfo[i].Hash
	// 			}
	// 		}
	// 	}
	// }

	return systemEvents, extrinsicsInfo, transferInfo, sminerRegInfo, newAccounts, blockhash.Hex(), block.Block.Header.ParentHash.Hex(), block.Block.Header.ExtrinsicsRoot.Hex(), block.Block.Header.StateRoot.Hex(), allGasFee.String(), timeUnixMilli, nil
}

func (c *ChainClient) ParseBlockData(blocknumber uint64) (event.BlockData, error) {
	var (
		ok             bool
		name           string
		err            error
		extBytes       []byte
		extrinsicIndex int
		blockdata      event.BlockData
		extInfo        event.ExtrinsicsInfo
		allGasFee      = new(big.Int)
	)

	blockdata.BlockId = uint32(blocknumber)

	blockhash, err := c.api.RPC.Chain.GetBlockHash(blocknumber)
	if err != nil {
		return blockdata, err
	}
	blockdata.BlockHash = blockhash.Hex()

	block, err := c.api.RPC.Chain.GetBlock(blockhash)
	if err != nil {
		return blockdata, err
	}
	blockdata.PreHash = block.Block.Header.ParentHash.Hex()
	blockdata.ExtHash = block.Block.Header.ExtrinsicsRoot.Hex()
	blockdata.StHash = block.Block.Header.StateRoot.Hex()

	blockdata.Extrinsics = make([]event.ExtrinsicsInfo, len(block.Block.Extrinsics))
	for k, v := range block.Block.Extrinsics {
		extBytes, err = codec.Encode(v)
		if err != nil {
			return blockdata, err
		}
		h := blake2b.Sum256(extBytes)
		blockdata.Extrinsics[k].Hash = hexutil.Encode(h[:])
	}

	if blocknumber == 0 {
		return blockdata, nil
	}

	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return blockdata, err
	}

	var eventsBuf = make([]string, 0)
	var signer string
	var fee string

	for _, e := range events {
		if e.Phase.IsApplyExtrinsic {
			if name, ok = ExtrinsicsName[block.Block.Extrinsics[e.Phase.AsApplyExtrinsic].Method.CallIndex]; ok {
				if extrinsicIndex >= len(blockdata.Extrinsics) {
					return blockdata, errors.New("The number of extrinsics hashes does not equal the number of extrinsics")
				}
				if name == ExtName_Timestamp_set {
					extInfo = event.ExtrinsicsInfo{}
					timestamp, err := scale.NewDecoder(bytes.NewReader(block.Block.Extrinsics[e.Phase.AsApplyExtrinsic].Method.Args)).DecodeUintCompact()
					if err != nil {
						return blockdata, err
					}
					blockdata.Timestamp = timestamp.Int64()
					blockdata.Extrinsics[extrinsicIndex].Name = name
					blockdata.Extrinsics[extrinsicIndex].Events = []string{e.Name}
					blockdata.Extrinsics[extrinsicIndex].Result = true
					extrinsicIndex++
					continue
				}
				eventsBuf = append(eventsBuf, e.Name)
				if e.Name == event.TransactionPaymentTransactionFeePaid ||
					e.Name == event.EvmAccountMappingTransactionFeePaid {
					signer, fee, err = parseSignerAndFeePaidFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					tmp, ok := new(big.Int).SetString(fee, 10)
					if ok {
						allGasFee = allGasFee.Add(allGasFee, tmp)
					}
				} else if e.Name == event.BalancesTransfer {
					from, to, amount, err := ParseTransferInfoFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					if to == c.treasuryAcc {
						blockdata.Punishment = append(blockdata.Punishment, event.Punishment{
							ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
							From:          from,
							To:            to,
							Amount:        amount,
						})
					}
					blockdata.TransferInfo = append(blockdata.TransferInfo, event.TransferInfo{
						ExtrinsicName: name,
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						From:          from,
						To:            to,
						Amount:        amount,
						Result:        true,
					})
				} else if e.Name == event.SminerRegistered {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.MinerReg = append(blockdata.MinerReg, event.MinerRegInfo{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Account:       acc,
					})
				} else if e.Name == event.SystemNewAccount {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.NewAccounts = append(blockdata.NewAccounts, acc)
				} else if e.Name == event.FileBankUploadDeclaration {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					fid, err := ParseStringFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.UploadDecInfo = append(blockdata.UploadDecInfo, event.UploadDecInfo{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Owner:         acc,
						Fid:           fid,
					})
				} else if e.Name == event.FileBankDeleteFile {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					fid, err := ParseStringFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.DeleteFileInfo = append(blockdata.DeleteFileInfo, event.DeleteFileInfo{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Owner:         acc,
						Fid:           fid,
					})
				} else if e.Name == event.FileBankCreateBucket {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					bucketname, err := ParseStringFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.CreateBucketInfo = append(blockdata.CreateBucketInfo, event.CreateBucketInfo{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Owner:         acc,
						BucketName:    bucketname,
					})
				} else if e.Name == event.FileBankDeleteBucket {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					bucketname, err := ParseStringFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.DeleteBucketInfo = append(blockdata.DeleteBucketInfo, event.DeleteBucketInfo{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Owner:         acc,
						BucketName:    bucketname,
					})
				} else if e.Name == event.AuditSubmitIdleProof {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.SubmitIdleProve = append(blockdata.SubmitIdleProve, event.SubmitIdleProve{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Miner:         acc,
					})
				} else if e.Name == event.AuditSubmitServiceProof {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.SubmitServiceProve = append(blockdata.SubmitServiceProve, event.SubmitServiceProve{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Miner:         acc,
					})
				} else if e.Name == event.AuditSubmitIdleVerifyResult {
					acc, result, err := ParseChallResultFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.SubmitIdleResult = append(blockdata.SubmitIdleResult, event.SubmitIdleResult{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Miner:         acc,
						Result:        result,
					})
				} else if e.Name == event.AuditSubmitServiceVerifyResult {
					acc, result, err := ParseChallResultFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.SubmitServiceResult = append(blockdata.SubmitServiceResult, event.SubmitServiceResult{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Miner:         acc,
						Result:        result,
					})
				} else if e.Name == event.SystemExtrinsicSuccess {
					extInfo.Events = append(make([]string, 0), eventsBuf...)
					extInfo.Name = name
					extInfo.Signer = signer
					extInfo.FeePaid = fee
					extInfo.Result = true
					extInfo.Hash = blockdata.Extrinsics[extrinsicIndex].Hash
					blockdata.Extrinsics[extrinsicIndex] = extInfo
					eventsBuf = make([]string, 0)
					extInfo = event.ExtrinsicsInfo{}
					extrinsicIndex++
				} else if e.Name == event.SystemExtrinsicFailed {
					for m := 0; m < len(blockdata.UploadDecInfo); m++ {
						if blockdata.UploadDecInfo[m].ExtrinsicHash == blockdata.Extrinsics[extrinsicIndex].Hash {
							if len(blockdata.UploadDecInfo) == 1 {
								blockdata.UploadDecInfo = nil
							} else {
								blockdata.UploadDecInfo = append(blockdata.UploadDecInfo[:m], blockdata.UploadDecInfo[m+1:]...)
							}
							break
						}
					}
					for m := 0; m < len(blockdata.DeleteFileInfo); m++ {
						if blockdata.DeleteFileInfo[m].ExtrinsicHash == blockdata.Extrinsics[extrinsicIndex].Hash {
							if len(blockdata.DeleteFileInfo) == 1 {
								blockdata.DeleteFileInfo = nil
							} else {
								blockdata.DeleteFileInfo = append(blockdata.DeleteFileInfo[:m], blockdata.DeleteFileInfo[m+1:]...)
							}
							break
						}
					}
					for m := 0; m < len(blockdata.MinerReg); m++ {
						if blockdata.MinerReg[m].ExtrinsicHash == blockdata.Extrinsics[extrinsicIndex].Hash {
							if len(blockdata.MinerReg) == 1 {
								blockdata.MinerReg = nil
							} else {
								blockdata.MinerReg = append(blockdata.MinerReg[:m], blockdata.MinerReg[m+1:]...)
							}
							break
						}
					}
					for m := 0; m < len(blockdata.CreateBucketInfo); m++ {
						if blockdata.CreateBucketInfo[m].ExtrinsicHash == blockdata.Extrinsics[extrinsicIndex].Hash {
							if len(blockdata.CreateBucketInfo) == 1 {
								blockdata.CreateBucketInfo = nil
							} else {
								blockdata.CreateBucketInfo = append(blockdata.CreateBucketInfo[:m], blockdata.CreateBucketInfo[m+1:]...)
							}
							break
						}
					}
					for m := 0; m < len(blockdata.DeleteBucketInfo); m++ {
						if blockdata.DeleteBucketInfo[m].ExtrinsicHash == blockdata.Extrinsics[extrinsicIndex].Hash {
							if len(blockdata.DeleteBucketInfo) == 1 {
								blockdata.DeleteBucketInfo = nil
							} else {
								blockdata.DeleteBucketInfo = append(blockdata.DeleteBucketInfo[:m], blockdata.DeleteBucketInfo[m+1:]...)
							}
							break
						}
					}
					for m := 0; m < len(blockdata.SubmitIdleProve); m++ {
						if blockdata.SubmitIdleProve[m].ExtrinsicHash == blockdata.Extrinsics[extrinsicIndex].Hash {
							if len(blockdata.SubmitIdleProve) == 1 {
								blockdata.SubmitIdleProve = nil
							} else {
								blockdata.SubmitIdleProve = append(blockdata.SubmitIdleProve[:m], blockdata.SubmitIdleProve[m+1:]...)
							}
							break
						}
					}
					for m := 0; m < len(blockdata.SubmitServiceProve); m++ {
						if blockdata.SubmitServiceProve[m].ExtrinsicHash == blockdata.Extrinsics[extrinsicIndex].Hash {
							if len(blockdata.SubmitServiceProve) == 1 {
								blockdata.SubmitServiceProve = nil
							} else {
								blockdata.SubmitServiceProve = append(blockdata.SubmitServiceProve[:m], blockdata.SubmitServiceProve[m+1:]...)
							}
							break
						}
					}
					for m := 0; m < len(blockdata.SubmitIdleResult); m++ {
						if blockdata.SubmitIdleResult[m].ExtrinsicHash == blockdata.Extrinsics[extrinsicIndex].Hash {
							if len(blockdata.SubmitIdleResult) == 1 {
								blockdata.SubmitIdleResult = nil
							} else {
								blockdata.SubmitIdleResult = append(blockdata.SubmitIdleResult[:m], blockdata.SubmitIdleResult[m+1:]...)
							}
							break
						}
					}
					for m := 0; m < len(blockdata.SubmitServiceResult); m++ {
						if blockdata.SubmitServiceResult[m].ExtrinsicHash == blockdata.Extrinsics[extrinsicIndex].Hash {
							if len(blockdata.SubmitServiceResult) == 1 {
								blockdata.SubmitServiceResult = nil
							} else {
								blockdata.SubmitServiceResult = append(blockdata.SubmitServiceResult[:m], blockdata.SubmitServiceResult[m+1:]...)
							}
							break
						}
					}
					for m := 0; m < len(blockdata.Punishment); m++ {
						if blockdata.Punishment[m].ExtrinsicHash == blockdata.Extrinsics[extrinsicIndex].Hash {
							if len(blockdata.Punishment) == 1 {
								blockdata.Punishment = nil
							} else {
								blockdata.Punishment = append(blockdata.Punishment[:m], blockdata.Punishment[m+1:]...)
							}
							break
						}
					}
					extInfo.Events = append(make([]string, 0), eventsBuf...)
					extInfo.Name = name
					extInfo.Signer = signer
					extInfo.FeePaid = fee
					extInfo.Result = false
					extInfo.Hash = blockdata.Extrinsics[extrinsicIndex].Hash
					blockdata.Extrinsics[extrinsicIndex] = extInfo
					eventsBuf = make([]string, 0)
					extInfo = event.ExtrinsicsInfo{}
					extrinsicIndex++
				}
			}
		} else {
			blockdata.SysEvents = append(blockdata.SysEvents, e.Name)
			if e.Name == event.StakingStakersElected {
				blockdata.IsNewEra = true
			}
			if e.Name == event.AuditGenerateChallenge {
				acc, err := ParseAccountFromEvent(e)
				if err != nil {
					return blockdata, err
				}
				blockdata.GenChallenge = append(blockdata.GenChallenge, acc)
			}
		}
	}
	blockdata.AllGasFee = allGasFee.String()
	return blockdata, nil
}

func parseSignerAndFeePaidFromEvent(e *parser.Event) (string, string, error) {
	if e == nil {
		return "", "", errors.New("event is nil")
	}
	if e.Name != event.TransactionPaymentTransactionFeePaid &&
		e.Name != event.EvmAccountMappingTransactionFeePaid {
		return "", "", fmt.Errorf("event is not %s or %s", event.TransactionPaymentTransactionFeePaid, event.EvmAccountMappingTransactionFeePaid)
	}
	var signAcc string
	var fee string
	for _, v := range e.Fields {
		val := reflect.ValueOf(v.Value)
		if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
			signAcc = parseAccount(val)
		}
		if reflect.TypeOf(v.Value).Kind() == reflect.Struct {
			if strings.Contains(v.Name, "actual") {
				fee = ExplicitBigInt(val, 0)
			}
		}
	}
	if signAcc == "" {
		return signAcc, fee, fmt.Errorf("failed to parse transaction signer")
	}
	return signAcc, fee, nil
}

func parseAccount(v reflect.Value) string {
	var acc string
	if v.Len() > 0 {
		allValue := fmt.Sprintf("%v", v.Index(0))
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
		acc, _ = utils.EncodePublicKeyAsCessAccount(puk)
	}
	return acc
}

func parseFidString(v reflect.Value) string {
	if v.Len() > 0 {
		allValue := fmt.Sprintf("%v", v.Index(0))
		temp := strings.Split(allValue, "] ")
		puk := make([]byte, pattern.FileHashLen)
		for _, v := range temp {
			if strings.Count(v, " ") == (pattern.FileHashLen - 1) {
				subValue := strings.TrimPrefix(v, "[")
				ids := strings.Split(subValue, " ")
				if len(ids) != pattern.FileHashLen {
					continue
				}
				for kk, vv := range ids {
					intv, _ := strconv.Atoi(vv)
					puk[kk] = byte(intv)
				}
			}
		}
		return string(puk)
	}
	return ""
}

func parseBucketNameString(v reflect.Value) string {
	var value []byte
	for i := 0; i < v.Len(); i++ {
		intv, _ := strconv.Atoi(fmt.Sprintf("%v", v.Index(i)))
		value = append(value, byte(intv))
	}
	return string(value)
}

func ExplicitBigInt(v reflect.Value, depth int) string {
	var fee string
	if v.CanInterface() {
		value, ok := v.Interface().(big.Int)
		if ok {
			return value.String()
		}
		t := v.Type()
		switch v.Kind() {
		case reflect.Ptr:
			fee = ExplicitBigInt(v.Elem(), depth)
		case reflect.Struct:
			//fmt.Printf(strings.Repeat("\t", depth)+"%v %v {\n", t.Name(), t.Kind())
			for i := 0; i < v.NumField(); i++ {
				f := v.Field(i)
				if f.Kind() == reflect.Struct || f.Kind() == reflect.Ptr {
					//fmt.Printf(strings.Repeat("\t", depth+1)+"%s %s : \n", t.Field(i).Name, f.Type())
					fee = ExplicitBigInt(f, depth+2)
				} else {
					if f.CanInterface() {
						//fmt.Printf(strings.Repeat("\t", depth+1)+"%s %s : %v \n", t.Field(i).Name, f.Type(), f.Interface())
					} else {
						//fmt.Printf(strings.Repeat("\t", depth+1)+"%s %s : %v \n", t.Field(i).Name, f.Type(), f)
						if t.Field(i).Name == "abs" {
							val := fmt.Sprintf("%v", f)
							val = strings.TrimPrefix(val, "[")
							val = strings.TrimSuffix(val, "]")
							return val
						}
					}
				}
			}
			//fmt.Println(strings.Repeat("\t", depth) + "}")
		}
	}
	// else {
	// 	  fmt.Printf(strings.Repeat("\t", depth)+"%+v\n", v)
	// }
	return fee
}

func (c *ChainClient) parseTransferInfoFromBlock(blockhash types.Hash) ([]event.TransferInfo, error) {
	var transferEvents = make([]event.TransferInfo, 0)
	var data types.StorageDataRaw
	ok, err := c.GetSubstrateAPI().RPC.State.GetStorage(c.GetKeyEvents(), &data, blockhash)
	if err != nil {
		return transferEvents, err
	}
	var from string
	var to string
	if ok {
		events := event.EventRecords{}
		err = types.EventRecordsRaw(data).DecodeEventRecords(c.GetMetadata(), &events)
		if err != nil {
			return transferEvents, err
		}
		for _, e := range events.Balances_Transfer {
			from, _ = utils.EncodePublicKeyAsCessAccount(e.From[:])
			to, _ = utils.EncodePublicKeyAsCessAccount(e.To[:])
			transferEvents = append(transferEvents, event.TransferInfo{
				From:   from,
				To:     to,
				Amount: e.Value.String(),
				Result: true,
			})
		}
	}
	return transferEvents, nil
}

func ParseTransferInfoFromEvent(e *parser.Event) (string, string, string, error) {
	if e == nil {
		return "", "", "", errors.New("event is nil")
	}
	if e.Name != event.BalancesTransfer {
		return "", "", "", fmt.Errorf("event is not %s", event.BalancesTransfer)
	}
	var from string
	var to string
	var amount string
	for _, v := range e.Fields {
		k := reflect.TypeOf(v.Value).Kind()
		val := reflect.ValueOf(v.Value)
		if k == reflect.Slice {
			if strings.Contains(v.Name, "from") {
				from = parseAccount(val)
			}
			if strings.Contains(v.Name, "to") {
				to = parseAccount(val)
			}
		}
		if k == reflect.Struct {
			if v.Name == "amount" {
				amount = ExplicitBigInt(val, 0)
			}
		}
	}
	if from == "" || to == "" {
		return from, to, amount, fmt.Errorf("failed to parse from or to in transfer transactions")
	}

	return from, to, amount, nil
}

func ParseAccountFromEvent(e *parser.Event) (string, error) {
	if e == nil {
		return "", errors.New("event is nil")
	}
	var acc string
	for _, v := range e.Fields {
		k := reflect.TypeOf(v.Value).Kind()
		val := reflect.ValueOf(v.Value)
		if k == reflect.Slice {
			if strings.Contains(v.Name, "acc") ||
				strings.Contains(v.Name, "account") ||
				strings.Contains(v.Name, "owner") ||
				strings.Contains(v.Name, "miner") {
				acc = parseAccount(val)
			}
		}
	}
	if acc == "" {
		return acc, fmt.Errorf("failed to parse owner from file storage order transaction")
	}
	return acc, nil
}

func ParseChallResultFromEvent(e *parser.Event) (string, bool, error) {
	if e == nil {
		return "", false, errors.New("event is nil")
	}
	var acc string
	var result bool
	for _, v := range e.Fields {
		k := reflect.TypeOf(v.Value).Kind()
		val := reflect.ValueOf(v.Value)
		fmt.Println("k: ", k)
		fmt.Println("name: ", v.Name)
		if k == reflect.Slice {
			if strings.Contains(v.Name, "miner") {
				acc = parseAccount(val)
			}
		}
		if k == reflect.Bool {
			if strings.Contains(v.Name, "result") {
				result, _ = v.Value.(bool)
			}
		}
	}
	if acc == "" {
		return acc, false, fmt.Errorf("failed to parse owner from file storage order transaction")
	}
	return acc, result, nil
}

func ParseStringFromEvent(e *parser.Event) (string, error) {
	if e == nil {
		return "", fmt.Errorf("ParseFidFromEvent: event is nil")
	}
	var value string
	for _, v := range e.Fields {
		k := reflect.TypeOf(v.Value).Kind()
		val := reflect.ValueOf(v.Value)
		if k == reflect.Slice {
			if strings.Contains(v.Name, "hash") {
				value = parseFidString(val)
			}
			if strings.Contains(v.Name, "bucket") {
				value = parseBucketNameString(val)
			}
		}
	}
	if value == "" {
		return value, fmt.Errorf("failed to parse string")
	}
	return value, nil
}
