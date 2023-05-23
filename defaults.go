/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdkgo

import (
	"github.com/CESSProject/sdk-go/core/rule"
)

// DefaultRpcAddrs configures client to use default RPC address.
var DefaultRpcAddrs = func(cfg *Config) error {
	rpcAddrs := []string{
		"wss://testnet-rpc0.cess.cloud/ws/",
		"wss://testnet-rpc1.cess.cloud/ws/",
	}
	return cfg.Apply(ConnectRpcAddrs(rpcAddrs))
}

// DefaultListenAddr configures client to use default listen address.
var DefaultListenAddr = func(cfg *Config) error {
	addrs := "0.0.0.0"
	return cfg.Apply(ListenAddrStrings(addrs))
}

// DefaultListenPort configures client to use default listen port.
var DefaultListenPort = func(cfg *Config) error {
	port := 15000
	return cfg.Apply(ListenPort(port))
}

// DefaultListenPort configures client to use default listen port.
var DefaultTimeout = func(cfg *Config) error {
	return cfg.Apply(TransactionTimeout(rule.BlockInterval))
}

// Complete list of default options and when to fallback on them.
//
// Please *DON'T* specify default options any other way. Putting this all here
// makes tracking defaults *much* easier.
var defaults = []struct {
	fallback func(cfg *Config) bool
	opt      Option
}{
	{
		fallback: func(cfg *Config) bool { return cfg.Rpc == nil },
		opt:      DefaultRpcAddrs,
	},
	{
		fallback: func(cfg *Config) bool { return cfg.Addr == "" },
		opt:      DefaultListenAddr,
	},
	{
		fallback: func(cfg *Config) bool { return cfg.Port == 0 },
		opt:      DefaultListenPort,
	},
	{
		fallback: func(cfg *Config) bool { return cfg.Timeout == 0 },
		opt:      DefaultTimeout,
	},
}

// FallbackDefaults applies default options to the libp2p node if and only if no
// other relevant options have been applied. will be appended to the options
// passed into New.
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
