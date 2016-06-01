package delmo

import (
	"fmt"

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
	return []byte(""), nil
}
