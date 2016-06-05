package main

import (
	"fmt"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

var (
	dockerClient *docker.Client
	initError    error
)

func init() {
	dockerClient, initError = docker.NewClientFromEnv()
}

type TaskFactory struct {
	configs map[string]TaskConfig
}

type Task struct {
	context TestContext
	client  *docker.Client
	Name    string
	config  TaskConfig
}

func NewTaskFactory(configs []TaskConfig) (*TaskFactory, error) {
	if initError != nil {
		return nil, fmt.Errorf("Could not initialize docker client! %s", initError)
	}

	configMap := map[string]TaskConfig{}
	for _, cfg := range configs {
		configMap[cfg.Name] = cfg
	}
	return &TaskFactory{
		configs: configMap,
	}, nil
}

func (t *TaskFactory) Task(context TestContext, taskName string) Task {
	config := t.configs[taskName]
	return Task{
		config:  config,
		context: context,
		client:  dockerClient,
		Name:    config.Name,
	}
}

type ReportWriter struct {
	reporter TaskReporter
	taskName string
}

func (r ReportWriter) Write(p []byte) (int, error) {
	r.reporter.TaskOutput(r.taskName, string(p))
	return len(p), nil
}

func (t Task) Execute(reporter TaskReporter) (int, error) {
	createOptions := docker.CreateContainerOptions{
		Name: t.containerName(),
		Config: &docker.Config{
			Image:      t.config.Image,
			Cmd:        append([]string{t.config.Run.Path}, t.config.Run.Args...),
			WorkingDir: "/delmo",
		},
	}
	container, err := t.client.CreateContainer(createOptions)
	if err != nil {
		if strings.Contains(err.Error(), "container already exists") {
			t.Cleanup()
			container, err = t.client.CreateContainer(createOptions)
			if err != nil {
				return 1, fmt.Errorf("Failed to re-create container %s; aborting", createOptions.Name)
			}
		} else {
			return 1, fmt.Errorf("Failed to create container from image %s: %s", t.config.Image, err)
		}
	}

	hostConfig := &docker.HostConfig{
		Binds: []string{fmt.Sprintf("%s:%s", t.context.DockerHostSyncDir, "/delmo")},
	}
	err = t.client.StartContainer(container.ID, hostConfig)
	if err != nil {
		return 1, err
	}

	returnValue, err := t.client.WaitContainer(container.ID)
	if err != nil {
		return 1, err
	}

	stream := ReportWriter{
		taskName: t.config.Name,
		reporter: reporter,
	}

	logOptions := docker.LogsOptions{
		Container:    container.ID,
		OutputStream: stream,
		ErrorStream:  stream,
		Stdout:       true,
		Stderr:       true,
	}
	err = t.client.Logs(logOptions)
	if err != nil {
		return 1, err
	}

	removeOptions := docker.RemoveContainerOptions{
		ID:            container.ID,
		RemoveVolumes: true,
		Force:         true,
	}
	t.client.RemoveContainer(removeOptions)

	return returnValue, nil
}

func (t *Task) Cleanup() error {
	containers, err := t.client.ListContainers(docker.ListContainersOptions{
		All: true,
		Filters: map[string][]string{
			"name": []string{t.containerName()},
		},
	})
	if err != nil {
		return fmt.Errorf("Failed to query list of containers: %s", err)
	}

	if len(containers) == 0 {
		return nil
	}

	for _, container := range containers {
		err = t.client.RemoveContainer(docker.RemoveContainerOptions{
			ID:    container.ID,
			Force: true,
		})
		if err != nil {
			return fmt.Errorf("Failed to purge container %s: %s", container.ID, err)
		}
	}
	return nil
}

func (t *Task) containerName() string {
	return fmt.Sprintf("%s__%s", t.context.TestName, t.config.Name)
}
