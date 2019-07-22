#!/usr/bin/env bash

CURRENT_DIR=`pwd`
OLD_GO_PATH="$GOPATH"
OLD_GO_BIN="$GOBIN"

export GOPATH="$CURRENT_DIR" 
export GOBIN="$CURRENT_DIR/bin"

SET CGO_ENABLE=0
SET GOOS=linux
SET GOARCH=amd64

# Specifies and organizes the current source path
gofmt -w src

go install src/$1

export GOPATH="$OLD_GO_PATH"
export GOBIN="$OLD_GO_BIN"

