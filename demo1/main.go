package main

import "fmt"

func main() {

	var x = 8
	var s = 8

	var y = x + s

	fmt.Println("Hi there", y)

	cards := [] string {"a", "b", "c", newCard()}

	cards = append(cards, newCard())

	print(cards)

	fmt.Println(cards)
}

func newCard() string {
	return "Ace of spades"
}