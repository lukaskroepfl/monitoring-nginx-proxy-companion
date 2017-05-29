package main

import "github.com/mssola/user_agent"

type MssolaUserAgentParser struct {
}

func (MssolaUserAgentParser) Parse(userAgentString string) UserAgent {
  ua := user_agent.New(userAgentString);

  mobile := ua.Mobile()
  browser, browserVersion := ua.Browser()
  os := ua.OS()

  userAgent := UserAgent{
    mobile: mobile,
    browser: browser,
    browserVersion: browserVersion,
    os: os,
  }

  return userAgent
}
