package main

import (
	"reflect"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		input    []any
		expected []int
	}{
		{
			name:     "right-skewed with left leaf",
			input:    []any{1, nil, 2, 3},
			expected: []int{3, 2, 1},
		},
		{
			name:     "complex tree",
			input:    []any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9},
			expected: []int{4, 6, 7, 5, 2, 9, 8, 3, 1},
		},
		{
			name:     "empty tree",
			input:    []any{},
			expected: []int{},
		},
		{
			name:     "single node",
			input:    []any{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			got := postorderTraversal(root)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.expected)
			}
		})
	}
}
