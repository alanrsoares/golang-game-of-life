package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

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

	grid.Render() // Render the grid

	// render line with instructions
	fmt.Println("\nPress Ctrl+C to quit.")
}
