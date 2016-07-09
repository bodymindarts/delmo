# delmo
DelMo is a framework to test node failure in distributed systems.

It is configured via a `delmo.yml` file which starts any number of containers defined in a `docker-compose.yml` file.
After starting containers via [docker-compose](https://docs.docker.com/compose/overview/)  _tasks_ can be executed to assert the state of the running containers.

## Installation
Find the [latest release](https://github.com/bodymindarts/delmo/releases) and download the binary for your environment.
Alternatively you can follow the [instructions below](#building-delmo) to build it yourself.

## Example

An example test suite is configures in `example/webapp/delmo.yml` and can be executed from the repo root via:
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
| `-m` | The docker-machine to run the tests on. |
| `--only-build-task` |

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
