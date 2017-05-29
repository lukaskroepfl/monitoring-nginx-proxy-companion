package main

type IPLocation struct {
  country string
  city string
}

type IIPLookup interface {
  Lookup(ip string)
}