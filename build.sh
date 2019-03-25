#!/usr/bin/env bash

CURRENT_DIR=`pwd`
OLD_GO_PATH="$GOPATH"
OLD_GO_BIN="$GOBIN"

export GOPATH="$CURRENT_DIR" 
export GOBIN="$CURRENT_DIR/bin"

#指定并整理当前的源码路径
gofmt -w src

go install src/test.go

export GOPATH="$OLD_GO_PATH"
export GOBIN="$OLD_GO_BIN"