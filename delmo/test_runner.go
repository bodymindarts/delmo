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
	report := NewTestReport(tr.config.Name, runtime.SystemOutput)

	tr.runtime.Cleanup()
	for _, step := range tr.beforeSteps {
		fmt.Fprintf(out.Stdout, "Executing - %s", step.Description())
		err := step.Execute(runtime, out)
		if err != nil {
			fmt.Fprintf(out.Stderr, "FAIL! Step - %s did not complete as expected.\nREASON - %s", step.Description(), err)
			report.Fail(err)
			return report
		}
	}

	fmt.Fprintf(out.Stdout, "Starting %s Runtime", tr.config.Name)
	err := runtime.StartAll()
	if err != nil {
		fmt.Fprintf(out.Stderr, "Could not start runtime for %s! %s", tr.config.Name, err)
		report.Fail(err)
		return report
	}

	for _, step := range tr.steps {
		fmt.Fprintf(out.Stdout, "Executing - %s", step.Description())
		err = step.Execute(runtime, out)
		if err != nil {
			fmt.Fprintf(out.Stderr, "FAIL! Step - %s did not complete as expected.\nREASON - %s", step.Description(), err)
			report.Fail(err)
			break
		}
	}

	fmt.Fprintf(out.Stdout, "Stoppinng %s Runtime", tr.config.Name)
	err = runtime.StopAll()
	if err != nil {
		fmt.Fprintf(out.Stderr, "Could not stop runtime for %s! %s", tr.config.Name, err)
		report.Fail(err)
	}
	return report
}

func (tr *TestRunner) Cleanup() error {
	return tr.runtime.Cleanup()
}

func initSteps(stepConfigs []StepConfig, tasks Tasks, env TaskEnvironment) []Step {
	steps := []Step{}
	for _, stepConfig := range stepConfigs {
		if len(stepConfig.Start) != 0 {
			steps = append(steps, NewStartStep(stepConfig))
		}
		if len(stepConfig.Stop) != 0 {
			steps = append(steps, NewStopStep(stepConfig))
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
