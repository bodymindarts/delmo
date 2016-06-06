package main

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type Suite struct {
	config *Config
}

func NewSuite(config *Config) (*Suite, error) {
	suite := &Suite{
		config: config,
	}
	return suite, nil
}

func (s *Suite) Run(ui cli.Ui) int {
	ui.Info(fmt.Sprintf("Running Test Suite for System %s", s.config.Suite.Name))

	failed := []*TestReport{}
	succeeded := []*TestReport{}

	for _, test := range s.config.Tests {
		runner := NewTestRunner(test)
		runtime, err := NewDockerCompose(s.config.Suite.System, test.Name)
		if err != nil {
			ui.Error(fmt.Sprintf("Error creating runtime! %s", err))
			return 1
		}

		ui.Info(fmt.Sprintf("Running test %s", test.Name))
		report := runner.RunTest(runtime, ui)
		if report.Success {
			succeeded = append(succeeded, report)
			runner.Cleanup()
			ui.Info(fmt.Sprintf("Test %s Succeeded!", test.Name))
		} else {
			failed = append(succeeded, report)
			ui.Info(fmt.Sprintf("Test %s Failed!\nRuntime Output:\n%s", test.Name, report.SystemOutput()))
		}
	}

	outputSummary(ui, failed, succeeded)
	if len(failed) != 0 {
		return 1
	}
	return 0
}

func outputSummary(ui cli.Ui, failed []*TestReport, succeeded []*TestReport) {
	ui.Output(
		fmt.Sprintf("\n\nSUMMARY:\n%d tests succeeded\n%d tests failed",
			len(succeeded),
			len(failed)))
}
