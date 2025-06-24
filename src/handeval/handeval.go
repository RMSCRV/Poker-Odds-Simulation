package handeval

import (
	"pocket-simulator/deck"
	"sort"
)

// NOTES:
// GET METHODS ALWAYS RETURN A HAND AS WE ALREADY KNOW THERE IS ONE
// GET METHODS ALWAYS RETURN HIGHEST HAND POSSIBLE
type HandRank int

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

func EvaluateHand(cards []deck.Card) ([]deck.Card, HandRank) {

	// Check for Royal Flush
	if hand := CheckRoyalFlush(cards); hand != nil {
		return hand, RoyalFlush
	}

	// Check for Straight Flush
	if hand := CheckStraightFlush(cards); hand != nil {
		return hand, StraightFlush
	}

	// Check for Four of a Kind or Full House
	if hand := CheckFourFull(cards); hand != nil {
		if hand[0].Rank == hand[3].Rank {
			return hand, FourOfAKind
		}
		return hand, FullHouse
	}

	// Check for Flush
	if hand := CheckFlush(cards); hand != nil {
		return hand, Flush
	}

	// Check for Straight
	if hand := CheckStraight(cards); hand != nil {
		return hand, Straight
	}

	// Check for Three of a Kind, Two Pair, or One Pair
	if hand := CheckPairDoubleTriple(cards); hand != nil {

		if hand[0].Rank == hand[2].Rank {
			return hand, ThreeOfAKind
		}

		if hand[2].Rank == hand[3].Rank {
			return hand, TwoPair
		}

		if hand[0].Rank == hand[1].Rank {
			return hand, OnePair
		}

	}

	// Default to High Card
	var hand []deck.Card
	AddHighCardUntilHand(&hand, cards)
	return hand, HighCard
}
func CheckPairDoubleTriple(cards []deck.Card) []deck.Card {
	// Count occurrences of each rank
	counter := make(map[int]int)
	for _, card := range cards {
		counter[card.Rank]++
	}
	// Collect triples and pairs
	var triples []int
	var pairs []int
	for rank, count := range counter {
		switch {
		case count == 3:
			triples = append(triples, rank)
		case count == 2:
			pairs = append(pairs, rank)
		}
	}

	// Sort descending to pick highest ranks
	sort.Sort(sort.Reverse(sort.IntSlice(triples)))
	sort.Sort(sort.Reverse(sort.IntSlice(pairs)))

	if len(triples) == 1 {
		return GetThreeOfKind(cards, triples[0])
	}

	if len(pairs) >= 2 {
		return GetDoublePair(cards, pairs[0], pairs[1])
	}

	if len(pairs) == 1 {
		return GetPair(cards, pairs[0])
	}

	// No full house or four of a kind
	return nil

}

func GetThreeOfKind(cards []deck.Card, cardNumber int) []deck.Card {
	var result []deck.Card
	var remainingCards []deck.Card
	for _, card := range cards {
		if card.Rank == cardNumber {
			result = append(result, card)
		} else {
			remainingCards = append(remainingCards, card)
		}
	}
	// Add the highest remaining card(s) to complete the hand to 5 cards
	AddHighCardUntilHand(&result, remainingCards)

	return result
}

func GetDoublePair(cards []deck.Card, pairOne int, pairTwo int) []deck.Card {
	var result []deck.Card
	var remainingCards []deck.Card
	//get pair one
	for _, card := range cards {
		if card.Rank == pairOne {
			result = append(result, card)
		} else {
			remainingCards = append(remainingCards, card)
		}
	}

	//get pair two
	var finalRemainingCards []deck.Card
	for _, card := range remainingCards {
		if card.Rank == pairTwo {
			result = append(result, card)
		} else {
			finalRemainingCards = append(finalRemainingCards, card)
		}
	}

	AddHighCardUntilHand(&result, finalRemainingCards)

	return result

}

func GetPair(cards []deck.Card, pairOne int) []deck.Card {
	var result []deck.Card
	var remainingCards []deck.Card
	//get pair
	for _, card := range cards {
		if card.Rank == pairOne {
			result = append(result, card)
		} else {
			remainingCards = append(remainingCards, card)
		}
	}

	AddHighCardUntilHand(&result, remainingCards)

	return result

}

