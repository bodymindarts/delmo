package delmo

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

type DockerCompose struct {
	rawCmd      string
	composeFile string
	prefix      string
	handle      *dockerComposeHandle
}

type dockerComposeHandle struct {
	cmd    *exec.Cmd
	output *bytes.Buffer
	stopCh chan int
	doneCh chan struct{}
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

func (d *DockerCompose) Start() {
	args := []string{
		"-f", d.composeFile, "-p", d.prefix, "up",
	}
	cmd := exec.Command(d.rawCmd, args...)
	buf := new(bytes.Buffer)
	d.handle = &dockerComposeHandle{
		cmd:    cmd,
		stopCh: make(chan int),
		doneCh: make(chan struct{}),
		output: buf,
	}
	go d.handle.run()
}

func (d *DockerCompose) Stop() {
	d.handle.stop()
}
func (d *DockerCompose) Output() io.Reader {
	return d.handle.output
}

func (h *dockerComposeHandle) run() {
	h.cmd.Stdout = h.output

	if err := h.cmd.Start(); err != nil {
		log.Fatal(err)
	}

	select {
	case <-h.stopCh:
	}

	h.cmd.Process.Signal(os.Interrupt)
	h.cmd.Wait()
	close(h.doneCh)
}

func (h *dockerComposeHandle) stop() {
	h.stopCh <- 1
	select {
	case <-h.doneCh:
	}
}
