package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/bodymindarts/delmo/delmo"
)

var Version = "(dev)"

func main() {
	os.Exit(Run(os.Args[1:]))
}

func Run(args []string) int {
	for _, arg := range args {
		if arg == "-v" || arg == "--version" || arg == "version" {
			fmt.Printf("delmo-v%s", Version)
			return 0
		}
	}

	options := delmo.ParseOptions(args)

	hostIp, err := setupDockerMachine(options.DockerMachine)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting up environment\n%s", err)
		return 2
	}

	config, err := delmo.LoadConfig(options.DelmoFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading configuration\n%s", err)
		return 2
	}

	globalTaskEnvironment := []string{fmt.Sprintf("DOCKER_HOST_IP=%s", hostIp)}
	os.Setenv("DOCKER_HOST_IP", hostIp)
	suite, err := delmo.NewSuite(options, config, globalTaskEnvironment)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not initialize suite %s")
		return 2
	}
	result := suite.Run()
	return result
}

func setupDockerMachine(machineName string) (string, error) {
	// get environment variables
	cmd := exec.Command("docker-machine", "env", machineName, "--shell", "sh")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("Couldn't export environment from docker-machine\n%s", err)
	}
	// set each variable of form: export DOCKER_HOST="tcp://192.168.99.100:2376"
	for _, l := range strings.Split(string(output), "\n") {
		if strings.HasPrefix(l, "export ") {
			assignment := strings.Split(strings.TrimPrefix(l, "export "), "=")
			key := assignment[0]
			value := strings.Replace(assignment[1], "\"", "", -1)
			os.Setenv(key, value)
		}
	}

	cmd = exec.Command("docker-machine", "ip", machineName)
	output, err = cmd.Output()
	if err != nil {
		return "", fmt.Errorf("Couldn't read docker host ip from docker-machine\n%s", err)
	}

	return strings.TrimSpace(string(output)), nil
}
