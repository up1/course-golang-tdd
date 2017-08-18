#!/bin/bash
export GOPATH=`pwd`
rm -rf bin

# go get -u github.com/jhillyerd/enmime
# git clone -b v2 git@github.com:go-mgo/mgo.git $GOPATH/src/gopkg.in/mgo.v2

gofmt -w src/readmail
go test -v --cover $(go list ./... | grep -v github.com | grep -v golang.org | grep -v gopkg.in)
go install readmail
./bin/readmail
