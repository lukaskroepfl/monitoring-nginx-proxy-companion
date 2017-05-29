package main

type UserAgent struct {
  browser        string
  browserVersion string
  os             string
  mobile         bool
}

type IUserAgentParser interface {
  Parse(userAgent string) UserAgent
}
