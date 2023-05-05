/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"errors"

	"github.com/CESSProject/sdk-go/core/chain"
)

func (c *Cli) RegisterRole(name string, income string, pledge uint64) (string, error) {
	var peerid string
	if len(c.Multiaddr()) > len(chain.PeerID{}) {
		index := len(c.Multiaddr()) - len(chain.PeerID{})
		peerid = c.Multiaddr()[index:]
	}
	if len(peerid) != len(chain.PeerID{}) {
		return "", errors.New("Invalid PeerId")
	}
	return c.Chain.Register(name, c.Multiaddr(), income, pledge)
}
