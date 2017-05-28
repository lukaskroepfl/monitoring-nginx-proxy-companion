package main

import (
  "github.com/fsouza/go-dockerclient"
  "time"
  "io"
  "bufio"
)

const DOCKER_DAEMON_SOCKET = "unix:///var/run/docker.sock"

func listenToPipes(stdout, stderr io.Reader, logCallback func(logLine string))  {
  listenToPipe := func(input io.Reader) {
    buf := bufio.NewReader(input)

    for {
      line, _ := buf.ReadString('\n')
      logCallback(line)
    }
  }
  go listenToPipe(stdout)
  go listenToPipe(stderr)
}

func AttachContainerLogListener(containerId string, logCallback func(logLine string)) {
  client, err := docker.NewClient(DOCKER_DAEMON_SOCKET)
  if err != nil {
    panic(err)
  }

  sinceTime := time.Now()

  stdoutReader, stdoutWriter := io.Pipe()
  stderrReader, stderrWriter := io.Pipe()

  listenToPipes(stdoutReader, stderrReader, logCallback)

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

func FindProxyContainerId() string {
  client, err := docker.NewClient(DOCKER_DAEMON_SOCKET)
  if err != nil {
    panic(err)
  }

  proxyContainerName := getProxyContainerName()

  filters := make(map[string][]string)
  filters["name"] = append(filters["name"], proxyContainerName)

  containers, err := client.ListContainers(docker.ListContainersOptions{Filters: filters})
  if err != nil {
    panic(err)
  }

  if len(containers) <= 0 {
    panic("No running container found with specified name.")
  }

  proxyContainer := containers[0]
  proxyContainerId := proxyContainer.ID

  return proxyContainerId
}
