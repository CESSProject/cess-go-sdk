/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package chain

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/utils"
	keyring "github.com/CESSProject/go-keyring"
	"github.com/btcsuite/btcutil/base58"
	"github.com/pkg/errors"
)

func (c *ChainClient) StoreFile(url, file, bucket string) (string, error) {
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

	kr, _ := keyring.FromURI(c.GetURI(), keyring.NetSubstrate{})

	// sign message
	message := utils.GetRandomcode(16)
	sig, _ := kr.Sign(kr.SigningContext([]byte(message)))

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
	req.Header.Set("Account", c.GetSignatureAcc())
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

func (c *ChainClient) StoreObject(url string, reader io.Reader, bucket string) (string, error) {
	if !utils.CheckBucketName(bucket) {
		return "", errors.New("invalid bucket name")
	}

	kr, _ := keyring.FromURI(c.GetURI(), keyring.NetSubstrate{})

	// sign message
	message := utils.GetRandomcode(16)
	sig, _ := kr.Sign(kr.SigningContext([]byte(message)))

	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return "", err
	}

	req.Header.Set("BucketName", bucket)
	req.Header.Set("Account", c.GetSignatureAcc())
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

func (c *ChainClient) RetrieveFile(url, fid, savepath string) error {
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

	kr, _ := keyring.FromURI(c.GetURI(), keyring.NetSubstrate{})

	// sign message
	message := utils.GetRandomcode(16)
	sig, _ := kr.Sign(kr.SigningContext([]byte(message)))
	req.Header.Set("Message", message)
	req.Header.Set("Signature", base58.Encode(sig[:]))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Operation", "download")
	req.Header.Set("Account", c.GetSignatureAcc())

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

func (c *ChainClient) RetrieveObject(url, fid string) (io.ReadCloser, error) {
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
	kr, _ := keyring.FromURI(c.GetURI(), keyring.NetSubstrate{})

	// sign message
	message := utils.GetRandomcode(16)
	sig, _ := kr.Sign(kr.SigningContext([]byte(message)))
	req.Header.Set("Message", message)
	req.Header.Set("Signature", base58.Encode(sig[:]))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Account", c.GetSignatureAcc())
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

func (c *chainClient) SplitFile(fpath, chunksDir string, chunkSize int64, filling bool) (int64, int, error) {
	fstat, err := os.Stat(fpath)
	if err != nil {
		return 0, 0, err
	}
	if fstat.IsDir() {
		return 0, 0, errors.New("not a file")
	}
	if fstat.Size() == 0 {
		return 0, 0, errors.New("empty file")
	}
	if fstat.Size() < chunkSize {
		chunkSize = fstat.Size()
	}
	count := fstat.Size() / chunkSize
	if fstat.Size()%chunkSize != 0 {
		count++
	}
	buf := make([]byte, chunkSize)
	f, err := os.Open(fpath)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	size := fstat.Size()
	for i := int64(0); i < count; i++ {
		n, err := reader.Read(buf)
		if err != nil {
			return 0, 0, err
		}
		if n <= 0 {
			return 0, 0, errors.New("read a empty block")
		}
		if n < int(chunkSize) && filling {
			if i+1 != count {
				return 0, 0, errors.New("read file err")
			}
			copy(buf[n:], make([]byte, chunkSize-int64(n)))
			size += chunkSize - int64(n)
			n = int(chunkSize)
		}
		err = utils.WriteBufToFile(buf[:n], filepath.Join(chunksDir, fmt.Sprintf("chunk-%d", i)))
		if err != nil {
			return 0, 0, err
		}
	}
	return size, int(count), nil
}

func (c *chainClient) SplitFileWithstandardSize(fpath, chunksDir string) (int64, int, error) {
	return c.SplitFile(fpath, chunksDir, pattern.SegmentSize, true)
}

func (c *chainClient) UploadFileChunks(url, chunksDir, bucket, fname string, chunksNum int, totalSize int64) (string, error) {
	entries, err := os.ReadDir(chunksDir)
	if err != nil {
		return "", errors.Wrap(err, "upload file chunk error")
	}
	if len(entries) == 0 {
		return "", errors.Wrap(errors.New("empty dir"), "upload file chunk error")
	}
	if len(entries) > chunksNum {
		return "", errors.Wrap(errors.New("bad chunks number"), "upload file chunk error")
	}
	var res string
	for i := chunksNum - len(entries); i < chunksNum; i++ {
		res, err = c.UploadFileChunk(url, chunksDir, bucket, fname, chunksNum, i, totalSize)
		if err != nil {
			return res, errors.Wrap(err, "upload file chunks error")
		}
		os.Remove(filepath.Join(chunksDir, fmt.Sprintf("chunk-%d", i)))
	}
	return res, nil
}

func (c *chainClient) UploadFileChunk(url, chunksDir, bucket, fname string, chunksNum, chunksId int, totalSize int64) (string, error) {

	file := filepath.Join(chunksDir, fmt.Sprintf("chunk-%d", chunksId))
	fstat, err := os.Stat(file)
	if err != nil {
		return "", errors.Wrap(err, "upload file chunk error")
	}

	if fstat.IsDir() {
		return "", errors.Wrap(errors.New("not a file"), "upload file chunk error")
	}

	if fstat.Size() == 0 {
		return "", errors.New("empty file")
	}

	if !utils.CheckBucketName(bucket) {
		return "", errors.New("invalid bucket name")
	}

	kr, _ := keyring.FromURI(c.GetURI(), keyring.NetSubstrate{})

	// sign message
	message := utils.GetRandomcode(16)
	sig, _ := kr.Sign(kr.SigningContext([]byte(message)))

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
	req.Header.Set("Account", c.GetSignatureAcc())
	req.Header.Set("Message", message)
	req.Header.Set("Signature", base58.Encode(sig[:]))
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("FileName", fname)
	req.Header.Set("BlockNumber", fmt.Sprint(chunksNum))
	req.Header.Set("BlockIndex", fmt.Sprint(chunksId))
	req.Header.Set("TotalSize", fmt.Sprint(totalSize))

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
