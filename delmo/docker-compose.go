package delmo

import "os/exec"

type DockerCompose struct {
	rawCmd      string
	composeFile string
	prefix      string
}

func NewDockerCompose(composeFile, prefix string) (*DockerCompose, error) {
	cmd, err := exec.LookPath("docker-compose")
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

func (d *DockerCompose) Start() error {
	args := d.makeArgs("up", "-d", "--force-recreate")
	cmd := exec.Command(d.rawCmd, args...)
	return cmd.Run()
}

func (d *DockerCompose) Stop() error {
	args := d.makeArgs("stop")
	cmd := exec.Command(d.rawCmd, args...)
	return cmd.Run()
}

func (d *DockerCompose) Output() ([]byte, error) {
	args := d.makeArgs("logs")
	cmd := exec.Command(d.rawCmd, args...)
	return cmd.Output()
}

func (d *DockerCompose) Cleanup() error {
	args := d.makeArgs("rm", "-f", "-v", "-a")
	cmd := exec.Command(d.rawCmd, args...)
	return cmd.Run()
}

func (d *DockerCompose) makeArgs(args ...string) []string {
	return append([]string{
		"--file", d.composeFile, "--project-name", d.prefix,
	}, args...)
}
