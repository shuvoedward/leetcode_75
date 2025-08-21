package main

func goodNodes(root *TreeNode) int {

	res := 0

	var dfs func(node *TreeNode, maxValue int, maxResult *int)
	dfs = func(node *TreeNode, maxValue int, maxResult *int) {
		if node == nil {
			return
		}

		if node.Val >= maxValue {
			maxValue = node.Val
			*maxResult++
		}

		dfs(node.Left, maxValue, maxResult)
		dfs(node.Right, maxValue, maxResult)
	}

	dfs(root, root.Val, &res)
	return res
}

// From neetcode
func goodNodes2(root *TreeNode) int {
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
			res = 1
		}

		maxVal = max(maxVal, node.Val)
		res += dfs(node.Left, maxVal)
		res += dfs(node.Right, maxVal)

		return res
	}

	return dfs(root, root.Val)

}

// Similar as first one
func goodNodes3(root *TreeNode) int {
	res := 0
	preOrder(root, root.Val, &res)
	return res
}

func preOrder(node *TreeNode, currMax int, res *int) {
	if node == nil {
		return
	}
	if node.Val >= currMax {
		currMax = node.Val
		*res++
	}
	preOrder(node.Left, currMax, res)
	preOrder(node.Right, currMax, res)
}
