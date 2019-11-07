package main

import "fmt"

// Create new type of deck which is a slice of strings

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Cubes"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}