package main

type Runtime interface {
	StartAll() error
	StopAll() error
	StopServices(...string) error
	StartServices(...string) error
	Output() ([]byte, error)
	Cleanup() error
}
