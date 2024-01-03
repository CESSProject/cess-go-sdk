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
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/utils"
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

func (c *chainClient) RetrieveEvent_Audit_SubmitIdleProof(blockhash types.Hash) (event.Event_SubmitIdleProof, error) {
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

func (c *chainClient) RetrieveEvent_Audit_SubmitServiceProof(blockhash types.Hash) (event.Event_SubmitServiceProof, error) {
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

func (c *chainClient) RetrieveEvent_Audit_SubmitIdleVerifyResult(blockhash types.Hash) (event.Event_SubmitIdleVerifyResult, error) {
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
							result.Tee = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.AuditSubmitIdleVerifyResult)
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
							result.Tee = *accid
							return result, nil
						}
					}
				}
			}
		}
	}
	return result, errors.Errorf("failed: no %s event found", event.AuditSubmitServiceVerifyResult)
}

func (c *chainClient) RetrieveEvent_Oss_OssUpdate(blockhash types.Hash) (event.Event_OssUpdate, error) {
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

func (c *chainClient) RetrieveEvent_Oss_OssRegister(blockhash types.Hash) (event.Event_OssRegister, error) {
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

func (c *chainClient) RetrieveEvent_Oss_OssDestroy(blockhash types.Hash) (event.Event_OssDestroy, error) {
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

func (c *chainClient) RetrieveEvent_Oss_Authorize(blockhash types.Hash) (event.Event_Authorize, error) {
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

func (c *chainClient) RetrieveEvent_Oss_CancelAuthorize(blockhash types.Hash) (event.Event_CancelAuthorize, error) {
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

func (c *chainClient) RetrieveEvent_FileBank_UploadDeclaration(blockhash types.Hash) (event.Event_UploadDeclaration, error) {
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

func (c *chainClient) RetrieveEvent_FileBank_CreateBucket(blockhash types.Hash) (event.Event_CreateBucket, error) {
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

func (c *chainClient) RetrieveEvent_FileBank_DeleteBucket(blockhash types.Hash) (event.Event_DeleteBucket, error) {
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

func (c *chainClient) RetrieveEvent_FileBank_DeleteFile(blockhash types.Hash) (event.Event_DeleteFile, error) {
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

func (c *chainClient) RetrieveEvent_FileBank_TransferReport(blockhash types.Hash) (event.Event_TransferReport, error) {
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

func (c *chainClient) RetrieveEvent_FileBank_RecoveryCompleted(blockhash types.Hash) (event.Event_RecoveryCompleted, error) {
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

func (c *chainClient) RetrieveEvent_FileBank_IdleSpaceCert(blockhash types.Hash) (event.Event_IdleSpaceCert, error) {
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

func (c *chainClient) RetrieveEvent_FileBank_ReplaceIdleSpace(blockhash types.Hash) (event.Event_ReplaceIdleSpace, error) {
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

func (c *chainClient) RetrieveEvent_FileBank_CalculateReport(blockhash types.Hash) (event.Event_CalculateReport, error) {
	var result event.Event_CalculateReport
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		fmt.Println(e.Name)
		if e.Name == event.FileBankCalculateReport {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
					fmt.Println("1")
					vf := reflect.ValueOf(v.Value)
					if vf.Len() > 0 {
						fmt.Println("2")
						allValue := fmt.Sprintf("%v", vf.Index(0))
						if strings.Contains(v.Name, "AccountId32") {
							temp := strings.Split(allValue, "] ")
							fmt.Println("3: ", temp)
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
							fmt.Println("4: ", puk)
							fmt.Println("4: ", c.GetSignatureAccPulickey())
							if !utils.CompareSlice(puk, c.GetSignatureAccPulickey()) {
								continue
							}
							fmt.Println("5")
							accid, err := types.NewAccountID(puk)
							if err != nil {
								continue
							}
							fmt.Println("6")
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

func (c *chainClient) RetrieveEvent_Sminer_UpdataIp(blockhash types.Hash) (event.Event_UpdataIp, error) {
	var result event.Event_UpdataIp
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

func (c *chainClient) RetrieveEvent_Sminer_UpdataBeneficiary(blockhash types.Hash) (event.Event_UpdataBeneficiary, error) {
	var result event.Event_UpdataBeneficiary
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

func (c *chainClient) RetrieveEvent_Sminer_Registered(blockhash types.Hash) (event.Event_Registered, error) {
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

func (c *chainClient) RetrieveEvent_Sminer_RegisterPoisKey(blockhash types.Hash) (event.Event_RegisterPoisKey, error) {
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

func (c *chainClient) RetrieveEvent_Sminer_MinerExitPrep(blockhash types.Hash) (event.Event_MinerExitPrep, error) {
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

func (c *chainClient) RetrieveEvent_Sminer_IncreaseCollateral(blockhash types.Hash) (event.Event_IncreaseCollateral, error) {
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

func (c *chainClient) RetrieveEvent_Sminer_IncreaseDeclarationSpace(blockhash types.Hash) (event.Event_IncreaseDeclarationSpace, error) {
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

func (c *chainClient) RetrieveEvent_Sminer_Receive(blockhash types.Hash) (event.Event_Receive, error) {
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

func (c *chainClient) RetrieveEvent_Sminer_Withdraw(blockhash types.Hash) (event.Event_Withdraw, error) {
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

func (c *chainClient) RetrieveEvent_StorageHandler_BuySpace(blockhash types.Hash) (event.Event_BuySpace, error) {
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

func (c *chainClient) RetrieveEvent_StorageHandler_ExpansionSpace(blockhash types.Hash) (event.Event_ExpansionSpace, error) {
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

func (c *chainClient) RetrieveEvent_StorageHandler_RenewalSpace(blockhash types.Hash) (event.Event_RenewalSpace, error) {
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

func (c *chainClient) RetrieveEvent_Balances_Transfer(blockhash types.Hash) (types.EventBalancesTransfer, error) {
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

func (c *chainClient) RetrieveEvent_FileBank_GenRestoralOrder(blockhash types.Hash) (event.Event_GenerateRestoralOrder, error) {
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

func (c *chainClient) RetrieveAllEvent_FileBank_UploadDeclaration(blockhash types.Hash) ([]event.AllUploadDeclarationEvent, error) {
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

func (c *chainClient) RetrieveAllEvent_FileBank_StorageCompleted(blockhash types.Hash) ([]string, error) {
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

func (c *chainClient) RetrieveAllEvent_FileBank_DeleteFile(blockhash types.Hash) ([]event.AllDeleteFileEvent, error) {
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
