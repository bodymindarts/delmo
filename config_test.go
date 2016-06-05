package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
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

	config := `
suite:
  name: test
  file: docker-compose.yml

tasks:
- name: redis_is_running
  image: redis
  run:
    path: echo
    args: [hello, world]

tests:
- name: simple
  spec:
  - {stop: [redis]}
  - {start: [redis]}`

	compose := `
version: '2'
services:
  redis:
    image: redis
    build: redis`

	tmpDir, configFile := writeFiles("TestSuite_Load", config, compose, t)
	defer os.Remove(tmpDir)

	suite, err := LoadConfig(configFile)
	if err != nil {
		t.Fatal("Load Suite Failed!", err)
	}

	if want, got := "test", suite.System.Name; want != got {
		t.Errorf("Name not correct. Want: %s, got: %s", want, got)
	}

	service, ok := suite.System.Services["redis"]
	if !ok {
		t.Errorf("Compose file not read correctly. Missing service %s", "redis")
	}
	if want, got := "redis", service.Image; want != got {
		t.Errorf("Image not set correctly. Want: %s, got: %s", want, got)
	}
	if want, got := 2, len(suite.Tests[0].Spec); want != got {
		t.Errorf("Spec not parsed correctly. Want: %d, got: %d", want, got)
	}
}
