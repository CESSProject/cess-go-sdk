/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package process

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

	"github.com/CESSProject/cess-go-sdk/config"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/btcsuite/btcutil/base58"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/pkg/errors"
)

// UploadFileChunks upload file chunks in the directory to the gateway as much as possible,
// chunks will be removed after being uploaded, if the chunks are not transferred successfuly, jus
//
// Receive parameter:
//   - url: the address of the gateway.
//   - chunksDir: directory path to store file chunks, please do not mix it elsewhere.
//   - bucket: the bucket name to store user data.
//   - fname: the name of the file.
//   - chunksNum: total number of file chunks.
//   - totalSize: chunks total size (byte), can be obtained from the first return value of SplitFile
//
// Return parameter:
//   - Reader: number of file chunks.
//   - error: error message.
func UploadFileChunks(url, mnemonic, chunksDir, bucket, fname string, chunksNum int, totalSize int64) (string, error) {
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
		res, err = UploadFileChunk(url, mnemonic, chunksDir, bucket, fname, chunksNum, i, totalSize)
		if err != nil {
			return res, errors.Wrap(err, "upload file chunks error")
		}
		os.Remove(filepath.Join(chunksDir, fmt.Sprintf("chunk-%d", i)))
	}
	return res, nil
}

// UploadFileChunk upload chunk of file to the gateway
//
// Receive parameter:
//   - url: the address of the gateway.
//   - chunksDir: directory path to store file chunks, please do not mix it elsewhere.
//   - bucket: the bucket name to store user data.
//   - fname: the name of the file.
//   - chunksNum: total number of file chunks.
//   - chunksId: index of the current chunk to be uploaded ([0,chunksNum)).
//   - totalSize: chunks total size (byte), can be obtained from the first return value of SplitFile
//
// Return parameter:
//   - Reader: number of file chunks.
//   - error: error message.
func UploadFileChunk(url, mnemonic, chunksDir, bucket, fname string, chunksNum, chunksId int, totalSize int64) (string, error) {

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

// Split File into Chunks with standard size.
// It split file into chunks of the default size and fills the last chunk that does not meet the size.
//
// Receive parameter:
//   - fpath: the path of the file to be split.
//   - chunksDir: directory path to store file chunks, please do not mix it elsewhere.
//
// Return parameter:
//   - int64: chunks total size (byte).
//   - int: number of file chunks.
//   - error: error message.
func SplitFileWithstandardSize(fpath, chunksDir string) (int64, int, error) {
	return SplitFile(fpath, chunksDir, config.SegmentSize, true)
}

// Split File into Chunks.
//
// Receive parameter:
//   - fpath: the path of the file to be split.
//   - chunksDir: directory path to store file chunks, please do not mix it elsewhere.
//   - chunkSize: the size of each chunk, it does not exceed the file size
//
// Return parameter:
//   - int64: chunks total size (byte).
//   - int: number of file chunks.
//   - error: error message.
func SplitFile(fpath, chunksDir string, chunkSize int64, filling bool) (int64, int, error) {
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
