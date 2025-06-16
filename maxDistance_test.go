package main

import (
	"reflect"
	"testing"
)

func TestMaxDistance(t *testing.T) {
	tests := []struct {
		items    []int
		expected int
	}{
		{
			items:    []int{1, 1, 1, 6, 1, 1, 1},
			expected: 3,
		},
		{
			items:    []int{1, 8, 3, 8, 3},
			expected: 4,
		},
		{
			items:    []int{0, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run("maxDistance test case", func(t *testing.T) {
			result := maxDistance(tt.items)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
