package main

import (
  "github.com/oschwald/geoip2-golang"
  "log"
  "net"
)

type GeoIp2IPLookupService struct {
}

func (geoIp2IPLookupService GeoIp2IPLookupService) Lookup(ip string) IPLocation {
  db, err := geoip2.Open("GeoLite2-City.mmdb")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  parsedIp := net.ParseIP("81.2.69.142")

  record, err := db.City(parsedIp)
  if err != nil {
    log.Fatal(err)
  }

  ipLocation := IPLocation{}

  ipLocation.country = record.Country.Names["en"]
  ipLocation.city = record.City.Names["en"]

  return ipLocation
}