package parser

import "time"

type Step struct {
	Method   string
	Args     []string
	Location string
	Line     int
}

type CreatedBy struct {
	Method   string
	Location string
	Line     int
}

type Routine struct {
	ID             int
	Duration       time.Duration
	Event          Event
	Stacktrace     []*Step
	CreatedBy      *CreatedBy
	LockedToThread bool
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
