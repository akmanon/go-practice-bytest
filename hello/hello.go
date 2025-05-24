package main

import "fmt"

const englishPrefix = "Hello, "

func Hello(s string) string {

	if s == "" {
		return "Hello, World!"
	}
	return englishPrefix + s + "!"
}

func main() {
	fmt.Println(Hello("Asis"))
}
