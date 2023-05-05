/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"github.com/CESSProject/sdk-go/core/chain"
	"github.com/CESSProject/sdk-go/core/utils"
)

func (c *Cli) Workspace() string {
	return c.Node.Workspace()
}

func (c *Cli) CheckSpaceUsageAuthorization(puk []byte) (bool, error) {
	grantor, err := c.Chain.QuaryAuthorizedAcc(puk)
	if err != nil {
		if err.Error() == chain.ERR_Empty {
			return false, nil
		}
		return false, err
	}
	account_chain, _ := utils.EncodePublicKeyAsCessAccount(grantor[:])
	account_local := c.Chain.GetSignatureAcc()

	return account_chain == account_local, nil
}

func (c *Cli) QueryUserSpaceSt(puk []byte) (UserSpaceSt, error) {
	var userSpaceSt UserSpaceSt
	spaceinfo, err := c.Chain.QueryUserSpaceInfo(puk)
	if err != nil {
		return userSpaceSt, err
	}
	userSpaceSt.Start = uint32(spaceinfo.Start)
	userSpaceSt.Deadline = uint32(spaceinfo.Deadline)
	userSpaceSt.TotalSpace = spaceinfo.TotalSpace.String()
	userSpaceSt.UsedSpace = spaceinfo.UsedSpace.String()
	userSpaceSt.RemainingSpace = spaceinfo.RemainingSpace.String()
	userSpaceSt.LockedSpace = spaceinfo.LockedSpace.String()
	userSpaceSt.State = string(spaceinfo.State)
	return userSpaceSt, nil
}

func (c *Cli) QueryTeeWorkerList() ([]TeeWorkerSt, error) {
	teelist, err := c.Chain.QueryTeeInfoList()
	if err != nil {
		return nil, err
	}
	var results = make([]TeeWorkerSt, len(teelist))
	for k, v := range teelist {
		results[k].Node_key = []byte(string(v.NodeKey.NodePublickey[:]))
		results[k].Peer_id = []byte(string(v.PeerId[:]))
		results[k].Controller_account, err = utils.EncodePublicKeyAsCessAccount(v.ControllerAccount[:])
		if err != nil {
			return results, err
		}
		results[k].Stash_account, err = utils.EncodePublicKeyAsCessAccount(v.StashAccount[:])
		if err != nil {
			return results, err
		}
	}
	return results, nil
}

func (c *Cli) QueryTeeWorkerPeerID(puk []byte) ([]byte, error) {
	peerid, err := c.Chain.QueryTeePeerID(puk)
	if err != nil {
		return nil, err
	}
	return []byte(string(peerid[:])), nil
}

func (c *Cli) QueryNodeSynchronizationSt() (bool, error) {
	return c.Chain.QueryNodeSynchronizationSt()
}

func (c *Cli) QueryNodeConnectionSt() bool {
	return c.Chain.QueryNodeConnectionSt()
}

func (c *Cli) QuaryAuthorizedAcc(puk []byte) (string, error) {
	acc, err := c.Chain.QuaryAuthorizedAcc(puk)
	if err != nil {
		return "", err
	}
	return utils.EncodePublicKeyAsCessAccount(acc[:])
}
