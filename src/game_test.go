package main

import (
	"testing"
)

func TestWrapCoordinate(t *testing.T) {
	gridSize := 20
	tests := []struct {
		name     string
		coord    Position
		expected Position
	}{
		{"Wrap both coordinates", Position{-1, -1}, Position{gridSize - 1, gridSize - 1}},
		{"Wrap X coordinate", Position{-1, 5}, Position{gridSize - 1, 5}},
		{"Wrap Y coordinate", Position{5, -1}, Position{5, gridSize - 1}},
		{"No wrap needed", Position{5, 5}, Position{5, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.coord.wrap(gridSize, gridSize)
			if result != tt.expected {
				t.Errorf("wrapCoordinate(%v) = %v, want %v", tt.coord, result, tt.expected)
			}
		})
	}
}

func TestCellShouldLive(t *testing.T) {
	tests := []struct {
		name     string
		position Position
		expected bool
	}{
		{
			"A live cell with fewer than two live neighbors dies",
			Position{2, 1},
			false,
		},
		{
			"A live cell with two or three live neighbors lives",
			Position{2, 2},
			true,
		},
		{
			"A dead cell with exactly three live neighbors becomes alive",
			Position{1, 2},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			game := NewGame(
				Dimensions{width: 5, height: 5},
			)

			game.SetCell(Position{2, 1}, true)
			game.SetCell(Position{2, 2}, true)
			game.SetCell(Position{2, 3}, true)

			// act
			result := game.ShouldLive(tt.position)

			// assert
			if result != tt.expected {
				t.Errorf("Cell at %v should live = %v, want %v", tt.position, result, tt.expected)
			}
		})
	}

}

func TestNextGeneration(t *testing.T) {
	game := NewGame(
		Dimensions{width: 5, height: 5},
	)

	game.SetCell(Position{x: 2, y: 1}, true)
	game.SetCell(Position{x: 2, y: 2}, true)
	game.SetCell(Position{x: 2, y: 3}, true)

	game.NextGeneration()

	// Check some expected alive cells after first generation of glider.
	expectedAlive := []Position{
		{x: 1, y: 2},
		{x: 2, y: 2},
		{x: 3, y: 2},
	}

	for _, coord := range expectedAlive {
		if !game.Cell(coord).Alive {
			t.Errorf("Expected cell at %v to be alive", coord)
		}
	}
}
