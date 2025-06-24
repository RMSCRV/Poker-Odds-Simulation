package simulator

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"pocket-simulator/deck"
	"pocket-simulator/game"
	"pocket-simulator/handeval"
	"runtime"
	"strconv"
	"sync"
)

type TwoDecimal float64

func (t TwoDecimal) MarshalJSON() ([]byte, error) {
	s := strconv.FormatFloat(float64(t), 'f', 2, 64)
	return []byte(s), nil
}

type SimulationResult struct {
	WinProbability     TwoDecimal            `json:"winProbability"`
	DrawProbability    TwoDecimal            `json:"drawProbability"`
	LossProbability    TwoDecimal            `json:"lossProbability"`
	WinningHandCounts  map[string]TwoDecimal `json:"winningHandCounts"`
	LosingToHandCounts map[string]TwoDecimal `json:"losingToHandCounts"`
}

func RunSingleRound(numPlayers int) {
	g := game.NewRound(numPlayers)
	g.DealHands()
	g.DealThree()
	g.DealOne()
	g.DealOne()

	results := g.Evaluate()

	g.PickWinner(results, false)

	// for id, rank := range results {
	// 	println("Player ", id, "has hand rank", rank)
	// }
}

// not optimize for paralelisim
func SimulatiorPlayer1RestUnknown() {
	initialDeck := deck.NewDeck()
	deckLength := len(initialDeck)
	maxNumberPlayers := 6
	minNumberPlayers := 7
	numIterationsPerSim := 100000

	results := make(map[int]map[string]SimulationResult)

	for k := minNumberPlayers; k <= maxNumberPlayers; k++ {
		fmt.Printf("Simulating for %d players...\n", k)
		results[k] = make(map[string]SimulationResult)
		for i := 0; i < deckLength; i++ {
			for j := i + 1; j < deckLength; j++ {
				card1 := initialDeck[i]
				card2 := initialDeck[j]
				combination := fmt.Sprintf("%s---%s", card1.String(), card2.String())

				fmt.Printf("Simulating combination: %s for %d players...\n", combination, k)

				// Run simulation and store results
				simulationResult := RunProbabilisticSimulation([]deck.Card{card1, card2}, k, numIterationsPerSim)
				results[k][combination] = simulationResult

				fmt.Printf("Finished simulation for combination: %s for %d players.\n", combination, k)
			}
		}
		fmt.Printf("Finished simulations for %d players.\n", k)
	}

	// Save results to JSON file
	file, err := os.Create("simulation_results.json")
	if err != nil {
		fmt.Printf("Error creating JSON file: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		fmt.Printf("Error encoding results to JSON: %v\n", err)
		return
	}

	fmt.Println("Simulation results saved to simulation_results.json")
}

func SimulatiorPlayer1RestUnknownParallel() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores

	initialDeck := deck.NewDeck()
	deckLength := len(initialDeck)
	maxNumberPlayers := 10
	minNumberPlayers := 7
	numIterationsPerSim := 100000

	results := make(map[int]map[string]SimulationResult)
	var mu sync.Mutex // To safely update the shared `results` map

	type Task struct {
		PlayerCount int
		Card1       deck.Card
		Card2       deck.Card
	}

	tasks := make(chan Task, 1000) // Channel to distribute tasks
	wg := sync.WaitGroup{}         // WaitGroup to wait for all goroutines to finish

	// Worker function
	worker := func() {
		defer wg.Done()
		for task := range tasks {
			combination := fmt.Sprintf("%s---%s", task.Card1.String(), task.Card2.String())
			fmt.Printf("Simulating combination: %s for %d players...\n", combination, task.PlayerCount)

			// Run simulation
			simulationResult := RunProbabilisticSimulation([]deck.Card{task.Card1, task.Card2}, task.PlayerCount, numIterationsPerSim)

			// Safely store the result
			mu.Lock()
			if _, exists := results[task.PlayerCount]; !exists {
				results[task.PlayerCount] = make(map[string]SimulationResult)
			}
			results[task.PlayerCount][combination] = simulationResult
			mu.Unlock()

			fmt.Printf("Finished simulation for combination: %s for %d players.\n", combination, task.PlayerCount)
		}
	}

	// Start worker goroutines
	numWorkers := runtime.NumCPU()
	println(numWorkers)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker()
	}

	// Generate tasks
	for k := minNumberPlayers; k <= maxNumberPlayers; k++ {
		for i := 0; i < deckLength; i++ {
			for j := i + 1; j < deckLength; j++ {
				tasks <- Task{
					PlayerCount: k,
					Card1:       initialDeck[i],
					Card2:       initialDeck[j],
				}
			}
		}
	}
	close(tasks) // Close the task channel to signal workers no more tasks are coming

	// Wait for all workers to finish
	wg.Wait()

	// Save results to JSON file
	file, err := os.Create("simulation_results.json")
	if err != nil {
		fmt.Printf("Error creating JSON file: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		fmt.Printf("Error encoding results to JSON: %v\n", err)
		return
	}

	fmt.Println("Simulation results saved to simulation_results.json")
}

