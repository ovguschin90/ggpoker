package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ovguschin90/ggpoker/deck"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	d := deck.Shuffle(deck.New())

	fmt.Println(d)

	card := deck.NewCard(deck.Spades, 1)
	fmt.Println(card)
}
