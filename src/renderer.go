package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"
	"os/exec"
	"runtime"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Renderer interface {
	Play()
}

type CLIRenderer struct {
	Renderer
	Game *Game
}

func clearConsole() {
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
func (cr *CLIRenderer) Render() {
	clearConsole()
	const (
		liveCell = "⬜"
		deadCell = "⬛"
	)

	grid := cr.Game.Grid
	width, height := cr.Game.width, cr.Game.height

	fmt.Println("Press Ctrl+C to quit.")

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := grid.GetCell(Position{x, y})

			cellToPrint := deadCell
			if cell.Alive {
				cellToPrint = liveCell
			}
			fmt.Print(cellToPrint)
		}
		fmt.Println()
	}
}

func (cr *CLIRenderer) Play() {
	generation := 0
	fps := 60

	fmt.Println("Press Ctrl+C to quit.")

	renderLoop(func() {
		generation++
		cr.Render()
		cr.Game.NextGeneration()
		fmt.Println("\nPress Ctrl+C to quit.")
		fmt.Printf("Generation: %d (%dfps)\n", generation, fps)
	}, fps)
}

type GUIRenderer struct {
	Renderer
	Game *Game
}

// Render the game grid with a fyne window
func (wr *GUIRenderer) Render(
	container *fyne.Container,
) {
	grid := wr.Game.Grid
	width, height := wr.Game.width, wr.Game.height

	var cellSize float32 = 10.0

	container.Objects = make([]fyne.CanvasObject, 0)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cell := grid.GetCell(Position{x, y})
			rect := canvas.NewRectangle(color.Gray{Y: 0x20})
			rect.SetMinSize(fyne.NewSize(cellSize, cellSize))

			if cell.Alive {
				rect.FillColor = color.RGBA{R: 0, G: 255, B: 255, A: 255}
			}

			container.Add(rect)
		}
	}

	container.Refresh()
}

func (wr *GUIRenderer) Play() {
	generation := 0

	app := app.New()
	window := app.NewWindow("Game of Life")
	container := container.NewGridWithColumns(wr.Game.width)

	startTime := time.Now()
	fps := 60
	frameDuration := time.Second / time.Duration(fps)

	go func() {
		renderLoop(func() {
			generation++

			elapsed := time.Since(startTime)

			effectiveFps := float64(generation) / elapsed.Seconds()

			// display the current generation and effective fps
			window.SetTitle(fmt.Sprintf(
				"Game of Life (Generation: %d, FPS: %.2f)",
				generation,
				effectiveFps,
			))

			wr.Render(container)
			wr.Game.NextGeneration()

			time.Sleep(frameDuration * 2)
		}, 60)
	}()

	fmt.Println("\nPress Ctrl+C to quit.")

	window.SetContent(container)
	window.ShowAndRun()
}

const (
	cli = "cli"
	gui = "gui"
)

var (
	cliFlag      = flag.Bool(cli, false, "Run in CLI mode")
	guiFlag      = flag.Bool(gui, false, "Run in GUI mode")
	rendererFlag = flag.String("renderer", "", "Choose renderer")
)

func ChooseRenderer(game *Game) Renderer {
	flag.Parse()

	if *cliFlag || *rendererFlag == cli {
		return &CLIRenderer{Game: game}
	} else if *guiFlag || *rendererFlag == gui {
		return &GUIRenderer{Game: game}
	}

	fmt.Println("Choose the rendering mode:")
	fmt.Println("1) CLI")
	fmt.Println("2) GUI")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		return &CLIRenderer{Game: game}
	case 2:
		return &GUIRenderer{Game: game}
	default:
		fmt.Println("Invalid choice, defaulting to Console.")
		return &CLIRenderer{Game: game}
	}
}

func renderLoop(render func(), fps int) {
	frameDuration := time.Second / time.Duration(fps)

	ticker := time.NewTicker(frameDuration)

	for range ticker.C {
		render()
	}
}
