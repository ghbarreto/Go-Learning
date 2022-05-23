package main

import "fmt"

var cards string

func main() {
	cards := []string{"Ace of diamonds, ", newCard()}
	cards = append(cards, "Six of spades")

	// iterating over the list of cards
	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds,"
}
