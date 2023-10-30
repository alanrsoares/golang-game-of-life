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
		// Add more test cases as necessary.
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
