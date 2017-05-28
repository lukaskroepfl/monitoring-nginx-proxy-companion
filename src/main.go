package main

import "fmt"

func logCallback(logLine string) {
  parsedLogline, err := ParseProxyLogLine(logLine)
  if err != nil {
    panic(err)
  }

  fmt.Println(parsedLogline)
}

func main() {
  AttachContainerLogListener("9a7f981eed04", logCallback)
}