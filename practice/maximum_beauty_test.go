package main

import (
	"reflect"
	"testing"
)

func TestMaximumBeauty(t *testing.T) {
	tests := []struct {
		items    [][]int
		queries  []int
		expected []int
	}{
		{
			items:    [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}},
			queries:  []int{1, 2, 3, 4, 5, 6},
			expected: []int{2, 4, 5, 5, 6, 6},
		},
	}

	for _, tt := range tests {
		t.Run("maximumBeauty test case", func(t *testing.T) {
			result := maximumBeauty(tt.items, tt.queries)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
