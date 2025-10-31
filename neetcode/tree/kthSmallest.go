package tree

func KthSmallestInorder(root *TreeNode, k int) int {

	var arr []int

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		arr = append(arr, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return arr[k-1]

	/*
		Time complexity: O(n)
		Space complexity: O(n)
	*/
}

func KthSmallestRecursiveDFS(root *TreeNode, k int) int {
	count, res := k, 0

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		count--
		if count == 0 {
			res = node.Val
			return
		}
		dfs(node.Right)
	}

	dfs(root)
	return res

	// Time complexity: O(n)
	// Space complexity: O(n)
}
