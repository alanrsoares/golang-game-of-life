package main

type Game struct {
	Grid
	Dimensions
}

// NewGame initializes and returns a new instance of a Game.
func NewGame(dimensions Dimensions) *Game {
	grid := NewGrid(dimensions) // Creates a new grid instance

	return &Game{
		Grid:       *grid, // Dereference to store by value
		Dimensions: dimensions,
	}
}

// NextGeneration computes the next generation of the game based on the current state,
// applying the rules of Conway's Game of Life.
// This is a facade method that calls the Grid's NextGeneration method.
func (g *Game) NextGeneration() {
	*g = Game{Grid: *g.Grid.NextGeneration(), Dimensions: g.Dimensions}
}
