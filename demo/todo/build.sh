#!/usr/bin/env bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src/

go install task
go install server

go test ./... -v -cover

export GOPATH="$OLDGOPATH"

echo 'finished'