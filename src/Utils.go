package main

import (
  "os"
  "fmt"
)

func GetEnvOrPanic(envName string) string {
  envValue := os.Getenv(envName)
  if envValue == "" {
    panic(fmt.Sprintf("Env %s not set", envName));
  }

  return envValue
}