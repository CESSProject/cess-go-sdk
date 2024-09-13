/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package process

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/CESSProject/cess-go-sdk/chain"
	"github.com/CESSProject/cess-go-sdk/core/crypte"
	"github.com/CESSProject/cess-go-sdk/core/erasure"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// StoreFileToMiners store a file to some miners
//
// Receive parameter:
//   - file: stored file
//   - mnemonic: account mnemonic
//   - territory: territory name
//   - bucket: bucket name
//   - timeout: timeout for waiting for block transaction to complete
//   - rpcs: rpc address list
//   - wantMiner: the wallet account of the miner you want to store. if it is empty, will be randomly selected.
//
// Return parameter:
//   - string: [fid] unique identifier for the file
//   - error: error message
//
// Preconditions:
//  1. your account needs to have money, and will be automatically created if the territory you specify does not exist.
//  2. if the number of miners you specify is less than 12, file storage will be exited if even one fails.
//  3. if the number of miners you specify is greater than 11, no other miners will be found for storage.
func StoreFileToMiners(file string, mnemonic string, territory string, bucket string, timeout time.Duration, rpcs []string, wantMiner []string) (string, error) {
	size, err := CheckFile(file)
	if err != nil {
		return "", err
	}

	if !chain.CheckBucketName(bucket) {
		return "", errors.New("invalid bucket name")
	}

	cacheDir := filepath.Join(filepath.Dir(file), fmt.Sprintf("%v", time.Now().UnixMilli()))
	err = os.MkdirAll(cacheDir, 0755)
	if err != nil {
		return "", err
	}
	defer func() {
		os.RemoveAll(cacheDir)
	}()

	segmentInfo, fid, err := FullProcessing(file, "", cacheDir)
	if err != nil {
		return "", err
	}

	if mnemonic == "" {
		return fid, errors.New("empty mnemonic")
	}

	cli, err := chain.NewChainClient(context.Background(), "", rpcs, mnemonic, timeout)
	if err != nil {
		return fid, err
	}
	defer cli.Close()

	err = CheckAccount(cli, territory, size)
	if err != nil {
		return fid, err
	}

	_, err = cli.PlaceStorageOrder(fid, filepath.Base(file), bucket, territory, segmentInfo, cli.GetSignatureAccPulickey(), uint64(size))
	if err != nil {
		return fid, err
	}

	segmentlength := len(segmentInfo)
	var fragmentGroup = make([][]string, chain.DataShards+chain.ParShards)
	for j := 0; j < chain.DataShards+chain.ParShards; j++ {
		fragmentGroup[j] = make([]string, segmentlength)
		for i := 0; i < segmentlength; i++ {
			fragmentGroup[j][i] = segmentInfo[i].FragmentHash[j]
		}
	}

	if len(wantMiner) >= (chain.DataShards + chain.ParShards) {
		return fid, StoreToAllDesignatedMiners(cli, fragmentGroup, fid, wantMiner)
	}
	err = StorageToMiners(cli, fragmentGroup, fid, wantMiner)
	return fid, err
}

// RetrieveFileFromMiners Retrieve a storaged file from storage miners
//
// Receive parameter:
//   - rpcs: rpc address list
//   - fid: [fid] unique identifier for the file
//   - cipher: decryption password, if any
//   - savedir: file save directory, final save location: <savedir>/<fid>
//
// Return parameter:
//   - error: error message
//
// Preconditions:
//  1. the file to be downloaded needs to have been stored in the miner
func RetrieveFileFromMiners(rpcs []string, fid string, cipher, savedir string) error {
	cli, err := chain.NewChainClient(context.Background(), "", rpcs, "", 0)
	if err != nil {
		return err
	}
	defer cli.Close()

	metaInfo, err := cli.QueryFile(fid, -1)
	if err != nil {
		if errors.Is(err, chain.ERR_RPC_EMPTY_VALUE) {
			return errors.New("not found")
		}
		return err
	}
	_, err = Retrievefile(cli, metaInfo, fid, savedir, cipher)
	return err
}

