package main

import (
	"testing"
)

func TestWrapCoordinate(t *testing.T) {
	tests := []struct {
		name     string
		coord    Coordinate
		expected Coordinate
	}{
		{"Wrap both coordinates", Coordinate{-1, -1}, Coordinate{gridSize - 1, gridSize - 1}},
		{"Wrap X coordinate", Coordinate{-1, 5}, Coordinate{gridSize - 1, 5}},
		{"Wrap Y coordinate", Coordinate{5, -1}, Coordinate{5, gridSize - 1}},
		{"No wrap needed", Coordinate{5, 5}, Coordinate{5, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.coord.wrap(gridSize)
			if result != tt.expected {
				t.Errorf("wrapCoordinate(%v) = %v, want %v", tt.coord, result, tt.expected)
			}
		})
	}
}

func TestNextGeneration(t *testing.T) {
	game := NewGame(
		5,
		5,
	)

	game.SetCell(Coordinate{X: 2, Y: 1}, true)
	game.SetCell(Coordinate{X: 2, Y: 2}, true)
	game.SetCell(Coordinate{X: 2, Y: 3}, true)

	game.NextGeneration()

	// Check some expected alive cells after first generation of glider.
	expectedAlive := []Coordinate{
		{X: 1, Y: 2},
		{X: 2, Y: 2},
		{X: 3, Y: 2},
	}

	for _, coord := range expectedAlive {
		if !game.GetCell(coord).Alive {
			t.Errorf("Expected cell at %v to be alive", coord)
		}
	}
}
