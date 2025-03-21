English | [简体中文](README_CN.md)

<p align="center">
<a href="https://cess.network/"><img src="https://github.com/CESSProject/doc-v2/blob/main/assets/introduction/banner.jpg"></a>
</p>

<h1 align="center">CESS Network SDK for Go</h1>

<div align="center">

[![GitHub license](https://img.shields.io/badge/license-Apache2-blue)](#LICENSE)
<a href=""><img src="https://img.shields.io/badge/golang-%3E%3D1.22-blue.svg" /></a>
[![Go Reference](https://pkg.go.dev/badge/github.com/CESSProject/cess-go-sdk.svg)](https://pkg.go.dev/github.com/CESSProject/cess-go-sdk)
[![build](https://github.com/CESSProject/cess-go-sdk/actions/workflows/build&test.yml/badge.svg)](https://github.com/CESSProject/cess-go-sdk/actions/workflows/build&test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/CESSProject/cess-go-sdk)](https://goreportcard.com/report/github.com/CESSProject/cess-go-sdk)

</div>

CESS Network SDK for Go allows you to access the CESS blockchain network, such as querying block data, transactions, and calling RPC methods. You don't need to deal with API related tasks such as signing and building requests to access the CESS blockchain network. It also allows you to access the storage network, such as uploading or downloading data directly from storage nodes, as well as implementations of file chunking, encryption, and redundancy.


## Bulletin
- Test Network RPC Endpoint
```
wss://testnet-rpc.cess.network/ws/
```

- Test Network Faucet
```
https://www.cess.network/faucet.html
```

## Requirements
Install a Go environment newer than 1.22.x.

## Installation

Use `go get` to install SDK：

```sh
go get -u "github.com/CESSProject/cess-go-sdk"
```

## Quick Use
Quickly create your SDK client:
```golang
cli, err := sdkgo.New(
    context.Background(),
    sdkgo.ConnectRpcAddrs("wss://testnet-rpc.cess.network/ws/"),
)
```


## Testing

To run test:

```sh
make check
```

## Documentation
- [Guidebook](https://doc.cess.network/developer/cess-sdk/sdk-golang)
- [Reference](https://pkg.go.dev/github.com/CESSProject/cess-go-sdk)

## Issues

If you find any system errors or you have better suggestions, please submit an [issue](https://github.com/CESSProject/cess-go-sdk/issues/new) or PR, or join the [CESS discord](https://discord.gg/mYHTMfBwNS) to communicate with us.

## License

Licensed under [Apache 2.0](https://github.com/CESSProject/cess-go-sdk/blob/main/LICENSE)
