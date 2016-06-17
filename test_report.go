package main

import "fmt"

type TestReport struct {
	Success      bool
	Error        error
	FailedSteps  []Step
	PassedSteps  []Step
	output       TestOutput
	name         string
	systemOutput SystemOutputFetcher
}

type SystemOutputFetcher func() ([]byte, error)

func NewTestReport(testName string, outputFetcher SystemOutputFetcher) *TestReport {
	return &TestReport{
		Success:      true,
		name:         testName,
		systemOutput: outputFetcher,
	}
}

func (r *TestReport) SystemOutput() string {
	output, err := r.systemOutput()
	if err != nil {
		return fmt.Sprintf("Couldn't fetch output! %s", err)
	}
	return string(output)
}

func (r *TestReport) Fail(err error) {
	r.Success = false
	r.Error = err
}
