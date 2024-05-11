<div align="center">

# Go sdk for CESS network

[![GitHub license](https://img.shields.io/badge/license-Apache2-blue)](#LICENSE)
<a href=""><img src="https://img.shields.io/badge/golang-%3E%3D1.20-blue.svg" /></a>
[![Go Reference](https://pkg.go.dev/badge/github.com/CESSProject/cess-go-sdk.svg)](https://pkg.go.dev/github.com/CESSProject/cess-go-sdk)
[![build](https://github.com/CESSProject/cess-go-sdk/actions/workflows/build&test.yml/badge.svg)](https://github.com/CESSProject/cess-go-sdk/actions/workflows/build&test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/CESSProject/cess-go-sdk)](https://goreportcard.com/report/github.com/CESSProject/cess-go-sdk)

</div>

The go sdk implementation of the CESS network, which provides RPC calls, status queries, block transactions and other functions.

## 📝 Reporting Vulnerability

If you find any system errors or you have better suggestions, please submit an issue or PR, or join the [CESS discord](https://discord.gg/mYHTMfBwNS) to communicate with us.

## 📢 Announcement
**CESS test network rpc endpoints**
```
wss://testnet-rpc0.cess.cloud/ws/
wss://testnet-rpc1.cess.cloud/ws/
wss://testnet-rpc2.cess.cloud/ws/
wss://testnet-rpc3.cess.cloud/ws/
```
**CESS test network bootstrap node**
```
_dnsaddr.boot-miner-testnet.cess.cloud
```

## 🚰 CESS test network faucet
```
https://testnet-faucet.cess.cloud/
```

## 🏗 Usage

To get the package use the standard:

```sh
go get "github.com/CESSProject/cess-go-sdk"
```

## ✅ Testing

To run test:

```sh
make check
```

## 📖 Document

- [Reference](https://pkg.go.dev/github.com/CESSProject/cess-go-sdk)
- [Guidebook](https://docs.cess.cloud/deoss/get-started/go-sdk)

## License

Licensed under [Apache 2.0](https://github.com/CESSProject/cess-go-sdk/blob/main/LICENSE)
