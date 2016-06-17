package delmo

import (
	"fmt"
	"io"
)

type TestOutput struct {
	Stdout io.Writer
	Stderr io.Writer
}

type TestRunner struct {
	config                TestConfig
	tasks                 []TaskConfig
	globalTaskEnvironment TaskEnvironment
	runtime               Runtime
	beforeSteps           []Step
	steps                 []Step
}

func NewTestRunner(config TestConfig, tasks Tasks, globalTaskEnvironment TaskEnvironment) *TestRunner {
	beforeSteps := []Step{}
	for _, taskName := range config.BeforeStartup {
		beforeSteps = append(beforeSteps, NewExecStep(tasks[taskName], globalTaskEnvironment))
	}

	return &TestRunner{
		config:                config,
		globalTaskEnvironment: globalTaskEnvironment,
		beforeSteps:           beforeSteps,
		steps:                 initSteps(config.Spec, tasks, globalTaskEnvironment),
	}
}

func (tr *TestRunner) RunTest(runtime Runtime, out TestOutput) *TestReport {
	tr.runtime = runtime
	report := NewTestReport(runtime.SystemOutput)

	tr.runtime.Cleanup()
	for _, step := range tr.beforeSteps {
		fmt.Fprintf(out.Stdout, "Executing - %s\n", step.Description())
		err := step.Execute(runtime, out)
		if err != nil {
			fmt.Fprintf(out.Stderr, "FAIL! Step - %s did not complete as expected.\nREASON - %s\n", step.Description(), err)
			report.Fail(err)
			return report
		}
	}

	fmt.Fprintf(out.Stdout, "Starting '%s' Runtime...\n", tr.config.Name)
	err := runtime.StartAll(out)
	if err != nil {
		fmt.Fprintf(out.Stderr, "Could not start runtime for %s! %s\n", tr.config.Name, err)
		report.Fail(err)
		return report
	}
	defer func() {
		if report.Success {
			runtime.Cleanup()
		}
	}()

	for _, step := range tr.steps {
		fmt.Fprintf(out.Stdout, "Executing - %s\n", step.Description())
		err = step.Execute(runtime, out)
		if err != nil {
			fmt.Fprintf(out.Stderr, "FAIL! Step - %s did not complete as expected.\nREASON - %s\n", step.Description(), err)
			report.Fail(err)
			break
		}
	}

	fmt.Fprintf(out.Stdout, "Stopping '%s' Runtime...\n", tr.config.Name)
	err = runtime.StopAll(out)
	if err != nil {
		fmt.Fprintf(out.Stderr, "Could not stop runtime for %s! %s\n", tr.config.Name, err)
		report.Fail(err)
		return report
	}
	return report
}

func initSteps(stepConfigs []StepConfig, tasks Tasks, env TaskEnvironment) []Step {
	steps := []Step{}
	for _, stepConfig := range stepConfigs {
		if len(stepConfig.Stop) != 0 {
			steps = append(steps, NewStopStep(stepConfig))
		}
		if len(stepConfig.Destroy) != 0 {
			steps = append(steps, NewDestroyStep(stepConfig))
		}
		if len(stepConfig.Start) != 0 {
			steps = append(steps, NewStartStep(stepConfig))
		}
		if len(stepConfig.Wait) != 0 {
			for _, taskName := range stepConfig.Wait {
				steps = append(steps, NewWaitStep(tasks[taskName], env))
			}
		}
		if len(stepConfig.Exec) != 0 {
			for _, taskName := range stepConfig.Exec {
				steps = append(steps, NewExecStep(tasks[taskName], env))
			}
		}
		if len(stepConfig.Assert) != 0 {
			for _, taskName := range stepConfig.Assert {
				steps = append(steps, NewAssertStep(tasks[taskName], env))
			}
		}
		if len(stepConfig.Fail) != 0 {
			for _, taskName := range stepConfig.Fail {
				steps = append(steps, NewFailStep(tasks[taskName], env))
			}
		}
	}
	return steps
}
