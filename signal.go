package server

const (
	SignalError = 1 << iota
	SignalStarting
	SignalNotLiveYet
	SignalListening
	SignalListeningTimeout
	SignalShutdown
)

type Signal interface {
	GetCode() int
	GetError() error
	GetString() string
}

type ServerSignal struct {
	code int
	data interface{}
}

func (serverSignal ServerSignal) GetCode() int {
	return serverSignal.code
}

func (serverSignal ServerSignal) GetError() error {
	return serverSignal.data.(error)
}

func (serverSignal ServerSignal) GetString() string {
	switch serverSignal.code {
	case SignalError:
		return "ERROR"
	case SignalListening:
		return "LISTENING"
	case SignalListeningTimeout:
		return "LISTENING_TIMEOUT"
	case SignalNotLiveYet:
		return "WAITING"
	case SignalShutdown:
		return "SHUTDOWN"
	case SignalStarting:
		return "STARTING"
	}
	return "UNKNOWN"
}
