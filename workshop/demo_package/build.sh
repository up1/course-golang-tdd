#!/usr/bin/env bash

CURDIR=`pwd`
export GOPATH="$CURDIR"

gofmt -w src/
go install mypackage
go install main
