/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/vedhavyas/go-subkey/sr25519"
)

// SignedSR25519WithMnemonic sign sr25519 with mnemonic
//   - mnemonic: polkadot account mnemonic
//   - msg: message
//
// Return:
//   - []byte: sr25519 signature
//   - error: error message
func SignedSR25519WithMnemonic(mnemonic string, msg string) ([]byte, error) {
	if len(msg) <= 0 {
		return nil, errors.New("SignedSR25519WithMnemonic: empty msg")
	}
	pri, err := sr25519.Scheme{}.FromPhrase(mnemonic, "")
	if err != nil {
		return nil, errors.New("SignedSR25519WithMnemonic: invalid mnemonic")
	}
	return pri.Sign([]byte(msg))
}

// VerifySR25519WithPublickey verify sr25519 signature with account public key
//   - msg: message
//   - sign: sr25519 signature
//   - account: polkadot account
//
// Return:
//   - bool: verification result
//   - error: error message
func VerifySR25519WithPublickey(msg string, sign []byte, account string) (bool, error) {
	if len(sign) <= 0 {
		return false, errors.New("VerifySR25519WithPublickey: empty sign")
	}
	pk, err := ParsingPublickey(account)
	if err != nil {
		return false, errors.New("VerifySR25519WithPublickey: invalid account")
	}
	public, err := sr25519.Scheme{}.FromPublicKey(pk)
	if err != nil {
		return false, err
	}
	ok := public.Verify([]byte(msg), sign)
	return ok, err
}

// VerifyPolkadotJsHexSign verify signature signed with polkadot.js
//   - account: polkadot account
//   - msg: message
//   - sign: signature
//
// Return:
//   - bool: verification result
//   - error: error message
//
// Tip:
//   - https://polkadot.js.org/apps/#/signing
func VerifyPolkadotJsHexSign(account, msg, signature string) (bool, error) {
	if len(msg) == 0 {
		return false, errors.New("msg is empty")
	}

	pkey, err := ParsingPublickey(account)
	if err != nil {
		return false, err
	}

	pub, err := sr25519.Scheme{}.FromPublicKey(pkey)
	if err != nil {
		return false, err
	}

	sign_bytes, err := hex.DecodeString(strings.TrimPrefix(signature, "0x"))
	if err != nil {
		return false, err
	}
	message := fmt.Sprintf("<Bytes>%s</Bytes>", msg)
	ok := pub.Verify([]byte(message), sign_bytes)
	return ok, nil
}

// ParseEthAccFromEthSign parsing eth account public key from eth account signature
//   - message: message
//   - sign: eth signature
//
// Return:
//   - string: eth account
//   - error: error message
func ParseEthAccFromEthSign(message string, sign string) (string, error) {
	// Hash the unsigned message using EIP-191
	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
	hash := crypto.Keccak256Hash(hashedMessage)

	// Get the bytes of the signed message
	decodedMessage, err := hexutil.Decode(sign)
	if err != nil {
		return "", err
	}

	// Handles cases where EIP-115 is not implemented (most wallets don't implement it)
	if decodedMessage[64] == 27 || decodedMessage[64] == 28 {
		decodedMessage[64] -= 27
	}

	// Recover a public key from the signed message
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), decodedMessage)
	if sigPublicKeyECDSA == nil {
		err = errors.New("could not get a public get from the message signature")
	}
	if err != nil {
		return "", err
	}

	return crypto.PubkeyToAddress(*sigPublicKeyECDSA).String(), nil
}
