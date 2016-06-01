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

func (s *Spec) Execute(runtime Runtime, reporter *TestReport) error {
	err := runtime.StartAll()
	if err != nil {
		reporter.ErrorStartingRuntime(err)
		return err
	}
	reporter.RuntimeStarted()

	for _, step := range s.steps {
		reporter.ExecutingStep(step)
		err = step.Execute(runtime)
		if err != nil {
			reporter.StepExecutionFailed(step, err)
			break
		}
		reporter.StepExecutionSucceeded(step)
	}

	err = runtime.StopAll()
	if err != nil {
		reporter.ErrorStoppingRuntime(err)
		return err
	}
	reporter.RuntimeStopped()
	return nil
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
