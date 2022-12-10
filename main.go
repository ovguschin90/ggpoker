package main

import (
	"fmt"

	"github.com/ovguschin90/ggpoker/deck"
)

func main() {
	deck := deck.New()
	fmt.Println(deck)
}
