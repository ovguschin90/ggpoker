package main

import (
	"fmt"

	"github.com/ovguschin90/ggpoker/deck"
)

func main() {
	card := deck.NewCard(deck.Spades, 1)
	fmt.Println(card)
}
