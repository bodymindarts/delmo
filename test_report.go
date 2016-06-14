package main

import (
	"fmt"
	"os"
	"strings"
)

type TestReport struct {
	Success     bool
	Error       error
	FailedSteps []Step
	PassedSteps []Step
	listeners   []Listener
	name        string
	output      SystemOutputFetcher
}

type Listener interface {
	Output(string)
	Error(string)
}

type SystemListener struct{}

func (s *SystemListener) Output(output string) {
	fmt.Println(output)
}

func (s *SystemListener) Error(output string) {
	fmt.Fprintln(os.Stderr, output)
}

type TaskReporter interface {
	TaskOutput(taskName, output string)
}

type SystemOutputFetcher func() ([]byte, error)

func NewTestReport(testName string, outputFetcher SystemOutputFetcher, listeners ...Listener) *TestReport {
	return &TestReport{
		Success:   true,
		name:      testName,
		listeners: listeners,
		output:    outputFetcher,
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
	output, err := r.output()
	if err != nil {
		return fmt.Sprintf("Couldn't fetch output! %s", err)
	}
	return string(output)
}

func (r *TestReport) reportError(msg string) {
	for _, l := range r.listeners {
		l.Error(msg)
	}
}

func (r *TestReport) reportOutput(msg string) {
	for _, l := range r.listeners {
		l.Output(msg)
	}
}
func (r *TestReport) Fail(msg string, err error) {
	r.reportError(msg)
	r.Success = false
	r.Error = err
}
