package delmo

//go:generate counterfeiter -o fakes/fake_runtime.go . Runtime
type Runtime interface {
	StartAll(TestOutput) error
	StopAll(TestOutput) error
	StopServices(TestOutput, ...string) error
	StartServices(TestOutput, ...string) error
	DestroyServices(TestOutput, ...string) error
	SystemOutput() ([]byte, error)
	ExecuteTask(string, TaskConfig, TaskEnvironment, TestOutput) error
	Cleanup() error
}
