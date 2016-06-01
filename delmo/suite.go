package delmo

import "github.com/mitchellh/cli"

type Suite struct {
	config *SuiteConfig
}

func NewSuite(config *SuiteConfig) *Suite {
	return &Suite{
		config: config,
	}
}

func (s *Suite) Run(ui cli.Ui) (int, error) {
	ui.Output("HELLO")
	return 0, nil
}
