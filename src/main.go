package main

import "log"

const PROXY_CONTAINER_NAME_ENV_NAME = "PROXY_CONTAINER_NAME"
const PROXY_CONTAINER_NAME_DEFAULT = "nginx"

const INFLUX_URL_ENV_NAME = "INFLUX_URL"
const INFLUX_URL_DEFAULT = "http://localhost:8086"

func logCallback(logLine string) {
  parsedLogline, err := ParseProxyLogLine(logLine)
  if err != nil {
    panic(err)
  }

  WriteToInflux(parsedLogline)
}

func getProxyContainerName() string {
  return GetEnvOrDefault(PROXY_CONTAINER_NAME_ENV_NAME, PROXY_CONTAINER_NAME_DEFAULT)
}

func getInfluxUrl() string {
  return GetEnvOrDefault(INFLUX_URL_ENV_NAME, INFLUX_URL_DEFAULT)
}

func main() {
  log.Println("Starting monitoring-nginx-proxy-companion")

  proxyContainerId := FindProxyContainerId()

  log.Println("Found proxy container with id: ", proxyContainerId)

  AttachContainerLogListener(proxyContainerId, logCallback)
}