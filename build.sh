#!/usr/bin/env bash

CURRENT_DIR=`pwd`
OLD_GO_PATH="$GOPATH"
OLD_GO_BIN="$GOBIN"
OLD_GO_ROOT="$GOROOT"

# export GOPATH="$CURRENT_DIR"
export GOPATH="/Users/wangjun/go" 
export GOBIN="$CURRENT_DIR/bin"
# export GOROOT="CURRENT_DIR/pkg"

# export http_proxy='http://localhost:1087'
# export https_proxy='http://localhost:1087'

# CGO_ENABLE=0
# GOOS=linux
# GOARCH=amd64

# Specifies and organizes the current source path
gofmt -w src

if [[ $2 == '1' ]]
then
    echo "build $1 for linux"
    CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build $1
else
    echo "build $1 for Mac"
    go install $1
fi

export GOPATH="$OLD_GO_PATH"
export GOBIN="$OLD_GO_BIN"
# export GOROOT="$OLD_GO_ROOT"

# unset http_proxy
# unset https_proxy
