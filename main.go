package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mitchellh/cli"
)

var Version = "(dev)"

func main() {
	os.Exit(Run(os.Args[1:]))
}

func Run(args []string) int {
	flags := flag.FlagSet{}

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	for _, arg := range args {
		if arg == "-v" || arg == "--version" || arg == "version" {
			ui.Output(fmt.Sprintf("delmo-v%s", Version))
			return 0
		}
	}

	var delmoFile, machine string
	flags.StringVar(&delmoFile, "f", "delmo.yml", "")
	flags.StringVar(&machine, "m", "default", "")
	if err := flags.Parse(args); err != nil {
		ui.Error(fmt.Sprintf("Error parsing arguments\n%s", err))
		return 2
	}

	hostIp, err := setupDockerMachine(machine)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting up environment\n%s", err))
		return 2
	}

	config, err := LoadConfig(delmoFile)
	if err != nil {
		ui.Error(fmt.Sprintf("Error reading configuration\n%s", err))
		return 2
	}

	globalTaskEnvironment := []string{fmt.Sprintf("DOCKER_HOST_IP=%s", hostIp)}
	os.Setenv("DOCKER_HOST_IP", hostIp)
	suite, err := NewSuite(config, globalTaskEnvironment)
	if err != nil {
		ui.Error(fmt.Sprintf("Could not initialize suite %s"))
		return 2
	}
	result := suite.Run(ui)
	return result
}

func Usage() string {
	helpText := `
Usage: delmo test [options]

  Run a test :-)
`
	return strings.TrimSpace(helpText)
}

func setupDockerMachine(machineName string) (string, error) {
	// get environment variables
	cmd := exec.Command("docker-machine", "env", machineName, "--shell", "sh")
	output, err := cmd.Output()
	if err != nil {
		return "", err
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
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
