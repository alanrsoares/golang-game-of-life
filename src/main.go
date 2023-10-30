package main

import (
	"fmt"
	"time"
)

const frameRate = 16 // Number of frames per second
const gridSize = 20

func main() {
	game := NewGame(gridSize, gridSize)

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
