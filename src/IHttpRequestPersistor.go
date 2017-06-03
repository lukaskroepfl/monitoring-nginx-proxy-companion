package main

type IHttpRequestPersistor interface {
  Persist(httpRequest HttpRequest)
}
