# golang grpc

[![go version](https://img.shields.io/github/go-mod/go-version/reggiepy/gogrpcTest?color=success&filename=go.mod&style=flat)](https://github.com/reggiepy/gogrpcTest)
[![release](https://img.shields.io/github/v/tag/reggiepy/gogrpcTest?color=success&label=release)](https://github.com/reggiepy/gogrpcTest)
[![build status](https://img.shields.io/badge/build-pass-success.svg?style=flat)](https://github.com/reggiepy/gogrpcTest)
[![License](https://img.shields.io/badge/license-GNU%203.0-success.svg?style=flat)](https://github.com/reggiepy/gogrpcTest)
[![Go Report Card](https://goreportcard.com/badge/github.com/reggiepy/gogrpcTest)](https://goreportcard.com/report/github.com/reggiepy/gogrpcTest)

## Installation

```bash
git clone https://github.com/reggiepy/gogrpcTest.git
cd gogrpcTest
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