// CheckFullOrFour returns the best four-of-a-kind or full house from the given cards, if any.
func CheckFourFull(cards []deck.Card) []deck.Card {
	// Count occurrences of each rank
	counter := make(map[int]int)
	for _, card := range cards {
		counter[card.Rank]++
	}

	// Check for four of a kind
	for rank, count := range counter {
		if count == 4 {
			return GetFourOfKind(cards, rank)
		}
	}

	// Collect triples and pairs
	var triples []int
	var pairs []int
	for rank, count := range counter {
		switch {
		case count == 3:
			triples = append(triples, rank)
		case count == 2:
			pairs = append(pairs, rank)
		}
	}

	// Sort descending to pick highest ranks
	sort.Sort(sort.Reverse(sort.IntSlice(triples)))
	sort.Sort(sort.Reverse(sort.IntSlice(pairs)))

	// Two triples: treat second triple as pair
	if len(triples) == 2 {
		return GetFullHouse(cards, triples[0], triples[1])
	}
	// One triple and at least one pair
	if len(triples) == 1 && len(pairs) >= 1 {
		return GetFullHouse(cards, triples[0], pairs[0])
	}

	// No full house or four of a kind
	return nil
}

func GetFourOfKind(cards []deck.Card, cardNumber int) []deck.Card {

	var result []deck.Card
	var remainingCards []deck.Card
	for _, card := range cards {
		if card.Rank == cardNumber {
			result = append(result, card)
		} else {
			remainingCards = append(remainingCards, card)
		}
	}
	// Add the highest remaining card(s) to complete the hand to 5 cards

	AddHighCardUntilHand(&result, remainingCards)

	return result
}

// it will always be the highest full house possibel thanks to how we are constructing the frequencyCounts
func GetFullHouse(cards []deck.Card, cardNumberThree int, cardNumberTwo int) []deck.Card {
	var result []deck.Card
	var remainingCards []deck.Card
	//get three
	for _, card := range cards {
		if card.Rank == cardNumberThree {
			result = append(result, card)
		} else {
			remainingCards = append(remainingCards, card)
		}
	}

	//get two
	for _, card := range remainingCards {
		if len(result) == 5 {
			return result

		}
		if card.Rank == cardNumberTwo {
			result = append(result, card)
		}
	}
	return result

}

// Complete hand with highest cards
func AddHighCardUntilHand(hand *[]deck.Card, remainingCards []deck.Card) {

	SortDescending(remainingCards)

	for i := 0; i < len(remainingCards) && len(*hand) < 5; i++ {
		*hand = append(*hand, remainingCards[i])
	}

}

func CheckFlush(cards []deck.Card) []deck.Card {
	counter := make(map[deck.Suit]int)
	for _, card := range cards {
		counter[card.Suit]++
		if counter[card.Suit] >= 5 {
			return GetHandFlush(cards, card.Suit)
		}
	}

	return nil
}

func GetHandFlush(cards []deck.Card, suit deck.Suit) []deck.Card {
	var flushCards []deck.Card
	// Collect all cards of the given suit
	for _, card := range cards {
		if card.Suit == suit {
			flushCards = append(flushCards, card)
		}
	}
	// Sort the flush cards in descending order of rank
	SortDescending(flushCards)
	// Return the top 5 cards as the flush hand
	if len(flushCards) >= 5 {
		flushHand := flushCards[:5]
		return flushHand
	}
	return nil
}

func CheckStraight(cards []deck.Card) []deck.Card {
	rankMap := make(map[int]bool)

	for _, card := range cards {
		rankMap[card.Rank] = true
	}

	if rankMap[14] {
		rankMap[1] = true
	}

	var ranks []int

	for rank := range rankMap {
		ranks = append(ranks, rank)
	}

	//sort it

	sort.Sort(sort.Reverse(sort.IntSlice(ranks)))

	consecutive := 1
	prev := -1
	currentBiggest := -1
	//check for five consecutive
	for _, rank := range ranks {
		if prev != -1 && rank == prev-1 {

			consecutive++
			if consecutive >= 5 {
				return GetHandStraight(cards, currentBiggest)
			}
		} else if rank != prev {
			consecutive = 1
			currentBiggest = rank
		}
		prev = rank
	}

	return nil
}

