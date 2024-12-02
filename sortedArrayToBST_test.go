package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Helper function to pre-order traverse the BST and collect values
func preOrderTraversal(node *TreeNode, result *[]int) {
	if node == nil {
		*result = append(*result, -1) // Representing nil as -1 for easier validation
		return
	}
	*result = append(*result, node.Val)
	preOrderTraversal(node.Left, result)
	preOrderTraversal(node.Right, result)
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
			expected: []int{0, -3, 9, -10, -1, 5}, // Pre-order: [root, left subtree, right subtree]
		},
		{
			name:     "Example 2",
			nums:     []int{1, 3},
			expected: []int{1, -1, 3, -1, -1}, // Another valid result could be: []int{3, 1, -1, -1, -1}
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
			fmt.Println(root)
			// Generate the pre-order traversal of the result tree
			var result []int
			preOrderTraversal(root, &result)

			// Compare the result with the expected value
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("sortedArrayToBST(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}
