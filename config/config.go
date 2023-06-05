/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"fmt"
	"time"

	"github.com/CESSProject/sdk-go/chain"
	"github.com/CESSProject/sdk-go/core/sdk"
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

// SDK available character name
const (
	CharacterName_Client = "client"
	CharacterName_Bucket = "bucket"
	CharacterName_Deoss  = "deoss"
)

// NewSDK constructs a new client from the Config.
//
// This function consumes the config. Do not reuse it (really!).
func (cfg *Config) NewSDK(characterName string) (sdk.SDK, error) {
	if characterName != CharacterName_Bucket &&
		characterName != CharacterName_Deoss &&
		characterName != CharacterName_Client {
		return nil, fmt.Errorf("invalid character name")
	}
	return chain.NewChainSDK(characterName, cfg.Rpc, cfg.Mnemonic, cfg.Timeout)
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
