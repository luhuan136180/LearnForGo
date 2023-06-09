package main

import "fmt"

const spanish = "Spanish"
const helloPrefix = "Hello,"
const spanishHelloPrefix = "Hola,"
const french = "French"
const frenchHelloPrefix = "Bonjour,"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case french:
		return frenchHelloPrefix + name
	case spanish:
		return spanishHelloPrefix + name
	default:
		return helloPrefix + name
	}

}

func main() {
	fmt.Println(Hello("world", ""))
}
