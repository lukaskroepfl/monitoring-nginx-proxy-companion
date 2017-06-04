package main

import (
  "regexp"
  "strconv"
  "errors"
  "time"
  "log"
)

type StandardLogParser struct {
  userAgentParser IUserAgentParser
  ipLookupService IIpLookupService
}

const LOG_LINE_REGEX = `\s*(\S+)\s+(\S+).+\[(.+)\]\s+"([^"]+)"\s+(\S+)\s+(\S+)\s+"([^"]+)"\s+"([^"]+)"($|\s+(\S+))`

const STANDARD_LOG_LINE_DATE_FORMAT = "02/Jan/2006:15:04:05 +0000"

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
  httpRequestHeader := logLineParserRegexResult[regexFieldIndex]
  var httpRequestRegex = regexp.MustCompile(`^(\S+)\s+(\S+)\s+(\S+)`)
  httpRequestRegexResult := httpRequestRegex.FindStringSubmatch(httpRequestHeader)
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

  regexFieldIndex += 2;
  latency := logLineParserRegexResult[regexFieldIndex]

  var latencyFloat float64
  if latency != "" {
    var err error
    latencyFloat, err = strconv.ParseFloat(latency, 64)
    if err != nil {
      log.Fatal("Could not parse float, reason: ", err)
    }
  }

  httpRequest := HttpRequest{}
  httpRequest.host = host
  httpRequest.sourceIp = remoteAddress
  httpRequest.timestamp = convertDateStringToTime(timestamp)
  httpRequest.requestMethod = requestType
  httpRequest.requestPath = requestPath
  httpRequest.httpVersion = httpVersion
  httpRequest.httpStatus = httpStatus
  httpRequest.bodyBytesSent = bodyBytesSent
  httpRequest.httpReferer = httpReferer
  httpRequest.userAgent = userAgent
  httpRequest.latency = latencyFloat

  parseUserAgentAndSetFields(standardLogParser.userAgentParser, userAgent, &httpRequest)
  lookupIpAndSetFields(standardLogParser.ipLookupService, remoteAddress, &httpRequest)

  return httpRequest, nil
}

func parseUserAgentAndSetFields(userAgentParser IUserAgentParser, userAgentString string, httpRequest *HttpRequest) {
  userAgent := userAgentParser.Parse(userAgentString)

  httpRequest.browser = userAgent.browser
  httpRequest.browserVersion = userAgent.browserVersion
  httpRequest.os = userAgent.os
  httpRequest.mobile = userAgent.mobile
}

func lookupIpAndSetFields(ipLookupService IIpLookupService, ip string, httpRequest *HttpRequest) {
  ipLocation := ipLookupService.Lookup(ip)

  httpRequest.country = ipLocation.country
  httpRequest.city = ipLocation.city
}

func convertDateStringToTime(dateString string) time.Time {
  t, err := time.Parse(STANDARD_LOG_LINE_DATE_FORMAT, dateString)

  if err != nil {
    log.Fatal("Could not parse date string, reason: ", err)
  }

  return t
}