package delmo

type Spec struct {
	name        string
	config      SpecConfig
	steps       []Step
	taskFactory *TaskFactory
}

func NewSpec(name string, config SpecConfig, taskFactory *TaskFactory) (*Spec, error) {
	spec := &Spec{
		name:        name,
		config:      config,
		taskFactory: taskFactory,
	}
	spec.steps = initSteps(name, config, taskFactory)
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
		err = step.Execute(runtime, reporter)
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

func initSteps(name string, stepConfigs []StepConfig, taskFactory *TaskFactory) []Step {
	steps := []Step{}
	for _, stepConfig := range stepConfigs {
		if len(stepConfig.Start) != 0 {
			steps = append(steps, NewStartStep(stepConfig))
		}
		if len(stepConfig.Stop) != 0 {
			steps = append(steps, NewStopStep(stepConfig))
		}
		if len(stepConfig.Assert) != 0 {
			for _, taskName := range stepConfig.Assert {
				task := taskFactory.Task(name, taskName)
				steps = append(steps, NewAssertStep(task))
			}
		}
	}
	return steps
}
