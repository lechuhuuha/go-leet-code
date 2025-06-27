package main

func isBalanced(root *TreeNode) bool {
	return checkHeight(root) != -1
}

func checkHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := checkHeight(node.Left)
	if left == -1 {
		return -1
	}

	right := checkHeight(node.Right)
	if right == -1 {
		return -1
	}

	if Abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}
