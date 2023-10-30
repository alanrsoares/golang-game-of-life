package main

import (
	"time"
)

const frameRate = 8 // Number of frames per second
const gridSize = 10

func main() {
	grid := NewGrid()
	// Seed the grid with an initial glider moving to the bottom right.
	grid.SetCell(Coordinate{X: 1, Y: 0}, true)
	grid.SetCell(Coordinate{X: 2, Y: 1}, true)
	grid.SetCell(Coordinate{X: 0, Y: 2}, true)
	grid.SetCell(Coordinate{X: 1, Y: 2}, true)
	grid.SetCell(Coordinate{X: 2, Y: 2}, true)

	// Game loop - runs for 10 generations for demonstration purposes.
	for i := 0; i < 100; i++ {
		render(grid)
		grid = grid.NextGeneration()
		time.Sleep(time.Second / frameRate)
	}
}
