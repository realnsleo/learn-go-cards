package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishbot struct{}

func main() {
	eb := englishBot{}
	sb := spanishbot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Custom logic for generating an english greeting
	return "Hello!"
}

func (spanishbot) getGreeting() string {
	// Custom logic for generating an english greeting
	return "Hola!"
}
