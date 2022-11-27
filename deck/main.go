package deck

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Suit int

func (s Suit) String() string {
	switch s {
	case Spades:
		return "SPADES"
	case Hearts:
		return "HEARTS"
	case Dimonds:
		return "DIAMONDS"
	case Clubs:
		return "CLUBS"
	default:
		panic("invalid card suit")
	}
}

const (
	Spades Suit = iota
	Hearts
	Dimonds
	Clubs
)

type Card struct {
	suit  Suit
	value int
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s %s", c.cardName(), c.suit, suitToUnicode(c.suit))
}

func (c Card) cardName() string {
	var value string
	switch c.value {
	case 1:
		value = "ACE"
	case 11:
		value = "JACK"
	case 12:
		value = "QUEEN"
	case 13:
		value = "KING"
	default:
		value = strconv.Itoa(c.value)
	}
	return value
}

func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("the value of the card cannot be higher then 13")
	}

	return Card{
		suit:  s,
		value: v,
	}
}

type Deck [52]Card

func New() Deck {
	var (
		nSuits = 4
		nCards = 13
		d      = [52]Card{}
	)

	x := 0
	for i := 0; i < nSuits; i++ {
		for j := 0; j < nCards; j++ {
			d[x] = NewCard(Suit(i), j+1)
			x++
		}
	}

	return d
}
func Shuffle(d Deck) Deck {
	for i := 0; i < len(d); i++ {
		r := rand.Intn(i + 1)
		if r != i {
			d[i], d[r] = d[r], d[i]
		}
	}

	return d
}

func suitToUnicode(s Suit) string {
	switch s {
	case Spades:
		return "♠"
	case Hearts:
		return "♥"
	case Dimonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		panic("invalid card suit")
	}
}
