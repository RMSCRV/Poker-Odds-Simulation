# Poker Simulator

This project is a poker simulation tool designed to calculate probabilities for various poker hands in games with 2 to 10 players. The simulations are implemented in Go and include both single-threaded and parallelized approaches for efficiency. The results are saved in JSON format for further analysis.

## Features

- **Single-threaded Simulation**: Simulates poker games for a range of players without parallelization.
- **Parallelized Simulation**: Utilizes all available CPU cores to run simulations concurrently, significantly improving performance for large-scale simulations.
- **Customizable Parameters**: Allows users to specify the number of players, iterations, and specific hands to simulate.
- **JSON Output**: Saves simulation results in a structured JSON format for easy analysis.

## How to Run


1. **Run the Simulation**:
   To run the parallelized simulation:
   ```sh
   go run src/main.go
   ```

2. **View the Output**:
   The simulation results will be saved in a JSON file named `simulation_results.json` in the project directory. For example:
   ```plaintext
   simulation_results.json
   ```

## Main Simulations

### 1. **Single-threaded Simulation**
   - Function: [`SimulatiorPlayer1RestUnknown`](src/simulator/simulator.go)
   - Description: Simulates poker games for a range of players without parallelization. This is not optimized for performance but serves as a baseline implementation.
   - Output: Results are saved in `simulation_results.json`.

### 2. **Parallelized Simulation**
   - Function: [`SimulatiorPlayer1RestUnknownParallel`](src/simulator/simulator.go)
   - Description: Utilizes all available CPU cores to run simulations concurrently. This approach is highly efficient for large-scale simulations.
   - Output: Results are saved in `simulation_results.json`.

## Future Enhancements

- Abstracting the `RunProbabilisticSimulation` function to handle any case.
- Using WebAssembly to run the simulation in a browser.

.