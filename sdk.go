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

// New constructs a new sdk client with the given options, falling back on
// reasonable defaults. The defaults are:
//
// - If no rpc address is provided, the sdk client uses the default address
// "wss://testnet-rpc0.cess.cloud/ws/"" or "wss://testnet-rpc1.cess.cloud/ws/";
//
// - If no transaction timeout is provided, the sdk client uses the default
// timeout: time.Duration(time.Second * 6)
//
// - The serviceName is used to specify the name of your service
func New(ctx context.Context, opts ...Option) (*chain.ChainClient, error) {
	return NewWithoutDefaults(ctx, append(opts, FallbackDefaults)...)
}

// NewWithoutDefaults constructs a new client with the given options but
// *without* falling back on reasonable defaults.
//
// Warning: This function should not be considered a stable interface. We may
// choose to add required services at any time and, by using this function, you
// opt-out of any defaults we may provide.
func NewWithoutDefaults(ctx context.Context, opts ...Option) (*chain.ChainClient, error) {
	var cfg Config
	if err := cfg.Apply(opts...); err != nil {
		return nil, err
	}
	return cfg.NewSDK(ctx)
}
