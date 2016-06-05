package delmo

type GlobalContext struct {
	DockerHostSyncDir string
}

type TestContext struct {
	DockerHostSyncDir string
	TestName          string
}
