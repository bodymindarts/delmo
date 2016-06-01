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

func (s *Suite) Run(ui cli.Ui) (int, error) {
	ui.Info(fmt.Sprintf("Running Test Suite for System %s", s.config.System.Name))

	for _, test := range s.config.Tests {
		runner := NewTestRunner(test)
		runtime, err := s.system.NewRuntime(test.Name)
		if err != nil {
			ui.Error(fmt.Sprintf("Error creating runtime! %s", err))
			continue
		}
		err = runner.RunTest(runtime)
		if err != nil {
			ui.Error(fmt.Sprintf("Test %s failed! %s", test.Name, err))
		}
		outputTest(ui, test.Name, runner)
		runner.Cleanup()
	}

	return 0, nil
}

func outputTest(ui cli.Ui, name string, runner *TestRunner) {
	runnerOut, _ := runner.Output()
	ui.Output(fmt.Sprintf("Output for %s:\n%s", name, runnerOut))
}
