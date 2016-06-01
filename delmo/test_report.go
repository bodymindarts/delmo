package delmo

import "fmt"

type TestReport struct {
	Success   bool
	Error     error
	listeners []Listener
	name      string
	output    OutputFetcher
}

type Listener interface {
	Output(string)
	Info(string)
	Error(string)
	Warn(string)
}

type TaskReporter interface {
}

type OutputFetcher func() ([]byte, error)

func NewTestReport(testName string, outputFetcher OutputFetcher, listeners ...Listener) *TestReport {
	return &TestReport{
		Success:   true,
		name:      testName,
		listeners: listeners,
		output:    outputFetcher,
	}
}

func (r *TestReport) ErrorStartingRuntime(err error) {
	r.Fail(fmt.Sprintf("Could not start runtime for %s! %s", r.name, err), err)
}

func (r *TestReport) ErrorStoppingRuntime(err error) {
	r.Fail(fmt.Sprintf("Could not stop runtime for %s! %s", r.name, err), err)
}

func (r *TestReport) RuntimeStarted() {
	r.reportInfo(fmt.Sprintf("Runtime for %s started", r.name))
}
func (r *TestReport) RuntimeStopped() {
	r.reportInfo(fmt.Sprintf("Runtime for %s stopped", r.name))
}

func (r *TestReport) ExecutingStep(step Step) {
	r.reportInfo(fmt.Sprintf("Executing step -  %s", step.Description()))
}

func (r *TestReport) StepExecutionFailed(step Step, err error) {
	r.Fail(fmt.Sprintf("FAIL! Could not execute step %s, REASON: %s", step.Description(), err), err)
}

func (r *TestReport) StepExecutionSucceeded(step Step) {
	r.reportInfo(fmt.Sprintf("Step succeeded"))
}

func (r *TestReport) Output() string {
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

func (r *TestReport) reportInfo(msg string) {
	for _, l := range r.listeners {
		l.Info(msg)
	}
}

func (r *TestReport) Fail(msg string, err error) {
	r.reportError(msg)
	r.Success = false
	r.Error = err
}
