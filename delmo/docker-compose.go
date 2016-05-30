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
	handle      *dockerComposeHandle
}

type dockerComposeHandle struct {
	cmd    *exec.Cmd
	output *bytes.Buffer
	stopCh chan int
}

func NewDockerCompose(composeFile string) (*DockerCompose, error) {
	cmd, err := exec.LookPath("docker-compose")
	if err != nil {
		return nil, err
	}
	dc := &DockerCompose{
		rawCmd:      cmd,
		composeFile: composeFile,
	}
	return dc, nil
}

func (d *DockerCompose) Start() {
	args := []string{
		"-f", d.composeFile, "up",
	}
	cmd := exec.Command(d.rawCmd, args...)
	buf := new(bytes.Buffer)
	d.handle = &dockerComposeHandle{
		cmd:    cmd,
		stopCh: make(chan int),
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
	// h.cmd.Stderr = h.Output

	if err := h.cmd.Start(); err != nil {
		log.Fatal(err)
	}

	select {
	case <-h.stopCh:
		os.Stdout.Write([]byte("recieved stop signal"))
	}

	h.cmd.Process.Signal(os.Interrupt)
}

func (h *dockerComposeHandle) stop() {
	h.stopCh <- 1
}
