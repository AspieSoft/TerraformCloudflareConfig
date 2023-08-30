#!/bin/bash

cd $(dirname "$0")

CGO_ENABLED=0 go build -o ../run

CGO_ENABLED=0 GOOS="windows" GOARCH="amd64" go build -o ../run.amd64.exe
CGO_ENABLED=0 GOOS="windows" GOARCH="386" go build -o ../run.386.exe
