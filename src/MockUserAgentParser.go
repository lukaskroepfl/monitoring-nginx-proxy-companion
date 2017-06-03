package main

type MockUserAgentParser struct {}

func (mockUserAgentParser MockUserAgentParser) Parse(userAgent string) UserAgent {
  return UserAgent{

  }
}