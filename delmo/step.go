package delmo

import "fmt"

type Step interface {
	Execute(Runtime) error
	Description() string
}

type StopStep struct {
	services []string
}

func NewStopStep(config StepConfig) Step {
	return &StopStep{
		services: config.Stop,
	}
}

func (s *StopStep) Execute(runtime Runtime) error {
	return runtime.StopServices(s.services...)
}

func (s *StopStep) Description() string {
	return fmt.Sprintf("Stop: %v", s.services)
}

type StartStep struct {
	services []string
}

func NewStartStep(config StepConfig) Step {
	return &StartStep{
		services: config.Start,
	}
}

func (s *StartStep) Execute(runtime Runtime) error {
	return runtime.StartServices(s.services...)
}

func (s *StartStep) Description() string {
	return fmt.Sprintf("Start: %v", s.services)
}
