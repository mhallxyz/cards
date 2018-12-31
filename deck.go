package main

import "fmt"

//Create a new type "deck" which is a slice of strings

type deck []string

func newDeck() deck {
	// cards := deck{"Ace of Spades", "Five of Diamonds", "Six of Spades", "Queen of Hearts"}
	cards := deck{}
	cardSuits := []string{"Diamonds", "Spades", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for i, suit := range cardSuits {
		for j, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
