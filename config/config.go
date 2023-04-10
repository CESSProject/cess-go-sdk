/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"time"

	"github.com/CESSProject/sdk-go/core/client"
)

// Config describes a set of settings for a client
type Config struct {
	Rpc       []string
	Phase     string
	Workspace string
	Addr      string
	Name      string
	Port      int
	Timeout   time.Duration
}

// Option is a client config option that can be given to the client constructor
type Option func(cfg *Config) error

const DefaultName = "SDK"

// BlockInterval is the time interval for generating blocks, in seconds
const BlockInterval = time.Second * time.Duration(6)

// NewNode constructs a new client from the Config.
//
// This function consumes the config. Do not reuse it (really!).
func (cfg *Config) NewClient(name string) (client.Client, error) {
	var serviceName = DefaultName
	if name != "" {
		serviceName = name
	}
	return client.NewBasicCli(cfg.Rpc, serviceName, cfg.Phase, cfg.Workspace, cfg.Addr, cfg.Port, cfg.Timeout)
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
