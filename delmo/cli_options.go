package delmo

import (
	"flag"
	"fmt"
	"os"
)

type CLIOptions struct {
	DelmoFile         string
	DockerMachine     string
	OnlyBuildTask     bool
	ParallelExecution bool
	Localhost         string
	Tests             []string
	Help              bool
	Usage             func()
}

func ParseOptions(args []string) CLIOptions {
	flags := flag.NewFlagSet("delmo", flag.ExitOnError)
	usage := func() {
		fmt.Fprintf(os.Stderr, `USAGE: delmo [--version] [--help] [options] [test...]

OPTIONS:
  -f                    path to the spec file (default: "delmo.yml").
  -m                    docker-machine to run the tests on (default: "default").
  --only-build-task     only build the task_image. All other images must be available via docker pull.
  --localhost           IP that will be set to DOCKER_HOST_IP environment variable when not running in a docker-machine.
  --parallel            execute tests in parallel.
`)
	}
	flags.Usage = usage
	var options CLIOptions
	options.Usage = usage
	flags.StringVar(&(options.DelmoFile), "f", "delmo.yml", "")
	flags.StringVar(&(options.DockerMachine), "m", "default", "")
	flags.BoolVar(&(options.OnlyBuildTask), "only-build-task", false, "")
	flags.BoolVar(&(options.ParallelExecution), "parallel", false, "")
	flags.StringVar(&(options.Localhost), "localhost", "", "")
	flags.BoolVar(&(options.Help), "help", false, "")

	flags.Parse(args)

	options.Tests = flags.Args()
	return options
}
