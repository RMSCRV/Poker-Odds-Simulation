package handeval

import (
	"pocket-simulator/deck"
	"reflect"
	"sort"
	"testing"
)

// Helper function to sort cards for consistent comparison
func sortCards(cards []deck.Card) {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Rank != cards[j].Rank {
			return cards[i].Rank > cards[j].Rank
		}
		return cards[i].Suit > cards[j].Suit
	})
}

func TestCheckRoyalFlush(t *testing.T) {
	testCases := []struct {
		name     string
		cards    []deck.Card
		expected []deck.Card
	}{

		{
			name: "Royal Flush in Spades with 7 cards",
			cards: []deck.Card{
				{Rank: 2, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 14, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades},
			},
		},
		{
			name: "Royal Flush in Spades with 7 cards",
			cards: []deck.Card{
				{Rank: 14, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Spades}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 14, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hand := CheckRoyalFlush(tc.cards)

			if !reflect.DeepEqual(hand, tc.expected) {
				t.Errorf("Expected hand %v, but got %v", tc.expected, hand)
			}
		})
	}
}

func TestFlush(t *testing.T) {
	testCases := []struct {
		name     string
		cards    []deck.Card
		expected []deck.Card
	}{

		{
			name: "Flush in Spades with 7 cards",
			cards: []deck.Card{
				{Rank: 2, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 14, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades},
			},
		},
		{
			name: "Flush in Spades with 7 cards",
			cards: []deck.Card{
				{Rank: 14, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Spades}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 14, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades},
			},
		},
		{
			name: "Flush in Spades with 7 cards",
			cards: []deck.Card{
				{Rank: 2, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 11, Suit: deck.Clubs}, {Rank: 10, Suit: deck.Clubs},
			},
			expected: []deck.Card{
				{Rank: 12, Suit: deck.Clubs}, {Rank: 11, Suit: deck.Clubs}, {Rank: 10, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Clubs}, {Rank: 2, Suit: deck.Clubs},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hand := CheckFlush(tc.cards)

			if !reflect.DeepEqual(hand, tc.expected) {
				t.Errorf("Expected hand %v, but got %v", tc.expected, hand)
			}
		})
	}
}

func TestStraight(t *testing.T) {
	testCases := []struct {
		name     string
		cards    []deck.Card
		expected []deck.Card
	}{
		{
			name: "Straight 6-high with 7 cards",
			cards: []deck.Card{
				{Rank: 6, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Clubs}, {Rank: 9, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 6, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Clubs},
			},
		},

		{
			name: "Straight 9-high with 7 cards",
			cards: []deck.Card{
				{Rank: 9, Suit: deck.Hearts}, {Rank: 8, Suit: deck.Clubs}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 6, Suit: deck.Spades}, {Rank: 5, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 9, Suit: deck.Hearts}, {Rank: 8, Suit: deck.Clubs}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 6, Suit: deck.Spades}, {Rank: 5, Suit: deck.Hearts},
			},
		},
		{
			name: "Straight from ace to 5",
			cards: []deck.Card{
				{Rank: 8, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Clubs}, {Rank: 9, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 5, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Clubs}, {Rank: 14, Suit: deck.Spades},
			},
		},
		{
			name: "Straight another",
			cards: []deck.Card{
				{Rank: 8, Suit: deck.Clubs}, {Rank: 2, Suit: deck.Diamonds}, {Rank: 10, Suit: deck.Spades}, {Rank: 6, Suit: deck.Hearts}, {Rank: 5, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Diamonds}, {Rank: 3, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 6, Suit: deck.Hearts}, {Rank: 5, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Diamonds}, {Rank: 3, Suit: deck.Spades}, {Rank: 2, Suit: deck.Diamonds},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hand := CheckStraight(tc.cards)

			if !reflect.DeepEqual(hand, tc.expected) {
				t.Errorf("Expected hand %v, but got %v", tc.expected, hand)
			}
		})
	}
}

