package main

import (
	"reflect"
	"testing"
)

// Helper function to convert BST to level-order traversal array
func levelOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{-1}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, -1)
			continue
		}

		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	// Trim trailing -1s
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != -1 {
			return result[:i+1]
		}
	}
	return result
}

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{-10, -3, 0, 5, 9},
			expected: []int{0, -3, 9, -10, -1, 5}, // Level-order traversal
		},
		{
			name:     "Example 2",
			nums:     []int{1, 3},
			expected: []int{3, 1}, // Level-order traversal
		},
		{
			name:     "Empty Input",
			nums:     []int{},
			expected: []int{-1}, // Representing an empty tree
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedArrayToBST(tt.nums)
			result := levelOrderTraversal(root)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sortedArrayToBST(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}
