package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	//custom logic
	return "Hi there!!!"
}

func (spanishBot) getGreeting() string {
	//custom logic
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {

	sb := spanishBot{}
	eg := englishBot{}

	printGreeting(sb)
	printGreeting(eg)

	fmt.Println(sb)
	fmt.Println(eg)
}
