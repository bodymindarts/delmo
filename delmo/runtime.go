package delmo

type Runtime interface {
	Start() error
	Stop() error
	Output() ([]byte, error)
	Cleanup() error
}
