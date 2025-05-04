package main

import (
	"testing"
)

func TestIsSortable(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected bool
	}{
		{
			name: "example 1 possible",
			matrix: [][]int{
				{1, 2},
				{2, 1},
			},
			expected: true,
		},
		{
			name: "example 1 impossible",
			matrix: [][]int{
				{10, 20, 30},
				{1, 1, 1},
				{0, 0, 1},
			},
			expected: false,
		},
		{
			name: "example 3 impossible",
			matrix: [][]int{
				{0, 2},
				{1, 1},
			},
			expected: false,
		},
		{
			name: "example 4 possible",
			matrix: [][]int{
				{1, 1},
				{1, 1},
			},
			expected: true,
		},
		{
			name: "example 5 possible",
			matrix: [][]int{
				{1000000000, 0},
				{0, 1000000000},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSortable(len(tt.matrix), tt.matrix)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
