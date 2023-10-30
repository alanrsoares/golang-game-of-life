package main

import (
	"fmt"
	"time"
)

const frameRate = 16 // Number of frames per second
const gridSize = 20

func main() {
	game := NewGame(gridSize, gridSize)

	// Seed the grid with an initial glider moving to the bottom right.
	game.SetCell(Coordinate{X: 1, Y: 0}, true)
	game.SetCell(Coordinate{X: 2, Y: 1}, true)
	game.SetCell(Coordinate{X: 0, Y: 2}, true)
	game.SetCell(Coordinate{X: 1, Y: 2}, true)
	game.SetCell(Coordinate{X: 2, Y: 2}, true)

	frames := 0

	// Game loop - runs for 10 generations for demonstration purposes.
	for i := 0; i < 1000000000; i++ {
		frames++
		render(game.Grid)
		game.NextGeneration()
		fmt.Printf("Generation: %d\n", i)
		time.Sleep(time.Second / frameRate)
	}
}
