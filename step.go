package main

import (
	"fmt"
	"time"
)

const defaultTimeout = time.Second * 30

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

func (s *WaitStep) Execute(runtime Runtime, reporter TaskReporter) error {
	timeout := time.After(defaultTimeout)
	select {
	case <-timeout:
		return fmt.Errorf("Task never completed successfully")
	default:
		if err := runtime.ExecuteTask(s.task, s.env, reporter); err == nil {
			return nil
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

func (s *AssertStep) Execute(runtime Runtime, reporter TaskReporter) error {
	return runtime.ExecuteTask(s.task, s.env, reporter)
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

func (s *FailStep) Execute(runtime Runtime, reporter TaskReporter) error {
	if err := runtime.ExecuteTask(s.task, s.env, reporter); err == nil {
		return fmt.Errorf("Expected task to fail!")
	}
	return nil
}

func (s *FailStep) Description() string {
	return fmt.Sprintf("<Fail: %s>", s.task.Name)
}
