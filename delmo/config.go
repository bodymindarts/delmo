package delmo

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Suite    SuiteConfig  `yaml:"suite"`
	TaskList []TaskConfig `yaml:"tasks"`
	Tasks    Tasks        `yaml:"-"`
	Tests    []TestConfig `yaml:"tests"`
}

type Tasks map[string]TaskConfig

type SuiteConfig struct {
	Name          string `yaml:"name"`
	RawSystemPath string `yaml:"system"`
	System        string `yaml:"-"`
	TaskService   string `yaml:"task_service"`
	OnlyBuildTask bool
}

type TaskConfig struct {
	Name    string `yaml:"name"`
	Service string
	Cmd     string `yaml:"command"`
}

type TestConfig struct {
	Name          string     `yaml:"name"`
	BeforeStartup []string   `yaml:"before_startup"`
	Spec          SpecConfig `yaml:"spec"`
}

type SpecConfig []StepConfig

type StepConfig struct {
	Start   []string `yaml:"start"`
	Stop    []string `yaml:"stop"`
	Destroy []string `yaml:"destroy"`
	Wait    []string `yaml:"wait"`
	Exec    []string `yaml:"exec"`
	Assert  []string `yaml:"assert"`
	Fail    []string `yaml:"fail"`
}

type ComposeConfig struct {
	Services map[string]ServiceConfig `yaml:"services"`
}
type ServiceConfig struct {
	Image string `yaml:"image"`
}

func LoadConfig(path string) (*Config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	config.Suite.System = filepath.Join(filepath.Dir(path), config.Suite.RawSystemPath)
	tasks := map[string]TaskConfig{}
	for _, t := range config.TaskList {
		t.Service = config.Suite.TaskService
		tasks[t.Name] = t
	}
	config.Tasks = tasks

	return &config, nil
}
