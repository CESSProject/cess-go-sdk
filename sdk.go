/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdkgo

import (
	"github.com/CESSProject/sdk-go/config"
	"github.com/CESSProject/sdk-go/core/sdk"
)

// Config describes a set of settings for a client.
type Config = config.Config

// Option is a client config option that can be given to the client constructor
type Option = config.Option

// New constructs a new client with the given options, falling back on
// reasonable defaults. The defaults are:
//
// - If no transport and listen addresses are provided, the node listens to
// the default addresses "/ip4/0.0.0.0/tcp/15000";
//
// - If no Rpc address is provided, the client uses the default address
// "wss://testnet-rpc0.cess.cloud/ws/"" or "wss://testnet-rpc1.cess.cloud/ws/";
//
// - If no working directory is provided, the client uses the current
// directory as the working directory;
func New(name string, opts ...Option) (sdk.SDK, error) {
	return NewWithoutDefaults(name, append(opts, FallbackDefaults)...)
}

// NewWithoutDefaults constructs a new client with the given options but
// *without* falling back on reasonable defaults.
//
// Warning: This function should not be considered a stable interface. We may
// choose to add required services at any time and, by using this function, you
// opt-out of any defaults we may provide.
func NewWithoutDefaults(name string, opts ...Option) (sdk.SDK, error) {
	var cfg Config
	if err := cfg.Apply(opts...); err != nil {
		return nil, err
	}
	return cfg.NewSDK(name)
}
