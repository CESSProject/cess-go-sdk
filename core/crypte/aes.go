/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package crypte

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"runtime"
)

var (
	ErrKeyLengthMax        = errors.New("key length cannot exceed 32")
	ErrKeyLengthEmpty      = errors.New("key length cannot be empty")
	ErrPaddingSize         = errors.New("padding size error please check the secret key or iv")
	ErrInvalidBlockSize    = errors.New("invalid blocksize")
	ErrInvalidPKCS7Data    = errors.New("invalid PKCS7 data (empty or not padded)")
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

// aes encrypt
func AesCbcEncrypt(plainText, secretKey []byte) (cipherText []byte, err error) {
	if len(secretKey) > 32 {
		return nil, ErrKeyLengthMax
	}
	if len(secretKey) <= 0 {
		return nil, ErrKeyLengthEmpty
	}
	keyPaddingLength := 32 - len(secretKey)
	usekey := string(secretKey)
	for i := 0; i < keyPaddingLength; i++ {
		usekey += "0"
	}

	block, err := aes.NewCipher([]byte(usekey))
	if err != nil {
		return nil, err
	}
	paddingText, err := pkcs7Pad(plainText, block.BlockSize())
	if err != nil {
		return nil, err
	}
	iv := usekey[:block.BlockSize()]
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	cipherText = make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherText, paddingText)
	return cipherText, nil
}

// aes decrypt
func AesCbcDecrypt(cipherText, secretKey []byte) (plainText []byte, err error) {
	if len(secretKey) > 32 {
		return nil, ErrKeyLengthMax
	}
	if len(secretKey) <= 0 {
		return nil, ErrKeyLengthEmpty
	}
	keyPaddingLength := 32 - len(secretKey)
	usekey := string(secretKey)
	for i := 0; i < keyPaddingLength; i++ {
		usekey += "0"
	}
	block, err := aes.NewCipher([]byte(usekey))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				fmt.Printf("runtime err=%v,Check that the key or text is correct", err)
			default:
				fmt.Printf("error=%v,check the cipherText ", err)
			}
		}
	}()
	iv := usekey[:block.BlockSize()]
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	paddingText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(paddingText, cipherText)

	plainText, err = pkcs7Unpad(paddingText, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

func pkcs5Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

func pkcs5UnPadding(plainText []byte, blockSize int) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number >= length || number > blockSize {
		return nil, ErrPaddingSize
	}
	return plainText[:length-number], nil
}

// pkcs7Pad right-pads the given byte slice with 1 to n bytes, where
// n is the block size. The size of the result is x times n, where x
// is at least 1.
func pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

// pkcs7Unpad validates and unpads data from the given bytes slice.
// The returned value will be 1 to n bytes smaller depending on the
// amount of padding, where n is the block size.
func pkcs7Unpad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	if len(b)%blocksize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}
	c := b[len(b)-1]
	n := int(c)
	if n == 0 || n > len(b) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if b[len(b)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return b[:len(b)-n], nil
}
