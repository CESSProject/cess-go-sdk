/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

func (c *Cli) RegisterRole(name string, earnings string, pledge uint64) (string, string, error) {
	return c.Chain.Register(name, c.Node.PeerId, earnings, pledge)
}
