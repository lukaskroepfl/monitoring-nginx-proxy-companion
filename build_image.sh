#!/usr/bin/env bash

./build.sh

docker build -t lukaskroepfl/monitoring-nginx-proxy-companion:`git describe` -t lukaskroepfl/monitoring-nginx-proxy-companion:latest .
