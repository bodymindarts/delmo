package delmo

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type Suite struct {
	config *SuiteConfig
	system *System
}

func NewSuite(config *SuiteConfig) *Suite {
	return &Suite{
		config: config,
		system: NewSystem(config.System),
	}
}

func (s *Suite) Run(ui cli.Ui) int {
	ui.Info(fmt.Sprintf("Running Test Suite for System %s", s.config.System.Name))

	failed := []*TestReport{}
	succeeded := []*TestReport{}

	for _, test := range s.config.Tests {
		runner := NewTestRunner(test)
		runtime, err := s.system.NewRuntime(test.Name)
		if err != nil {
			ui.Error(fmt.Sprintf("Error creating runtime! %s", err))
			continue
		}

		ui.Info(fmt.Sprintf("Running test %s", test.Name))
		report := runner.RunTest(runtime, ui)
		if report.Success {
			succeeded = append(succeeded, report)
			ui.Info(fmt.Sprintf("Test %s Succeeded!", test.Name))
		} else {
			failed = append(succeeded, report)
			ui.Info(fmt.Sprintf("Test %s Failed!\nRuntime Output:\n%s", test.Name, report.Output()))
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
		fmt.Sprintf("SUMMARY:\n%d tests succeeded\n%d tests failed",
			len(succeeded),
			len(failed)))
}
