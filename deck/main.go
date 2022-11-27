package deck

type Suit int

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

func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("the value of the card cannot be higher then 13")
	}
	return Card{
		suit:  s,
		value: v,
	}
}
