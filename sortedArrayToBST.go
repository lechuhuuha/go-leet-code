package main

import "fmt"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	// Key Idea: Let us recursively build the trees
	n := len(nums)

	// Base cases
	if n == 0 {
		return nil
	}

	if n == 1 {
		return &TreeNode{Val: nums[0]}
	}

	l, r := 0, n-1
	mid := l + (r-l)/2
	fmt.Println(nums[:mid], nums[mid+1:], nums[mid])
	newRoot := &TreeNode{Val: nums[mid]}
	newRoot.Left, newRoot.Right = sortedArrayToBST(nums[:mid]), sortedArrayToBST(nums[mid+1:])

	return newRoot
}
