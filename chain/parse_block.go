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

	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/registry/parser"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/vedhavyas/go-subkey/scale"
	"golang.org/x/crypto/blake2b"
)

func (c *ChainClient) ParseBlockData(blocknumber uint64) (BlockData, error) {
	var (
		ok             bool
		name           string
		err            error
		extBytes       []byte
		extrinsicIndex int
		blockdata      BlockData
		extInfo        ExtrinsicsInfo
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

	blockdata.Extrinsics = make([]ExtrinsicsInfo, len(block.Block.Extrinsics))
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
					extInfo = ExtrinsicsInfo{}
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
				if e.Name == TransactionPaymentTransactionFeePaid ||
					e.Name == EvmAccountMappingTransactionFeePaid {
					signer, fee, err = parseSignerAndFeePaidFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					tmp, ok := new(big.Int).SetString(fee, 10)
					if ok {
						allGasFee = allGasFee.Add(allGasFee, tmp)
					}
				} else if e.Name == BalancesTransfer {
					from, to, amount, err := ParseTransferInfoFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					if to == TreasuryAccount {
						blockdata.Punishment = append(blockdata.Punishment, Punishment{
							ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
							From:          from,
							To:            to,
							Amount:        amount,
						})
					}
					blockdata.TransferInfo = append(blockdata.TransferInfo, TransferInfo{
						ExtrinsicName: name,
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						From:          from,
						To:            to,
						Amount:        amount,
						Result:        true,
					})
				} else if e.Name == SminerRegistered {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.MinerReg = append(blockdata.MinerReg, MinerRegInfo{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Account:       acc,
					})
				} else if e.Name == SystemNewAccount {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.NewAccounts = append(blockdata.NewAccounts, acc)
				} else if e.Name == FileBankUploadDeclaration {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					fid, err := ParseStringFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.UploadDecInfo = append(blockdata.UploadDecInfo, UploadDecInfo{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Owner:         acc,
						Fid:           fid,
					})
				} else if e.Name == FileBankDeleteFile {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					fid, err := ParseStringFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.DeleteFileInfo = append(blockdata.DeleteFileInfo, DeleteFileInfo{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Owner:         acc,
						Fid:           fid,
					})
				} else if e.Name == FileBankCreateBucket {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					bucketname, err := ParseStringFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.CreateBucketInfo = append(blockdata.CreateBucketInfo, CreateBucketInfo{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Owner:         acc,
						BucketName:    bucketname,
					})
				} else if e.Name == FileBankDeleteBucket {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					bucketname, err := ParseStringFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.DeleteBucketInfo = append(blockdata.DeleteBucketInfo, DeleteBucketInfo{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Owner:         acc,
						BucketName:    bucketname,
					})
				} else if e.Name == AuditSubmitIdleProof {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.SubmitIdleProve = append(blockdata.SubmitIdleProve, SubmitIdleProve{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Miner:         acc,
					})
				} else if e.Name == AuditSubmitServiceProof {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.SubmitServiceProve = append(blockdata.SubmitServiceProve, SubmitServiceProve{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Miner:         acc,
					})
				} else if e.Name == AuditSubmitIdleVerifyResult {
					acc, result, err := ParseChallResultFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.SubmitIdleResult = append(blockdata.SubmitIdleResult, SubmitIdleResult{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Miner:         acc,
						Result:        result,
					})
				} else if e.Name == AuditSubmitServiceVerifyResult {
					acc, result, err := ParseChallResultFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.SubmitServiceResult = append(blockdata.SubmitServiceResult, SubmitServiceResult{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Miner:         acc,
						Result:        result,
					})
				} else if e.Name == SminerRegisterPoisKey {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.MinerRegPoiskeys = append(blockdata.MinerRegPoiskeys, MinerRegPoiskey{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Miner:         acc,
					})
				} else if e.Name == OssOssRegister {
					acc, err := ParseAccountFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.GatewayReg = append(blockdata.GatewayReg, GatewayReg{
						ExtrinsicHash: blockdata.Extrinsics[extrinsicIndex].Hash,
						Account:       acc,
					})
				} else if e.Name == FileBankStorageCompleted {
					fid, err := ParseStringFromEvent(e)
					if err != nil {
						return blockdata, err
					}
					blockdata.StorageCompleted = append(blockdata.StorageCompleted, fid)
				} else if e.Name == SystemExtrinsicSuccess {
					extInfo.Events = append(make([]string, 0), eventsBuf...)
					extInfo.Name = name
					extInfo.Signer = signer
					extInfo.FeePaid = fee
					extInfo.Result = true
					extInfo.Hash = blockdata.Extrinsics[extrinsicIndex].Hash
					blockdata.Extrinsics[extrinsicIndex] = extInfo
					eventsBuf = make([]string, 0)
					extInfo = ExtrinsicsInfo{}
					extrinsicIndex++
				} else if e.Name == SystemExtrinsicFailed {
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
					for m := 0; m < len(blockdata.MinerRegPoiskeys); m++ {
						if blockdata.MinerRegPoiskeys[m].ExtrinsicHash == blockdata.Extrinsics[extrinsicIndex].Hash {
							if len(blockdata.MinerRegPoiskeys) == 1 {
								blockdata.MinerRegPoiskeys = nil
							} else {
								blockdata.MinerRegPoiskeys = append(blockdata.MinerRegPoiskeys[:m], blockdata.MinerRegPoiskeys[m+1:]...)
							}
							break
						}
					}
					for m := 0; m < len(blockdata.GatewayReg); m++ {
						if blockdata.GatewayReg[m].ExtrinsicHash == blockdata.Extrinsics[extrinsicIndex].Hash {
							if len(blockdata.GatewayReg) == 1 {
								blockdata.GatewayReg = nil
							} else {
								blockdata.GatewayReg = append(blockdata.GatewayReg[:m], blockdata.GatewayReg[m+1:]...)
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
					extInfo = ExtrinsicsInfo{}
					extrinsicIndex++
				}
			}
		} else {
			blockdata.SysEvents = append(blockdata.SysEvents, e.Name)
			if e.Name == StakingStakersElected {
				blockdata.IsNewEra = true
			}
			if e.Name == AuditGenerateChallenge {
				acc, err := ParseAccountFromEvent(e)
				if err != nil {
					return blockdata, err
				}
				blockdata.GenChallenge = append(blockdata.GenChallenge, acc)
			}
			if e.Name == StakingEraPaid {
				eraIndex, validatorPayout, remainder, err := ParseStakingEraPaidFromEvent(e)
				if err != nil {
					return blockdata, err
				}
				blockdata.EraPaid = EraPaid{
					HaveValue:       true,
					EraIndex:        eraIndex,
					ValidatorPayout: validatorPayout,
					Remainder:       remainder,
				}
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
	if e.Name != TransactionPaymentTransactionFeePaid &&
		e.Name != EvmAccountMappingTransactionFeePaid {
		return "", "", fmt.Errorf("event is not %s or %s", TransactionPaymentTransactionFeePaid, EvmAccountMappingTransactionFeePaid)
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
		puk := make([]byte, FileHashLen)
		for _, v := range temp {
			if strings.Count(v, " ") == (FileHashLen - 1) {
				subValue := strings.TrimPrefix(v, "[")
				ids := strings.Split(subValue, " ")
				if len(ids) != FileHashLen {
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

func ParseTransferInfoFromEvent(e *parser.Event) (string, string, string, error) {
	if e == nil {
		return "", "", "", errors.New("event is nil")
	}
	if e.Name != BalancesTransfer {
		return "", "", "", fmt.Errorf("event is not %s", BalancesTransfer)
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

func ParseStakingEraPaidFromEvent(e *parser.Event) (uint32, string, string, error) {
	var eraIndex uint32
	var validatorPayout string
	var remainder string
	for _, v := range e.Fields {
		k := reflect.TypeOf(v.Value).Kind()
		val := reflect.ValueOf(v.Value)
		if k == reflect.Uint32 {
			if strings.Contains(v.Name, "era_index") {
				eraid, err := strconv.ParseUint(fmt.Sprintf("%v", val), 10, 32)
				if err != nil {
					return 0, "", "", err
				}
				eraIndex = uint32(eraid)
			}

		}
		if k == reflect.Struct {
			if strings.Contains(v.Name, "validator_payout") {
				validatorPayout = ExplicitBigInt(val, 0)
			}
			if strings.Contains(v.Name, "remainder") {
				remainder = ExplicitBigInt(val, 0)
			}
		}
	}
	return eraIndex, validatorPayout, remainder, nil
}
