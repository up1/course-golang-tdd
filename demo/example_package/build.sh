#!/usr/bin/env bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src/

go install demo

export GOPATH="$OLDGOPATH"

echo 'finished'