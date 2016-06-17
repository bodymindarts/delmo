package delmo

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func writeFiles(name, config, compose string, t *testing.T) (string, string) {
	tmpDir, err := ioutil.TempDir("", name)
	if err != nil {
		t.Fatal("Couldn't create temp dir", err)
	}

	configFile := fmt.Sprintf("%s/%s", tmpDir, "delmo.yml")
	err = ioutil.WriteFile(configFile, []byte(config), os.ModePerm)
	if err != nil {
		t.Fatal("Couldn't write config", err)
	}
	t.Logf("Written config file %s", configFile)

	composeFile := fmt.Sprintf("%s/%s", tmpDir, "docker-compose.yml")
	err = ioutil.WriteFile(composeFile, []byte(compose), os.ModePerm)
	if err != nil {
		t.Fatal("Couldn't write compose", err)
	}
	t.Logf("Written compose files %s", composeFile)

	return tmpDir, configFile
}

func TestConfig_Load(t *testing.T) {
	t.Parallel()

	rawConfig := `
suite:
  name: test
  system: docker-compose.yml
  task_service: redis

tasks:
- name: echo
  command: echo hello, world

tests:
- name: simple
  spec:
  - {wait: [echo]}
  - {stop: [redis]}
  - {start: [redis]}`

	compose := `
version: '2'
services:
  redis:
    image: redis
    build: redis`

	tmpDir, configFile := writeFiles("TestSuite_Load", rawConfig, compose, t)
	defer os.Remove(tmpDir)

	config, err := LoadConfig(configFile)
	if err != nil {
		t.Fatal("Load Suite Failed!", err)
	}

	if want, got := "test", config.Suite.Name; want != got {
		t.Errorf("Name not correct. Want: %s, got: %s", want, got)
	}

	if want, got := tmpDir+"/docker-compose.yml", config.Suite.System; want != got {
		t.Errorf("Path to docker-compose.yml not correct. Want: %s, got: %s", want, got)
	}

	if want, got := "redis", config.Tasks["echo"].Service; want != got {
		t.Errorf("Service not set correctly. Want: %d, got: %d", want, got)
	}

	if want, got := "echo hello, world", config.Tasks["echo"].Cmd; want != got {
		t.Errorf("Command not set correctly. Want: %d, got: %d", want, got)
	}

	if want, got := 3, len(config.Tests[0].Spec); want != got {
		t.Errorf("Spec not parsed correctly. Want: %d, got: %d", want, got)
	}

	if want, got := 60*time.Second, config.Tests[0].Spec[0].Timeout; want != got {
		t.Errorf("Timeout not set correctly. Want: %d, got: %d", want, got)
	}
}
