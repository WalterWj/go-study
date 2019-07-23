#!/usr/bin/env bash

CURRENT_DIR=`pwd`
OLD_GO_PATH="$GOPATH"
OLD_GO_BIN="$GOBIN"
OLD_GO_ROOT="$GOROOT"

export GOPATH="$CURRENT_DIR" 
export GOBIN="$CURRENT_DIR/bin"
# export GOROOT="$CURRENT_DIR/src"

# export http_proxy='http://localhost:1087'
# export https_proxy='http://localhost:1087'

# SET CGO_ENABLE=0
# SET GOOS=linux
# SET GOARCH=amd64

# Specifies and organizes the current source path
gofmt -w src

go install $1

export GOPATH="$OLD_GO_PATH"
export GOBIN="$OLD_GO_BIN"
# export GOROOT="$OLD_GO_ROOT"

# unset http_proxy
# unset https_proxy
