package bellman

type Signal int

const (
	SIGTERM Signal = iota
	SIGPUB
	SIGSUB
	SIGUNSUB
)

type Event struct {
	Signal  Signal
	Payload interface{}
}
