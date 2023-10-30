package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const deadCell = "⬛"
const liveCell = "⬜"

type ConsoleRenderer struct {
	Game *Game
}

// Clear the screen
func (cr ConsoleRenderer) clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Render the game grid
func (cr *ConsoleRenderer) Render() {
	cr.clearScreen()
	grid := cr.Game.Grid
	width, height := cr.Game.width, cr.Game.height

	for Y := 0; Y < height; Y++ {
		for X := 0; X < width; X++ {
			cell := grid.GetCell(Coordinate{X, Y})
			if cell.Alive {
				fmt.Print(liveCell)
			} else {
				fmt.Print(deadCell)
			}
		}
		fmt.Println()
	}

	fmt.Println("\nPress Ctrl+C to quit.")
}
