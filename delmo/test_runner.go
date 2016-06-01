package delmo

type TestRunner struct {
	testConfig TestConfig
	runtime    Runtime
	report     *TestReport
	spec       *Spec
}

func NewTestRunner(testConfig TestConfig, taskFactory *TaskFactory) *TestRunner {
	spec, _ := NewSpec(testConfig.Name, testConfig.Spec, taskFactory)
	return &TestRunner{
		testConfig: testConfig,
		spec:       spec,
	}
}

func (tr *TestRunner) RunTest(runtime Runtime, listener Listener) *TestReport {
	tr.runtime = runtime
	outputFetcher := func() ([]byte, error) {
		return runtime.Output()
	}
	tr.report = NewTestReport(tr.testConfig.Name, outputFetcher, listener)

	tr.spec.Execute(runtime, tr.report)

	return tr.report
}

func (tr *TestRunner) Cleanup() error {
	for _, step := range tr.report.FailedSteps {
		step.Cleanup()
	}
	return tr.runtime.Cleanup()
}
