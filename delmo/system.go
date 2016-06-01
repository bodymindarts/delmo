package delmo

type System struct {
	config SystemConfig
}

func NewSystem(config SystemConfig) *System {
	return &System{config: config}
}

func (s *System) NewRuntime(name string) (Runtime, error) {
	runtime, err := NewDockerCompose(s.config.File, name)
	if err != nil {
		return nil, err
	}
	return runtime, nil
}