// it will always return the highest straight possible
func GetHandStraight(cards []deck.Card, highestRank int) []deck.Card {
	//delete duplicates by rank
	var result []deck.Card

	SortDescending(cards)
	toPickCard := highestRank
	for _, card := range cards {

		//if it has entered previus if it is never going to enter this if as there is no card Rank 1
		if card.Rank == toPickCard {
			result = append(result, card)
			toPickCard = toPickCard - 1
		}

		if len(result) == 5 {
			return result
		}

	}

	if toPickCard == 1 {
		result = append(result, cards[0])
		return result
	}
	return nil

}

func CheckStraightFlush(cards []deck.Card) []deck.Card {

	flushCards := CheckFlush(cards)
	if flushCards == nil {
		return nil
	}

	return CheckStraight(flushCards)
}

func CheckRoyalFlush(cards []deck.Card) []deck.Card {
	straightFlushCards := CheckStraightFlush(cards)
	if straightFlushCards == nil {
		return nil
	}

	if (straightFlushCards)[0].Rank == 14 {
		return straightFlushCards
	}
	return nil
}

func SortDescending(cards []deck.Card) {
	// Sort the flush cards in descending order of rank
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank > cards[j].Rank
	})
}

// Returns the player id of those who have won plus the winning handType. If more than one player then is a draw
func EvaluateDraw(highestRankPlayersHands map[int][]deck.Card, highestRank HandRank) ([]int, HandRank) {

	//check if only one player

	if len(highestRankPlayersHands) == 1 {
		var keys []int
		for key := range highestRankPlayersHands {
			keys = append(keys, key)
		}
		return keys, highestRank
	}
	// get players with highest rank
	//identify type of draw

	if highestRank == RoyalFlush {
		var keys []int
		for key := range highestRankPlayersHands {
			keys = append(keys, key)
		}
		return keys, RoyalFlush
	}

	if highestRank == StraightFlush {
		//check if someones straight is bigger
		var winnerIds []int
		var currentBiggestInitialCard deck.Card

		for playerId, hand := range highestRankPlayersHands {
			if hand[0].Rank > currentBiggestInitialCard.Rank {
				winnerIds = []int{playerId}
				currentBiggestInitialCard = hand[0]
			} else if hand[0].Rank == currentBiggestInitialCard.Rank {
				winnerIds = append(winnerIds, playerId)
			}

		}

		return winnerIds, StraightFlush
	}

	//kicker is always on fourth position
	if highestRank == FourOfAKind {
		var winnerIds []int
		var currentBiggestInitialCard deck.Card
		var kicker deck.Card

		for playerId, hand := range highestRankPlayersHands {
			if hand[0].Rank > currentBiggestInitialCard.Rank {
				winnerIds = []int{playerId}
				currentBiggestInitialCard = hand[0]
				kicker = hand[4]
			} else if hand[0].Rank == currentBiggestInitialCard.Rank {
				if hand[4].Rank > kicker.Rank {
					winnerIds = []int{playerId}
					kicker = hand[4]
				} else if hand[4].Rank == kicker.Rank {
					winnerIds = append(winnerIds, playerId)

				}
			}

		}

		return winnerIds, FourOfAKind
	}

	//Always the triple comes first and after the pair
	if highestRank == FullHouse {
		var winnerIds []int
		var currentBiggestInitialTriple deck.Card
		var currentBiggestDouble deck.Card

		for playerId, hand := range highestRankPlayersHands {
			if hand[0].Rank > currentBiggestInitialTriple.Rank {
				winnerIds = []int{playerId}
				currentBiggestInitialTriple = hand[0]
				currentBiggestDouble = hand[4]
			} else if hand[0].Rank == currentBiggestInitialTriple.Rank {
				if hand[4].Rank > currentBiggestDouble.Rank {
					winnerIds = []int{playerId}
					currentBiggestDouble = hand[4]
				} else if hand[4].Rank == currentBiggestDouble.Rank {
					winnerIds = append(winnerIds, playerId)

				}
			}

		}

		return winnerIds, FullHouse
	}

	//highest card first always
	if highestRank == Flush {
		var winnerIds []int
		var currentBiggestCards []deck.Card

		for playerId, hand := range highestRankPlayersHands {
			if len(currentBiggestCards) == 0 {
				winnerIds = []int{playerId}
				currentBiggestCards = hand
				continue
			}
			for i := 0; i < len(hand); i++ {
				if hand[i].Rank > currentBiggestCards[i].Rank {
					winnerIds = []int{playerId}
					currentBiggestCards = hand
					break
				} else if hand[i].Rank < currentBiggestCards[i].Rank {
					break
				}
				if i == len(hand)-1 {
					winnerIds = append(winnerIds, playerId)
				}
			}
		}

		return winnerIds, Flush
	}

	//highest card first always
	if highestRank == Straight {
		var winnerIds []int
		var currentBiggestInitialCard deck.Card

		for playerId, hand := range highestRankPlayersHands {
			if hand[0].Rank > currentBiggestInitialCard.Rank {
				winnerIds = []int{playerId}
				currentBiggestInitialCard = hand[0]
			} else if hand[0].Rank == currentBiggestInitialCard.Rank {
				winnerIds = append(winnerIds, playerId)
			}

		}

		return winnerIds, Straight
	}

	//tiple and then kickers from biggest to smalest
	if highestRank == ThreeOfAKind {
		var winnerIds []int
		var currentBiggestTriple deck.Card
		var kickers []deck.Card

		for playerId, hand := range highestRankPlayersHands {
			if hand[0].Rank > currentBiggestTriple.Rank {
				winnerIds = []int{playerId}
				currentBiggestTriple = hand[0]
				kickers = hand[3:5]
			} else if hand[0].Rank == currentBiggestTriple.Rank {
				isTie := true
				for i := 0; i < len(kickers); i++ {
					if hand[3+i].Rank > kickers[i].Rank {
						winnerIds = []int{playerId}
						kickers = hand[3:5]
						isTie = false

						break
					} else if hand[3+i].Rank < kickers[i].Rank {
						isTie = false
						break
					}
				}
				if isTie {
					winnerIds = append(winnerIds, playerId)
				}
			}
		}
		return winnerIds, ThreeOfAKind
	}
	//pairs ordered then kicker
	if highestRank == TwoPair {
		var winnerIds []int
		var currentBiggestFirstPair deck.Card
		var currentBiggestSecondPair deck.Card
		var kicker deck.Card

		for playerId, hand := range highestRankPlayersHands {
			if hand[0].Rank > currentBiggestFirstPair.Rank {
				winnerIds = []int{playerId}
				currentBiggestFirstPair = hand[0]
				currentBiggestSecondPair = hand[2]
				kicker = hand[4]
			} else if hand[0].Rank == currentBiggestFirstPair.Rank {
				if hand[2].Rank > currentBiggestSecondPair.Rank {
					winnerIds = []int{playerId}
					currentBiggestSecondPair = hand[2]
					kicker = hand[4]
				} else if hand[2].Rank == currentBiggestSecondPair.Rank {
					if hand[4].Rank > kicker.Rank {
						winnerIds = []int{playerId}
						kicker = hand[4]
					} else if hand[4].Rank == kicker.Rank {
						winnerIds = append(winnerIds, playerId)
					}
				}
			}
		}
		return winnerIds, TwoPair
	}

	//pair and then kickers from biggest to smalest
	if highestRank == OnePair {
		var winnerIds []int
		var currentBiggestPair deck.Card
		var kickers []deck.Card

		for playerId, hand := range highestRankPlayersHands {
			if hand[0].Rank > currentBiggestPair.Rank {
				winnerIds = []int{playerId}
				currentBiggestPair = hand[0]
				kickers = hand[2:5]
			} else if hand[0].Rank == currentBiggestPair.Rank {
				isTie := true
				for i := 0; i < len(kickers); i++ {
					if hand[2+i].Rank > kickers[i].Rank {
						winnerIds = []int{playerId}
						kickers = hand[2:5]
						isTie = false
						break
					} else if hand[2+i].Rank < kickers[i].Rank {
						isTie = false
						break
					}
				}
				if isTie {
					winnerIds = append(winnerIds, playerId)
				}
			}
		}
		return winnerIds, OnePair
	}

	// If we reach this point, it is a High Card draw
	var winnerIds []int
	var currentBiggestCards []deck.Card

	for playerId, hand := range highestRankPlayersHands {
		if len(currentBiggestCards) == 0 {
			winnerIds = []int{playerId}
			currentBiggestCards = hand
			continue
		}
		for i := 0; i < len(hand); i++ {
			if hand[i].Rank > currentBiggestCards[i].Rank {
				winnerIds = []int{playerId}
				currentBiggestCards = hand
				break
			} else if hand[i].Rank < currentBiggestCards[i].Rank {
				break
			}
			if i == len(hand)-1 {
				winnerIds = append(winnerIds, playerId)
			}
		}
	}
	return winnerIds, HighCard
}
