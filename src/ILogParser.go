package main

import "time"

type HttpRequest struct {
  host           string
  sourceIp       string
  timestamp      time.Time
  requestMethod  string
  requestPath    string
  httpVersion    string
  httpStatus     int
  bodyBytesSent  int
  httpReferer    string
  userAgent      string
  browser        string
  browserVersion string
  os             string
  mobile         bool
  country        string
  city           string
  latency        float64
}

type ILogParser interface {
  Parse(logLine string) (HttpRequest, error)
}