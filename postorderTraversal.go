package main

func postorderTraversal(root *TreeNode) []int {
	result = []int{}
	dfsPostOrder(root)

	return result
}

func dfsPostOrder(g *TreeNode) {
	if g == nil {
		return
	}

	dfsPostOrder(g.Left)
	dfsPostOrder(g.Right)
	result = append(result, g.Val)
}
