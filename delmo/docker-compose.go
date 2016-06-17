package delmo

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
	args := d.makeArgs("up", append([]string{"-d"}, name...)...)
	cmd := exec.Command(d.rawCmd, args...)
	return cmd.Run()
}

func (d *DockerCompose) SystemOutput() ([]byte, error) {
	args := d.makeArgs("logs")
	cmd := exec.Command(d.rawCmd, args...)
	return cmd.Output()
}

type TaskEnvironment []string

func (d *DockerCompose) ExecuteTask(prefix string, task TaskConfig, env TaskEnvironment, output TestOutput) error {
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

	// need to hookup stdin for docker-compose to have correct output
	cmd.Stdin = os.Stdin
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
	stdoutCh := make(chan struct{})
	stderrCh := make(chan struct{})
	// make sure all output has been streemed before returning control
	defer func() {
		<-stdoutCh
		<-stderrCh
	}()
	go func() {
		for outScanner.Scan() {
			fmt.Fprintf(output.Stdout, "%s | %s\n", prefix, outScanner.Text())
		}
		close(stdoutCh)
	}()
	go func() {
		for errScanner.Scan() {
			fmt.Fprintf(output.Stderr, "%s | %s\n", task.Name, outScanner.Text())
		}
		close(stderrCh)
	}()
	return cmd.Run()
}

func (d *DockerCompose) Cleanup() error {
	args := d.makeArgs("kill")
	cmd := exec.Command(d.rawCmd, args...)
	cmd.Run()
	args = d.makeArgs("rm", "-f", "-v", "-a")
	cmd = exec.Command(d.rawCmd, args...)
	cmd.Run()
	args = d.makeArgs("down", "--volumes", "--remove-orphans")
	cmd = exec.Command(d.rawCmd, args...)
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
