package main

// Position represents a point on the grid with X and Y positions.
type Position struct {
	x, y int
}

// Dimensions represents the width and height of the grid.
type Dimensions struct {
	width, height int
}

// Wrap adjusts the Position to Wrap around the grid based on the given gridSize,
// ensuring that the grid behaves as if it were toroidal (top connects to bottom and left connects to right).
func (c Position) Wrap(
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

func (c Cell) Toggle() Cell {
	return Cell{
		Position: c.Position,
		Alive:    !c.Alive,
	}
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

// RandomGrid initializes and returns a new instance of a Grid with random cells.
func RandomGrid(dimensions Dimensions) *Grid {
	grid := NewGrid(dimensions)

	for x := 0; x < dimensions.width; x++ {
		for y := 0; y < dimensions.height; y++ {
			grid.SetCell(Position{x, y}, RandomBool())
		}
	}

	return grid
}

// GliderGrid initializes and returns a new instance of a Grid with a glider.
func GliderGrid(dimensions Dimensions) *Grid {
	grid := NewGrid(dimensions)

	gliderPositions := []Position{
		{x: 1, y: 0},
		{x: 2, y: 1},
		{x: 0, y: 2},
		{x: 1, y: 2},
		{x: 2, y: 2},
	}

	for _, pos := range gliderPositions {
		grid.SetCell(pos, true)
	}

	return grid
}

// SetCell updates the grid at the specified Position to be alive or dead.
func (g Grid) SetCell(pos Position, alive bool) {
	g.Cells[pos] = Cell{Alive: alive, Position: pos}
}

// Cell retrieves the Cell at the specified Position.
func (g Grid) Cell(coord Position) Cell {
	return g.Cells[coord]
}

func (g Grid) ToggleCell(coord Position) {
	g.Cells[coord] = g.Cells[coord].Toggle()
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

			wrappedCoord := Position{x: pos.x + dx, y: pos.y + dy}.Wrap(
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
