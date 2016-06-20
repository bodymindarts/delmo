package delmo

import "flag"

type CLIOptions struct {
	DelmoFile     string
	DockerMachine string
	OnlyBuildTask bool
	Tests         []string
}

func ParseOptions(args []string) CLIOptions {
	flags := flag.NewFlagSet("delmo", flag.ExitOnError)
	var options CLIOptions
	flags.StringVar(&(options.DelmoFile), "f", "delmo.yml", "Path to the delmo.yml file.")
	flags.StringVar(&(options.DockerMachine), "m", "default", "The docker-machine to use.")
	flags.BoolVar(&(options.OnlyBuildTask), "only-build-task", false, "Only build the task_image. All other images must be available via docker pull.")
	flags.Parse(args)

	options.Tests = flags.Args()
	return options
}
