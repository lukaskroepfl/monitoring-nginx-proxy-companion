package main

import (
  "regexp"
  "strconv"
  "errors"
)

type StandardLogParser struct {
}

const LOG_LINE_REGEX = `\s*(\S+)\s+(\S+).+\[(.+)\]\s+"([^"]+)"\s+(\S+)\s+(\S+)\s+"([^"]+)"\s+"([^"]+)"`

func (standardLogParser StandardLogParser) Parse(logLine string) (HttpRequest, error) {
  var logLineParserRegex = regexp.MustCompile(LOG_LINE_REGEX)

  logLineParserRegexResult := logLineParserRegex.FindStringSubmatch(logLine)
  if len(logLineParserRegexResult) <= 0 {
    return HttpRequest{}, errors.New("Log line did not match nginx log line.")
  }

  regexFieldIndex := 1
  host := logLineParserRegexResult[regexFieldIndex]
  regexFieldIndex++
  remoteAddress := logLineParserRegexResult[regexFieldIndex]
  regexFieldIndex++
  timestamp := logLineParserRegexResult[regexFieldIndex]

  regexFieldIndex++
  httpRequest := logLineParserRegexResult[regexFieldIndex]
  var httpRequestRegex = regexp.MustCompile(`^(\S+)\s+(\S+)\s+(\S+)`)
  httpRequestRegexResult := httpRequestRegex.FindStringSubmatch(httpRequest)
  requestType := httpRequestRegexResult[1]
  requestPath := httpRequestRegexResult[2]
  httpVersion := httpRequestRegexResult[3]

  regexFieldIndex++
  httpStatus, _ := strconv.Atoi(logLineParserRegexResult[regexFieldIndex])
  regexFieldIndex++
  bodyBytesSent, _ := strconv.Atoi(logLineParserRegexResult[regexFieldIndex])
  regexFieldIndex++
  httpReferer := logLineParserRegexResult[regexFieldIndex]
  regexFieldIndex++
  userAgent := logLineParserRegexResult[regexFieldIndex]

  parsedLogLine := HttpRequest{}
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

  parseUserAgentAndSetFields(userAgent, &parsedLogLine)
  lookupIpAndSetFields(remoteAddress, &parsedLogLine)

  return parsedLogLine, nil
}

func parseUserAgentAndSetFields(userAgentString string, parsedLogLine *HttpRequest) {
  mssolaUserAgentParser := MssolaUserAgentParser{}
  userAgent := mssolaUserAgentParser.Parse(userAgentString)

  parsedLogLine.browser = userAgent.browser
  parsedLogLine.browserVersion = userAgent.browserVersion
  parsedLogLine.os = userAgent.os
  parsedLogLine.mobile = userAgent.mobile
}

func lookupIpAndSetFields(ip string, parsedLogLine *HttpRequest) {
  geoIp2IPLookupService := GeoIp2IPLookupService{}

  ipLocation := geoIp2IPLookupService.Lookup(ip)

  parsedLogLine.country = ipLocation.country
  parsedLogLine.city = ipLocation.city
}