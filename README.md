# delmo
DelMo is a tool to test systems running within multiple docker containers.

In perticular it is possible to start and stop individual services to test how the system behaves when facing outages. It was written and is still being used to test automatic failover capabilities of a clustered postgresql deployment. It is well suited to run integration tests of any system that runs as collaborating docker containers.

After [installation](#installation) you can run the [example](#examples) to quickly see delmo in action.

## Installation
Find the [latest release](https://github.com/bodymindarts/delmo/releases) and download the binary for your environment.
Alternatively you can follow the [instructions below](#building-delmo) to build it yourself.

## Usage
```
$ delmo -h
USAGE: delmo [--version] [--help] [options] [test...]

OPTIONS:
  -f                    path to the spec file (default: "delmo.yml").
  -m <machine-name>     docker-machine to run the tests on. DOCKER_HOST_IP will
                        be set to the ip returned by 'docker-machine ip <machine>'.
  --localhost <ip>      an IP that DOCKER_HOST_IP will be set to when not using -m.
  --parallel            execute tests in parallel.
  --only-build-task     only build the task_image. All other images must be
                        available via 'docker pull'.
  --skip-pull           don't pull the images before building.
```

Omitting `[test...]` will result in all tests being run.

## Files

For delmo to work 2 files must be provided.

A docker-compose.yml (complete file reference can be found [here](https://docs.docker.com/compose/compose-file/)) that defines the way your containers will start.
```yaml
version: '2'

services:
  tests:
    image: busybox
```
A delmo.yml that defines the tests that will run and the tasks that will be executed during the tests.
```yaml
---
suite:
  name: Webapp                   # Name of the test suite
  system: docker-compose.yml     # Path to a file that docker-compose can read
  task_service: tests            # The name of a service in the docker-compose.yml
                                 # that will be used to execute tasks

tasks:                           # Definition of tasks that can be run during a test.
                                 # These tasks get run by calling
                                 # 'docker-compose run <task_service> <command>'

- name: hello_world              # The name of a task to refer to it in a test.
  command: echo hello world      # The command to run within the <task_service> image.
- name: list_root
  command: ls /
- name: failing_task
  command: exit 1

tests:                           # Definition of tests to run.

- name: example_test             # Name of the test
  before_startup: [hello_world]  # Tasks to run before running 'docker-compose up'
  spec:                          # The steps to be performed during the test
                                 # Before running any steps 'docker-compose up' is run                                 # to start all the defined containers
  - exec:
    - hello_world
    - list_root
  - failing: [failing_task]

```

### Available steps
The `spec` key is an array of hashes that define the steps which should be executed during a test.
The type of step is determined by which keys are present.

| step | meaning |
| ---- | ---- |
| `- assert: [<task>...]` | Array of tasks to run. The test will fail if a task returns a non-0 exit status |
| `- exec: [<task>...]` |  Same as `assert`. Provided for differentiating steps that are preparatory in nature. |
| `- {wait: <task>, timeout: 120}` | A Task to repeat as long as the exit status is non-0. The test fails if the `<task>` doesn't return '0' within `<timeout>` seconds. `timeout:` key is optional, default is 60. |
| `- fail: [<task>...]` | Opposite of assert. These tasks are expected to return non-0. The test fails if a '0' exit status is returned. |
| `- stop: [<service>...]` | An array of service names. The services must be defined in the docker-compose.yml and will be stopped via `docker-compose stop [<service>...]` |
| `- start: [<service>...]` | An array of service names. The services must be defined in the docker-compose.yml and will be started via `docker-compose up [<service>...]` |
| `- destroy: [<service>...]` | An array of service names. The services must be defined in the docker-compose.yml. The containers will be killed and then removed via `docker-compose kill [<service>...] && docker-compose rm -f -v [<service>...]` |

### Examples

An [example](./example/webapp) test suite is configured in `example/webapp/delmo.yml` and can be executed from the repo root via:
```
delmo -f example/webapp/delmo.yml
```

Delmo was originally written to test clustering and automatic failover of postgresql nodes. It is still being used and you can [look here](https://github.com/dingotiles/dingo-postgresql-release/tree/master/images) to see real-world usage of delmo.

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
