package main

type TestRunner struct {
	testConfig TestConfig
	runtime    Runtime
	steps      []Step
	report     *TestReport
}

func NewTestRunner(testConfig TestConfig, taskFactory *TaskFactory, globals GlobalContext) *TestRunner {
	context := TestContext{
		DockerHostSyncDir: globals.DockerHostSyncDir,
		TestName:          testConfig.Name,
	}
	return &TestRunner{
		testConfig: testConfig,
		steps:      initSteps(context, testConfig.Spec, taskFactory),
	}
}

func (tr *TestRunner) RunTest(runtime Runtime, listener Listener) *TestReport {
	tr.runtime = runtime
	outputFetcher := func() ([]byte, error) {
		return runtime.Output()
	}
	tr.report = NewTestReport(tr.testConfig.Name, outputFetcher, listener)

	tr.report.StartingRuntime()
	err := runtime.StartAll()
	if err != nil {
		tr.report.ErrorStartingRuntime(err)
		return tr.report
	}

	for _, step := range tr.steps {
		tr.report.ExecutingStep(step)
		err = step.Execute(runtime, tr.report)
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
	for _, step := range tr.report.FailedSteps {
		step.Cleanup()
	}
	return tr.runtime.Cleanup()
}

func initSteps(context TestContext, stepConfigs []StepConfig, taskFactory *TaskFactory) []Step {
	steps := []Step{}
	for _, stepConfig := range stepConfigs {
		if len(stepConfig.Start) != 0 {
			steps = append(steps, NewStartStep(stepConfig))
		}
		if len(stepConfig.Stop) != 0 {
			steps = append(steps, NewStopStep(stepConfig))
		}
		if len(stepConfig.Assert) != 0 {
			for _, taskName := range stepConfig.Assert {
				task := taskFactory.Task(context, taskName)
				steps = append(steps, NewAssertStep(task))
			}
		}
	}
	return steps
}
