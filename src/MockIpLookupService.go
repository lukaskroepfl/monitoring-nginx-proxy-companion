package main

type MockIpLookupService struct {}

func (mockIpLookupService MockIpLookupService) Lookup(ip string) IPLocation {
  return IPLocation{

  }
}
