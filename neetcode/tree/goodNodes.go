package tree

func GoodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(node *TreeNode, maxVal int) int
	dfs = func(node *TreeNode, maxVal int) int {
		if node == nil {
			return 0
		}

		res := 0
		if node.Val >= maxVal {
			res++
		}

		maxVal = max(node.Val, maxVal)
		res += dfs(node.Left, maxVal)
		res += dfs(node.Right, maxVal)

		return res
	}

	return dfs(root, root.Val)
}
