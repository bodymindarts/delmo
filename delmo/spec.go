package delmo

type Spec struct {
	config SpecConfig
	steps  []Step
}

func NewSpec(config SpecConfig) (*Spec, error) {
	spec := &Spec{
		config: config,
	}
	spec.steps = initSteps(config)
	return spec, nil
}

func initSteps(stepConfigs []StepConfig) []Step {
	steps := []Step{}
	for _, stepConfig := range stepConfigs {
		if len(stepConfig.Start) != 0 {
			steps = append(steps, NewStartStep(stepConfig))
		}
		if len(stepConfig.Stop) != 0 {
			steps = append(steps, NewStopStep(stepConfig))
		}
	}
	return steps
}
