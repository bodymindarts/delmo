package main

type Runtime interface {
	StartAll() error
	StopAll() error
	StopServices(...string) error
	StartServices(...string) error
	SystemOutput() ([]byte, error)
	Cleanup() error
}
