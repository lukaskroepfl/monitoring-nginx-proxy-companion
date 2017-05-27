package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestParsesContainerName(t *testing.T) {
  line := "nginx                                | kroepfl.io 54.89.89.135 - - [27/May/2017:13:00:57 +0000] \"GET / HTTP/1.1\" 200 2211 \"-\" \"Java/1.8.0_121\""

  parsedLogLine := ParseLogLine(line)

  assert.Equal(t, "nginx", parsedLogLine.containerName)
  assert.Equal(t, "kroepfl.io", parsedLogLine.host)
  assert.Equal(t, "54.89.89.135", parsedLogLine.sourceIp)
  assert.Equal(t, "27/May/2017:13:00:57 +0000", parsedLogLine.timestamp)
}