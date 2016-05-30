package delmo

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Suite struct {
	System      string        `yaml:"system"`
	Assertions  []Assertion   `yaml:"assertions"`
	Tests       []interface{} `yaml:"tests"`
	ComposeFile string
}

type Assertion struct {
	Name     string   `yaml:"name"`
	Image    string   `yaml:"image"`
	Networks []string `yaml:"networks"`
	Run      string   `yaml:"run"`
}

func Load(path string) (*Suite, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	suite := Suite{}
	err = yaml.Unmarshal(bytes, &suite)
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(path)
	composeFile := filepath.Join(dir, suite.System)
	suite.ComposeFile = composeFile

	return &suite, nil
}
