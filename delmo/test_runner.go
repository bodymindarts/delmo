package delmo

type TestRunner struct {
	testConfig TestConfig
}

func NewTestRunner(testConfig TestConfig) *TestRunner {
	return &TestRunner{
		testConfig: testConfig,
	}
}

func (tr *TestRunner) RunTest(runtime Runtime) error {
	return nil
}
