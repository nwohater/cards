package main

import "strings"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// new deck function
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// receiver print function
func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}

// receiver deal function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// receiver toString function
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
