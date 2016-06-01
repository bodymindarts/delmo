package delmo

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type Suite struct {
	config *SuiteConfig
}

func NewSuite(config *SuiteConfig) *Suite {
	return &Suite{
		config: config,
	}
}

func (s *Suite) Run(ui cli.Ui) (int, error) {
	ui.Info(fmt.Sprintf("Running Test Suite for System %s", s.config.System.Name))

	return 0, nil
}
