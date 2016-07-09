# delmo
DelMo is a framework to test node failure in distributed systems.

It is configured via a `delmo.yml` file which starts any number of containers defined in a `docker-compose.yml` file.
After starting containers via [docker-compose](https://docs.docker.com/compose/overview/)  _tasks_ can be executed to assert the state of the running containers.

## Installation
Find the [latest release](https://github.com/bodymindarts/delmo/releases) and download the binary for your environment.
Alternatively you can follow the [instructions below](#building-delmo) to build it yourself.

## Example

An example test suite is configured in `example/webapp/delmo.yml` and can be executed from the repo root via:
```
delmo -f example/webapp/delmo.yml
```

## Usage
```
USAGE: delmo [--version] [--help] [options] [test...]
```

| Option | meaning |
|-----|---|
| `-f` | Path to the spec file (default delmo.yml) |
| `-m <machine-name>` | The docker-machine to run the tests on. With this option you can specify the name of a docker host that is managed by `docker-machine`. All tests will then be executed on that machine and additionally the environment variable `DOCKER_HOST_IP` will be available and contain the ip returned by `docker-machine ip <machine>`|
| `--only-build-task` | Only build the image required for running tasks. Other images will be pulled but not built. Use this when you want to test images downloaded from a registry without any local changes. Omit this when you are making local changes to the images under test that should be picked up. |
| `--skip-pull` | Skip pulling the images entirely. This is usefull when your don't have network connectivity but all images are already present locally. |
| `--localhost <your-local-ip>` | Used to set `DOCKER_HOST_IP` when running the tests on a local docker daemon (as opposed to one managed by docker-machine). |
| `--parallel` | to execute tests in parallel |

## Building delmo

 For local dev first make sure Go is properly installed, including setting up a GOPATH. After setting up Go, clone this repository into $GOPATH/src/github.com/bodymindarts/delmo.

To install dev dependencies:
```
$ make bootstrap
...
```

To see if the tests are working:
```
$ make test
...
```

To build delmo for your environment
```
$ make dev
...
$ bin/delmo help
...
```

For cross compilation run `make build` this will compile delmo for multiple platforms and place the resulting binaries into the ./pkg directory:
```
$ make build
...
```

## Contributing

Any feedback is welcome. Feel free to open issues or PRs.
