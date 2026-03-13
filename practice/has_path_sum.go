package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	// pass the value from the root to the branch
	// if the value is equal to target sum and is a leaf node (current node left and right value is 0)
	if root == nil {
		return false
	}

	return checkSum(root, targetSum, 0) == -2
}

func checkSum(node *TreeNode, targetSum int, parentValue int) int {
	if node == nil {
		return 0
	}

	left := checkSum(node.Left, targetSum, node.Val+parentValue)

	right := checkSum(node.Right, targetSum, node.Val+parentValue)
	if left == -2 {
		return -2
	}
	if right == -2 {
		return -2
	}

	if (node.Val+parentValue) == targetSum && (left == 0 && right == 0) {
		return -2 // is correct
	}

	return -1
}
