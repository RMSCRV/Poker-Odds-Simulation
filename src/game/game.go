package game

import (
	"pocket-simulator/deck"
	"pocket-simulator/handeval"
)

type Player struct {
	ID   int
	Hand []deck.Card
}

type Game struct {
	Players        []Player
	CommunityCards []deck.Card
	Deck           []deck.Card
}

type HandResult struct {
	HandRank handeval.HandRank
	Hand     []deck.Card
}

func NewRound(numPlayers int) *Game {
	d := deck.Shuffle(deck.NewDeck())
	players := make([]Player, numPlayers)
	for i := range players {
		players[i] = Player{ID: i}
	}
	return &Game{Players: players, Deck: d}
}

func NewProbabilisticRound(numPlayers int, playerOneHand []deck.Card) *Game {
	fullDeck := deck.NewDeck()
	remainingDeck := removeCards(fullDeck, playerOneHand)
	shuffledDeck := deck.Shuffle(remainingDeck)

	players := make([]Player, numPlayers)
	for i := range players {
		players[i] = Player{ID: i}
	}
	players[0].Hand = playerOneHand

	return &Game{Players: players, Deck: shuffledDeck}
}

func (g *Game) DealHands() {
	for i := range g.Players {
		g.Players[i].Hand = append(g.Players[i].Hand, g.Deck[:2]...)
		g.Deck = g.Deck[2:]
	}
}

// This method is used for the simulation
func (g *Game) DealHandsToOpponents() {
	for i := 1; i < len(g.Players); i++ {
		g.Players[i].Hand = append(g.Players[i].Hand, g.Deck[:2]...)
		g.Deck = g.Deck[2:]
	}
}

func (g *Game) DealThree() {
	g.Deck = g.Deck[1:]
	g.CommunityCards = append(g.CommunityCards, g.Deck[:3]...)
	g.Deck = g.Deck[3:]
}

func (g *Game) DealOne() {
	g.Deck = g.Deck[1:]
	g.CommunityCards = append(g.CommunityCards, g.Deck[0])
	g.Deck = g.Deck[1:]
}

func (g *Game) Evaluate() map[int]HandResult {
	results := make(map[int]HandResult)
	for _, p := range g.Players {
		handWithCom := append(p.Hand, g.CommunityCards...)

		hand, handType := handeval.EvaluateHand(handWithCom)
		results[p.ID] = HandResult{
			HandRank: handType,
			Hand:     hand,
		}
		// println("Player", p.ID, "rank:", results[p.ID])

		// for _, card := range hand {

		// 	println("Suit: ", card.Suit, " Value: ", card.Rank)

		// }
	}
	return results
}

func (g *Game) PickWinner(results map[int]HandResult, simulation bool) ([]int, handeval.HandRank) {
	var highestRank handeval.HandRank
	highestRankPlayers := []int{}

	// Determine the highest rank and collect players with that rank
	for id, handResult := range results {
		if handResult.HandRank > highestRank {
			highestRank = handResult.HandRank
			highestRankPlayers = []int{id} // Reset to the new highest rank player
		} else if handResult.HandRank == highestRank {
			highestRankPlayers = append(highestRankPlayers, id)
		}
	}

	// Handle single winner case
	if len(highestRankPlayers) == 1 {
		if !simulation {
			println("Player", highestRankPlayers[0], "is the winner with hand type:", highestRank)
		}
		return []int{highestRankPlayers[0]}, highestRank
	}

	// Evaluate draw for players with the highest rank
	highestRankPlayersHands := make(map[int][]deck.Card)
	for _, id := range highestRankPlayers {
		highestRankPlayersHands[id] = append(g.Players[id].Hand, g.CommunityCards...)
	}

	winnersIds, handType := handeval.EvaluateDraw(highestRankPlayersHands, highestRank)

	// Print winners and hand type
	if !simulation {
		if len(winnersIds) == 1 {
			println("Player", winnersIds[0], "is the winner with hand type:", handType)
		} else {
			print("Players")
			for _, id := range winnersIds {
				print(" ", id)
			}
			println(" are the winners with hand type:", handType)
		}
	}

	return winnersIds, handType
}

func removeCards(d []deck.Card, toRemove []deck.Card) []deck.Card {
	var result []deck.Card

	removeMap := make(map[deck.Card]bool)

	for _, card := range toRemove {
		removeMap[card] = true
	}

	for _, card := range d {
		if !removeMap[card] {
			result = append(result, card)
		}
	}
	return result

}
