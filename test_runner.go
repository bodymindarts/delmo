package main

type TestRunner struct {
	config                TestConfig
	tasks                 []TaskConfig
	globalTaskEnvironment TaskEnvironment
	runtime               Runtime
	steps                 []Step
	report                *TestReport
}

func NewTestRunner(config TestConfig, tasks Tasks, globalTaskEnvironment TaskEnvironment) *TestRunner {
	return &TestRunner{
		config:                config,
		globalTaskEnvironment: globalTaskEnvironment,
		steps: initSteps(config.Spec, tasks, globalTaskEnvironment),
	}
}

func (tr *TestRunner) RunTest(runtime Runtime, out TestOutput) *TestReport {
	tr.runtime = runtime
	systemOutputFetcher := func() ([]byte, error) {
		return runtime.SystemOutput()
	}
	tr.report = NewTestReport(tr.config.Name, systemOutputFetcher, out)

	tr.report.StartingRuntime()
	err := runtime.StartAll()
	if err != nil {
		tr.report.ErrorStartingRuntime(err)
		return tr.report
	}

	for _, step := range tr.steps {
		tr.report.ExecutingStep(step)
		err = step.Execute(runtime, out)
		if err != nil {
			tr.report.StepExecutionFailed(step, err)
			break
		}
	}

	tr.report.StoppingRuntime()
	err = runtime.StopAll()
	if err != nil {
		tr.report.ErrorStoppingRuntime(err)
	}
	return tr.report
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
