/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdkgo

import (
	"time"
)

// DefaultRpcAddrs configures the default rpc address
var DefaultRpcAddrs = func(cfg *Config) error {
	rpcAddrs := []string{
		"wss://testnet-rpc.cess.network/ws/",
	}
	return cfg.Apply(ConnectRpcAddrs(rpcAddrs))
}

// DefaultTimeout configures the default transaction waiting timeout
var DefaultTimeout = func(cfg *Config) error {
	return cfg.Apply(TransactionTimeout(time.Second * 18))
}

// DefaultTimeout configures the default transaction waiting timeout
var DefaultSDKName = func(cfg *Config) error {
	return cfg.Apply(Name(DefaultName))
}

// Complete list of default options and when to fallback on them.
var defaults = []struct {
	fallback func(cfg *Config) bool
	opt      Option
}{
	{
		fallback: func(cfg *Config) bool { return cfg.Rpc == nil },
		opt:      DefaultRpcAddrs,
	},
	{
		fallback: func(cfg *Config) bool { return cfg.Timeout == 0 },
		opt:      DefaultTimeout,
	},
	{
		fallback: func(cfg *Config) bool { return cfg.Name == "" },
		opt:      DefaultSDKName,
	},
}

// FallbackDefaults Default options are used if and only
// if no other relevant options apply.
var FallbackDefaults Option = func(cfg *Config) error {
	for _, def := range defaults {
		if !def.fallback(cfg) {
			continue
		}
		if err := cfg.Apply(def.opt); err != nil {
			return err
		}
	}
	return nil
}
