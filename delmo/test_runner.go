package delmo

import (
	"errors"
	"fmt"
)

type TestRunner struct {
	testConfig TestConfig
	runtime    Runtime
	spec       *Spec
}

func NewTestRunner(testConfig TestConfig) *TestRunner {
	spec, _ := NewSpec(testConfig.Spec)
	return &TestRunner{
		testConfig: testConfig,
		spec:       spec,
	}
}

func (tr *TestRunner) RunTest(runtime Runtime) error {
	tr.runtime = runtime

	err := runtime.Start()
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't start runtime\n%s", err))
	}

	err = runtime.Stop()
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't stop runtime\n%s", err))
	}

	return errors.New("whoops")
}

func (tr *TestRunner) Output() ([]byte, error) {
	return tr.runtime.Output()
}

func (tr *TestRunner) Cleanup() error {
	return tr.runtime.Cleanup()
}
