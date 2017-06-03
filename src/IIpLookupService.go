package main

type IPLocation struct {
  country string
  city    string
}

type IIpLookupService interface {
  Lookup(ip string) IPLocation
}