package main

import (
  "log"
)

const PROXY_CONTAINER_NAME_ENV_NAME = "PROXY_CONTAINER_NAME"
const PROXY_CONTAINER_NAME_DEFAULT = "nginx"

func getProxyContainerName() string {
  return GetEnvOrDefault(PROXY_CONTAINER_NAME_ENV_NAME, PROXY_CONTAINER_NAME_DEFAULT)
}

func getInfluxUrl() string {
  return GetEnvOrDefault(INFLUX_URL_ENV_NAME, INFLUX_URL_DEFAULT)
}

func main() {
  log.Println("Starting monitoring-nginx-proxy-companion")

  log.Println("Setting up influx client.")
  influxdbLogPersistor := InfluxdbLogPersistor{}
  influxdbLogPersistor.Setup()

  log.Println("Creating docker container log miner.")
  dockerContainerLogMiner := DockerContainerLogMiner{}
  dockerContainerLogMiner.SetLogPersistor(influxdbLogPersistor)
  logParser := StandardLogParser{}
  dockerContainerLogMiner.SetLogParser(logParser)

  log.Println("Start log mining.")
  dockerContainerLogMiner.Mine()
}