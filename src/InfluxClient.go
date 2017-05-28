package main

import (
  "github.com/influxdata/influxdb/client/v2"
  "log"
  "time"
  "strconv"
)

func WriteToInflux(parsedLogLine ParsedLogLine) {
  dbClient, err := client.NewHTTPClient(client.HTTPConfig{
    Addr: getInfluxUrl(),
  })

  if err != nil {
    log.Fatal(err)
  }

  batchPoints, err := client.NewBatchPoints(client.BatchPointsConfig{
    Database:  "monitoring",
  })
  if err != nil {
    log.Fatal(err)
  }

  tags := map[string]string {
    "container_name": parsedLogLine.containerName,
    "host": parsedLogLine.host,
    "request_method": parsedLogLine.requestMethod,
    "http_version": parsedLogLine.httpVersion,
    "http_status": strconv.Itoa(parsedLogLine.httpStatus),
  }

  fields := map[string]interface{} {
    "source_ip": parsedLogLine.sourceIp,
    "request_path": parsedLogLine.requestPath,
    "body_bytes_sent": parsedLogLine.bodyBytesSent,
    "http_referer": parsedLogLine.httpReferer,
    "user_agent": parsedLogLine.userAgent,
  }

  point, err := client.NewPoint("http_status", tags, fields, time.Now())
  if err != nil {
    log.Fatal(err)
  }

  batchPoints.AddPoint(point)

  if err := dbClient.Write(batchPoints); err != nil {
    log.Println("Could not insert into influx .")
    return
  }
}