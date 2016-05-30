package delmo

import (
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
	t.Ui.Output("Hellohotesnht")
	return 0
}

func (t *TestCommand) Synopsis() string {
	return "Run some tests"
}
