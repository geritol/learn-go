package main

import "fmt"

const helloEnglish = "Hello, "
const helloSpanish = "Hola, "
const helloFrench = "Bonjour, "

func greetingPrefix(language string) string {
	switch language {
	case "french":
		return helloFrench
	case "spanish":
		return helloSpanish
	default:
		return helloEnglish
	}
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", "english"))
}
