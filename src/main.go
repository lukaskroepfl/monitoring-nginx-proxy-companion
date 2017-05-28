package main

import (
  "os"
)

const PROXY_CONTAINER_NAME_ENV_NAME = "PROXY_CONTAINER_NAME"
const PROXY_CONTAINER_NAME_DEFAULT = "nginx"

const INFLUX_URL = "http://localhost:8086"

func logCallback(logLine string) {
  parsedLogline, err := ParseProxyLogLine(logLine)
  if err != nil {
    panic(err)
  }

  WriteToInflux(parsedLogline)
}

func getProxyContainerName() string {
  envProxyContainerName := os.Getenv(PROXY_CONTAINER_NAME_ENV_NAME)
  if envProxyContainerName == "" {
    return PROXY_CONTAINER_NAME_DEFAULT
  }

  return envProxyContainerName
}

func main() {
  proxyContainerId := FindProxyContainerId()

  AttachContainerLogListener(proxyContainerId, logCallback)
}