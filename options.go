/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdkgo

import (
	"time"

	"github.com/CESSProject/sdk-go/core/rule"
)

// ConnectRpcAddrs configures client to connect to the given RPC
// addresses.
func ConnectRpcAddrs(s []string) Option {
	return func(cfg *Config) error {
		cfg.Rpc = s
		return nil
	}
}

// Workspace configures client to use the given workspace.
func Mnemonic(mnemonic string) Option {
	return func(cfg *Config) error {
		cfg.Mnemonic = mnemonic
		return nil
	}
}

// TransactionTimeout configures the transaction timeout period.
func TransactionTimeout(timeout time.Duration) Option {
	return func(cfg *Config) error {
		if timeout < rule.BlockInterval {
			cfg.Timeout = rule.BlockInterval
		} else {
			cfg.Timeout = timeout
		}
		return nil
	}
}
