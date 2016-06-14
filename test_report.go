package main

import (
	"fmt"
	"io"
	"strings"
)

type TestReport struct {
	Success      bool
	Error        error
	FailedSteps  []Step
	PassedSteps  []Step
	output       TestOutput
	name         string
	systemOutput SystemOutputFetcher
}

type TestOutput struct {
	Stdout io.Writer
	Stderr io.Writer
}

type SystemOutputFetcher func() ([]byte, error)

func NewTestReport(testName string, outputFetcher SystemOutputFetcher, output TestOutput) *TestReport {
	return &TestReport{
		Success:      true,
		name:         testName,
		output:       output,
		systemOutput: outputFetcher,
	}
}

func (r *TestReport) TaskOutput(taskName, output string) {
	r.reportOutput(fmt.Sprintf("%s >| %s", taskName, strings.TrimSpace(output)))
}

func (r *TestReport) StartingRuntime() {
	r.reportOutput(fmt.Sprintf("Starting %s Runtime", r.name))
}
func (r *TestReport) StoppingRuntime() {
	r.reportOutput(fmt.Sprintf("Stopping %s Runtime", r.name))
}

func (r *TestReport) ErrorStartingRuntime(err error) {
	r.Fail(fmt.Sprintf("Could not start runtime for %s!\n%s", r.name, err), err)
}

func (r *TestReport) ErrorStoppingRuntime(err error) {
	r.Fail(fmt.Sprintf("Could not stop runtime for %s! %s", r.name, err), err)
}

func (r *TestReport) ExecutingStep(step Step) {
	r.reportOutput(fmt.Sprintf("Executing - %s", step.Description()))
}

func (r *TestReport) StepExecutionFailed(step Step, err error) {
	r.Fail(fmt.Sprintf("FAIL! Step - %s did not complete as expected.\nREASON - %s", step.Description(), err), err)
}

func (r *TestReport) SystemOutput() string {
	output, err := r.systemOutput()
	if err != nil {
		return fmt.Sprintf("Couldn't fetch output! %s", err)
	}
	return string(output)
}

func (r *TestReport) reportError(msg string) {
	fmt.Fprintln(r.output.Stderr, msg)
}

func (r *TestReport) reportOutput(msg string) {
	fmt.Fprintln(r.output.Stdout, msg)
}
func (r *TestReport) Fail(msg string, err error) {
	r.reportError(msg)
	r.Success = false
	r.Error = err
}
