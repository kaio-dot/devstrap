package ui

import (
	"fmt"
	"time"
)

type Spinner struct {
	msg  string
	done chan struct{}
}

func NewSpinner(msg string) *Spinner {
	return &Spinner{
		msg:  msg,
		done: make(chan struct{}),
	}
}

func (s *Spinner) Start() {
	go func() {
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0

		for {
			select {
			case <-s.done:
				return
			default:
				fmt.Printf("\r%s %s", frames[i], s.msg)
				time.Sleep(100 * time.Millisecond)
				i = (i + 1) % len(frames)
			}
		}
	}()
}

func (s *Spinner) Stop() {
	close(s.done)
	fmt.Printf("\r✔ %s\n", s.msg)
}

func (s *Spinner) StopError(err error) {
	close(s.done)
	fmt.Printf("\r✖ %s: %v\n", s.msg, err)
}
