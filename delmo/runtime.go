package delmo

type Runtime interface {
	StartAll() error
	StopAll() error
	Output() ([]byte, error)
	Cleanup() error
}