func RunProbabilisticSimulation(playerOneHand []deck.Card, numPlayers int, numSimulations int) SimulationResult {
	playerOneWinsCount := 0
	playerOneLossCount := 0
	playerOneDrawCount := 0
	winningHandCounts := make(map[string]int)
	losingToHandCounts := make(map[string]int)

	for i := 0; i < numSimulations; i++ {
		g := game.NewProbabilisticRound(numPlayers, playerOneHand)
		g.DealHandsToOpponents()
		g.DealThree()
		g.DealOne()
		g.DealOne()

		results := g.Evaluate()
		winners, handType := g.PickWinner(results, true)

		playerOneWins := false
		for _, winnerID := range winners {
			if winnerID == 0 {
				playerOneWins = true
				break
			}
		}
		handTypeStr := Conversor(handType)

		if playerOneWins {
			if len(winners) > 1 {
				playerOneDrawCount++
			} else {
				winningHandCounts[handTypeStr]++
				playerOneWinsCount++
			}
		} else {
			playerOneLossCount++
			losingToHandCounts[handTypeStr]++
		}
	}

	// compute probabilities, rounded to two decimals, then cast to TwoDecimal
	winProbability := TwoDecimal(math.Round(float64(playerOneWinsCount)/float64(numSimulations)*100*100) / 100)
	drawProbability := TwoDecimal(math.Round(float64(playerOneDrawCount)/float64(numSimulations)*100*100) / 100)
	lossProbability := TwoDecimal(math.Round(float64(playerOneLossCount)/float64(numSimulations)*100*100) / 100)

	// convert hand-count maps to probabilities
	winningHandProbabilities := make(map[string]TwoDecimal, len(winningHandCounts))
	for hand, count := range winningHandCounts {
		p := math.Round(float64(count)/float64(numSimulations)*100*100) / 100
		winningHandProbabilities[hand] = TwoDecimal(p)
	}

	losingHandProbabilities := make(map[string]TwoDecimal, len(losingToHandCounts))
	for hand, count := range losingToHandCounts {
		p := math.Round(float64(count)/float64(numSimulations)*100*100) / 100
		losingHandProbabilities[hand] = TwoDecimal(p)
	}

	return SimulationResult{
		WinProbability:     winProbability,
		DrawProbability:    drawProbability,
		LossProbability:    lossProbability,
		WinningHandCounts:  winningHandProbabilities,
		LosingToHandCounts: losingHandProbabilities,
	}
}

func Conversor(h handeval.HandRank) string {
	switch h {
	case handeval.HighCard:
		return "HighCard"
	case handeval.OnePair:
		return "OnePair"
	case handeval.TwoPair:
		return "TwoPair"
	case handeval.ThreeOfAKind:
		return "ThreeOfAKind"
	case handeval.Straight:
		return "Straight"
	case handeval.Flush:
		return "Flush"
	case handeval.FullHouse:
		return "FullHouse"
	case handeval.FourOfAKind:
		return "FourOfAKind"
	case handeval.StraightFlush:
		return "StraightFlush"
	case handeval.RoyalFlush:
		return "RoyalFlush"
	default:
		return "Unknown"
	}
}
