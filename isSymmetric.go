package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Left != nil && root.Right == nil {
		return false
	}

	if root.Left == nil && root.Right != nil {
		return false
	}

	return isSymmetricDFS(root.Left, root.Right)
}

func isSymmetricDFS(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSymmetricDFS(p.Left, q.Right) && isSymmetricDFS(p.Right, q.Left)
}
