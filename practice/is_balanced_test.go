package main

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name: "balanced tree [3,9,20,null,null,15,7]",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: true,
		},
		{
			name: "unbalanced tree [1,2,2,3,3,null,null,4,4]",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			name:     "empty tree []",
			root:     nil,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBalanced(tt.root)
			if got != tt.expected {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, got)
			}
		})
	}
}
