package delmo

import (
	"fmt"
	"os"

	"github.com/fsouza/go-dockerclient"
)

var (
	dockerClient *docker.Client
	initError    error
)

func init() {
	dockerClient, initError = docker.NewClientFromEnv()
}

type Tasks map[string]Task

type Task struct {
	client *docker.Client
	config TaskConfig
}

func NewTasks(configs []TaskConfig) (Tasks, error) {
	if initError != nil {
		return nil, fmt.Errorf("Could not initialize docker client! %s", initError)
	}
	tasks := Tasks{}
	for _, taskConfig := range configs {
		tasks[taskConfig.Name] = newTask(taskConfig)
	}
	return tasks, nil
}

func newTask(config TaskConfig) Task {
	return Task{
		config: config,
		client: dockerClient,
	}
}

func (t Task) Execute() ([]byte, error) {
	testName := "webapp"
	createOptions := docker.CreateContainerOptions{
		Name: fmt.Sprintf("%s_%s", testName, t.config.Name),
		Config: &docker.Config{
			Image: t.config.Image,
			Cmd:   append([]string{t.config.Run.Path}, t.config.Run.Args...),
		},
	}
	container, err := t.client.CreateContainer(createOptions)
	if err != nil {
		fmt.Printf("ERROR creating container: %s\n", err)
		return nil, err
	}

	hostConfig := &docker.HostConfig{}
	err = t.client.StartContainer(container.ID, hostConfig)
	if err != nil {
		fmt.Printf("ERROR starting container: %s\n", err)
		return nil, err
	}

	ret, err := t.client.WaitContainer(container.ID)
	if err != nil {
		fmt.Printf("ERROR waiting container: %s\n", err)
		return nil, err
	}

	logOptions := docker.LogsOptions{
		Container:    container.ID,
		OutputStream: os.Stdout,
		ErrorStream:  os.Stderr,
		Stdout:       true,
		Stderr:       true,
	}
	err = t.client.Logs(logOptions)
	if err != nil {
		fmt.Printf("ERROR logging container: %s\n", err)
		return nil, err
	}
	fmt.Printf("Return value for container: %d\n", ret)

	removeOptions := docker.RemoveContainerOptions{
		ID:            container.ID,
		RemoveVolumes: true,
		Force:         true,
	}
	t.client.RemoveContainer(removeOptions)

	return []byte(""), nil
}
