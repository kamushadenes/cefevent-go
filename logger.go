package cefevent

type Logger interface {
	Connect() error
	Send([]byte) error
	Close() error
	IsConnected() bool
}
