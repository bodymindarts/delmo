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

type TaskFactory struct {
	configs map[string]TaskConfig
}

type Task struct {
	context string
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

func (t *TaskFactory) Task(context, taskName string) Task {
	config := t.configs[taskName]
	return Task{
		config:  config,
		context: context,
		client:  dockerClient,
		Name:    config.Name,
	}
}

func (t Task) Execute() ([]byte, error) {
	createOptions := docker.CreateContainerOptions{
		Name: t.containerName(),
		Config: &docker.Config{
			Image: t.config.Image,
			Cmd:   append([]string{t.config.Run.Path}, t.config.Run.Args...),
		},
	}
	container, err := t.client.CreateContainer(createOptions)
	// if err != nil {
	// 	if strings.Contains(err.Error(), "container already exists") {
	// 		// Get the ID of the existing container so we can delete it
	// 		containers, err := t.client.ListContainers(docker.ListContainersOptions{
	// 			// The image might be in use by a stopped container, so check everything
	// 			All: true,
	// 			Filters: map[string][]string{
	// 				"name": []string{createOptions.Name},
	// 			},
	// 		})
	// 		if err != nil {
	// 			return nil, fmt.Errorf("Failed to query list of containers: %s", err)
	// 		}

	// 		if len(containers) == 0 {
	// 			return nil, fmt.Errorf("Failed to get id for container %s", createOptions.Name)
	// 		}

	// 		for _, container := range containers {
	// 			err = t.client.RemoveContainer(docker.RemoveContainerOptions{
	// 				ID:    container.ID,
	// 				Force: true,
	// 			})
	// 			if err != nil {
	// 				return nil, fmt.Errorf("Failed to purge container %s: %s", container.ID, err)
	// 			}
	// 		}

	// 		container, err = t.client.CreateContainer(createOptions)
	// 		if err != nil {
	// 			return nil, fmt.Errorf("Failed to re-create container %s; aborting", createOptions.Name)
	// 		}
	// 	} else {
	// 		return nil, fmt.Errorf("Failed to create container from image %s: %s", t.config.Image, err)
	// 	}
	// }

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

func (t *Task) containerName() string {
	return fmt.Sprintf("%s__%s", t.context, t.config.Name)
}
