package parser

type Step struct {
	Method   string
	Args     []string
	Location string
	Line     int
}

type CreatedBy struct {
	Method string
	Location string
	Line int
}

type Routine struct {
	ID         int
	Event      Event
	Stacktrace []*Step
	CreatedBy *CreatedBy
}

type Event int

const (
	EventIOWait Event = iota
	EventRunning
	EventChanReceive
	EventChanSend
	EventSyscall
	EventSelect
	EventSleep
	EventSemAcquire
	EventRunnable
)