func TestCheckFourFull(t *testing.T) {
	testCases := []struct {
		name     string
		cards    []deck.Card
		expected []deck.Card
	}{
		{
			name: "Four of a Kind (3)",
			cards: []deck.Card{
				{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 14, Suit: deck.Spades}, {Rank: 3, Suit: deck.Diamonds}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades},
			},
		},
		{
			name: "Four of a Kind (Aces)",
			cards: []deck.Card{
				{Rank: 14, Suit: deck.Clubs}, {Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 14, Suit: deck.Clubs}, {Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Hearts}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades},
			},
		},

		{
			name: "Full House (Kings over Queens)",
			cards: []deck.Card{
				{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Clubs}, {Rank: 2, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Hearts}, {Rank: 13, Suit: deck.Spades}, {Rank: 3, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Clubs}, {Rank: 12, Suit: deck.Hearts},
			},
		},
		{
			name: "Full House from two triples (Kings and Queens)",
			cards: []deck.Card{
				{Rank: 2, Suit: deck.Spades}, {Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Hearts}, {Rank: 12, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Diamonds},
			},
			expected: []deck.Card{
				{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Hearts}, {Rank: 12, Suit: deck.Clubs},
			},
		},
		{
			name: "Full House with one triple and two pairs (Kings, Queens, Jacks)",
			cards: []deck.Card{
				{Rank: 13, Suit: deck.Clubs}, {Rank: 11, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Hearts}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Clubs}, {Rank: 11, Suit: deck.Diamonds},
			},
			expected: []deck.Card{
				{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Hearts}, {Rank: 12, Suit: deck.Clubs},
			},
		},
		{
			name: "Full House with one triple and a pairs (Kings, Queens)",
			cards: []deck.Card{
				{Rank: 13, Suit: deck.Clubs}, {Rank: 11, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Hearts}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Clubs}, {Rank: 10, Suit: deck.Diamonds},
			},
			expected: []deck.Card{
				{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Hearts}, {Rank: 12, Suit: deck.Clubs},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hand := CheckFourFull(tc.cards)

			if !reflect.DeepEqual(hand, tc.expected) {
				t.Errorf("Expected hand %v, but got %v", tc.expected, hand)
			}
		})
	}
}

func TestCheckStraightFlush(t *testing.T) {
	testCases := []struct {
		name     string
		cards    []deck.Card
		expected []deck.Card
	}{

		{
			name: "Straight Flush 9-high",
			cards: []deck.Card{
				{Rank: 9, Suit: deck.Spades}, {Rank: 8, Suit: deck.Spades}, {Rank: 7, Suit: deck.Spades}, {Rank: 6, Suit: deck.Spades}, {Rank: 5, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Diamonds},
			},
			expected: []deck.Card{
				{Rank: 9, Suit: deck.Spades}, {Rank: 8, Suit: deck.Spades}, {Rank: 7, Suit: deck.Spades}, {Rank: 6, Suit: deck.Spades}, {Rank: 5, Suit: deck.Spades},
			},
		},
		{
			name: "Ace-low Straight Flush",
			cards: []deck.Card{
				{Rank: 14, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Hearts}, {Rank: 4, Suit: deck.Hearts}, {Rank: 5, Suit: deck.Hearts}, {Rank: 10, Suit: deck.Clubs}, {Rank: 11, Suit: deck.Diamonds},
			},
			expected: []deck.Card{
				{Rank: 5, Suit: deck.Hearts}, {Rank: 4, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Hearts}, {Rank: 14, Suit: deck.Hearts},
			},
		},
		{
			name: "Flush but not a Straight",
			cards: []deck.Card{
				{Rank: 10, Suit: deck.Spades}, {Rank: 8, Suit: deck.Spades}, {Rank: 7, Suit: deck.Spades}, {Rank: 6, Suit: deck.Spades}, {Rank: 2, Suit: deck.Spades}, {Rank: 3, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Diamonds},
			},
			expected: nil,
		},
		{
			name: "Straight but not a Flush",
			cards: []deck.Card{
				{Rank: 9, Suit: deck.Spades}, {Rank: 8, Suit: deck.Clubs}, {Rank: 7, Suit: deck.Spades}, {Rank: 6, Suit: deck.Spades}, {Rank: 5, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Diamonds},
			},
			expected: nil,
		},
		{
			name: "No Straight Flush",
			cards: []deck.Card{
				{Rank: 10, Suit: deck.Spades}, {Rank: 8, Suit: deck.Clubs}, {Rank: 7, Suit: deck.Spades}, {Rank: 6, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Spades}, {Rank: 3, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Diamonds},
			},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hand := CheckStraightFlush(tc.cards)
			sortCards(hand)
			sortCards(tc.expected)
			if !reflect.DeepEqual(hand, tc.expected) {
				t.Errorf("Expected hand %v, but got %v", tc.expected, hand)
			}
		})
	}
}

