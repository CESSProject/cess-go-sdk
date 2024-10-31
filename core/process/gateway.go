/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package process

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/btcsuite/btcutil/base58"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/pkg/errors"
)

// StoreFile stores files to the gateway
//
// Receive parameter:
//   - url: gateway url
//   - file: stored file
//   - territory: territory name
//   - mnemonic: polkadot account mnemonic
//
// Return parameter:
//   - string: [fid] unique identifier for the file.
//   - error: error message.
//
// Preconditions:
//  1. Account requires purchasing territory, refer to [MintTerritory] interface.
//  2. Authorize the space usage rights of the account to the gateway account,
//     refer to the [AuthorizeSpace] interface.
func StoreFile(url, file, territory, mnemonic string) (string, error) {
	fstat, err := os.Stat(file)
	if err != nil {

		return "", err
	}

	if fstat.IsDir() {
		return "", errors.New("not a file")
	}

	if fstat.Size() == 0 {
		return "", errors.New("empty file")
	}

	keyringPair, err := signature.KeyringPairFromSecret(mnemonic, 0)
	if err != nil {
		return "", fmt.Errorf("[KeyringPairFromSecret] %v", err)
	}

	acc, err := utils.EncodePublicKeyAsCessAccount(keyringPair.PublicKey)
	if err != nil {
		return "", fmt.Errorf("[EncodePublicKeyAsCessAccount] %v", err)
	}

	// sign message
	message := utils.GetRandomcode(16)
	sig, err := utils.SignedSR25519WithMnemonic(keyringPair.URI, message)
	if err != nil {
		return "", fmt.Errorf("[SignedSR25519WithMnemonic] %v", err)
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	//
	formFile, err := writer.CreateFormFile("file", fstat.Name())
	if err != nil {
		return "", err
	}

	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(formFile, f)
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Territory", territory)
	req.Header.Set("Account", acc)
	req.Header.Set("Message", message)
	req.Header.Set("Signature", base58.Encode(sig[:]))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		if len(respbody) > 0 {
			return "", errors.New(string(respbody))
		}
		return "", errors.New(fmt.Sprintf("upload failed, code: %d", resp.StatusCode))
	}

	var respValue RespType
	err = json.Unmarshal(respbody, &respValue)
	if err != nil {
		return "", nil
	}
	reapdata, ok := respValue.Data.(map[string]string)
	if !ok {
		return "", nil
	}
	return reapdata["fid"], nil
}

// RangeUploadFast quickly upload a file using range request
//
// Receive parameter:
//   - url: gateway url
//   - file: upload file
//   - territory: territory name
//   - mnemonic: polkadot account mnemonic
//   - start: starting position of file
//   - rangeSize: range size
//
// Return parameter:
//   - string: [fid] unique identifier for the file.
//   - error: error message.
//
// Preconditions:
//  1. Account requires purchasing territory, refer to [MintTerritory] interface.
//  2. Authorize the space usage rights of the account to the gateway account,
//     refer to the [AuthorizeSpace] interface.
func RangeUploadFast(url, file, territory, mnemonic string, start int64, rangeSize int64) (string, error) {
	fstat, err := os.Stat(file)
	if err != nil {
		return "", err
	}
	if fstat.IsDir() {
		return "", errors.New("not a file")
	}
	if rangeSize <= 0 {
		return "", errors.New("invalid range size")
	}
	totalSize := fstat.Size()
	if totalSize <= 0 {
		return "", errors.New("empty file")
	}
	var fid string

	// sign message
	keyringPair, err := signature.KeyringPairFromSecret(mnemonic, 0)
	if err != nil {
		return "", fmt.Errorf("[KeyringPairFromSecret] %v", err)
	}
	message := utils.GetRandomcode(16)
	sig, err := utils.SignedSR25519WithMnemonic(keyringPair.URI, message)
	if err != nil {
		return "", fmt.Errorf("[SignedSR25519WithMnemonic] %v", err)
	}
	account, err := utils.EncodePublicKeyAsCessAccount(keyringPair.PublicKey)
	if err != nil {
		return "", err
	}
	for start := int64(0); start < totalSize; start += rangeSize {
		fid, err = RangeUpload(url, file, territory, account, message, base58.Encode(sig[:]), start, rangeSize)
		if err != nil {
			return fid, err
		}
	}
	return fid, nil
}

// RangeUpload upload a file using range request
//
// Receive parameter:
//   - url: gateway url
//   - file: upload file
//   - territory: territory name
//   - account: user account
//   - msg: signed message
//   - sign: signature
//   - start: starting position of file
//   - rangeSize: range size
//
// Return parameter:
//   - string: [fid] unique identifier for the file.
//   - error: error message.
//
// Preconditions:
//  1. Account requires purchasing territory, refer to [MintTerritory] interface.
//  2. Authorize the space usage rights of the account to the gateway account,
//     refer to the [AuthorizeSpace] interface.
func RangeUpload(url, file, territory, account, msg, sign string, start int64, rangeSize int64) (string, error) {
	fd, err := os.Open(file)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer fd.Close()

	fdStat, err := fd.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to get stat: %v", err)
	}

	totalSize := fdStat.Size()
	_, err = fd.Seek(start, 0)
	if err != nil {
		return "", fmt.Errorf("failed to seek file: %w", err)
	}

	end := start + rangeSize - 1
	if end >= totalSize {
		end = totalSize - 1
	}
	bytesToUpload := end - start + 1

	req, err := http.NewRequest("PUT", url, io.LimitReader(fd, bytesToUpload))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, totalSize))
	req.Header.Set("Content-Length", fmt.Sprintf("%d", bytesToUpload))
	req.Header.Set("Territory", territory)
	req.Header.Set("Account", account)
	req.Header.Set("Message", msg)
	req.Header.Set("Signature", sign)

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to upload chunk: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusPermanentRedirect {
		fmt.Printf("Uploaded bytes %d-%d\n", start, end)
		return "", nil
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("upload failed: %s", resp.Status)
	}
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	var respValue RespType
	err = json.Unmarshal(buf, &respValue)
	if err != nil {
		return "", nil
	}
	reapdata, ok := respValue.Data.(map[string]string)
	if !ok {
		return "", nil
	}
	return reapdata["fid"], nil
}

