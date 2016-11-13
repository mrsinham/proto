package parser

type Frame struct {
}

type Routine struct {
	ID         int
	Event      Event
	Stacktrace []*Frame
}

type Event int

const (
	EventRunning Event = iota
	EventIOWait
	EventChanReceive
	EventChanSend
	EventSyscall
)
