package tree

// For each individual node, get the max height between left and right.
func diameterDepthFirst(root *TreeNode) int {

	res := 0

	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)
		return 1 + max(res, left+right)
	}

	dfs(root)
	return res

	// time complexity O(n)
	// space complexity: O(h)
	// Best case (balanced tree): O(log(n))
	// worst case (degenerate tree): O(n)

}
