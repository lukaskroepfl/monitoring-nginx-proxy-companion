package main

import (
  "github.com/influxdata/influxdb/client/v2"
  "strconv"
  "log"
  "time"
)

type InfluxdbLogPersistor struct {
  influxClient client.Client
}

func (influxdbLogPersistor *InfluxdbLogPersistor) Setup() {
  dbClient, err := client.NewHTTPClient(client.HTTPConfig{
    Addr: getInfluxUrl(),
  })

  if err != nil {
    log.Fatal("Could not setup influx client, reason: ", err)
  }

  influxdbLogPersistor.influxClient = dbClient
}

func (influxLogPersistor InfluxdbLogPersistor) Persist(parsedLogLine ParsedLogLine) {
  batchPoints, err := client.NewBatchPoints(client.BatchPointsConfig{
    Database: getInfluxDbName(),
  })
  if err != nil {
    log.Fatal(err)
  }

  tags := map[string]string{
    "host": parsedLogLine.host,
    "request_method": parsedLogLine.requestMethod,
    "http_version": parsedLogLine.httpVersion,
    "http_status": strconv.Itoa(parsedLogLine.httpStatus),
    "browser": parsedLogLine.browser,
    "browser_version": parsedLogLine.browserVersion,
    "os": parsedLogLine.os,
    "mobile": strconv.FormatBool(parsedLogLine.mobile),
    "country": parsedLogLine.country,
    "city": parsedLogLine.city,
  }

  fields := map[string]interface{}{
    "source_ip": parsedLogLine.sourceIp,
    "request_path": parsedLogLine.requestPath,
    "body_bytes_sent": parsedLogLine.bodyBytesSent,
    "http_referer": parsedLogLine.httpReferer,
    "user_agent": parsedLogLine.userAgent,
    "latency": parsedLogLine.latency,
  }

  point, err := client.NewPoint("http_status", tags, fields, time.Now())
  if err != nil {
    log.Fatal(err)
  }

  batchPoints.AddPoint(point)

  if err := influxLogPersistor.influxClient.Write(batchPoints); err != nil {
    log.Println("Could not insert into influx .")
    return
  }
}

func getInfluxUrl() string {
  return GetEnvOrPanic(INFLUX_URL_ENV_NAME)
}

func getInfluxDbName() string {
  return GetEnvOrPanic(INFLUX_DB_ENV_NAME)
}