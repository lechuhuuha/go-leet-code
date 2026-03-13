package main

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name     string
		numRows  int
		expected [][]int
	}{
		{
			name:     "Zero rows",
			numRows:  0,
			expected: [][]int{},
		},
		{
			name:    "One row",
			numRows: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:    "Two rows",
			numRows: 2,
			expected: [][]int{
				{1},
				{1, 1},
			},
		},
		{
			name:    "Three rows",
			numRows: 3,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
			},
		},
		{
			name:    "Five rows",
			numRows: 5,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pascalTriangle(tt.numRows)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("generate(%d) = %v, want %v", tt.numRows, result, tt.expected)
			}
		})
	}
}
