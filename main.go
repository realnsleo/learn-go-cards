package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	fmt.Println([]byte("Arnold Babasa"))
}

func newCard() string {
	return "Five of Diamonds"
}
