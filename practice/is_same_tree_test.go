package main

import (
	"reflect"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name       string
		root       *TreeNode
		secondRoot *TreeNode
		expected   bool
	}{
		{
			name: "p = [1,2,3], q = [1,2,3]",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3},
				Left:  &TreeNode{Val: 2},
			},
			secondRoot: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3},
				Left:  &TreeNode{Val: 2},
			},
			expected: true,
		},
		{
			name: "p = [1,2], q = [1,null,2]",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			secondRoot: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			name: "p = [1,2,1], q = [1,1,2]",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 1},
				Left:  &TreeNode{Val: 2},
			},
			secondRoot: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
				Left:  &TreeNode{Val: 1},
			},
			expected: false,
		},
		{
			name: "[2,2,2,null,2,null,null,2] [2,2,2,2,null,2,null]",
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			secondRoot: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			expected: false,
		},
		{
			name: "p = [], q = [0]",
			root: nil,
			secondRoot: &TreeNode{
				Val: 0,
			},
			expected: false,
		},
		{
			name: "p = [0], q = [0]",
			root: &TreeNode{
				Val: 0,
			},
			secondRoot: &TreeNode{
				Val: 0,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isSameTree(tt.root, tt.secondRoot)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
