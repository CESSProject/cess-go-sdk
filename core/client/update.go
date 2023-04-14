/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import "github.com/CESSProject/sdk-go/core/utils"

func (c *Cli) UpdateAddress(name string) (string, error) {
	return c.Chain.UpdateAddress(name, c.Node.Multiaddr())
}

func (c *Cli) UpdateIncomeAccount(income string) (string, error) {
	pubkey, err := utils.ParsingPublickey(income)
	if err != nil {
		return "", err
	}
	return c.Chain.UpdateIncomeAccount(pubkey)
}
