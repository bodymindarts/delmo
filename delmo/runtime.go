package delmo

type Runtime interface {
	StartAll() error
	StopAll() error
	StopServices(...string) error
	StartServices(...string) error
	SystemOutput() ([]byte, error)
	ExecuteTask(string, TaskConfig, TaskEnvironment, TestOutput) error
	Cleanup() error
}
