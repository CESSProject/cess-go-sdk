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
	"github.com/CESSProject/p2p-go/protocol"
	"github.com/CESSProject/sdk-go/core/chain"
	"github.com/CESSProject/sdk-go/core/rule"
	"github.com/CESSProject/sdk-go/core/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
)

type Client interface {
	Workspace() string
	RegisterRole(name string, income string, pledge uint64) (string, error)
	QueryAllBucketName(owner []byte) ([]string, error)
	CheckSpaceUsageAuthorization(puk []byte) (bool, error)
	QueryUserSpaceSt(puk []byte) (UserSpaceSt, error)
	QueryChallengeSt() (ChallengeSnapshot, error)
	QueryChallenge(puk []byte) (ChallengeInfo, error)
	QueryTeeWorkerPeerID(puk []byte) ([]byte, error)
	QueryTeeWorkerList() ([]TeeWorkerSt, error)
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
	*protocol.Protocol
}

func NewBasicCli(rpc []string, name, phase, workspace, addr string, port int, timeout time.Duration) (Client, error) {
	var err error
	var cli = &Cli{}

	cli.Chain, err = chain.NewChainClient(rpc, phase, timeout)
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

	cli.Protocol = protocol.NewProtocol(p2pnode)
	cli.Protocol.WriteFileProtocol = protocol.NewWriteFileProtocol(p2pnode)
	cli.Protocol.ReadFileProtocol = protocol.NewReadFileProtocol(p2pnode)
	cli.Protocol.AggrProofProtocol = protocol.NewAggrProofProtocol(p2pnode)
	cli.Protocol.CustomDataTagProtocol = protocol.NewCustomDataTagProtocol(p2pnode)
	cli.Protocol.IdleDataTagProtocol = protocol.NewIdleDataTagProtocol(p2pnode)
	cli.Protocol.FileProtocol = protocol.NewFileProtocol(p2pnode)
	cli.Protocol.PushTagProtocol = protocol.NewPushTagProtocol(p2pnode)

	return cli, nil
}

func (c *Cli) Sign(msg []byte) ([]byte, error) {
	return signature.Sign(msg, c.GetSignatureURI())
}
