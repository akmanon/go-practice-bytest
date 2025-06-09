package main

import (
	"os"

	"github.com/akmanon/go-practice-bytest/mocking"
)

func main() {
	sleeper := mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, &sleeper)
}
