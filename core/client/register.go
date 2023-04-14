/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

func (c *Cli) Register(name string, income string, pledge uint64) (string, error) {
	return c.Chain.Register(name, c.Multiaddr(), income, pledge)
}
