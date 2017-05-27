package main

import (
  "strings"
)

type ParsedLogLine struct {
  containerName string
  host string
  sourceIp string
  timestamp string
  requestType string
  requestPath string
  responseCode int
  latency int
  userAgent string
}

const CONTAINER_SPLIT = "|"

func ParseLogLine(line string) ParsedLogLine {

  containerSplit := strings.Split(line, CONTAINER_SPLIT)

  containerNameWithSpaces := containerSplit[0]
  containerName := strings.TrimRight(containerNameWithSpaces, " ")


  logRest := containerSplit[1]

  logRestParts := strings.Split(logRest, " ")

  hostName := logRestParts[1]
  sourceIp := logRestParts[2]

  timeStart := strings.Index(logRest, "[") + 1
  timeEnd := strings.Index(logRest, "]")

  time := logRest[timeStart:timeEnd]

  parsedLogLine := ParsedLogLine{}
  parsedLogLine.containerName = containerName
  parsedLogLine.host = hostName
  parsedLogLine.sourceIp = sourceIp
  parsedLogLine.timestamp = time

  return parsedLogLine
}
