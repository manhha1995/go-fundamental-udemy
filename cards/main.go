package main

func main() {
	cards := newDeck()
	// cards.print()
	// cards.saveToFile("my cards")
	cards.shuffle()
	cards.print()
}
