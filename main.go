package main

import (
	"fmt"

	"github.com/ovguschin90/ggpoker/deck"
)

func main() {
	d := deck.New()

	fmt.Println(d)
	
	card := deck.NewCard(deck.Spades, 1)
	fmt.Println(card)
}
