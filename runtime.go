package main

type Runtime interface {
	StartAll() error
	StopAll() error
	StopServices(...string) error
	StartServices(...string) error
	SystemOutput() ([]byte, error)
	ExecuteTask(TaskConfig, TaskEnvironment, TaskReporter) error
	Cleanup() error
}
