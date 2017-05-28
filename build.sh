#!/usr/bin/env bash

CGO=0 GOARCH=amd64 GOOS=Linux go build -o ./build/main ./src