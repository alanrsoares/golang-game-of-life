package main

func main() {
	const gridSize = 20 // Number of cells in each dimension of the grid

	game := NewGame(
		Dimensions{width: gridSize * 2, height: gridSize},
	)

	renderer := ChooseRenderer(game)

	renderer.Play()
}
