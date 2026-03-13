package main

import (
	"testing"
)

func TestMaxManhattanDistance(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			s:        "NWSE",
			k:        1,
			expected: 3,
		},
		{
			name:     "Example 2",
			s:        "NSWWEW",
			k:        3,
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxManhattanDistance(tt.s, tt.k)
			if result != tt.expected {
				t.Errorf("maxManhattanDistance(%q, %d) = %d; want %d", tt.s, tt.k, result, tt.expected)
			}
		})
	}
}