func TestCheckPairDoubleTriple(t *testing.T) {
	testCases := []struct {
		name     string
		cards    []deck.Card
		expected []deck.Card
	}{
		{
			name: "Three of a Kind",
			cards: []deck.Card{
				{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades}, {Rank: 4, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs},
			},
			expected: []deck.Card{
				{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 14, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds},
			},
		},
		{
			name: "Two Pairs",
			cards: []deck.Card{
				{Rank: 4, Suit: deck.Clubs}, {Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs},
			},
			expected: []deck.Card{
				{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds},
			},
		},
		{
			name: "One Pair",
			cards: []deck.Card{
				{Rank: 3, Suit: deck.Clubs}, {Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Hearts},
			},
			expected: []deck.Card{
				{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Hearts}, {Rank: 4, Suit: deck.Spades},
			},
		},
		{
			name: "No Pairs or Triples",
			cards: []deck.Card{
				{Rank: 3, Suit: deck.Clubs}, {Rank: 14, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Hearts}, {Rank: 9, Suit: deck.Diamonds},
			},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hand := CheckPairDoubleTriple(tc.cards)
			if !reflect.DeepEqual(hand, tc.expected) {
				t.Errorf("Expected hand %v, but got %v", tc.expected, hand)
			}
		})
	}
}

func TestEvaluateHand(t *testing.T) {
	testCases := []struct {
		name     string
		cards    []deck.Card
		expected []deck.Card
		category HandRank
	}{
		{
			name: "Royal Flush",
			cards: []deck.Card{
				{Rank: 14, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Diamonds},
			},
			expected: []deck.Card{
				{Rank: 14, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades},
			},
			category: RoyalFlush,
		},
		{
			name: "Straight Flush",
			cards: []deck.Card{
				{Rank: 9, Suit: deck.Hearts}, {Rank: 8, Suit: deck.Hearts}, {Rank: 7, Suit: deck.Hearts}, {Rank: 6, Suit: deck.Hearts}, {Rank: 5, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Diamonds},
			},
			expected: []deck.Card{
				{Rank: 9, Suit: deck.Hearts}, {Rank: 8, Suit: deck.Hearts}, {Rank: 7, Suit: deck.Hearts}, {Rank: 6, Suit: deck.Hearts}, {Rank: 5, Suit: deck.Hearts},
			},
			category: StraightFlush,
		},
		{
			name: "Four of a Kind",
			cards: []deck.Card{
				{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades},
			},
			category: FourOfAKind,
		},
		{
			name: "Full House",
			cards: []deck.Card{
				{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Hearts}, {Rank: 12, Suit: deck.Clubs}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Hearts}, {Rank: 12, Suit: deck.Clubs},
			},
			category: FullHouse,
		},
		{
			name: "Flush",
			cards: []deck.Card{
				{Rank: 7, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 5, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Diamonds},
			},
			expected: []deck.Card{
				{Rank: 13, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades}, {Rank: 7, Suit: deck.Spades}, {Rank: 5, Suit: deck.Spades},
			},
			category: Flush,
		},
		{
			name: "Straight",
			cards: []deck.Card{
				{Rank: 6, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Clubs}, {Rank: 9, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades},
			},
			expected: []deck.Card{
				{Rank: 6, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Clubs},
			},
			category: Straight,
		},
		{
			name: "Three of a Kind",
			cards: []deck.Card{
				{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 14, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs},
			},
			expected: []deck.Card{
				{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 14, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds},
			},
			category: ThreeOfAKind,
		},
		{
			name: "Two Pair",
			cards: []deck.Card{
				{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs},
			},
			expected: []deck.Card{
				{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds},
			},
			category: TwoPair,
		},
		{
			name: "One Pair",
			cards: []deck.Card{
				{Rank: 13, Suit: deck.Spades}, {Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Hearts},
			},
			expected: []deck.Card{
				{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 13, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Hearts},
			},
			category: OnePair,
		},
		{
			name: "High Card",
			cards: []deck.Card{
				{Rank: 3, Suit: deck.Clubs}, {Rank: 12, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 2, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Hearts}, {Rank: 9, Suit: deck.Diamonds},
			},
			expected: []deck.Card{
				{Rank: 12, Suit: deck.Spades}, {Rank: 9, Suit: deck.Diamonds}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Hearts}, {Rank: 4, Suit: deck.Spades},
			},
			category: HighCard,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hand, category := EvaluateHand(tc.cards)
			if !reflect.DeepEqual(hand, tc.expected) || category != tc.category {
				t.Errorf("Expected hand %v and category %v, but got hand %v and category %v", tc.expected, tc.category, hand, category)
			}
		})
	}
}

