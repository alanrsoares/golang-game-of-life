package main

import (
	"fmt"
	"time"
)

func main() {
	const frameRate = 60           // Number of frames per second
	const gridSize = 20            // Number of cells in each dimension of the grid
	const maxGenerations = 1000000 // Number of generations to run the game for

	game := NewGame(
		Dimensions{width: gridSize * 2, height: gridSize},
	)

	// Seed the grid with an initial glider moving to the bottom right.

	gliderPositions := []Position{
		{x: 1, y: 0},
		{x: 2, y: 1},
		{x: 0, y: 2},
		{x: 1, y: 2},
		{x: 2, y: 2},
	}

	for _, pos := range gliderPositions {
		game.SetCell(pos, true)
	}

	frames := 0
	renderer := ConsoleRenderer{game}

	// Game loop - runs for 10 generations for demonstration purposes.
	for i := 0; i < maxGenerations; i++ {
		frames++
		renderer.Render()
		game.NextGeneration()
		fmt.Printf("Generation: %d/%d (%dfps)\n", i, maxGenerations, frameRate)
		time.Sleep(time.Second / frameRate)
	}
}
