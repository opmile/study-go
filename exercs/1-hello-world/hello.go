// anatomia de um executavel
package main

import "fmt"

const (
	spanish = "spanish"
	french = "french"

	englishHelloPrefix = "hello, "
	spanishHelloPrefix = "hola, "
	frenchHelloPrefix = "bonjour, "
)

func Hello(name, language string) string {
	if name == "" { name = "world" }
	
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("milena", "french"))
}