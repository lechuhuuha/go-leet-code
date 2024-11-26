package main

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "Example 2",
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			name:     "Example 3",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeSortedArray(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.expected) {
				t.Errorf("got %v, want %v", tt.nums1, tt.expected)
			}
		})
	}
}
