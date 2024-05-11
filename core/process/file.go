/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package process

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/btcsuite/btcutil/base58"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/pkg/errors"
)

var globalTransport = &http.Transport{
	DisableKeepAlives: true,
}

// StoreFile stores files to the gateway
//   - url: gateway url
//   - file: stored file
//   - bucket: bucket for storing file
//   - mnemonic: polkadot account mnemonic
//
// Return:
//   - string: fid
//   - error: error message
func StoreFile(url, file, bucket, mnemonic string) (string, error) {
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

	if !utils.CheckBucketName(bucket) {
		return "", errors.New("invalid bucket name")
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

	req.Header.Set("BucketName", bucket)
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

	return strings.TrimPrefix(strings.TrimSuffix(string(respbody), "\""), "\""), nil
}

// StoreObject stores object to the gateway
//   - url: gateway url
//   - bucket: bucket for storing file
//   - mnemonic: polkadot account mnemonic
//   - reader: reader
//
// Return:
//   - string: fid
//   - error: error message
func StoreObject(url string, bucket, mnemonic string, reader io.Reader) (string, error) {
	if !utils.CheckBucketName(bucket) {
		return "", errors.New("invalid bucket name")
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

	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return "", err
	}

	req.Header.Set("BucketName", bucket)
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

	return strings.TrimPrefix(strings.TrimSuffix(string(respbody), "\""), "\""), nil
}

// RetrieveFile downloads files from the gateway
//   - url: gateway url
//   - fid: fid
//   - mnemonic: polkadot account mnemonic
//   - savepath: file save path
//
// Return:
//   - string: fid
//   - error: error message
func RetrieveFile(url, fid, mnemonic, savepath string) error {
	fstat, err := os.Stat(savepath)
	if err == nil {
		if fstat.IsDir() {
			savepath = filepath.Join(savepath, fid)
		}
		if fstat.Size() > 0 {
			return nil
		}
	}

	if url == "" {
		return errors.New("empty url")
	}

	if url[len(url)-1] != byte(47) {
		url += "/"
	}

	f, err := os.Create(savepath)
	if err != nil {
		return err
	}
	defer f.Close()

	req, err := http.NewRequest(http.MethodGet, url+fid, nil)
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
	req.Header.Set("Operation", "download")
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
//   - fid: fid
//   - mnemonic: polkadot account mnemonic
//
// Return:
//   - io.ReadCloser: object
//   - error: error message
func RetrieveObject(url, fid, mnemonic string) (io.ReadCloser, error) {
	if url == "" {
		return nil, errors.New("empty url")
	}

	if url[len(url)-1] != byte(47) {
		url += "/"
	}

	req, err := http.NewRequest(http.MethodGet, url+fid, nil)
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
	req.Header.Set("Operation", "download")

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
