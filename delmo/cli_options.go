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
	SkipPull          bool
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
  -m <machine-name>     docker-machine to run the tests on. DOCKER_HOST_IP will
                        be set to the ip returned by 'docker-machine ip <machine>'.
  --localhost <ip>      an IP that DOCKER_HOST_IP will be set to when not using -m.
  --parallel            execute tests in parallel.
  --only-build-task     only build the task_image. All other images must be
                        available via 'docker pull'.
  --skip-pull           don't pull the images before building.
`)
	}
	flags.Usage = usage
	var options CLIOptions
	options.Usage = usage
	flags.StringVar(&(options.DelmoFile), "f", "delmo.yml", "")
	flags.StringVar(&(options.DockerMachine), "m", "", "")
	flags.BoolVar(&(options.OnlyBuildTask), "only-build-task", false, "")
	flags.BoolVar(&(options.ParallelExecution), "parallel", false, "")
	flags.BoolVar(&(options.SkipPull), "skip-pull", false, "")
	flags.StringVar(&(options.Localhost), "localhost", "", "")
	flags.BoolVar(&(options.Help), "help", false, "")

	flags.Parse(args)

	options.Tests = flags.Args()
	return options
}
