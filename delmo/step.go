package delmo

type Step interface {
	Execute(Runtime) error
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
	return nil
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
	return nil
}
