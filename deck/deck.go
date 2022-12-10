package deck

import (
	"math/rand"
	"time"
)

type Deck [52]Card

func New() Deck {
	var (
		nSuits = 4
		nCards = 13
		deck   = [52]Card{}
	)

	x := 0
	for i := 0; i < nSuits; i++ {
		for j := 0; j < nCards; j++ {
			deck[x] = NewCard(Suit(i), j+1)
			x++
		}
	}

	return deck
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
