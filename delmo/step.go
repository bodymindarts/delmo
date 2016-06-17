package delmo

import (
	"fmt"
	"time"
)

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
	return runtime.StopServices(output, s.services...)
}

func (s *StopStep) Description() string {
	return fmt.Sprintf("<Stop: %v>", s.services)
}

type DestroyStep struct {
	services []string
}

func NewDestroyStep(config StepConfig) Step {
	return &DestroyStep{
		services: config.Destroy,
	}
}

func (s *DestroyStep) Execute(runtime Runtime, output TestOutput) error {
	return runtime.DestroyServices(output, s.services...)
}

func (s *DestroyStep) Description() string {
	return fmt.Sprintf("<Destroy: %v>", s.services)
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
	return runtime.StartServices(output, s.services...)
}

func (s *StartStep) Description() string {
	return fmt.Sprintf("<Start: %v>", s.services)
}

type WaitStep struct {
	task    TaskConfig
	env     TaskEnvironment
	timeout time.Duration
}

func NewWaitStep(timeout time.Duration, task TaskConfig, env TaskEnvironment) Step {
	return &WaitStep{
		task:    task,
		env:     env,
		timeout: timeout,
	}
}

func (s *WaitStep) Execute(runtime Runtime, output TestOutput) error {
	timeout := time.After(s.timeout)
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
	return fmt.Sprintf("<Wait: %s, Timeout: %ds>", s.task.Name, int(s.timeout.Seconds()))
}

type ExecStep struct {
	task TaskConfig
	env  TaskEnvironment
}

func NewExecStep(task TaskConfig, env TaskEnvironment) Step {
	return &ExecStep{
		task: task,
		env:  env,
	}
}

func (s *ExecStep) Execute(runtime Runtime, output TestOutput) error {
	return runtime.ExecuteTask(s.task.Name, s.task, s.env, output)
}

func (s *ExecStep) Description() string {
	return fmt.Sprintf("<Exec: %s>", s.task.Name)
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
