package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/cli"
)

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

	var delmoFile, machine string
	flags.StringVar(&delmoFile, "f", "delmo.yml", "")
	flags.StringVar(&machine, "m", "", "")
	if err := flags.Parse(args); err != nil {
		ui.Error(fmt.Sprintf("Error parsing arguments\n%s", err))
		return 2
	}

	config, err := LoadConfig(delmoFile)
	if err != nil {
		ui.Error(fmt.Sprintf("Error reading configuration\n%s", err))
		return 2
	}

	hostDir, err := prepareDockerHost(ui, delmoFile, machine, config.Suite.Name)
	if err != nil {
		ui.Error(fmt.Sprintf("Cloud not setup docker-machine\n%s", err))
		return 2
	}

	context := GlobalContext{
		DockerHostSyncDir: hostDir,
	}
	suite, err := NewSuite(config, context)
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

func prepareDockerHost(ui cli.Ui, delmoFile, machine, suiteName string) (string, error) {
	if machine == "" {
		absPath, err := filepath.Abs(delmoFile)
		if err != nil {
			return "", err
		}
		delmoDir := filepath.Dir(absPath)
		return delmoDir, nil
	}

	rawCmd, err := exec.LookPath("docker-machine")
	if err != nil {
		return "", err
	}
	hostDir := fmt.Sprintf(".delmo/%s", suiteName)

	ui.Info("Preparing host machine")
	args := []string{
		"ssh",
		machine,
		"rm",
		"-rf",
		hostDir,
	}
	cmd := exec.Command(rawCmd, args...)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Could not delete dir %s\n%s", hostDir, err)
	}

	args = []string{
		"ssh",
		machine,
		"mkdir",
		"-p",
		hostDir,
	}
	cmd = exec.Command(rawCmd, args...)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Could not create dir %s\n%s", hostDir, err)
	}

	ui.Info("Uploading files")
	dir := filepath.Dir(delmoFile)
	files, err := ioutil.ReadDir(dir)
	for _, f := range files {
		file := filepath.Join(dir, f.Name())
		ui.Info(fmt.Sprintf("file: %s", file))
		args = []string{
			"scp",
			"-r",
			file,
			fmt.Sprintf("%s:%s", machine, hostDir),
		}
		cmd = exec.Command(rawCmd, args...)
		out, err := cmd.Output()
		if err != nil {
			return "", fmt.Errorf("Could not upload file %s\n%s\n%s", f.Name(), out, err)
		}
	}

	args = []string{
		"ssh",
		machine,
		"pwd",
	}
	cmd = exec.Command(rawCmd, args...)
	hostWD, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("Could not determin home dir on host\n%s", hostDir, err)
	}

	return filepath.Join(strings.TrimSpace(string(hostWD)), hostDir), nil
}
