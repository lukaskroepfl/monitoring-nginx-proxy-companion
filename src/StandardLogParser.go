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
  layout := "02/Jan/2006:15:04:05 +0000"
  t, err := time.Parse(layout, dateString)

  if err != nil {
    log.Fatal("Could not parse date string, reason: ", err)
  }

  return t
}