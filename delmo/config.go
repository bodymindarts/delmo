package delmo

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	dockerConfig "github.com/docker/libcompose/config"
	"github.com/docker/libcompose/lookup"
	"gopkg.in/yaml.v2"
)

type SuiteConfig struct {
	System SystemConfig `yaml:"system"`
}

type SystemConfig struct {
	Name     string `yaml:"name"`
	File     string `yaml:"file"`
	Services map[string]*dockerConfig.ServiceConfig
	Volumes  map[string]*dockerConfig.VolumeConfig
	Networks map[string]*dockerConfig.NetworkConfig
}

func LoadConfig(path string) (*SuiteConfig, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config SuiteConfig
	yaml.Unmarshal(bytes, &config)
	err = loadComposeFile(path, &config.System)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func loadComposeFile(path string, systemConfig *SystemConfig) error {
	composePath := fmt.Sprintf("%s/%s", filepath.Dir(path), systemConfig.File)
	bytes, err := ioutil.ReadFile(composePath)
	if err != nil {
		return err
	}

	services, volumes, networks, err := dockerConfig.Merge(dockerConfig.NewServiceConfigs(), &lookup.OsEnvLookup{}, &lookup.FileConfigLookup{}, "", bytes, &dockerConfig.ParseOptions{})
	systemConfig.Services = services
	systemConfig.Volumes = volumes
	systemConfig.Networks = networks
	return err
}
