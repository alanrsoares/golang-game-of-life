package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const frameRate = 8 // Number of frames per second
const gridSize = 10

const deadCell = "□"
const liveCell = "■"

func main() {
	grid := make(Grid)
	// Seed the grid with an initial glider moving to the bottom right.
	grid[Coordinate{X: 1, Y: 0}] = Cell{Alive: true}
	grid[Coordinate{X: 2, Y: 1}] = Cell{Alive: true}
	grid[Coordinate{X: 0, Y: 2}] = Cell{Alive: true}
	grid[Coordinate{X: 1, Y: 2}] = Cell{Alive: true}
	grid[Coordinate{X: 2, Y: 2}] = Cell{Alive: true}

	// Game loop - runs for 10 generations for demonstration purposes.
	for i := 0; i < 100; i++ {
		topLeft := Coordinate{X: 0, Y: 0}
		bottomRight := Coordinate{X: gridSize - 1, Y: gridSize - 1}
		render(grid, topLeft, bottomRight)
		grid = grid.NextGeneration()
		time.Sleep(time.Second / frameRate)
	}
}

type Coordinate struct {
	X, Y int
}

type Cell struct {
	Alive bool
	coord Coordinate
}

func (c Cell) GetNeighbors(g Grid) (neighbors []Cell) {
	for _, neighbor := range getNeighbors(Coordinate{X: 0, Y: 0}) {
		neighbors = append(neighbors, g[neighbor])
	}
	return
}

func (c Cell) GetLiveNeighbors(g Grid) (liveNeighbors int) {
	for _, neighbor := range c.GetNeighbors(g) {
		if neighbor.Alive {
			liveNeighbors++
		}
	}
	return
}

type Grid map[Coordinate]Cell

func (g Grid) SetCell(coord Coordinate, alive bool) {
	g[coord] = Cell{Alive: alive, coord: coord}
}

func (g Grid) GetCell(coord Coordinate) Cell {
	return g[coord]
}

func (g Grid) GetLiveNeighbors(coord Coordinate) (liveNeighbors int) {
	for _, neighbor := range getNeighbors(coord) {
		if cell, ok := g[neighbor]; ok && cell.Alive {
			liveNeighbors++
		}
	}
	return
}

func (g Grid) NextGeneration() Grid {
	newGrid := make(Grid)
	considered := make(map[Coordinate]bool)

	for coord := range g {
		for _, neighbor := range getNeighbors(coord) {
			considered[neighbor] = true
		}
	}

	for coord := range considered {
		if shouldLive(coord, g) {
			newGrid.SetCell(coord, true)
		}
	}

	return newGrid
}

// wrapCoordinate wraps the coordinate around the grid if it goes out of bounds.
func wrapCoordinate(coord Coordinate) Coordinate {
	return Coordinate{
		X: (coord.X + gridSize) % gridSize,
		Y: (coord.Y + gridSize) % gridSize,
	}
}

// clearScreen uses platform-specific commands to clear the terminal screen.
func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func render(grid Grid, topLeft Coordinate, bottomRight Coordinate) {
	clearScreen() // Clear the screen before rendering the new state

	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		for x := topLeft.X; x <= bottomRight.X; x++ {
			coord := Coordinate{X: x, Y: y}
			if cell, ok := grid[coord]; ok && cell.Alive {
				fmt.Print(liveCell)
			} else {
				fmt.Print(deadCell)
			}
		}
		fmt.Println()
	}
}

func getNeighbors(coord Coordinate) (neighbors []Coordinate) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx != 0 || dy != 0 {
				wrappedCoord := wrapCoordinate(Coordinate{X: coord.X + dx, Y: coord.Y + dy})
				neighbors = append(neighbors, wrappedCoord)
			}
		}
	}

	return
}

func shouldLive(coord Coordinate, grid Grid) bool {
	_, isAlive := grid[coord]
	liveNeighbors := grid.GetLiveNeighbors(coord)
	return (isAlive && liveNeighbors == 2) || liveNeighbors == 3
}
