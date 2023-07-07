/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"context"
	"fmt"
	"time"

	"github.com/CESSProject/cess-go-sdk/chain"
	"github.com/CESSProject/cess-go-sdk/core/sdk"
)

// Config describes a set of settings for a client
type Config struct {
	Rpc            []string
	Bootnodes      []string
	Mnemonic       string
	Name           string
	Workspace      string
	ProtocolPrefix string
	P2pPort        int
	Timeout        time.Duration
}

// Option is a client config option that can be given to the client constructor
type Option func(cfg *Config) error

// default service name
const (
	CharacterName_Client = "client"
	CharacterName_Bucket = "bucket"
	CharacterName_Deoss  = "deoss"
)

// cess network protocol prefix
const (
	DevnetProtocolPrefix  = "/kldr-devnet"
	TestnetProtocolPrefix = "/kldr-testnet"
	MainnetProtocolPrefix = "/kldr-mainnet"
)

// NewSDK constructs a new client from the Config.
//
// This function consumes the config. Do not reuse it (really!).
func (cfg *Config) NewSDK(ctx context.Context, serviceName string) (sdk.SDK, error) {
	if serviceName == "" {
		return nil, fmt.Errorf("empty service name")
	}
	return chain.NewChainSDK(ctx, serviceName, cfg.Rpc, cfg.Mnemonic, cfg.Timeout, cfg.Workspace, cfg.P2pPort, cfg.Bootnodes, cfg.ProtocolPrefix)
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
