/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdkgo

import (
	"context"

	"github.com/CESSProject/cess-go-sdk/chain"
)

// NewDefault constructs a new SDK client using the given options and default options.
//
// If no rpc endpoint are provided, use the default: “wss://testnet-rpc.cess.network/ws/”
//
// If no transaction timeout are provided, use the default timeout: 18s
//
// If no name are provided, use the default name: cess-sdk-go
func New(ctx context.Context, opts ...Option) (chain.Chainer, error) {
	return NewWithoutDefaults(ctx, append(opts, FallbackDefaults)...)
}

// New constructs a new sdk client with the given options.
//
// If no rpc endpoint are provided, the CESS blockchain network cannot be accessed.
//
// If no account mnemonic are provided, block transactions cannot be conducted.
func NewWithoutDefaults(ctx context.Context, opts ...Option) (chain.Chainer, error) {
	var cfg Config
	if err := cfg.Apply(opts...); err != nil {
		return nil, err
	}
	return cfg.NewSDK(ctx)
}
