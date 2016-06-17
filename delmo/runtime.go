package delmo

//go:generate counterfeiter -o fakes/fake_runtime.go . Runtime
type Runtime interface {
	StartAll() error
	StopAll() error
	StopServices(...string) error
	StartServices(...string) error
	SystemOutput() ([]byte, error)
	ExecuteTask(string, TaskConfig, TaskEnvironment, TestOutput) error
	Cleanup() error
}
