package delmo

import (
	"flag"
	"fmt"
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
		return 1
	}

	config, err := LoadConfig(path)
	if err != nil {
		t.Ui.Error(fmt.Sprintf("Error reading file %s\n%s", path, err))
		return 1
	}
	suite := NewSuite(config)
	result, _ := suite.Run(t.Ui)
	return result

}

func (t *TestCommand) Synopsis() string {
	return "Run some tests"
}
