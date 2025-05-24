package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Asis")
	want := "Hello, Asis!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
