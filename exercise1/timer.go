package main

import (
	"time"
)

// Timer type
type Timer struct {
	Duration     time.Duration
	TimerChannel chan bool
}

// NewTimer Creates a new timer object
func NewTimer(duration time.Duration) *Timer {
	return &Timer{
		Duration:     duration,
		TimerChannel: make(chan bool),
	}
}

// Run runs the timer
func (t *Timer) Run() {
	timer := time.NewTimer(t.Duration)
	<-timer.C
	t.TimerChannel <- true
}
