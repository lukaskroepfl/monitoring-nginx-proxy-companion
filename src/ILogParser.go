package main

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

type ILogParser interface {
  Parse(logLine string) (ParsedLogLine, error)
}