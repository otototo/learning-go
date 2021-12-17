package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const exclamationSufix = "!"
const spanishLang = "Spanish"
const frenchLang = "French"
const defaultName = "World"

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}
	return langPrefix(language) + name + exclamationSufix
}

func langPrefix(language string) (prefix string) {
	switch language {
	case spanishLang:
		prefix = spanishHelloPrefix
	case frenchLang:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
