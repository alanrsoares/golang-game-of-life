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

// Cell retrieves the Cell at the specified Position.
func (g Grid) Cell(coord Position) Cell {
	return g.Cells[coord]
}

// Neighbors computes and returns a slice of Positions that surround a given Position on the grid.
// This includes diagonals, so each cell has eight neighbors.
func (g Grid) Neighbors(pos Position) (neighbors []Position) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// early continue if we're looking at the center cell
			if dx == 0 && dy == 0 {
				continue
			}

			wrappedCoord := Position{x: pos.x + dx, y: pos.y + dy}.wrap(
				g.height,
				g.width,
			)
			neighbors = append(neighbors, wrappedCoord)
		}
	}
	return
}

// LiveNeighbors counts and returns the number of live neighbors around a given Position.
func (g Grid) LiveNeighbors(coord Position) (liveNeighbors int) {
	for _, coord := range g.Neighbors(coord) {
		if cell, ok := g.Cells[coord]; ok && cell.Alive {
			liveNeighbors++
		}
	}
	return
}

// ShouldLive determines whether a cell at a given Position should be alive in the next generation,
// based on Conway's Game of Life rules.
func (g Grid) ShouldLive(coord Position) bool {
	switch cell, liveNeighbors := g.Cell(coord), g.LiveNeighbors(coord); {
	// rule: Any live cell with fewer than two live neighbours
	// rule: Any live cell with more than three live neighbours dies, as if by overpopulation.
	case cell.Alive && liveNeighbors < 2 || liveNeighbors > 3:
		return false
	// rule: Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	case !cell.Alive && liveNeighbors == 3:
		return true
	// rule: Any live cell with two or three live neighbours lives on to the next generation..
	default:
		return cell.Alive
	}
}

// NextGeneration computes the next generation of the grid based on the current state,
// applying the rules of Conway's Game of Life and returns the new grid.
func (g Grid) NextGeneration() *Grid {
	newGrid := NewGrid(g.Dimensions)
	considered := make(map[Position]bool)

	// Consider the state of each cell and its neighbors.
	for gridPos := range g.Cells {
		for _, neighborPos := range g.Neighbors(gridPos) {
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
