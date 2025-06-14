package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("Say Hello to people in english", func(t *testing.T) {
		got := Hello("Asis", "english")
		want := "Hello, Asis"
		assertCorrectMessage(t, got, want)
	})
	t.Run("If no people, default to world", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello to People in Hindi", func(t *testing.T) {
		got := Hello("Asis", "Hindi")
		want := "Namaste, Asis"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
