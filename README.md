# golang grpc

[![go version](https://img.shields.io/github/go-mod/go-version/reggiepy/grpcTest?color=success&filename=go.mod&style=flat)](https://github.com/reggiepy/grpcTest)
[![release](https://img.shields.io/github/v/tag/reggiepy/grpcTest?color=success&label=release)](https://github.com/reggiepy/grpcTest)
[![build status](https://img.shields.io/badge/build-pass-success.svg?style=flat)](https://github.com/reggiepy/grpcTest)
[![License](https://img.shields.io/badge/license-GNU%203.0-success.svg?style=flat)](https://github.com/reggiepy/grpcTest)
[![Go Report Card](https://goreportcard.com/badge/github.com/reggiepy/grpcTest)](https://goreportcard.com/report/github.com/reggiepy/grpcTest)

## Installation

```bash
git clone https://github.com/reggiepy/grpcTest.git
cd grpcTest
go mod tidy
```

## Usage

```bash
go run main.go
```

## generator proto

```bash
protoc -Iproto --go_out=plugins=grpc:. common.proto
```

## Architecture
