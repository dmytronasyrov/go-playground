package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	cards = newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