func TestEvaluateDraw(t *testing.T) {
	testCases := []struct {
		name                    string
		highestRank             HandRank
		highestRankPlayersHands map[int][]deck.Card
		expectedWinners         []int
		expectedRank            HandRank
	}{
		{
			name:        "Royal Flush with 1 players",
			highestRank: RoyalFlush,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 14, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades}},
			},
			expectedWinners: []int{1},
			expectedRank:    RoyalFlush,
		},
		{
			name:        "Royal Flush with 2 players",
			highestRank: RoyalFlush,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 14, Suit: deck.Spades}, {Rank: 13, Suit: deck.Spades}, {Rank: 12, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades}},
				2: {{Rank: 14, Suit: deck.Hearts}, {Rank: 13, Suit: deck.Hearts}, {Rank: 12, Suit: deck.Hearts}, {Rank: 11, Suit: deck.Hearts}, {Rank: 10, Suit: deck.Hearts}},
			},
			expectedWinners: []int{1, 2},
			expectedRank:    RoyalFlush,
		},
		{
			name:        "Straight Flush with 3 players",
			highestRank: StraightFlush,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 9, Suit: deck.Spades}, {Rank: 8, Suit: deck.Spades}, {Rank: 7, Suit: deck.Spades}, {Rank: 6, Suit: deck.Spades}, {Rank: 5, Suit: deck.Spades}},
				2: {{Rank: 10, Suit: deck.Hearts}, {Rank: 9, Suit: deck.Hearts}, {Rank: 8, Suit: deck.Hearts}, {Rank: 7, Suit: deck.Hearts}, {Rank: 6, Suit: deck.Hearts}},
				3: {{Rank: 9, Suit: deck.Clubs}, {Rank: 8, Suit: deck.Clubs}, {Rank: 7, Suit: deck.Clubs}, {Rank: 6, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Clubs}},
			},
			expectedWinners: []int{2},
			expectedRank:    StraightFlush,
		},
		{
			name:        "Straight Flush with edge 1",
			highestRank: StraightFlush,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 10, Suit: deck.Hearts}, {Rank: 9, Suit: deck.Hearts}, {Rank: 8, Suit: deck.Hearts}, {Rank: 7, Suit: deck.Hearts}, {Rank: 6, Suit: deck.Hearts}},
				2: {{Rank: 10, Suit: deck.Hearts}, {Rank: 9, Suit: deck.Hearts}, {Rank: 8, Suit: deck.Hearts}, {Rank: 7, Suit: deck.Hearts}, {Rank: 6, Suit: deck.Hearts}},
			},
			expectedWinners: []int{1, 2},
			expectedRank:    StraightFlush,
		},
		{
			name:        "Four of a Kind with 4 players",
			highestRank: FourOfAKind,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades}},
				2: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Diamonds}, {Rank: 13, Suit: deck.Spades}},
				3: {{Rank: 4, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Spades}, {Rank: 4, Suit: deck.Hearts}, {Rank: 4, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Spades}},
				4: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Hearts}},
			},
			expectedWinners: []int{3},
			expectedRank:    FourOfAKind,
		},
		{
			name:        "Four of a Kind with edge 1",
			highestRank: FourOfAKind,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades}},
				2: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades}},
			},
			expectedWinners: []int{1, 2},
			expectedRank:    FourOfAKind,
		},
		{
			name:        "Four of a Kind with edge 2",
			highestRank: FourOfAKind,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Diamonds}, {Rank: 2, Suit: deck.Spades}},
				2: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Diamonds}, {Rank: 14, Suit: deck.Spades}},
			},
			expectedWinners: []int{2},
			expectedRank:    FourOfAKind,
		},
		{
			name:        "Full House with 5 players",
			highestRank: FullHouse,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Hearts}, {Rank: 12, Suit: deck.Clubs}},
				2: {{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 11, Suit: deck.Hearts}, {Rank: 11, Suit: deck.Clubs}},
				3: {{Rank: 12, Suit: deck.Clubs}, {Rank: 12, Suit: deck.Spades}, {Rank: 12, Suit: deck.Diamonds}, {Rank: 11, Suit: deck.Hearts}, {Rank: 11, Suit: deck.Clubs}},
				4: {{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Hearts}, {Rank: 12, Suit: deck.Clubs}},
				5: {{Rank: 14, Suit: deck.Clubs}, {Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 13, Suit: deck.Hearts}, {Rank: 13, Suit: deck.Clubs}},
			},
			expectedWinners: []int{5},
			expectedRank:    FullHouse,
		},
		{
			name:        "Full House with edge 1",
			highestRank: FullHouse,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 11, Suit: deck.Hearts}, {Rank: 11, Suit: deck.Clubs}},

				2: {{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Hearts}, {Rank: 12, Suit: deck.Clubs}},
			},
			expectedWinners: []int{2},
			expectedRank:    FullHouse,
		},
		{
			name:        "Full House with edge 2",
			highestRank: FullHouse,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 11, Suit: deck.Hearts}, {Rank: 11, Suit: deck.Clubs}},
				2: {{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 11, Suit: deck.Hearts}, {Rank: 11, Suit: deck.Clubs}},
			},
			expectedWinners: []int{1, 2},
			expectedRank:    FullHouse,
		},
		{
			name:        "Flush with 3 players",
			highestRank: Flush,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 13, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades}, {Rank: 7, Suit: deck.Spades}, {Rank: 5, Suit: deck.Spades}},
				2: {{Rank: 14, Suit: deck.Hearts}, {Rank: 12, Suit: deck.Hearts}, {Rank: 10, Suit: deck.Hearts}, {Rank: 8, Suit: deck.Hearts}, {Rank: 6, Suit: deck.Hearts}},
				3: {{Rank: 13, Suit: deck.Clubs}, {Rank: 11, Suit: deck.Clubs}, {Rank: 9, Suit: deck.Clubs}, {Rank: 7, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Clubs}},
			},
			expectedWinners: []int{2},
			expectedRank:    Flush,
		},
		{
			name:        "Flush edge 1",
			highestRank: Flush,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 13, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades}, {Rank: 7, Suit: deck.Spades}, {Rank: 5, Suit: deck.Spades}},
				2: {{Rank: 13, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades}, {Rank: 7, Suit: deck.Spades}, {Rank: 5, Suit: deck.Spades}},
			},
			expectedWinners: []int{1, 2},
			expectedRank:    Flush,
		},
		{
			name:        "Flush edge 2",
			highestRank: Flush,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 13, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades}, {Rank: 7, Suit: deck.Spades}, {Rank: 5, Suit: deck.Spades}},
				2: {{Rank: 13, Suit: deck.Spades}, {Rank: 11, Suit: deck.Spades}, {Rank: 10, Suit: deck.Spades}, {Rank: 7, Suit: deck.Spades}, {Rank: 4, Suit: deck.Spades}},
			},
			expectedWinners: []int{1},
			expectedRank:    Flush,
		},
		{
			name:        "Straight with 2 players",
			highestRank: Straight,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 6, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Clubs}},
				2: {{Rank: 7, Suit: deck.Hearts}, {Rank: 6, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Spades}, {Rank: 4, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Diamonds}},
			},
			expectedWinners: []int{2},
			expectedRank:    Straight,
		},
		{
			name:        "Straight with 2 players edge 1",
			highestRank: Straight,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 6, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Clubs}},
				2: {{Rank: 6, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 2, Suit: deck.Clubs}},
			},
			expectedWinners: []int{1, 2},
			expectedRank:    Straight,
		},
		{
			name:        "Three of a Kind with 4 players",
			highestRank: ThreeOfAKind,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 14, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}},
				2: {{Rank: 4, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Spades}, {Rank: 4, Suit: deck.Hearts}, {Rank: 13, Suit: deck.Spades}, {Rank: 6, Suit: deck.Diamonds}},
				3: {{Rank: 5, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Spades}, {Rank: 5, Suit: deck.Hearts}, {Rank: 12, Suit: deck.Spades}, {Rank: 8, Suit: deck.Diamonds}},
				4: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 14, Suit: deck.Hearts}, {Rank: 9, Suit: deck.Diamonds}},
			},
			expectedWinners: []int{3},
			expectedRank:    ThreeOfAKind,
		},
		{
			name:        "Three of a Kind with 2 players edge 1",
			highestRank: ThreeOfAKind,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 14, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}},
				2: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 14, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}},
			},
			expectedWinners: []int{1, 2},
			expectedRank:    ThreeOfAKind,
		},
		{
			name:        "Three of a Kind with 2 players edge 2",
			highestRank: ThreeOfAKind,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 14, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}},
				2: {{Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 3, Suit: deck.Hearts}, {Rank: 14, Suit: deck.Spades}, {Rank: 6, Suit: deck.Diamonds}},
			},
			expectedWinners: []int{1},
			expectedRank:    ThreeOfAKind,
		},
		{
			name:        "Two Pair with 3 players",
			highestRank: TwoPair,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}},
				2: {{Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Spades}, {Rank: 8, Suit: deck.Diamonds}},
				3: {{Rank: 12, Suit: deck.Spades}, {Rank: 12, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Spades}, {Rank: 9, Suit: deck.Diamonds}},
			},
			expectedWinners: []int{1},
			expectedRank:    TwoPair,
		},
		{
			name:        "Two Pair with 3 players edge 1",
			highestRank: TwoPair,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}},
				2: {{Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Spades}, {Rank: 8, Suit: deck.Diamonds}},
				3: {{Rank: 12, Suit: deck.Spades}, {Rank: 12, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Spades}, {Rank: 9, Suit: deck.Diamonds}},
				4: {{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}},
			},
			expectedWinners: []int{1, 4},
			expectedRank:    TwoPair,
		},
		{
			name:        "Two Pair with 3 players edge 2",
			highestRank: TwoPair,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}},
				2: {{Rank: 13, Suit: deck.Spades}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Spades}, {Rank: 8, Suit: deck.Diamonds}},
				3: {{Rank: 12, Suit: deck.Spades}, {Rank: 12, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Clubs}, {Rank: 5, Suit: deck.Spades}, {Rank: 9, Suit: deck.Diamonds}},
				4: {{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 3, Suit: deck.Clubs}, {Rank: 3, Suit: deck.Spades}, {Rank: 6, Suit: deck.Diamonds}},
			},
			expectedWinners: []int{1},
			expectedRank:    TwoPair,
		},
		{
			name:        "One Pair with 3 players",
			highestRank: OnePair,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 13, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Hearts}},
				2: {{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Spades}, {Rank: 8, Suit: deck.Diamonds}, {Rank: 6, Suit: deck.Hearts}},
				3: {{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 13, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Hearts}},
			},
			expectedWinners: []int{1, 3},
			expectedRank:    OnePair,
		},
		{
			name:        "One Pair with 2 players edge 1",
			highestRank: OnePair,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 13, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Hearts}},
				2: {{Rank: 13, Suit: deck.Clubs}, {Rank: 13, Suit: deck.Diamonds}, {Rank: 12, Suit: deck.Spades}, {Rank: 8, Suit: deck.Diamonds}, {Rank: 6, Suit: deck.Hearts}},
			},
			expectedWinners: []int{1},
			expectedRank:    OnePair,
		},
		{
			name:        "One Pair with 2 players edge 2",
			highestRank: OnePair,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 13, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Hearts}},
				2: {{Rank: 14, Suit: deck.Spades}, {Rank: 14, Suit: deck.Diamonds}, {Rank: 13, Suit: deck.Spades}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 4, Suit: deck.Hearts}},
			},
			expectedWinners: []int{1},
			expectedRank:    OnePair,
		},
		{
			name:        "High Card with 3 players",
			highestRank: HighCard,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 12, Suit: deck.Spades}, {Rank: 9, Suit: deck.Diamonds}, {Rank: 7, Suit: deck.Diamonds}, {Rank: 5, Suit: deck.Hearts}, {Rank: 4, Suit: deck.Spades}},
				2: {{Rank: 13, Suit: deck.Clubs}, {Rank: 10, Suit: deck.Diamonds}, {Rank: 8, Suit: deck.Diamonds}, {Rank: 6, Suit: deck.Hearts}, {Rank: 3, Suit: deck.Spades}},
				3: {{Rank: 14, Suit: deck.Hearts}, {Rank: 11, Suit: deck.Clubs}, {Rank: 9, Suit: deck.Spades}, {Rank: 7, Suit: deck.Clubs}, {Rank: 2, Suit: deck.Diamonds}},
			},
			expectedWinners: []int{3},
			expectedRank:    HighCard,
		},
		{
			name:        "High Card with edge 1",
			highestRank: HighCard,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 14, Suit: deck.Hearts}, {Rank: 11, Suit: deck.Clubs}, {Rank: 9, Suit: deck.Spades}, {Rank: 7, Suit: deck.Clubs}, {Rank: 2, Suit: deck.Diamonds}},
				2: {{Rank: 14, Suit: deck.Hearts}, {Rank: 11, Suit: deck.Clubs}, {Rank: 9, Suit: deck.Spades}, {Rank: 7, Suit: deck.Clubs}, {Rank: 2, Suit: deck.Diamonds}},
			},
			expectedWinners: []int{1, 2},
			expectedRank:    HighCard,
		},
		{
			name:        "High Card with edge 2",
			highestRank: HighCard,
			highestRankPlayersHands: map[int][]deck.Card{
				1: {{Rank: 14, Suit: deck.Hearts}, {Rank: 11, Suit: deck.Clubs}, {Rank: 9, Suit: deck.Spades}, {Rank: 7, Suit: deck.Clubs}, {Rank: 2, Suit: deck.Diamonds}},
				2: {{Rank: 14, Suit: deck.Hearts}, {Rank: 11, Suit: deck.Clubs}, {Rank: 9, Suit: deck.Spades}, {Rank: 7, Suit: deck.Clubs}, {Rank: 4, Suit: deck.Diamonds}},
			},
			expectedWinners: []int{2},
			expectedRank:    HighCard,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			winners, rank := EvaluateDraw(tc.highestRankPlayersHands, tc.highestRank)
			sort.Ints(winners)
			sort.Ints(tc.expectedWinners)
			if !reflect.DeepEqual(winners, tc.expectedWinners) || rank != tc.expectedRank {
				t.Errorf("Expected winners %v and rank %v, but got winners %v and rank %v", tc.expectedWinners, tc.expectedRank, winners, rank)
			}
		})
	}
}
