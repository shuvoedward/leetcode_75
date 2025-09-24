package leetcode75

func leafSimilar(root1, root2 *TreeNode) bool {

	leaves1 := getLeafSequence(root1)
	leaves2 := getLeafSequence(root2)

	if len(leaves1) != len(leaves2) {
		return false
	}

	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}
	return true
}

func getLeafSequence(root *TreeNode) []int {

	var leaves []int
	collectLeaf(root, &leaves)
	return leaves
}

func collectLeaf(root *TreeNode, leaves *[]int) {

	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
		return
	}
	collectLeaf(root.Left, leaves)
	collectLeaf(root.Right, leaves)
}

func leafSimilar2(root1, root2 *TreeNode) bool {
	l := []int{}
	r := []int{}

	var dfs func(node *TreeNode, arr *[]int)
	dfs = func(node *TreeNode, arr *[]int) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			*arr = append(*arr, node.Val)
			return
		}

		dfs(node.Left, arr)
		dfs(node.Right, arr)
	}

	dfs(root1, &l)
	dfs(root2, &r)

	if len(l) != len(r) {
		return false
	}

	for i := 0; i < len(l); i++ {
		if l[i] != r[i] {
			return false
		}
	}
	return true
}
