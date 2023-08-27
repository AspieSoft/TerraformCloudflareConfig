#!/bin/bash

cd $(dirname "$0")

go build -o ../run

GOOS=windows go build -o ../run.exe
