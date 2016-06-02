package delmo

type Spec struct {
	context     TestContext
	config      SpecConfig
	steps       []Step
	taskFactory *TaskFactory
}

func NewSpec(context TestContext, config SpecConfig, taskFactory *TaskFactory) (*Spec, error) {
	spec := &Spec{
		context:     context,
		config:      config,
		taskFactory: taskFactory,
	}
	spec.steps = initSteps(context, config, taskFactory)
	return spec, nil
}

func (s *Spec) Execute(runtime Runtime, reporter *TestReport) error {
	reporter.StartingRuntime()
	err := runtime.StartAll()
	if err != nil {
		reporter.ErrorStartingRuntime(err)
		return err
	}

	for _, step := range s.steps {
		reporter.ExecutingStep(step)
		err = step.Execute(runtime, reporter)
		if err != nil {
			reporter.StepExecutionFailed(step, err)
			break
		}
	}

	reporter.StoppingRuntime()
	err = runtime.StopAll()
	if err != nil {
		reporter.ErrorStoppingRuntime(err)
		return err
	}
	return nil
}

func initSteps(context TestContext, stepConfigs []StepConfig, taskFactory *TaskFactory) []Step {
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
				task := taskFactory.Task(context, taskName)
				steps = append(steps, NewAssertStep(task))
			}
		}
	}
	return steps
}
