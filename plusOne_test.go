package main

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "Example 1",
			input:  []int{1, 2, 3},
			output: []int{1, 2, 4},
		},
		{
			name:   "Example 2",
			input:  []int{4, 3, 2, 1},
			output: []int{4, 3, 2, 2},
		},
		{
			name:   "Example 3",
			input:  []int{9},
			output: []int{1, 0},
		},
		{
			name:   "Example 4",
			input:  []int{9, 9},
			output: []int{1, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := plusOne(tt.input)
			if !reflect.DeepEqual(result, tt.output) {
				t.Errorf("plusOne(%v) = %v; want %v", tt.input, result, tt.output)
			}
		})
	}
}
