package delmo

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

type Suite struct {
	options               CLIOptions
	config                *Config
	globalTaskEnvironment TaskEnvironment
	tests                 []TestConfig
}

func NewSuite(options CLIOptions, config *Config, globalTaskEnvironment TaskEnvironment) (*Suite, error) {
	suite := &Suite{
		options:               options,
		config:                config,
		globalTaskEnvironment: globalTaskEnvironment,
	}
	tests, err := suite.testsToRun(options.Tests, config.Tests)
	if err != nil {
		return nil, err
	}
	suite.tests = tests
	return suite, nil
}

func (s *Suite) Run() int {
	fmt.Printf("Running Test Suite for System %s\n", s.config.Suite.Name)
	executing := "Tests to execute: "
	for _, t := range s.tests {
		executing += t.Name + ", "
	}
	fmt.Printf(executing + "\n")

	err := s.initializeSystem()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 1
	}

	failed := []*TestReport{}
	succeeded := []*TestReport{}

	var wg sync.WaitGroup
	for _, test := range s.tests {
		wg.Add(1)
		fmt.Printf("Running test '%s'...\n", test.Name)
		runner := NewTestRunner(test, s.config.Tasks, s.globalTaskEnvironment)
		runtime, err := NewDockerCompose(s.config.Suite.System, test.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating runtime for '%s'\n%s\n", test.Name, err)
			continue
		}

		go func() {
			defer wg.Done()

			// Capture the name because the test variable will change
			testName := test.Name

			var output TestOutput
			var outputBytes bytes.Buffer
			if s.options.ParallelExecution {
				output.Stdout = &outputBytes
				output.Stderr = &outputBytes
			} else {
				output = TestOutput{
					Stdout: os.Stdout,
					Stderr: os.Stderr,
				}
			}

			report := runner.RunTest(runtime, output)
			if report.Success {
				succeeded = append(succeeded, report)
				fmt.Printf("Test '%s' completed sucessfully!\n", testName)
			} else {
				failed = append(failed, report)
				fmt.Printf("Test '%s' Failed!\n%s\n", testName, report.Error)
				if s.options.ParallelExecution {
					fmt.Printf("Output from test '%s'\n%s\n", testName, outputBytes)
				}
			}
		}()

		if s.options.ParallelExecution == false {
			wg.Wait()
			fmt.Println("")
		}
	}

	if s.options.ParallelExecution == true {
		wg.Wait()
		fmt.Println("")
	}

	outputSummary(failed, succeeded)
	if len(failed) != 0 {
		return 1
	}
	return 0
}

func outputSummary(failed []*TestReport, succeeded []*TestReport) {
	fmt.Printf("SUMMARY:\n%d tests succeeded\n%d tests failed\n",
		len(succeeded),
		len(failed))
}

func (s *Suite) initializeSystem() error {
	dc, err := NewDockerCompose(s.config.Suite.System, "")
	if err != nil {
		return fmt.Errorf("Could not initialize docker-compose\n%s", err)
	}

	if s.options.SkipPull != true {
		fmt.Printf("\nPulling images for system %s...\n", s.config.Suite.Name)
		err = dc.Pull()
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("Error pulling images\n%s\n", err))
		}
	}

	fmt.Printf("\nBuiling images for system %s...\n", s.config.Suite.Name)
	if s.options.OnlyBuildTask {
		err = dc.Build(s.config.Suite.TaskService)
	} else {
		err = dc.Build()
	}
	if err != nil {
		return fmt.Errorf("Could not build system\n%s", err)
	}

	fmt.Println("")

	return nil
}

func (s *Suite) testsToRun(testNames []string, allTests []TestConfig) ([]TestConfig, error) {
	if len(testNames) == 0 {
		return allTests, nil
	}
	var tests []TestConfig
	for _, n := range testNames {
		found := false
		for _, t := range allTests {
			if n == t.Name {
				tests = append(tests, t)
				found = true
				break
			}
		}
		if found == false {
			return nil, fmt.Errorf("Couldn't find test named: %s", n)
		}
	}
	return tests, nil
}
