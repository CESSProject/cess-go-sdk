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

	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
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

func (c *ChainClient) RetrieveEvent_FileBank_ClaimRestoralOrder(blockhash types.Hash) (Event_ClaimRestoralOrder, error) {
	var result Event_ClaimRestoralOrder
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == FileBankClaimRestoralOrder {
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
	return result, errors.Errorf("failed: no %s event found", FileBankClaimRestoralOrder)
}

func (c *ChainClient) RetrieveEvent_Audit_SubmitIdleProof(blockhash types.Hash) (Event_SubmitIdleProof, error) {
	var result Event_SubmitIdleProof
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == AuditSubmitIdleProof {
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
	return result, errors.Errorf("failed: no %s event found", AuditSubmitIdleProof)
}

func (c *ChainClient) RetrieveEvent_Audit_SubmitServiceProof(blockhash types.Hash) (Event_SubmitServiceProof, error) {
	var result Event_SubmitServiceProof
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == AuditSubmitServiceProof {
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
	return result, errors.Errorf("failed: no %s event found", AuditSubmitServiceProof)
}

func (c *ChainClient) RetrieveEvent_Audit_SubmitIdleVerifyResult(blockhash types.Hash) (Event_SubmitIdleVerifyResult, error) {
	var result Event_SubmitIdleVerifyResult
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == AuditSubmitIdleVerifyResult {
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
	return result, errors.Errorf("failed: no %s event found", AuditSubmitIdleVerifyResult)
}

func (c *ChainClient) RetrieveEvent_Audit_SubmitServiceVerifyResult(blockhash types.Hash) (Event_SubmitServiceVerifyResult, error) {
	var result Event_SubmitServiceVerifyResult
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == AuditSubmitServiceVerifyResult {
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
	return result, errors.Errorf("failed: no %s event found", AuditSubmitServiceVerifyResult)
}

func (c *ChainClient) RetrieveEvent_Oss_OssUpdate(blockhash types.Hash) (Event_OssUpdate, error) {
	var result Event_OssUpdate
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == OssOssUpdate {
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
	return result, errors.Errorf("failed: no %s event found", OssOssUpdate)
}

func (c *ChainClient) RetrieveEvent_Oss_OssRegister(blockhash types.Hash) (Event_OssRegister, error) {
	var result Event_OssRegister
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == OssOssRegister {
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
	return result, errors.Errorf("failed: no %s event found", OssOssRegister)
}

func (c *ChainClient) RetrieveEvent_Oss_OssDestroy(blockhash types.Hash) (Event_OssDestroy, error) {
	var result Event_OssDestroy
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == OssOssDestroy {
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
	return result, errors.Errorf("failed: no %s event found", OssOssDestroy)
}

func (c *ChainClient) RetrieveEvent_Oss_Authorize(blockhash types.Hash) (Event_Authorize, error) {
	var result Event_Authorize
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == OssAuthorize {
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
	return result, errors.Errorf("failed: no %s event found", OssAuthorize)
}

func (c *ChainClient) RetrieveEvent_Oss_CancelAuthorize(blockhash types.Hash) (Event_CancelAuthorize, error) {
	var result Event_CancelAuthorize
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == OssCancelAuthorize {
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
	return result, errors.Errorf("failed: no %s event found", OssCancelAuthorize)
}

func (c *ChainClient) RetrieveEvent_FileBank_UploadDeclaration(blockhash types.Hash) (Event_UploadDeclaration, error) {
	var result Event_UploadDeclaration
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == FileBankUploadDeclaration {
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
	return result, errors.Errorf("failed: no %s event found", FileBankUploadDeclaration)
}

func (c *ChainClient) RetrieveEvent_FileBank_CreateBucket(blockhash types.Hash) (Event_CreateBucket, error) {
	var result Event_CreateBucket
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == FileBankCreateBucket {
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
	return result, errors.Errorf("failed: no %s event found", FileBankCreateBucket)
}

func (c *ChainClient) RetrieveEvent_FileBank_DeleteBucket(blockhash types.Hash) (Event_DeleteBucket, error) {
	var result Event_DeleteBucket
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == FileBankDeleteBucket {
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
	return result, errors.Errorf("failed: no %s event found", FileBankDeleteBucket)
}

func (c *ChainClient) RetrieveEvent_FileBank_DeleteFile(blockhash types.Hash) (Event_DeleteFile, error) {
	var result Event_DeleteFile
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == FileBankDeleteFile {
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
	return result, errors.Errorf("failed: no %s event found", FileBankDeleteFile)
}

func (c *ChainClient) RetrieveEvent_FileBank_TransferReport(blockhash types.Hash) (Event_TransferReport, error) {
	var result Event_TransferReport
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == FileBankTransferReport {
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
	return result, errors.Errorf("failed: no %s event found", FileBankTransferReport)
}

func (c *ChainClient) RetrieveEvent_FileBank_RecoveryCompleted(blockhash types.Hash) (Event_RecoveryCompleted, error) {
	var result Event_RecoveryCompleted
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == FileBankRecoveryCompleted {
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
	return result, errors.Errorf("failed: no %s event found", FileBankRecoveryCompleted)
}

func (c *ChainClient) RetrieveEvent_FileBank_IdleSpaceCert(blockhash types.Hash) (Event_IdleSpaceCert, error) {
	var result Event_IdleSpaceCert
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == FileBankIdleSpaceCert {
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
	return result, errors.Errorf("failed: no %s event found", FileBankIdleSpaceCert)
}

func (c *ChainClient) RetrieveEvent_FileBank_ReplaceIdleSpace(blockhash types.Hash) (Event_ReplaceIdleSpace, error) {
	var result Event_ReplaceIdleSpace
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == FileBankReplaceIdleSpace {
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
	return result, errors.Errorf("failed: no %s event found", FileBankReplaceIdleSpace)
}

func (c *ChainClient) RetrieveEvent_FileBank_CalculateReport(blockhash types.Hash) (Event_CalculateReport, error) {
	var result Event_CalculateReport
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == FileBankCalculateReport {
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
	return result, errors.Errorf("failed: no %s event found", FileBankCalculateReport)
}

func (c *ChainClient) RetrieveEvent_Sminer_UpdataIp(blockhash types.Hash) (Event_UpdatePeerId, error) {
	var result Event_UpdatePeerId
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == SminerUpdatePeerId {
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
	return result, errors.Errorf("failed: no %s event found", SminerUpdatePeerId)
}

func (c *ChainClient) RetrieveEvent_Sminer_UpdataBeneficiary(blockhash types.Hash) (Event_UpdateBeneficiary, error) {
	var result Event_UpdateBeneficiary
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == SminerUpdateBeneficiary {
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
	return result, errors.Errorf("failed: no %s event found", SminerUpdateBeneficiary)
}

func (c *ChainClient) RetrieveEvent_Sminer_Registered(blockhash types.Hash) (Event_Registered, error) {
	var result Event_Registered
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == SminerRegistered {
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
	return result, errors.Errorf("failed: no %s event found", SminerRegistered)
}

func (c *ChainClient) RetrieveEvent_Sminer_RegisterPoisKey(blockhash types.Hash) (Event_RegisterPoisKey, error) {
	var result Event_RegisterPoisKey
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == SminerRegisterPoisKey {
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
	return result, errors.Errorf("failed: no %s event found", SminerRegisterPoisKey)
}

func (c *ChainClient) RetrieveEvent_Sminer_MinerExitPrep(blockhash types.Hash) (Event_MinerExitPrep, error) {
	var result Event_MinerExitPrep
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == SminerMinerExitPrep {
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
	return result, errors.Errorf("failed: no %s event found", SminerMinerExitPrep)
}

func (c *ChainClient) RetrieveEvent_Sminer_IncreaseCollateral(blockhash types.Hash) (Event_IncreaseCollateral, error) {
	var result Event_IncreaseCollateral
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == SminerIncreaseCollateral {
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
	return result, errors.Errorf("failed: no %s event found", SminerIncreaseCollateral)
}

func (c *ChainClient) RetrieveEvent_Sminer_IncreaseDeclarationSpace(blockhash types.Hash) (Event_IncreaseDeclarationSpace, error) {
	var result Event_IncreaseDeclarationSpace
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == SminerIncreaseDeclarationSpace {
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
	return result, errors.Errorf("failed: no %s event found", SminerIncreaseDeclarationSpace)
}

func (c *ChainClient) RetrieveEvent_Sminer_Receive(blockhash types.Hash) (Event_Receive, error) {
	var result Event_Receive

	block, err := c.api.RPC.Chain.GetBlock(blockhash)
	if err != nil {
		return result, err
	}

	fmt.Println("block number: ", block.Block.Header.Number)
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}

	var signer string
	var earningsAcc string
	fmt.Println("c.GetSignatureAcc():", c.GetSignatureAcc())
	for _, e := range events {
		fmt.Println("e.Name: ", e.Name)
		if e.Phase.IsApplyExtrinsic {
			if name, ok := ExtrinsicsName[block.Block.Extrinsics[e.Phase.AsApplyExtrinsic].Method.CallIndex]; ok {
				if name == ExtName_Sminer_receive_reward {
					switch e.Name {
					case SminerReceive:
						earningsAcc, _ = ParseAccountFromEvent(e)
						fmt.Println("earningsAcc: ", earningsAcc)
						result.Acc = earningsAcc
					case TransactionPaymentTransactionFeePaid, EvmAccountMappingTransactionFeePaid:
						signer, _, _ = parseSignerAndFeePaidFromEvent(e)
						fmt.Println("signer: ", signer)
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

func (c *ChainClient) RetrieveEvent_Sminer_Withdraw(blockhash types.Hash) (Event_Withdraw, error) {
	var result Event_Withdraw
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == SminerWithdraw {
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
	return result, errors.Errorf("failed: no %s event found", SminerWithdraw)
}

func (c *ChainClient) RetrieveEvent_StorageHandler_BuySpace(blockhash types.Hash) (Event_BuySpace, error) {
	var result Event_BuySpace
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == StorageHandlerBuySpace {
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
	return result, errors.Errorf("failed: no %s event found", StorageHandlerBuySpace)
}

func (c *ChainClient) RetrieveEvent_StorageHandler_ExpansionSpace(blockhash types.Hash) (Event_ExpansionSpace, error) {
	var result Event_ExpansionSpace
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == StorageHandlerExpansionSpace {
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
	return result, errors.Errorf("failed: no %s event found", StorageHandlerExpansionSpace)
}

func (c *ChainClient) RetrieveEvent_StorageHandler_RenewalSpace(blockhash types.Hash) (Event_RenewalSpace, error) {
	var result Event_RenewalSpace
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == StorageHandlerRenewalSpace {
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
	return result, errors.Errorf("failed: no %s event found", StorageHandlerRenewalSpace)
}

func (c *ChainClient) RetrieveEvent_Balances_Transfer(blockhash types.Hash) (types.EventBalancesTransfer, error) {
	var result types.EventBalancesTransfer
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == BalancesTransfer {
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
	return result, errors.Errorf("failed: no %s event found", BalancesTransfer)
}

func (c *ChainClient) RetrieveEvent_FileBank_GenRestoralOrder(blockhash types.Hash) (Event_GenerateRestoralOrder, error) {
	var result Event_GenerateRestoralOrder
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == FileBankGenerateRestoralOrder {
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
	return result, errors.Errorf("failed: no %s event found", FileBankGenerateRestoralOrder)
}

func (c *ChainClient) RetrieveAllEvent_FileBank_UploadDeclaration(blockhash types.Hash) ([]AllUploadDeclarationEvent, error) {
	var result = make([]AllUploadDeclarationEvent, 0)
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}

	for _, e := range events {
		if e.Name == FileBankUploadDeclaration {
			var ele AllUploadDeclarationEvent
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
								if strings.Count(v, " ") == (FileHashLen - 1) {
									subValue := strings.TrimPrefix(v, "[")
									ids := strings.Split(subValue, " ")
									if len(ids) != FileHashLen {
										continue
									}
									var fhash FileHash
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
		if e.Name == FileBankStorageCompleted {
			for _, v := range e.Fields {
				if reflect.TypeOf(v.Value).Kind() == reflect.Slice {
					vf := reflect.ValueOf(v.Value)
					if vf.Len() > 0 {
						allValue := fmt.Sprintf("%v", vf.Index(0))
						if strings.Contains(v.Name, "file_hash") {
							temp := strings.Split(allValue, "] ")
							for _, v := range temp {
								if strings.Count(v, " ") == (FileHashLen - 1) {
									subValue := strings.TrimPrefix(v, "[")
									ids := strings.Split(subValue, " ")
									if len(ids) != FileHashLen {
										continue
									}
									var fhash FileHash
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

func (c *ChainClient) RetrieveAllEvent_FileBank_DeleteFile(blockhash types.Hash) ([]AllDeleteFileEvent, error) {
	var result = make([]AllDeleteFileEvent, 0)
	events, err := c.eventRetriever.GetEvents(blockhash)
	if err != nil {
		return result, err
	}
	for _, e := range events {
		if e.Name == FileBankDeleteFile {
			var ele AllDeleteFileEvent
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
								if strings.Count(v, " ") == (FileHashLen - 1) {
									subValue := strings.TrimPrefix(v, "[")
									ids := strings.Split(subValue, " ")
									if len(ids) != FileHashLen {
										continue
									}
									var fhash FileHash
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
