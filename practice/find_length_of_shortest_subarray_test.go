package main

import (
	"testing"
)

func TestFindLengthOfShortestSubarray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "example 1",
			input:    []int{1, 2, 3, 10, 4, 2, 3, 5},
			expected: 3,
		},
		{
			name:     "example 2",
			input:    []int{5, 4, 3, 2, 1},
			expected: 4,
		},
		{
			name:     "example 3",
			input:    []int{1, 2, 3},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findLengthOfShortestSubarray(tt.input)
			if result != tt.expected {
				t.Errorf("findLengthOfShortestSubarray(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
