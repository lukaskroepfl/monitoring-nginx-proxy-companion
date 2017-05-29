package main

type ILogPersistor interface {
  Persist(parsedLogLine ParsedLogLine)
}
