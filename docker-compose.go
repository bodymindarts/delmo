package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type DockerCompose struct {
	rawCmd      string
	composeFile string
	prefix      string
}

func NewDockerCompose(composeFile, prefix string) (*DockerCompose, error) {
	cmd, err := assertExecPreconditions()
	if err != nil {
		return nil, err
	}
	dc := &DockerCompose{
		rawCmd:      cmd,
		prefix:      prefix,
		composeFile: composeFile,
	}
	return dc, nil
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

func (d *DockerCompose) ExecuteTask(task TaskConfig, reporter TaskReporter) error {
	args := append([]string{task.Service}, strings.Split(task.Cmd, " ")...)
	args = d.makeArgs("run", args...)
	fmt.Println("cmd: docker-compose ", args)
	cmd := exec.Command(d.rawCmd, args...)
	wrapper := &OutputWrapper{
		taskName: task.Name,
		reporter: reporter,
	}
	cmd.Stderr = wrapper
	cmd.Stdout = wrapper
	return cmd.Run()
}

func (d *DockerCompose) Cleanup() error {
	args := d.makeArgs("rm", "-f", "-v", "-a")
	cmd := exec.Command(d.rawCmd, args...)
	return cmd.Run()
}

func (d *DockerCompose) makeArgs(command string, args ...string) []string {
	return append([]string{
		"--file", d.composeFile, "--project-name", d.prefix, command,
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
