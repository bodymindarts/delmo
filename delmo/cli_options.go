package delmo

import "flag"

type CLIOptions struct {
	DelmoFile         string
	DockerMachine     string
	OnlyBuildTask     bool
	ParallelExecution bool
	Localhost         string
	Tests             []string
}

func ParseOptions(args []string) CLIOptions {
	flags := flag.NewFlagSet("delmo", flag.ExitOnError)
	var options CLIOptions
	flags.StringVar(&(options.DelmoFile), "f", "delmo.yml", "Path to the delmo config file.")
	flags.StringVar(&(options.DockerMachine), "m", "default", "The docker-machine to use.")
	flags.BoolVar(&(options.OnlyBuildTask), "only-build-task", false, "Only build the task_image. All other images must be available via docker pull.")
	flags.BoolVar(&(options.ParallelExecution), "parallel", false, "Execute tests in parallel.")
	flags.StringVar(&(options.Localhost), "localhost", "", "Run containers on local machine passing IP as DOCKER_HOST_IP")

	flags.Parse(args)

	options.Tests = flags.Args()
	return options
}
