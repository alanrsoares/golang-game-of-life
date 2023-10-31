package main

func main() {
	const gridSize = 20 // Number of cells in each dimension of the grid

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

	renderer := ChooseRenderer(game)

	renderer.Play()

}
