package main

type System struct {
	config SuiteConfig
}

func NewSystem(config SuiteConfig) *System {
	return &System{config: config}
}

func (s *System) NewRuntime(name string) (Runtime, error) {
	return NewDockerCompose(s.config.CompleteFilePath, name)
}
