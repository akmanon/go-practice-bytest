package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	s := SpySleeper{}
	Countdown(buffer, &s)

	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if s.Calls != 3 {
		t.Errorf("got %d want %d", s.Calls, 3)
	}
}