func Retrievefile(cli chain.Chainer, fmeta chain.FileMetadata, fid, savedir, cipher string) (string, error) {
	userfile := filepath.Join(savedir, fid)
	fstat, err := os.Stat(userfile)
	if err == nil {
		if fstat.Size() > 0 {
			return userfile, nil
		}
	}
	err = os.MkdirAll(savedir, 0755)
	if err != nil {
		return userfile, err
	}

	defer func(basedir string) {
		for _, segment := range fmeta.SegmentList {
			os.Remove(filepath.Join(basedir, string(segment.Hash[:])))
			for _, fragment := range segment.FragmentList {
				os.Remove(filepath.Join(basedir, string(fragment.Hash[:])))
			}
		}
	}(savedir)

	var segmentspath = make([]string, 0)
	fragmentpaths := make([]string, chain.DataShards+chain.ParShards)

	for _, segment := range fmeta.SegmentList {
		for k, fragment := range segment.FragmentList {
			fragmentpath := filepath.Join(savedir, string(fragment.Hash[:]))
			fragmentpaths[k] = fragmentpath
			if string(fragment.Hash[:]) != chain.ZeroFileHash_8M {
				account, err := utils.EncodePublicKeyAsCessAccount(fragment.Miner[:])
				if err != nil {
					return userfile, err
				}
				buf, err := DownloadFragmentFromMiner(cli, fragment.Miner[:], fid, string(fragment.Hash[:]))
				if err != nil {
					return userfile, fmt.Errorf("download from [%s] failed: %v", account, err)
				}
				err = utils.WriteBufToFile(buf, fragmentpath)
				if err != nil {
					return userfile, err
				}
			} else {
				_, err = os.Stat(fragmentpath)
				if err != nil {
					ff, err := os.Create(fragmentpath)
					if err != nil {
						return userfile, err
					}
					_, err = ff.Write(make([]byte, chain.FragmentSize))
					if err != nil {
						return userfile, err
					}
					err = ff.Sync()
					if err != nil {
						return userfile, err
					}
					err = ff.Close()
					if err != nil {
						return userfile, err
					}
				}
			}
		}
		segmentpath := filepath.Join(savedir, string(segment.Hash[:]))
		err = erasure.RSRestore(segmentpath, fragmentpaths)
		if err != nil {
			return "", err
		}
		segmentspath = append(segmentspath, segmentpath)
	}

	if len(segmentspath) != len(fmeta.SegmentList) {
		return "", errors.New("download failed")
	}

	fd, err := os.Create(userfile)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	var writecount = 0
	for i := 0; i < len(segmentspath); i++ {
		buf, err := os.ReadFile(segmentspath[i])
		if err != nil {
			return "", err
		}
		if cipher != "" {
			buffer, err := crypte.AesCbcDecrypt(buf, []byte(cipher))
			if err != nil {
				return "", err
			}
			if (writecount + 1) >= len(fmeta.SegmentList) {
				fd.Write(buffer[:(fmeta.FileSize.Uint64() - uint64(writecount*chain.SegmentSize))])
			} else {
				fd.Write(buffer)
			}
		} else {
			if (writecount + 1) >= len(fmeta.SegmentList) {
				fd.Write(buf[:(fmeta.FileSize.Uint64() - uint64(writecount*chain.SegmentSize))])
			} else {
				fd.Write(buf)
			}
		}
		writecount++
	}
	if writecount != len(fmeta.SegmentList) {
		return "", errors.New("write failed")
	}
	err = fd.Sync()
	return userfile, err
}

func DownloadFragmentFromMiner(cli chain.Chainer, minerpuk []byte, fid, fragment string) ([]byte, error) {
	minerInfo, err := cli.QueryMinerItems(minerpuk, -1)
	if err != nil {
		return nil, err
	}

	url := string(minerInfo.PeerId[:])

	if strings.HasSuffix(url, "/") {
		url = url + "fragment"
	} else {
		url = url + "/fragment"
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	message := utils.GetRandomcode(16)
	sig, err := utils.SignedSR25519WithMnemonic(cli.GetURI(), message)
	if err != nil {
		return nil, fmt.Errorf("[SignedSR25519WithMnemonic] %v", err)
	}
	req.Header.Set("Fid", fid)
	req.Header.Set("Fragment", fragment)
	req.Header.Set("Account", cli.GetSignatureAcc())
	req.Header.Set("Message", message)
	req.Header.Set("Signature", string(sig))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed code: %d", resp.StatusCode)
	}
	return respbody, nil
}

func StorageToMiners(cli chain.Chainer, fragmentGroup [][]string, fid string, wantMiner []string) error {
	var ok bool
	var err error
	var sucMiner = make(map[string]struct{}, 0)
	for i := 0; i < len(wantMiner); i++ {
		_, ok = sucMiner[wantMiner[i]]
		if ok {
			continue
		}
		err = StoreBatchFragmentsToMiner(cli, fragmentGroup[i], fid, wantMiner[i])
		if err != nil {
			return fmt.Errorf("[%s] failed: %v\n", wantMiner[i], err)
		}
		sucMiner[wantMiner[i]] = struct{}{}
		if (i + 1) >= (chain.DataShards + chain.ParShards) {
			return nil
		}
	}
	fragmentGroup = fragmentGroup[len(wantMiner):]

	allminers, err := cli.QueryAllMiner(-1)
	if err != nil {
		return err
	}
	length := len(allminers)
	account := ""
	for i := 0; i < length; i++ {
		account, err = utils.EncodePublicKeyAsCessAccount(allminers[i][:])
		if err != nil {
			continue
		}
		_, ok = sucMiner[account]
		if ok {
			continue
		}
		err = StoreBatchFragmentsToMiner(cli, fragmentGroup[0], fid, account)
		if err != nil {
			continue
		}
		if len(fragmentGroup) == 1 {
			return nil
		}
		fragmentGroup = fragmentGroup[1:]
	}
	return errors.New("storage failed")
}

