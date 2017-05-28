package main

import "os"

func GetEnvOrDefault(envName string, envDefault string) string {
  envValue := os.Getenv(envName)
  if envValue == "" {
    return envDefault
  }

  return envValue
}
