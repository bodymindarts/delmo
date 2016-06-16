package main

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
	systemOutputFetcher := func() ([]byte, error) {
		return runtime.SystemOutput()
	}
	report := NewTestReport(tr.config.Name, systemOutputFetcher, out)

	tr.runtime.Cleanup()
	for _, step := range tr.beforeSteps {
		report.ExecutingStep(step)
		err := step.Execute(runtime, out)
		if err != nil {
			report.StepExecutionFailed(step, err)
			return report
		}
	}

	report.StartingRuntime()
	err := runtime.StartAll()
	if err != nil {
		report.ErrorStartingRuntime(err)
		return report
	}

	for _, step := range tr.steps {
		report.ExecutingStep(step)
		err = step.Execute(runtime, out)
		if err != nil {
			report.StepExecutionFailed(step, err)
			break
		}
	}

	report.StoppingRuntime()
	err = runtime.StopAll()
	if err != nil {
		report.ErrorStoppingRuntime(err)
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
