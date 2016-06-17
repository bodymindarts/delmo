package delmo

type TestReport struct {
	Success bool
	Error   error
}

func NewTestReport() *TestReport {
	return &TestReport{
		Success: true,
	}
}

func (r *TestReport) Fail(err error) {
	r.Success = false
	r.Error = err
}