func StoreToAllDesignatedMiners(cli chain.Chainer, fragmentGroup [][]string, fid string, wantMiner []string) error {
	var ok bool
	var err error
	var rntMsg string
	var sucMiner = make(map[string]struct{}, 0)
	minerlength := len(wantMiner)
	for i := 0; i < chain.DataShards+chain.ParShards; i++ {
		for j := 0; j < minerlength; j++ {
			_, ok = sucMiner[wantMiner[j]]
			if ok {
				continue
			}
			err = StoreBatchFragmentsToMiner(cli, fragmentGroup[i], fid, wantMiner[j])
			if err != nil {
				rntMsg += fmt.Sprintf("[%s] failed: %v\n", wantMiner[j], err)
			} else {
				sucMiner[wantMiner[j]] = struct{}{}
				rntMsg += fmt.Sprintf("[%s] suc\n", wantMiner[j])
			}
		}
	}
	if len(sucMiner) == chain.DataShards+chain.ParShards {
		return nil
	}
	return errors.New(rntMsg)
}

func StoreBatchFragmentsToMiner(cli chain.Chainer, fragments []string, fid, account string) error {
	puk, err := utils.ParsingPublickey(account)
	if err != nil {
		return err
	}
	minerInfo, err := cli.QueryMinerItems(puk, -1)
	if err != nil {
		return err
	}

	if string(minerInfo.State) != chain.MINER_STATE_POSITIVE {
		return errors.New("not positive state")
	}

	if minerInfo.IdleSpace.Uint64() < uint64(len(fragments)*chain.FragmentSize) {
		return errors.New("insufficient space")
	}

	length := len(fragments)
	for i := 0; i < length; i++ {
		err = UploadFragmentToMiner(cli, string(minerInfo.PeerId[:]), fid, fragments[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func UploadFragmentToMiner(cli chain.Chainer, ip string, fid string, file string) error {
	message := utils.GetRandomcode(16)
	sig, err := utils.SignedSR25519WithMnemonic(cli.GetURI(), message)
	if err != nil {
		return fmt.Errorf("[SignedSR25519WithMnemonic] %v", err)
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile("file", filepath.Base(file))
	if err != nil {
		return err
	}

	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	_, err = io.Copy(formFile, fd)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	url := ip
	if strings.HasSuffix(url, "/") {
		url = url + "fragment"
	} else {
		url = url + "/fragment"
	}

	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Fid", fid)
	req.Header.Set("Account", cli.GetSignatureAcc())
	req.Header.Set("Message", message)
	req.Header.Set("Signature", string(sig))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed code: %d", resp.StatusCode)
	}
	var respinfo RespType
	err = json.Unmarshal(respbody, &respinfo)
	if err != nil {
		return errors.New("server returns invalid data")
	}
	if respinfo.Code != 200 {
		return fmt.Errorf("server returns code: %d", respinfo.Code)
	}
	return nil
}

func CheckAccount(cli chain.Chainer, territory string, size int64) error {
	useSpace := CalcUsedSpace(size)
	territoryInfo, err := cli.QueryTerritory(cli.GetSignatureAccPulickey(), territory, -1)
	if err != nil {
		if !errors.Is(err, chain.ERR_RPC_EMPTY_VALUE) {
			return err
		}
		gibs := useSpace / chain.SIZE_1GiB
		if useSpace%chain.SIZE_1GiB != 0 {
			gibs += 1
		}
		_, err = cli.MintTerritory(uint32(gibs), territory, uint32(30))
		if err != nil {
			return err
		}
		time.Sleep(chain.BlockInterval)
		territoryInfo, err = cli.QueryTerritory(cli.GetSignatureAccPulickey(), territory, -1)
		if err != nil {
			return err
		}
	}
	header, err := cli.GetSubstrateAPI().RPC.Chain.GetHeaderLatest()
	if err != nil {
		return err
	}

	if territoryInfo.Deadline <= types.U32(header.Number) {
		return errors.New("expired territory")
	}

	if territoryInfo.RemainingSpace.Uint64() < useSpace {
		return errors.New("insufficient territorial space")
	}
	return nil
}

func CheckFile(file string) (int64, error) {
	fstat, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	if fstat.IsDir() {
		return 0, errors.New("not a file")
	}
	if fstat.Size() <= 0 {
		return 0, errors.New("empty file")
	}
	return fstat.Size(), nil
}

func CalcUsedSpace(size int64) uint64 {
	count := size / chain.SegmentSize
	if size%chain.SegmentSize != 0 {
		count += 1
	}
	return uint64(count*chain.SegmentSize) * chain.NumberOfDataCopies
}
