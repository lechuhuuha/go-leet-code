package main

func preorderTraversal(root *TreeNode) []int {
	result = []int{}
	dfsPreOrder(root)

	return result
}

func dfsPreOrder(g *TreeNode) {
	if g == nil {
		return
	}

	result = append(result, g.Val)
	dfsPreOrder(g.Left)
	dfsPreOrder(g.Right)
}

// Helper to build tree from level-order input
func buildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
