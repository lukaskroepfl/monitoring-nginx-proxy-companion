package main

import (
  "github.com/fsouza/go-dockerclient"
  "time"
  "io"
  "bufio"
)

const DOCKER_DAEMON_SOCKET = "unix:///var/run/docker.sock"

func listenToPipes(stdout, stderr io.Reader)  {
  listenToPipe := func(input io.Reader) {
    buf := bufio.NewReader(input)

    for {
      line, _ := buf.ReadString('\n')
      println(line)
    }
  }
  go listenToPipe(stdout)
  go listenToPipe(stderr)
}

func getLogsOfContainer(containerId string) {
  client, err := docker.NewClient(DOCKER_DAEMON_SOCKET)
  if err != nil {
    panic(err)
  }

  sinceTime := time.Now()

  stdoutReader, stdoutWriter := io.Pipe()
  stderrReader, stderrWriter := io.Pipe()

  listenToPipes(stdoutReader, stderrReader)

  for {
    dockerLogErr := client.Logs(docker.LogsOptions{
      Container:         containerId,
      OutputStream:      stdoutWriter,
      ErrorStream:       stderrWriter,
      Stdout:            true,
      Stderr:            true,
      Follow:            true,
      Tail:              "all",
      Since:             sinceTime.Unix(),
      InactivityTimeout: 0,
    })

    sinceTime = time.Now()

    if dockerLogErr != nil {
      panic(dockerLogErr)
    }
  }
}
