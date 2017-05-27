package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestParsesContainerName(t *testing.T) {
  line := `nginx                                | blog.kroepfl.io 193.80.91.32 - - [27/May/2017:19:26:27 +0000] "GET /wp-content/uploads/2017/04/Untitled.png HTTP/1.1" 404 18000 "https://blog.kroepfl.io/" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"`

  parsedLogLine := ParseProxyLogLine(line)

  assert.Equal(t, "nginx", parsedLogLine.containerName)
  assert.Equal(t, "blog.kroepfl.io", parsedLogLine.host)
  assert.Equal(t, "193.80.91.32", parsedLogLine.sourceIp)
  assert.Equal(t, "27/May/2017:19:26:27 +0000", parsedLogLine.timestamp)
  assert.Equal(t, "GET", parsedLogLine.requestMethod)
  assert.Equal(t, "/wp-content/uploads/2017/04/Untitled.png", parsedLogLine.requestPath)
  assert.Equal(t, "HTTP/1.1", parsedLogLine.httpVersion)
  assert.Equal(t, 404, parsedLogLine.httpStatus)
  assert.Equal(t, 18000, parsedLogLine.bodyBytesSent)
  assert.Equal(t, "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36", parsedLogLine.userAgent)
}