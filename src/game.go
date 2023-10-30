package main

// Position represents a point on the grid with X and Y positions.
type Position struct {
	x, y int
}

type Dimensions struct {
	width, height int
}

// wrap adjusts the Position to wrap around the grid based on the given gridSize,
// ensuring that the grid behaves as if it were toroidal (top connects to bottom and left connects to right).
func (c Position) wrap(
	height int,
	width int,
) Position {
	return Position{
		x: (c.x + width) % width,
		y: (c.y + height) % height,
	}
}

// Cell represents a single cell on the grid. It holds its aliveness state and its position.
type Cell struct {
	Position
	Alive bool
}

// Grid is a map that associates each Position with a Cell. It represents the entire game state.
type Grid struct {
	Dimensions
	Cells map[Position]Cell
}

// NewGrid initializes and returns a new instance of an empty Grid.
func NewGrid(dimensions Dimensions) *Grid {
	return &Grid{
		Dimensions: dimensions,
		Cells:      make(map[Position]Cell),
	}
}

// SetCell updates the grid at the specified Position to be alive or dead.
func (g Grid) SetCell(pos Position, alive bool) {
	g.Cells[pos] = Cell{Alive: alive, Position: pos}
}

// GetCell retrieves the Cell at the specified Position.
func (g Grid) GetCell(coord Position) Cell {
	return g.Cells[coord]
}

// GetNeighbors computes and returns a slice of Positions that surround a given Position on the grid.
// This includes diagonals, so each cell has eight neighbors.
func (g Grid) GetNeighbors(pos Position) (neighbors []Position) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx != 0 || dy != 0 {
				wrappedCoord := Position{x: pos.x + dx, y: pos.y + dy}.wrap(
					g.height,
					g.width,
				)
				neighbors = append(neighbors, wrappedCoord)
			}
		}
	}
	return
}

// GetLiveNeighbors counts and returns the number of live neighbors around a given Position.
func (g Grid) GetLiveNeighbors(coord Position) (liveNeighbors int) {
	for _, coord := range g.GetNeighbors(coord) {
		if cell, ok := g.Cells[coord]; ok && cell.Alive {
			liveNeighbors++
		}
	}
	return
}

// ShouldLive determines whether a cell at a given Position should be alive in the next generation,
// based on Conway's Game of Life rules.
func (g Grid) ShouldLive(coord Position) bool {
	cell := g.GetCell(coord)
	liveNeighbors := g.GetLiveNeighbors(coord)
	return (cell.Alive && liveNeighbors == 2) || liveNeighbors == 3
}

// NextGeneration computes the next generation of the grid based on the current state,
// applying the rules of Conway's Game of Life and returns the new grid.
func (g Grid) NextGeneration() *Grid {
	newGrid := NewGrid(g.Dimensions)
	considered := make(map[Position]bool)

	// Consider the state of each cell and its neighbors.
	for gridPos := range g.Cells {
		for _, neighborPos := range g.GetNeighbors(gridPos) {
			considered[neighborPos] = true
		}
	}

	// Update the cells based on whether they should live or die.
	for pos := range considered {
		if g.ShouldLive(pos) {
			newGrid.SetCell(pos, true)
		}
	}

	return newGrid
}

type Game struct {
	Grid
	Dimensions
}

func NewGame(dimensions Dimensions) *Game {
	grid := NewGrid(dimensions) // Creates a new grid instance

	return &Game{
		Grid:       *grid, // Dereference to store by value
		Dimensions: dimensions,
	}
}

func (g *Game) NextGeneration() {
	*g = Game{Grid: *g.Grid.NextGeneration(), Dimensions: g.Dimensions}
}
