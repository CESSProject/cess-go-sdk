/*
	Copyright (C) CESS. All rights reserved.
	Copyright (C) Cumulus Encrypted Storage System. All rights reserved.

	SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	p2pgo "github.com/CESSProject/p2p-go"
	"github.com/CESSProject/p2p-go/core"
	"github.com/CESSProject/sdk-go/core/chain"
	"github.com/CESSProject/sdk-go/core/rule"
	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
)

type Client interface {
	chain.Chain
	Workspace() string
	RegisterRole(name string, earnings string, pledge uint64) (string, string, error)
	QueryAllBucketName(owner []byte) ([]string, error)
	CheckSpaceUsageAuthorization(puk []byte) (bool, error)
	QueryUserSpaceSt(puk []byte) (UserSpaceSt, error)
	QueryChallengeSt() (ChallengeSnapshot, error)
	QueryChallenge(puk []byte) (ChallengeInfo, error)
	QueryTeeWorkerPeerID(puk []byte) ([]byte, error)
	QueryTeeWorkerList() ([]TeeWorkerSt, error)
	QuaryRewards(puk []byte) (RewardsType, error)
	CheckBucketName(bucketname string) bool
	ProcessingData(path string) ([]SegmentInfo, string, error)
	PutFile(owner []byte, segmentInfo []SegmentInfo, roothash, filename, bucketname string) (uint8, error)
	GetFile(roothash, dir string) (string, error)
	UpdateRoleAddress(name string) (string, error)
	UpdateIncomeAccount(income string) (string, error)
	SubmitIdleFile(teeAcc []byte, idlefiles []IdleFileMeta) (string, error)
	ReportFiles(roothash []string) (string, []string, error)
	ReplaceFile(roothash []string) (string, []string, error)
	IncreaseSminerStakes(token string) (string, error)
}

type Cli struct {
	chain.Chain
}

func NewBasicCli(rpc []string, name, phase, workspace, addr string, port int, timeout time.Duration) (Client, error) {
	var err error
	var cli = &Cli{}

	cli.Chain, err = chain.NewChainSDK(rpc, phase, timeout)
	if err != nil {
		return cli, err
	}

	keyring, err := signature.KeyringPairFromSecret(phase, 0)
	if err != nil {
		return nil, err
	}

	account, err := utils.EncodePublicKeyAsCessAccount(keyring.PublicKey)
	if err != nil {
		return cli, err
	}

	if workspace == "" {
		return cli, nil
	}

	workspaceActual := filepath.Join(workspace, account, name)
	fstat, err := os.Stat(workspaceActual)
	if err != nil {
		err = os.MkdirAll(workspaceActual, rule.DirMode)
		if err != nil {
			return cli, err
		}
	} else {
		if !fstat.IsDir() {
			return cli, fmt.Errorf("%s is not a directory", workspaceActual)
		}
	}

	privatekeyPath := filepath.Join(workspaceActual, PrivatekeyFile)

	// To construct a simple host with all the default settings, just use `New`
	p2phost, err := p2pgo.New(
		privatekeyPath,
		p2pgo.ListenAddrStrings(addr, port), // regular tcp connections
		p2pgo.Workspace(workspaceActual),
	)
	if err != nil {
		return cli, err
	}

	p2pnode, ok := p2phost.(*core.Node)
	if !ok {
		return cli, errors.New("p2p host type error")
	}

	return cli, nil
}

func (c *Cli) Sign(msg []byte) ([]byte, error) {
	return signature.Sign(msg, c.GetSignatureURI())
}
