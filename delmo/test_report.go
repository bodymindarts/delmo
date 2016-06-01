package delmo

type TestReport struct {
	Success bool
	Error   error
}

func NewTestReport() *TestReport {
	return &TestReport{}
}

func (r *TestReport) Output() string {
	return ""
}
