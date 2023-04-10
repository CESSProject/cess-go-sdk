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
)

type Client interface {
	Workspace() string
	Register(name string, income string, pledge uint64) (string, error)
	QueryStorageMiner(pubkey []byte) (chain.MinerInfo, error)
	QueryDeoss(pubkey []byte) (string, error)
	QueryFile(roothash string) (chain.FileMetaInfo, error)
	QueryBucket(owner []byte, bucketname string) (chain.BucketInfo, error)
	QueryBuckets(owner []byte) ([]string, error)
	QueryGrantor(pubkey []byte) (bool, error)
	CheckBucketName(bucketname string) bool
	CreateBucket(owner []byte, bucketname string) (string, error)
	ProcessingData(path string) ([]SegmentInfo, string, error)
	PutFile(owner []byte, segmentInfo []SegmentInfo, roothash, filename, bucketname string) (string, error)
	DeleteFile(owner []byte, roothash []string) (string, []chain.FileHash, error)
	DeleteBucket(owner []byte, bucketName string) (string, error)
	Update(name string) (string, error)
	SubmitIdleFile(size uint64, blockNum, blocksize, scansize uint32, pubkey []byte, hash string) (string, error)
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

	//
	os.MkdirAll(filepath.Join(workspaceActual, rule.FileDir), rule.DirMode)
	os.MkdirAll(filepath.Join(workspaceActual, rule.TempDir), rule.DirMode)
	os.MkdirAll(filepath.Join(workspaceActual, rule.DbDir), rule.DirMode)

	return cli, nil
}
