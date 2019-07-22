#!/usr/bin/env bash

CURRENT_DIR=`pwd`
OLD_GO_PATH="$GOPATH"
OLD_GO_BIN="$GOBIN"

export GOPATH="$CURRENT_DIR" 
export GOBIN="$CURRENT_DIR/bin"

# Specifies and organizes the current source path
gofmt -w src

go install src/$1

export GOPATH="$OLD_GO_PATH"
export GOBIN="$OLD_GO_BIN"

