package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type deck []string 


// NewDeck returns a type deck which is a slice of strings
// describing the cards in a 52 playing card deck.
func NewDeck() deck {
	cards := deck{}

	suits := []string{" of Spades", " of Clubs", " of Hearts", " of Diamonds"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+suit)
		}
	} 

	return cards
}

// Print prints out the elements of the deck
func (d deck) Print() {
	for i, v := range d {
	fmt.Println(i, v)	
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) 
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}