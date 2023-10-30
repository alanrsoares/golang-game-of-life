package main

// Coordinate represents a point on the grid with X and Y positions.
type Coordinate struct {
	X, Y int
}

// wrap adjusts the Coordinate to wrap around the grid based on the given gridSize,
// ensuring that the grid behaves as if it were toroidal (top connects to bottom and left connects to right).
func (c Coordinate) wrap(gridSize int) Coordinate {
	return Coordinate{
		X: (c.X + gridSize) % gridSize,
		Y: (c.Y + gridSize) % gridSize,
	}
}

// Cell represents a single cell on the grid. It holds its aliveness state and its position.
type Cell struct {
	Alive bool
	Coord Coordinate
}

// Grid is a map that associates each Coordinate with a Cell. It represents the entire game state.
type Grid map[Coordinate]Cell

// NewGrid initializes and returns a new instance of an empty Grid.
func NewGrid() Grid {
	return make(Grid)
}

// SetCell updates the grid at the specified Coordinate to be alive or dead.
func (g Grid) SetCell(coord Coordinate, alive bool) {
	g[coord] = Cell{Alive: alive, Coord: coord}
}

// GetCell retrieves the Cell at the specified Coordinate.
func (g Grid) GetCell(coord Coordinate) Cell {
	return g[coord]
}

// WrapCoordinate wraps a given Coordinate around the grid to ensure it doesn't go out of bounds.
func (g Grid) WrapCoordinate(coord Coordinate) Coordinate {
	return coord.wrap(gridSize)
}

// GetNeighbors computes and returns a slice of Coordinates that surround a given Coordinate on the grid.
// This includes diagonals, so each cell has eight neighbors.
func (g Grid) GetNeighbors(coord Coordinate) (neighbors []Coordinate) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx != 0 || dy != 0 {
				wrappedCoord := Coordinate{X: coord.X + dx, Y: coord.Y + dy}.wrap(gridSize)
				neighbors = append(neighbors, wrappedCoord)
			}
		}
	}
	return
}

// GetLiveNeighbors counts and returns the number of live neighbors around a given Coordinate.
func (g Grid) GetLiveNeighbors(coord Coordinate) (liveNeighbors int) {
	for _, neighborCoord := range g.GetNeighbors(coord) {
		if cell, ok := g[neighborCoord]; ok && cell.Alive {
			liveNeighbors++
		}
	}
	return
}

// ShouldLive determines whether a cell at a given Coordinate should be alive in the next generation,
// based on Conway's Game of Life rules.
func (g Grid) ShouldLive(coord Coordinate) bool {
	cell := g.GetCell(coord)
	liveNeighbors := g.GetLiveNeighbors(coord)
	return (cell.Alive && liveNeighbors == 2) || liveNeighbors == 3
}

// NextGeneration computes the next generation of the grid based on the current state,
// applying the rules of Conway's Game of Life and returns the new grid.
func (g Grid) NextGeneration() Grid {
	newGrid := NewGrid()
	considered := make(map[Coordinate]bool)

	// Consider the state of each cell and its neighbors.
	for coord := range g {
		for _, neighbor := range g.GetNeighbors(coord) {
			considered[neighbor] = true
		}
	}

	// Update the cells based on whether they should live or die.
	for coord := range considered {
		if g.ShouldLive(coord) {
			newGrid.SetCell(coord, true)
		}
	}

	return newGrid
}
