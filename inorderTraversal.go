package main

var (
	result = []int{}
)

func inorderTraversal(root *TreeNode) []int {
	result = []int{}
	dfs(root)

	return result
}

func dfs(g *TreeNode) {
	if g == nil {
		return
	}

	dfs(g.Left)
	result = append(result, g.Val)
	dfs(g.Right)
}
