package main

type GlobalContext struct {
	DockerHostSyncDir string
}

type TestContext struct {
	DockerHostSyncDir string
	TestName          string
}
