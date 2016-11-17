package parser

type Step struct {
	Method   string
	Args     []string
	Location string
	Line     int
}

type Routine struct {
	ID         int
	Event      Event
	Stacktrace []*Step
}

type Event int

const (
	EventIOWait Event = iota
	EventRunning
	EventChanReceive
	EventChanSend
	EventSyscall
	EventSelect
)
