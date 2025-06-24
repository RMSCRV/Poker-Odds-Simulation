package main

import (
	"fmt"
	"pocket-simulator/simulator"
	"time"
)

func main() {

	// for i := 0; i < 100; i++ {
	// 	simulator.RunSingleRound(10)

	// }

	start := time.Now()
	simulator.SimulatiorPlayer1RestUnknownParallel()
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}
