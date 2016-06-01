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
		return 2
	}

	config, err := LoadConfig(path)
	if err != nil {
		t.Ui.Error(fmt.Sprintf("Error reading configuration\n%s", err))
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
