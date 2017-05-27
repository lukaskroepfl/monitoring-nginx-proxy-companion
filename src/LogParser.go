package main

import (
  "regexp"
  "strconv"
)

type ParsedLogLine struct {
  containerName string
  host          string
  sourceIp      string
  timestamp     string
  requestMethod string
  requestPath   string
  httpVersion   string
  httpStatus    int
  bodyBytesSent int
  httpReferer   string
  userAgent     string
}

func ParseLogLine(line string) ParsedLogLine {
  var logLineParserRegex = regexp.MustCompile(`^(\S+) *\|\s+(\S+)\s+(\S+).+\[(.+)\]\s+"([^"]+)"\s+(\S+)\s+(\S+)\s+"([^"]+)"\s+"([^"]+)"`)

  logLineParserRegexResult := logLineParserRegex.FindStringSubmatch(line)

  containerName := logLineParserRegexResult[1]
  host := logLineParserRegexResult[2]
  remoteAddress := logLineParserRegexResult[3]
  timestamp := logLineParserRegexResult[4]
  httpRequest := logLineParserRegexResult[5]
  var httpRequestRegex = regexp.MustCompile(`^(\S+)\s+(\S+)\s+(\S+)`)
  httpRequestRegexResult := httpRequestRegex.FindStringSubmatch(httpRequest)
  requestType := httpRequestRegexResult[1]
  requestPath := httpRequestRegexResult[2]
  httpVersion := httpRequestRegexResult[3]

  httpStatus, _ := strconv.Atoi(logLineParserRegexResult[6])
  bodyBytesSent, _ := strconv.Atoi(logLineParserRegexResult[7])
  httpReferer := logLineParserRegexResult[8]
  userAgent := logLineParserRegexResult[9]

  parsedLogLine := ParsedLogLine{}
  parsedLogLine.containerName = containerName
  parsedLogLine.host = host
  parsedLogLine.sourceIp = remoteAddress
  parsedLogLine.timestamp = timestamp
  parsedLogLine.requestMethod = requestType
  parsedLogLine.requestPath = requestPath
  parsedLogLine.httpVersion = httpVersion
  parsedLogLine.httpStatus = httpStatus
  parsedLogLine.bodyBytesSent = bodyBytesSent
  parsedLogLine.httpReferer = httpReferer
  parsedLogLine.userAgent = userAgent

  return parsedLogLine
}
