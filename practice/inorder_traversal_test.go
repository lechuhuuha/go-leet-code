package main

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "simple tree [1,null,2,3]",
			root:     &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			expected: []int{1, 3, 2},
		},
		{
			name: "complex tree [1,2,3,4,5,null,8,null,null,6,7,9]",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 7},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val:   8,
						Right: &TreeNode{Val: 9},
					},
				},
			},
			expected: []int{4, 2, 6, 5, 7, 1, 3, 8, 9},
		},
		{
			name:     "empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "single node [1]",
			root:     &TreeNode{Val: 1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := inorderTraversal(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// Helper to build a balanced binary tree of depth `d` (full binary tree)
func buildBalancedTree(val, depth int) *TreeNode {
	if depth == 0 {
		return nil
	}
	return &TreeNode{
		Val:   val,
		Left:  buildBalancedTree(val*2, depth-1),
		Right: buildBalancedTree(val*2+1, depth-1),
	}
}

func BenchmarkInorderTraversal(b *testing.B) {
	root := buildBalancedTree(1, 15) // 32,767 nodes

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inorderTraversal(root)
	}
}
