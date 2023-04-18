package sdkgo

import (
	"fmt"
	"time"

	"github.com/CESSProject/sdk-go/config"
)

func Example_newClient() {
	cli, err := New(
		config.DefaultName,
		ConnectRpcAddrs([]string{""}),
		ListenPort(15000),
		Workspace("/"),
		Mnemonic("xxx xxx ... xxx"),
		TransactionTimeout(time.Duration(time.Second*10)),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Client created successfully and the workspace is %s\n", cli.Workspace())
}

func Example_registerOss() {
	cli, err := New(
		"oss",
		ConnectRpcAddrs([]string{""}),
		ListenPort(15000),
		Workspace("/"),
		Mnemonic(""),
		TransactionTimeout(time.Duration(time.Second*10)),
	)
	if err != nil {
		panic(err)
	}
	txhash, err := cli.Register("oss", "", 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("OSS registration successful, transaction hash is %s\n", txhash)
}

func Example_registerMiner() {
	cli, err := New(
		"bucket",
		ConnectRpcAddrs([]string{""}),
		ListenPort(15000),
		Workspace("/"),
		Mnemonic(""),
		TransactionTimeout(time.Duration(time.Second*10)),
	)
	if err != nil {
		panic(err)
	}
	txhash, err := cli.Register("bucket", "cXxxx...xxx", 100000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Miner registration successful, transaction hash is %s\n", txhash)
}
