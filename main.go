package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"os"
)

func main() {
	os.Exit(Run(os.Args[1:]))
}

func Run(args []string) int {
	// Get the command line args. We shortcut "--version" and "-v" to
	// just show the version.
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	commands := map[string]cli.CommandFactory{
		"test": func() (cli.Command, error) {
			return &TestCommand{
				Ui: ui,
			}, nil
		},
	}

	cli := &cli.CLI{
		Args:     args,
		Commands: commands,
		HelpFunc: cli.BasicHelpFunc("delmo"),
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
