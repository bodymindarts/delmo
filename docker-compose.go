package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type DockerCompose struct {
	rawCmd      string
	composeFile string
	scope       string
}

func NewDockerCompose(composeFile, scope string) (*DockerCompose, error) {
	cmd, err := assertExecPreconditions()
	if err != nil {
		return nil, err
	}
	dc := &DockerCompose{
		rawCmd:      cmd,
		scope:       scope,
		composeFile: composeFile,
	}
	return dc, nil
}

func (d *DockerCompose) Pull() error {
	args := d.makeArgs("pull", "--ignore-pull-failures")
	cmd := exec.Command(d.rawCmd, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (d *DockerCompose) Build(services ...string) error {
	args := d.makeArgs("build", services...)
	cmd := exec.Command(d.rawCmd, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (d *DockerCompose) StartAll() error {
	args := d.makeArgs("up", "-d", "--force-recreate")
	cmd := exec.Command(d.rawCmd, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s\n%s", strings.TrimSpace(string(out)), err)
	}
	return nil
}

func (d *DockerCompose) StopAll() error {
	args := d.makeArgs("stop")
	cmd := exec.Command(d.rawCmd, args...)
	return cmd.Run()
}

func (d *DockerCompose) StopServices(name ...string) error {
	args := d.makeArgs("stop", name...)
	cmd := exec.Command(d.rawCmd, args...)
	return cmd.Run()
}

func (d *DockerCompose) StartServices(name ...string) error {
	args := d.makeArgs("start", name...)
	cmd := exec.Command(d.rawCmd, args...)
	return cmd.Run()
}

func (d *DockerCompose) SystemOutput() ([]byte, error) {
	args := d.makeArgs("logs")
	cmd := exec.Command(d.rawCmd, args...)
	return cmd.Output()
}

type OutputWrapper struct {
	taskName string
	reporter TaskReporter
}

func (o *OutputWrapper) Write(p []byte) (int, error) {
	o.reporter.TaskOutput(o.taskName, string(p))
	return len(p), nil
}

type TaskEnvironment []string

func (d *DockerCompose) ExecuteTask(task TaskConfig, env TaskEnvironment, reporter TaskReporter) error {
	args := []string{
		"-e",
		"DELMO_TEST_NAME=" + d.scope,
	}
	for _, variable := range env {
		args = append(args, "-e", variable)
	}
	args = append(args, task.Service)
	args = append(args, strings.Split(task.Cmd, " ")...)
	args = d.makeArgs("run", args...)
	cmd := exec.Command(d.rawCmd, args...)
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return err
	}
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating Stderr for Cmd", err)
		return err
	}

	outScanner := bufio.NewScanner(stdOut)
	errScanner := bufio.NewScanner(stdErr)
	go func() {
		for outScanner.Scan() {
			reporter.TaskOutput(task.Name, outScanner.Text())
		}
	}()
	go func() {
		for errScanner.Scan() {
			reporter.TaskOutput(task.Name, errScanner.Text())
		}
	}()
	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return err
}

func (d *DockerCompose) Cleanup() error {
	args := d.makeArgs("rm", "-f", "-v", "-a")
	cmd := exec.Command(d.rawCmd, args...)
	return cmd.Run()
}

func (d *DockerCompose) makeArgs(command string, args ...string) []string {
	return append([]string{
		"--file", d.composeFile, "--project-name", d.scope, command,
	}, args...)
}

func assertExecPreconditions() (string, error) {
	if host := os.Getenv("DOCKER_HOST"); host == "" {
		return "", fmt.Errorf("Environment not setup correctly! DOCKER_HOST is not set")
	}

	cmd, err := exec.LookPath("docker-compose")
	if err != nil {
		return "", err
	}
	return cmd, nil
}
