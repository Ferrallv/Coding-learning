package main


func main() {
	// cards := newDeck()
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// cards.saveToFile("my_cards")

	cards := NewDeck()
	cards.shuffle()
	cards.print()
}

func NewCard() deck {
	return deck{"Five of Diamonds", "Ace of Hearts", "Jack of Clubs"}
}