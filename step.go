package main

import (
	"fmt"
	"time"
)

const defaultTimeout = time.Second * 60

type Step interface {
	Execute(Runtime, TestOutput) error
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

func (s *StopStep) Execute(runtime Runtime, output TestOutput) error {
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

func (s *StartStep) Execute(runtime Runtime, output TestOutput) error {
	return runtime.StartServices(s.services...)
}

func (s *StartStep) Description() string {
	return fmt.Sprintf("<Start: %v>", s.services)
}

type WaitStep struct {
	task TaskConfig
	env  TaskEnvironment
}

func NewWaitStep(task TaskConfig, env TaskEnvironment) Step {
	return &WaitStep{
		task: task,
		env:  env,
	}
}

func (s *WaitStep) Execute(runtime Runtime, output TestOutput) error {
	timeout := time.After(defaultTimeout)
	i := 0
	for {
		select {
		case <-timeout:
			return fmt.Errorf("Task never completed successfully")
		default:
			i++
			if err := runtime.ExecuteTask(fmt.Sprintf("(%d) %s", i, s.task.Name), s.task, s.env, output); err == nil {
				return nil
			}
		}
	}
}

func (s *WaitStep) Description() string {
	return fmt.Sprintf("<Wait: %s>", s.task.Name)
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

func (s *AssertStep) Execute(runtime Runtime, output TestOutput) error {
	return runtime.ExecuteTask(s.task.Name, s.task, s.env, output)
}

func (s *AssertStep) Description() string {
	return fmt.Sprintf("<Assert: %s>", s.task.Name)
}

type FailStep struct {
	task TaskConfig
	env  TaskEnvironment
}

func NewFailStep(task TaskConfig, env TaskEnvironment) Step {
	return &FailStep{
		task: task,
		env:  env,
	}
}

func (s *FailStep) Execute(runtime Runtime, output TestOutput) error {
	if err := runtime.ExecuteTask(s.task.Name, s.task, s.env, output); err == nil {
		return fmt.Errorf("Expected task to fail!")
	}
	return nil
}

func (s *FailStep) Description() string {
	return fmt.Sprintf("<Fail: %s>", s.task.Name)
}
