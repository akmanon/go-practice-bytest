package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second * 1)
}

const (
	CountdownStart = 3
	FinalMessage   = "Go!"
)

func Countdown(out io.Writer, s Sleeper) {
	for i := CountdownStart; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)
		s.Sleep()
	}
	fmt.Fprint(out, FinalMessage)
}
