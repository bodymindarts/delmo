package main

import "fmt"

type Step interface {
	Execute(Runtime, TaskReporter) error
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

func (s *StopStep) Execute(runtime Runtime, reporter TaskReporter) error {
	return runtime.StopServices(s.services...)
}

func (s *StopStep) Description() string {
	return fmt.Sprintf("<Stop: %v>", s.services)
}

type StartStep struct {
	services []string
}

func NewStartStep(config StepConfig) Step {
	return &StartStep{
		services: config.Start,
	}
}

func (s *StartStep) Execute(runtime Runtime, reporter TaskReporter) error {
	return runtime.StartServices(s.services...)
}

func (s *StartStep) Description() string {
	return fmt.Sprintf("<Start: %v>", s.services)
}

type AssertStep struct {
	task TaskConfig
	env  TaskEnvironment
}

func NewAssertStep(task TaskConfig, env TaskEnvironment) Step {
	return &AssertStep{
		task: task,
		env:  env,
	}
}

func (s *AssertStep) Execute(runtime Runtime, reporter TaskReporter) error {
	exit, err := runtime.ExecuteTask(s.task, s.env, reporter)
	if err != nil {
		return err
	}
	if exit != 0 {
		return fmt.Errorf("Task exited with non 0 exit status!")
	}
	return nil
}

func (s *AssertStep) Description() string {
	return fmt.Sprintf("<Assert: %s>", s.task.Name)
}