// StoreObject stores object to the gateway
//
// Receive parameter:
//   - url: gateway url
//   - territory: territory name
//   - mnemonic: polkadot account mnemonic
//   - reader: strings, byte data, file streams, network streams, etc
//
// Return parameter:
//   - string: [fid] unique identifier for the file
//   - error: error message
//
// Preconditions:
//  1. Account requires purchasing space, refer to [BuySpace] interface.
//  2. Authorize the space usage rights of the account to the gateway account,
//     refer to the [AuthorizeSpace] interface.
func StoreObject(url string, territory, mnemonic string, reader io.Reader) (string, error) {
	keyringPair, err := signature.KeyringPairFromSecret(mnemonic, 0)
	if err != nil {
		return "", fmt.Errorf("[KeyringPairFromSecret] %v", err)
	}

	acc, err := utils.EncodePublicKeyAsCessAccount(keyringPair.PublicKey)
	if err != nil {
		return "", fmt.Errorf("[EncodePublicKeyAsCessAccount] %v", err)
	}

	// sign message
	message := utils.GetRandomcode(16)
	sig, err := utils.SignedSR25519WithMnemonic(keyringPair.URI, message)
	if err != nil {
		return "", fmt.Errorf("[SignedSR25519WithMnemonic] %v", err)
	}

	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return "", err
	}

	req.Header.Set("Territory", territory)
	req.Header.Set("Account", acc)
	req.Header.Set("Message", message)
	req.Header.Set("Signature", base58.Encode(sig[:]))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		if len(respbody) > 0 {
			return "", errors.New(string(respbody))
		}
		return "", errors.New(fmt.Sprintf("upload failed, code: %d", resp.StatusCode))
	}

	var respValue RespType
	err = json.Unmarshal(respbody, &respValue)
	if err != nil {
		return "", nil
	}
	reapdata, ok := respValue.Data.(map[string]string)
	if !ok {
		return "", nil
	}
	return reapdata["fid"], nil
}

// RetrieveFile downloads files from the gateway
//   - url: gateway url
//   - mnemonic: polkadot account mnemonic
//   - savepath: file save path
//
// Return:
//   - error: error message
func RetrieveFile(url, mnemonic, savepath string) error {
	fid := filepath.Base(url)
	fstat, err := os.Stat(savepath)
	if err != nil {
		err = os.MkdirAll(savepath, 0755)
		if err != nil {
			return err
		}
	} else {
		if fstat.IsDir() {
			savepath = filepath.Join(savepath, fid)
			fstat, err = os.Stat(savepath)
			if err == nil {
				if fstat.Size() > 0 {
					return nil
				}
			}
		} else {
			if fstat.Size() > 0 {
				return nil
			}
		}
	}

	f, err := os.Create(savepath)
	if err != nil {
		return err
	}
	defer f.Close()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	keyringPair, err := signature.KeyringPairFromSecret(mnemonic, 0)
	if err != nil {
		return fmt.Errorf("[KeyringPairFromSecret] %v", err)
	}

	acc, err := utils.EncodePublicKeyAsCessAccount(keyringPair.PublicKey)
	if err != nil {
		return fmt.Errorf("[EncodePublicKeyAsCessAccount] %v", err)
	}

	// sign message
	message := utils.GetRandomcode(16)
	sig, err := utils.SignedSR25519WithMnemonic(keyringPair.URI, message)
	if err != nil {
		return fmt.Errorf("[SignedSR25519WithMnemonic] %v", err)
	}

	req.Header.Set("Message", message)
	req.Header.Set("Signature", base58.Encode(sig[:]))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Account", acc)

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed")
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// RetrieveObject gets the object from the gateway
//   - url: gateway url
//   - mnemonic: polkadot account mnemonic
//
// Return:
//   - io.ReadCloser: object
//   - error: error message
func RetrieveObject(url, mnemonic string) (io.ReadCloser, error) {
	if url == "" {
		return nil, errors.New("empty url")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	keyringPair, err := signature.KeyringPairFromSecret(mnemonic, 0)
	if err != nil {
		return nil, fmt.Errorf("[KeyringPairFromSecret] %v", err)
	}

	acc, err := utils.EncodePublicKeyAsCessAccount(keyringPair.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("[EncodePublicKeyAsCessAccount] %v", err)
	}

	// sign message
	message := utils.GetRandomcode(16)
	sig, err := utils.SignedSR25519WithMnemonic(keyringPair.URI, message)
	if err != nil {
		return nil, fmt.Errorf("[SignedSR25519WithMnemonic] %v", err)
	}

	req.Header.Set("Message", message)
	req.Header.Set("Signature", base58.Encode(sig[:]))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Account", acc)

	client := &http.Client{}
	client.Transport = globalTransport
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Retrieve failed, code: %d", resp.StatusCode))
	}

	return resp.Body, nil
}
