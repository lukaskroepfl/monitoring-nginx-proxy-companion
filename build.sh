#!/usr/bin/env bash

go build --ldflags '-extldflags "-static"' -o ./build/main ./src