/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package sdkgo

import (
	"context"
	"time"

	"github.com/CESSProject/cess-go-sdk/chain"
	"github.com/CESSProject/cess-go-sdk/config"
)

// Config describes a set of settings for a client
type Config struct {
	Rpc      []string
	Mnemonic string
	Name     string
	Timeout  time.Duration
}

// Option is a client config option that can be given to the client constructor
type Option func(cfg *Config) error

// // Option is a client config option that can be given to the client constructor
// type Option func(cfg *Config) error

// NewSDK constructs a new client from the Config.
//
// This function consumes the config. Do not reuse it (really!).
func (cfg *Config) NewSDK(ctx context.Context) (*chain.ChainClient, error) {
	if cfg.Name == "" {
		cfg.Name = config.CharacterName_Default
	}
	return chain.NewChainClient(ctx, cfg.Name, cfg.Rpc, cfg.Mnemonic, cfg.Timeout)
}

// Apply applies the given options to the config, returning the first error
// encountered (if any).
func (cfg *Config) Apply(opts ...Option) error {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(cfg); err != nil {
			return err
		}
	}
	return nil
}
