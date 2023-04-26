/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import "github.com/CESSProject/sdk-go/core/utils"

func (c *Cli) UpdateRoleAddress(name string) (string, error) {
	return c.Chain.UpdateAddress(name, c.Node.Multiaddr())
}

func (c *Cli) UpdateIncomeAccount(income string) (string, error) {
	puk, err := utils.ParsingPublickey(income)
	if err != nil {
		return "", err
	}
	return c.Chain.UpdateIncomeAcc(puk)
}
