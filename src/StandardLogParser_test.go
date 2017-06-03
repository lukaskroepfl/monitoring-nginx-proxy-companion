package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "time"
)

func TestParsesSomeLogLines(t *testing.T) {
  line := `blog.kroepfl.io 193.80.91.32 - - [27/May/2017:19:26:27 +0000] "GET /wp-content/uploads/2017/04/Untitled.png HTTP/1.1" 404 18000 "https://blog.kroepfl.io/" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"`

  userAgentParser := MssolaUserAgentParser{}
  mockIpLookupService := MockIpLookupService{}

  standardLogParser := StandardLogParser{
    userAgentParser: userAgentParser,
    ipLookupService: mockIpLookupService,
  }

  httpRequest, err := standardLogParser.Parse(line)

  if err != nil {
    t.Fail()
  }

  ti := time.Unix(1495913187, 0)

  assert.Equal(t, "blog.kroepfl.io", httpRequest.host)
  assert.Equal(t, "193.80.91.32", httpRequest.sourceIp)
  assert.Equal(t, ti.Unix(), httpRequest.timestamp.Unix())
  assert.Equal(t, "GET", httpRequest.requestMethod)
  assert.Equal(t, "/wp-content/uploads/2017/04/Untitled.png", httpRequest.requestPath)
  assert.Equal(t, "HTTP/1.1", httpRequest.httpVersion)
  assert.Equal(t, 404, httpRequest.httpStatus)
  assert.Equal(t, 18000, httpRequest.bodyBytesSent)
  assert.Equal(t, "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36", httpRequest.userAgent)
  assert.Equal(t, "Chrome", httpRequest.browser)
  assert.Equal(t, "58.0.3029.110", httpRequest.browserVersion)
  assert.Equal(t, "Linux x86_64", httpRequest.os)
  assert.Equal(t, false, httpRequest.mobile)
}