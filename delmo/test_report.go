package delmo

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

func (r *TestReport) Output() string {
	return ""
}
