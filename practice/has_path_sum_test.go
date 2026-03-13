package main

import "testing"

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      *TreeNode
		targetSum int
		expected  bool
	}{
		{
			name: "Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val:  8,
					Left: &TreeNode{Val: 13},
					Right: &TreeNode{
						Val:   4,
						Right: &TreeNode{Val: 1},
					},
				},
			},
			targetSum: 22,
			expected:  true,
		},
		{
			name: "Input: root = [1,2,3], targetSum = 5",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			targetSum: 5,
			expected:  false,
		},
		{
			name:      "Input: root = [], targetSum = 0",
			root:      nil,
			targetSum: 0,
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasPathSum(tt.root, tt.targetSum)
			if got != tt.expected {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, got)
			}
		})
	}
}
