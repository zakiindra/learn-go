package main

import "fmt"

const indonesian = "Indonesian"
const french = "French"
const englishPrefix = "Hello, "
const indonesianPrefix = "Halo, "
const frenchPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := GreetingPrefix(language)

	return prefix + name
}

func GreetingPrefix(language string) (prefix string) {
	switch language {
	case indonesian:
		prefix = indonesianPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Dude", "Indonesian"))
}
