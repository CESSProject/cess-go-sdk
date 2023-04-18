<div align="center">

# Go sdk for CESS network

[![GitHub license](https://img.shields.io/badge/license-Apache2-blue)](#LICENSE) <a href=""><img src="https://img.shields.io/badge/golang-%3E%3D1.19-blue.svg" /></a> [![Go Reference](https://pkg.go.dev/badge/github.com/CESSProject/sdk-go.svg)](https://pkg.go.dev/github.com/CESSProject/sdk-go)

</div>

The go sdk implementation of the CESS network, which provides RPC calls, status queries, and access to the p2p storage network of the CESS chain.

## Reporting a Vulnerability
If you find out any vulnerability, Please send an email to frode@cess.one, we are happy to communicate with you.

## Installation
To get the package use the standard:
```
go get -u "github.com/CESSProject/sdk-go"
```
Using Go modules is recommended.

## Documentation & Examples
Please refer to https://pkg.go.dev/github.com/CESSProject/sdk-go

## Usage
Usually, you only care about how to access your data in the CESS network, you need to build such a web service yourself, this sdk will help you quickly realize data access.

#### Create an sdk client instance
The following is an example of creating an sdk client:
```
cli, err := New(
    config.DefaultName,
    ConnectRpcAddrs([]string{""}),
    ListenPort(15000),
    Workspace("/"),
    Mnemonic("xxx xxx ... xxx"),
    TransactionTimeout(time.Duration(time.Second*10)),
)
```
Creating a client requires you to configure some information, you can refer to the following configuration files:
```
# The rpc endpoint of the chain node
Rpc:
  - "ws://127.0.0.1:9948/"
  - "wss://testnet-rpc0.cess.cloud/ws/"
  - "wss://testnet-rpc1.cess.cloud/ws/"
# Account mnemonic
Mnemonic: "xxx xxx xxx"
# Service workspace
Workspace: /
# Service running address
Address: "127.0.0.1"
# P2P communication port
P2P_Port: 8088
# Service listening port
HTTP_Port: 15000
```

#### Register as an oss role
Call the Register method to register. Note that the first parameter specifies that the role you register is oss, and its value can be any one of "oss","OSS","Deoss","DEOSS".
```
txhash, err := cli.Register("oss", "", 0)
```

#### Process your documents according to the specifications of CESS
Call the ProcessingData method to process your file. You need to specify the file path. The method returns a segment list and the unique identifier hash of the file in the CESS network.
```
segmentInfo, roothash, err := cli.ProcessingData(filepath)
```

#### Store your files
Call the PutFile method to store your files. You need to specify a bucket to store the files. CESS will automatically create the bucket for you, provided that the name of the bucket is legal. You can call the CheckBucketName method in advance to check whether the name of the bucket meets the requirements.
After the storage is successful, it will return `count`. The `count` indicates the number of times your file is stored. A file can only be stored up to 5 times. If your file is not stored successfully after 5 times, you need to upload your file again.
```
count, err := cli.PutFile(publickey, segmentInfo, roothash, filename, buckname)
```

#### Download file
Call the GetFile method to download the file you want, and this method will save the file you downloaded under the roothash name in the directory you specify.
```
filepath, err := n.Cli.GetFile(roothash, dir)
```

## License
Licensed under [Apache 2.0](https://github.com/CESSProject/sdk-go/blob/main/LICENSE)