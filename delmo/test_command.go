package delmo

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/cli"
)

type TestCommand struct {
	Ui cli.Ui
}

func (t *TestCommand) Help() string {
	helpText := `
Usage: delmo test [options]

  Run a test :-)
`
	return strings.TrimSpace(helpText)
}

func (t *TestCommand) Run(args []string) int {
	flags := flag.FlagSet{
		Usage: func() { t.Help() },
	}

	var path string
	flags.StringVar(&path, "f", "delmo.yml", "")
	if err := flags.Parse(args); err != nil {
		t.Ui.Error(fmt.Sprintf("Error parsing arguments\n%s", err))
		return 2
	}

	config, err := LoadConfig(path)
	if err != nil {
		t.Ui.Error(fmt.Sprintf("Error reading configuration\n%s", err))
		return 2
	}

	_, err = t.prepareDockerHost(path, config.System)
	if err != nil {
		t.Ui.Error(fmt.Sprintf("Cloud not setup docker-machine\n%s", err))
		return 2
	}

	suite, err := NewSuite(config)
	if err != nil {
		t.Ui.Error(fmt.Sprintf("Could not initialize suite %s"))
		return 2
	}
	result := suite.Run(t.Ui)
	return result

}

func (t *TestCommand) Synopsis() string {
	return "Run some tests"
}

func (t *TestCommand) prepareDockerHost(path string, system SystemConfig) (string, error) {
	rawCmd, err := exec.LookPath("docker-machine")
	if err != nil {
		return "", err
	}
	hostDir := fmt.Sprintf(".delmo/%s", system.Name)

	args := []string{
		"ssh",
		system.MachineName,
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
		system.MachineName,
		"mkdir",
		"-p",
		hostDir,
	}
	cmd = exec.Command(rawCmd, args...)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Could not create dir %s\n%s", hostDir, err)
	}

	dir := filepath.Dir(path)
	files, err := ioutil.ReadDir(dir)
	for _, f := range files {
		file := filepath.Join(dir, f.Name())
		args = []string{
			"scp",
			"-r",
			file,
			fmt.Sprintf("%s:%s", system.MachineName, hostDir),
		}
		cmd = exec.Command(rawCmd, args...)
		err = cmd.Run()
		if err != nil {
			return "", fmt.Errorf("Could not upload file %s\n%s", f.Name(), err)
		}
	}

	return hostDir, nil
}
