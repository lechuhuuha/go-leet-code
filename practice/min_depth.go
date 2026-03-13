package main

// if current node left and right value is 0 mean that is a leaf node\
// return the track node of that node
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	totalLeftNode := minDepth(root.Left)
	totalRightNode := minDepth(root.Right)

	if totalLeftNode == 0 && totalRightNode == 0 {
		return 1
	}

	// if one child is missing, return the depth of the other
	if totalLeftNode == 0 || totalRightNode == 0 {
		return totalLeftNode + totalRightNode + 1
	}

	if totalLeftNode < totalRightNode {
		return totalLeftNode + 1
	}

	return totalRightNode + 1
}
