#!/bin/sh
export GOCACHE=/tmp
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
go get github.com/gorilla/mux
go get github.com/tidwall/gjson
go build -o main .