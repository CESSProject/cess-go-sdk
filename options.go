/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdkgo

// ConnectRpcAddrs configures client to connect to the given RPC
// addresses.
func ConnectRpcAddrs(s []string) Option {
	return func(cfg *Config) error {
		cfg.Rpc = s
		return nil
	}
}

// ListenAddrStrings configures client to listen on the given (unparsed)
// addresses.
func ListenAddrStrings(s string) Option {
	return func(cfg *Config) error {
		cfg.Addr = s
		return nil
	}
}

// ListenPort configures client to listen on the given (unparsed)
// port.
func ListenPort(port int) Option {
	return func(cfg *Config) error {
		cfg.Port = port
		return nil
	}
}

// Workspace configures client to use the given workspace.
func Workspace(workspace string) Option {
	return func(cfg *Config) error {
		cfg.Workspace = workspace
		return nil
	}
}

// Workspace configures client to use the given workspace.
func Mnemonic(mnemonic string) Option {
	return func(cfg *Config) error {
		cfg.Phase = mnemonic
		return nil
	}
}
