package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const deadCell = "□"
const liveCell = "■"

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

func render(grid Grid) {
	clearScreen() // Clear the screen before rendering the new state
	topLeft := Coordinate{X: 0, Y: 0}
	bottomRight := Coordinate{X: gridSize - 1, Y: gridSize - 1}

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
