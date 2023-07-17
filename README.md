<div align="center">

# Go sdk for CESS network

[![GitHub license](https://img.shields.io/badge/license-Apache2-blue)](#LICENSE)
<a href=""><img src="https://img.shields.io/badge/golang-%3E%3D1.19-blue.svg" /></a>
[![Go Reference](https://pkg.go.dev/badge/github.com/CESSProject/cess-go-sdk.svg)](https://pkg.go.dev/github.com/CESSProject/cess-go-sdk)
[![build](https://github.com/CESSProject/cess-go-sdk/actions/workflows/build&test.yml/badge.svg)](https://github.com/CESSProject/cess-go-sdk/actions/workflows/build&test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/CESSProject/cess-go-sdk)](https://goreportcard.com/report/github.com/CESSProject/cess-go-sdk)

</div>

The go sdk implementation of the CESS network, which provides RPC calls, status queries, block transactions and other functions.

## 📝 Reporting Vulnerability

If you find out any system bugs or you have a better suggestions, please send an email to frode@cess.one or join [CESS discord](https://discord.gg/mYHTMfBwNS) to communicate with us.

## 📢 Announcement
**CESS test network rpc endpoints**
```
wss://testnet-rpc0.cess.cloud/ws/
wss://testnet-rpc1.cess.cloud/ws/
wss://testnet-rpc2.cess.cloud/ws/
```
**CESS test network bootstrap node**
```
_dnsaddr.boot-kldr-testnet.cess.cloud
```

## 🚰 CESS test network faucet
```
https://testnet-faucet.cess.cloud/
```

## 🏗 Usage

To get the package use the standard:

```sh
go get -u "github.com/CESSProject/cess-go-sdk"
```

## ✅ Testing

To run test:

1. Run a [CESS node](https://github.com/CESSProject/cess) locally.
2. Run the command

	```sh
	go test -v
	```

## 📖 Documentation & Examples

Please refer to: https://pkg.go.dev/github.com/CESSProject/cess-go-sdk

## 💡 Example

Usually, you only care about how to access your data in the CESS network, you need to build such a web service yourself, this sdk will help you quickly realize data access. Note that [p2p-go](https://github.com/CESSProject/p2p-go) library needs to be used to enable the data transmission.


### Create an sdk client

To create an sdk client, you need to provide some configuration information: your rpc address (if not, use the rpc address disclosed by CESS), your wallet private key, and transaction timeout. Please refer to the following examples:

```go
sdk, err := cess.New(
	context.Background(),
	config.CharacterName_Client,
	cess.ConnectRpcAddrs([]string{"<rpc addr>"}),
	cess.Mnemonic("<your account mnmonic>"),
	cess.TransactionTimeout(time.Second * 10),
)
```

### Create an sdk client with p2p functionality

When you need to store data or download data you need to initialize an sdk with p2p network, refer to the following code:
```go
sdk, err := cess.New(
	context.Background(),
	config.CharacterName_Client,
	cess.ConnectRpcAddrs([]string{"<rpc addr>"}),
	cess.Mnemonic("<your account mnmonic>"),
	cess.TransactionTimeout(time.Second * 10),
	cess.Workspace("<work space>"),
	cess.P2pPort(<port>),
	cess.Bootnodes([]string{"<bootstrap node>"}),
	cess.ProtocolPrefix("<protocol prefix>"),
)
```

### Create storage data bucket
cess as an object storage service, the data are stored in buckets, which can be created automatically when uploading data, or separately, refer to the following code:
```go
sdk, err := cess.New(
	context.Background(),
	config.CharacterName_Client,
	cess.ConnectRpcAddrs([]string{"<rpc addr>"}),
	cess.Mnemonic("<your account mnmonic>"),
	cess.TransactionTimeout(time.Second*10),
)
if err != nil {
	panic(err)
}
fmt.Println(sdk.CreateBucket(sdk.GetSignatureAccPulickey(), "<your bucket name>"))
```

### Store data
You need to purchase space with your account before uploading files, please refer to [Buy Space](https://github.com/CESSProject/W3F-illustration/blob/4995c1584006823990806b9d30fa7d554630ec14/deoss/buySpace.png).
The following is an example of uploading a file:
```go
sdk, err := cess.New(
	context.Background(),
	config.CharacterName_Client,
	cess.ConnectRpcAddrs([]string{"<rpc addr>"}),
	cess.Mnemonic("<your account mnmonic>"),
	cess.TransactionTimeout(time.Second * 10),
	cess.Workspace("<work space>"),
	cess.P2pPort(<port>),
	cess.Bootnodes([]string{"<bootstrap node>"}),
	cess.ProtocolPrefix("<protocol prefix>"),
)
if err != nil {
	panic(err)
}
fmt.Println(sdk.StoreFile("<your file>", "<your bucket name>"))
```

### Retrieve data
To retrieve the data, you need to provide the unique hash of the data, which will be returned to you when the data is uploaded successfully, here is the sample code to retrieve the data:
```go
sdk, err := cess.New(
	context.Background(),
	config.CharacterName_Client,
	cess.ConnectRpcAddrs([]string{"<rpc addr>"}),
	cess.Mnemonic("<your account mnmonic>"),
	cess.TransactionTimeout(time.Second * 10),
	cess.Workspace("<work space>"),
	cess.P2pPort(<port>),
	cess.Bootnodes([]string{"<bootstrap node>"}),
	cess.ProtocolPrefix("<protocol prefix>"),
)
if err != nil {
	panic(err)
}
fmt.Println(sdk.RetrieveFile("<file hash>", "<save path>"))
```

## License

Licensed under [Apache 2.0](https://github.com/CESSProject/cess-go-sdk/blob/main/LICENSE)
