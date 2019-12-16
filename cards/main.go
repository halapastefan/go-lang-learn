package main

import (
	"fmt"
)

func main() {

	cards := newDeck()
	cards.saveToFile("cards.txt")

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	fmt.Println(hand.toString())

	err := hand.saveToFile("hand.txt")

	if err != nil {
		fmt.Println(err)
	}

	deck := newDeckFromFile("cards.txt")
	fmt.Println(deck.toString())

	deck.shuffleDeck()
	deck.print()

}
