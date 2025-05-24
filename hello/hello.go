package main

import "fmt"

const (
	englishPrefix = "Hello, "
	hindiPrefix   = "Namaste, "
	spanishPrefix = "Hola, "
	hindi         = "Hindi"
	spanish       = "Spanish"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetPrefix(lang) + name
}
func greetPrefix(lang string) (prefix string) {
	switch lang {
	case hindi:
		prefix = hindiPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return

}

func main() {
	fmt.Println(Hello("Asis", "hindi"))
}
