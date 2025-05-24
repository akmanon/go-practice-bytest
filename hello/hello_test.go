package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("Say Hello to people", func(t *testing.T) {
		got := Hello("Asis")
		want := "Hello, Asis!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("If no people, default to world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
