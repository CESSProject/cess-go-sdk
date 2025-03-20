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

CESS Network SDK for Go 允许您访问CESS区块链网络，如查询区块数据、交易、调用RPC方法。您无需处理 API 相关任务（如签名和构建请求）即可访问CESS区块链网络。还允许您访问存储网络，如直接从存储节点上传或下载数据，以及文件分块，加密，冗余的实现。

## 公告
- 测试网RPC端点
```
wss://testnet-rpc.cess.network/ws/
```

- 测试网水龙头地址
```
https://www.cess.network/faucet.html
```

## 环境要求

安装比 1.22.x 更新的Go环境。

## 安装

使用 `go get` 安装SDK：

```sh
go get -u "github.com/CESSProject/cess-go-sdk"
```

## 快速使用

快速创建您的 SDK 客户端：
```golang
cli, err := sdkgo.New(
    context.Background(),
    sdkgo.ConnectRpcAddrs("wss://testnet-rpc.cess.network/ws/"),
)
```

## 测试

运行测试：

```sh
make check
```

## 文档

- [Guidebook](https://doc.cess.network/developer/cess-sdk/sdk-golang)
- [Reference](https://pkg.go.dev/github.com/CESSProject/cess-go-sdk)

## 问题

如果你发现任何系统错误或者有更好的建议，请提交[issue](https://github.com/CESSProject/cess-go-sdk/issues/new)或者PR，或者加入[CESS discord](https://discord.gg/mYHTMfBwNS)与我们交流。

## License

许可依据 [Apache 2.0](https://github.com/CESSProject/cess-go-sdk/blob/main/LICENSE)








