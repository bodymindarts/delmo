package delmo

import "fmt"

type TestReport struct {
	Success   bool
	Error     error
	listeners []Listener
	name      string
}

type Listener interface {
	Output(string)
	Info(string)
	Error(string)
	Warn(string)
}

func NewTestReport(testName string, listeners ...Listener) *TestReport {
	return &TestReport{name: testName, listeners: listeners}
}

func (r *TestReport) ErrorStartingRuntime(err error) {
	r.reportError(fmt.Sprintf("Could not start runtime for %s! %s", r.name, err))
}

func (r *TestReport) ErrorStoppingRuntime(err error) {
	r.reportError(fmt.Sprintf("Could not stop runtime for %s! %s", r.name, err))
}

func (r *TestReport) RuntimeStarted() {
	r.reportInfo(fmt.Sprintf("Runtime for %s started", r.name))
}
func (r *TestReport) RuntimeStopped() {
	r.reportInfo(fmt.Sprintf("Runtime for %s stopped", r.name))
}

func (r *TestReport) Output() string {
	return ""
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
