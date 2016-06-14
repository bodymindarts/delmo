package main

import (
	"fmt"
	"os"
)

type Suite struct {
	config                *Config
	globalTaskEnvironment TaskEnvironment
}

func NewSuite(config *Config, globalTaskEnvironment TaskEnvironment) (*Suite, error) {
	suite := &Suite{
		config:                config,
		globalTaskEnvironment: globalTaskEnvironment,
	}
	return suite, nil
}

func (s *Suite) Run() int {
	err := s.initializeSystem()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 1
	}

	fmt.Printf("\nRunning Test Suite for System %s\n", s.config.Suite.Name)

	failed := []*TestReport{}
	succeeded := []*TestReport{}
	output := TestOutput{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	for _, test := range s.config.Tests {
		runner := NewTestRunner(test, s.config.Tasks, s.globalTaskEnvironment)
		runtime, err := NewDockerCompose(s.config.Suite.System, test.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating runtime! %s\n", err)
			return 1
		}

		fmt.Printf("\nRunning test %s\n", test.Name)
		report := runner.RunTest(runtime, output)
		if report.Success {
			succeeded = append(succeeded, report)
			runner.Cleanup()
			fmt.Printf("Test %s Succeeded!\n", test.Name)
		} else {
			failed = append(succeeded, report)
			fmt.Printf("Test %s Failed!\nRuntime Output:\n%s\n", test.Name, report.SystemOutput())
		}
	}

	outputSummary(failed, succeeded)
	if len(failed) != 0 {
		return 1
	}
	return 0
}

func outputSummary(failed []*TestReport, succeeded []*TestReport) {
	fmt.Printf("\nSUMMARY:\n%d tests succeeded\n%d tests failed\n",
		len(succeeded),
		len(failed))
}

func (s *Suite) initializeSystem() error {
	dc, err := NewDockerCompose(s.config.Suite.System, "")
	if err != nil {
		return fmt.Errorf("Could not initialize docker-compose\n%s", err)
	}

	fmt.Printf("Pulling images for system %s\n", s.config.Suite.Name)
	err = dc.Pull()
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Error pulling images\n%s\n", err))
	}

	fmt.Printf("Builing images for system %s\n", s.config.Suite.Name)
	if s.config.Suite.OnlyBuildTask {
		err = dc.Build(s.config.Suite.TaskService)
	} else {
		err = dc.Build()
	}
	if err != nil {
		return fmt.Errorf("Could not build system\n%s", err)
	}

	return nil
}
