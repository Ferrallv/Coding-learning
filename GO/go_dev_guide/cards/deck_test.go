package main

import (
	"testing"
	"fmt"
	"os"
)

func BenchmarkNewDeck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDeck()
	}
}

func ExampleNewDeck() {
	c := NewDeck()
	fmt.Println(c[0])
	// Output:
	// Ace of Spades
}

func TestNewDeck(t *testing.T) {
	d := NewDeck()	

	if len(d) != 52 {
		t.Error("Expected deck length of 52, got", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades', got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	
	deck := NewDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 entries, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}