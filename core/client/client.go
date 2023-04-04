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
)

type SegmentInfo struct {
	SegmentHash  string
	FragmentHash []string
}

type Client interface {
	PutFile(path, owner, filename, bucketname string) error
	DeleteFile(roothash string) error
	DeleteBucket(bucketName string) error
}

type Cli struct {
	chain.Chain
	*protocol.Protocol
}

func NewBasicCli(rpc []string, name, phase, workspace, addr string, port, timeout int) (Client, error) {
	var err error
	var cli = &Cli{}

	cli.Chain, err = chain.NewChainClient(rpc, phase, time.Duration(time.Second*time.Duration(timeout)))
	if err != nil {
		return cli, err
	}
	account, err := cli.Chain.GetCessAccount()
	if err != nil {
		return cli, err
	}
	workspaceActual := filepath.Join(workspace, account, name)
	fstat, err := os.Stat(workspaceActual)
	if err != nil {
		err = os.MkdirAll(workspaceActual, 0755)
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
		p2pgo.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/%d", addr, port), // regular tcp connections
		),
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

	return cli, nil
}
