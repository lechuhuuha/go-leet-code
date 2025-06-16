package main

import (
	"reflect"
	"testing"
)

func TestResultsArray(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		k      int
		output []int
	}{
		{
			name:   "test case positive - increasing sequence",
			nums:   []int{1, 2, 3, 4, 3, 2, 5},
			k:      3,
			output: []int{3, 4, -1, -1, -1},
		},
		{
			name:   "test case negative - all same numbers",
			nums:   []int{2, 2, 2, 2, 2},
			k:      4,
			output: []int{-1, -1},
		},
		{
			name:   "test case negative - alternating numbers",
			nums:   []int{3, 2, 3, 2, 3, 2},
			k:      2,
			output: []int{-1, 3, -1, 3, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resultsArray(tt.nums, tt.k)
			if !reflect.DeepEqual(result, tt.output) {
				t.Errorf("resultsArray(%v, %d) = %v; want %v", tt.nums, tt.k, result, tt.output)
			}
		})
	}
}
