#!/bin/bash
set -euxo pipefail
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
go build -o director -trimpath .
