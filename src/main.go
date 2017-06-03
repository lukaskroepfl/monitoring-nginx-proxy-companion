package main

import (
  "log"
)

const PROXY_CONTAINER_NAME_ENV_NAME = "PROXY_CONTAINER_NAME"
const INFLUX_URL_ENV_NAME = "INFLUX_URL"
const INFLUX_DB_ENV_NAME = "INFLUX_DB_NAME"

func main() {
  log.Println("Starting monitoring-nginx-proxy-companion.")

  log.Println("Creating dependencies of log miner.")

  log.Println("Creating influx client.")
  influxdbHttpRequestPersistor := InfluxdbHttpRequestPersistor{}
  influxdbHttpRequestPersistor.Setup()

  log.Println("Creating user agent parser, ip lookup service and log parser.")
  mssolaUserAgentParser := MssolaUserAgentParser{}
  geoIp2IpLookupService := GeoIp2IpLookupService{}
  logParser := StandardLogParser{
    userAgentParser: mssolaUserAgentParser,
    ipLookupService: geoIp2IpLookupService,
  }

  log.Println("Creating docker container log miner.")
  dockerContainerLogMiner := DockerContainerLogMiner{
    logParser: logParser,
    httpRequestPersistor: influxdbHttpRequestPersistor,
  }

  log.Println("Start log mining.")
  dockerContainerLogMiner.Mine()
}