package delmo_test

import (
	"bytes"
	"strings"
	"testing"

	. "github.com/bodymindarts/delmo/delmo"
	"github.com/bodymindarts/delmo/delmo/fakes"
)

func TestTestRunner_RunTest_NoSteps(t *testing.T) {
	config := TestConfig{}
	tasks := Tasks{}
	runner := NewTestRunner(config, tasks, TaskEnvironment{})
	var b bytes.Buffer
	out := TestOutput{
		Stdout: &b,
		Stderr: &b,
	}
	runtime := new(fakes.FakeRuntime)
	runner.RunTest(runtime, out)

	if want, got := 2, runtime.CleanupCallCount(); want != got {
		t.Errorf("Cleanup not called correctly! Want: %d, got: %d", want, got)
	}

	if want, got := 1, runtime.StartAllCallCount(); want != got {
		t.Errorf("StartAll not called correctly! Want: %d, got: %d", want, got)
	}

	if want, got := 1, runtime.StopAllCallCount(); want != got {
		t.Errorf("StopAll not called correctly! Want: %d, got: %d", want, got)
	}
}

func TestTestRunner_RunTest_WithSteps(t *testing.T) {
	config := TestConfig{
		Name: "test",
		Spec: SpecConfig{
			StepConfig{
				Start:  []string{"service"},
				Stop:   []string{"service"},
				Wait:   []string{"fake_task"},
				Exec:   []string{"fake_task"},
				Assert: []string{"fake_task"},
			},
		},
	}
	tasks := Tasks{
		"fake_task": TaskConfig{
			Name: "fake_task",
		},
	}
	runner := NewTestRunner(config, tasks, TaskEnvironment{})
	var b bytes.Buffer
	out := TestOutput{
		Stdout: &b,
		Stderr: &b,
	}
	runtime := new(fakes.FakeRuntime)
	runner.RunTest(runtime, out)
	t.Logf(b.String())
	outputLines := strings.Split(b.String(), "\n")
	step := 0
	if want, got := "Starting 'test' Runtime", outputLines[step]; want != got {
		t.Errorf("Bad step execution line %d!\nWant: '%s', got: '%s'", step, want, got)
	}

	step++
	if want, got := "Executing - <Start: [service]>", outputLines[step]; want != got {
		t.Errorf("Bad step execution line %d!\nWant: '%s', got: '%s'", step, want, got)
	}

	step++
	if want, got := "Executing - <Stop: [service]>", outputLines[step]; want != got {
		t.Errorf("Bad step execution line %d!\nWant: '%s', got: '%s'", step, want, got)
	}

	step++
	if want, got := "Executing - <Wait: fake_task>", outputLines[step]; want != got {
		t.Errorf("Bad step execution line %d!\nWant: '%s', got: '%s'", step, want, got)
	}

	step++
	if want, got := "Executing - <Exec: fake_task>", outputLines[step]; want != got {
		t.Errorf("Bad step execution line %d!\nWant: '%s', got: '%s'", step, want, got)
	}

	step++
	if want, got := "Executing - <Assert: fake_task>", outputLines[step]; want != got {
		t.Errorf("Bad step execution line %d!\nWant: '%s', got: '%s'", step, want, got)
	}

	step++
	if want, got := "Stopping test Runtime", outputLines[step]; want != got {
		t.Errorf("Bad step execution line %d!\nWant: '%s', got: '%s'", step, want, got)
	}
}

func TestTestRunner_NoCleanupOnFailure(t *testing.T) {
	config := TestConfig{
		Name: "test",
		Spec: SpecConfig{
			StepConfig{
				Fail: []string{"fake_task"},
			},
		},
	}
	tasks := Tasks{
		"fake_task": TaskConfig{
			Name: "fake_task",
		},
	}
	runner := NewTestRunner(config, tasks, TaskEnvironment{})
	var b bytes.Buffer
	out := TestOutput{
		Stdout: &b,
		Stderr: &b,
	}
	runtime := new(fakes.FakeRuntime)
	runner.RunTest(runtime, out)

	if want, got := 1, runtime.StartAllCallCount(); want != got {
		t.Errorf("StartAll not called correctly! Want: %d, got: %d", want, got)
	}
	if want, got := 1, runtime.StopAllCallCount(); want != got {
		t.Errorf("StopAll not called correctly! Want: %d, got: %d", want, got)
	}

	// Cleanup should only be called once at beginning
	if want, got := 1, runtime.CleanupCallCount(); want != got {
		t.Errorf("Cleanup not called correctly! Want: %d, got: %d", want, got)
	}
}
