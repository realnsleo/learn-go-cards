package main

func main() {
	cards := newDeckFromFile("my_cards")
	// cards.shuffle()
	// cards.print()

	dealtHand, remainingHand := deal(cards, 4)

	dealtHand.print()
	remainingHand.print()
}
