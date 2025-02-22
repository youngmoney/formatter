#!/usr/bin/env bash

GOOS=darwin GOARCH=arm64 go build -o bin/formatter-darwin-arm64 .
GOOS=linux GOARCH=amd64 go build -o bin/formatter-linux-amd64 .
