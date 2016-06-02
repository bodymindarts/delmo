package delmo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type SuiteConfig struct {
	System SystemConfig `yaml:"system"`
	Tasks  []TaskConfig `yaml:"tasks"`
	Tests  []TestConfig `yaml:"tests"`
}

type SystemConfig struct {
	Name             string `yaml:"name"`
	File             string `yaml:"file"`
	CompleteFilePath string
	Services         map[string]ServiceConfig
	MachineName      string `yaml:"machine-name"`
}

type TaskConfig struct {
	Name  string    `yaml:"name"`
	Image string    `yaml:"image"`
	Run   RunConfig `yaml:"run"`
}

type RunConfig struct {
	Path string   `yaml:"path"`
	Args []string `yaml:"args"`
}

type TestConfig struct {
	Name string     `yaml:"name"`
	Spec SpecConfig `yaml:"spec"`
}

type SpecConfig []StepConfig

type StepConfig struct {
	Stop   []string `yaml:"stop"`
	Start  []string `yaml:"start"`
	Assert []string `yaml:"assert"`
}

type ComposeConfig struct {
	Services map[string]ServiceConfig `yaml:"services"`
}
type ServiceConfig struct {
	Image string `yaml:"image"`
}

func LoadConfig(path string) (*SuiteConfig, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config SuiteConfig
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	err = loadComposeConfig(path, &config.System)
	return &config, nil
}

func loadComposeConfig(path string, systemConfig *SystemConfig) error {
	composePath := fmt.Sprintf("%s/%s", filepath.Dir(path), systemConfig.File)
	systemConfig.CompleteFilePath = composePath
	bytes, err := ioutil.ReadFile(composePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error loading file '%s'\n%s", composePath, err))
	}

	var composeConfig ComposeConfig
	err = yaml.Unmarshal(bytes, &composeConfig)
	if err != nil {
		return err
	}
	systemConfig.Services = composeConfig.Services
	return nil
}
