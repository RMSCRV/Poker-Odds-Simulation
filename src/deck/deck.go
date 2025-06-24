package deck

import (
	"fmt"
	"math/rand"
)

type Suit string

const (
	Spades   Suit = "S"
	Hearts   Suit = "H"
	Diamonds Suit = "D"
	Clubs    Suit = "C"
)

type Card struct {
	Rank int
	Suit Suit
}

func (c Card) String() string {
	rankStr := ""
	switch c.Rank {
	case 11:
		rankStr = "J"
	case 12:
		rankStr = "Q"
	case 13:
		rankStr = "K"
	case 14:
		rankStr = "A"
	default:
		rankStr = fmt.Sprintf("%d", c.Rank)
	}
	return rankStr + "_" + string(c.Suit)
}

func NewDeck() []Card {
	var deck []Card
	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		for rank := 2; rank <= 14; rank++ {
			deck = append(deck, Card{Rank: rank, Suit: suit})
		}
	}
	return deck
}

func Shuffle(deck []Card) []Card {

	shuffled := make([]Card, len(deck))
	copy(shuffled, deck)
	rand.Shuffle(len(shuffled), func(i int, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled

}
