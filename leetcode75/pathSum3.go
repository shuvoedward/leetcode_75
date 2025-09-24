package leetcode75

func pathSum3(root *TreeNode, targetSum int) int {
	prefixSumFreq := map[int]int{0: 1}

	count := 0

	findPaths(root, 0, targetSum, prefixSumFreq, &count)

	return count
}

func findPaths(node *TreeNode, currentSum int, targetSum int, prefixSumFreq map[int]int, count *int) {
	if node == nil {
		return
	}

	currentSum += node.Val

	*count += prefixSumFreq[currentSum-targetSum]

	prefixSumFreq[currentSum]++

	findPaths(node.Left, currentSum, targetSum, prefixSumFreq, count)
	findPaths(node.Right, currentSum, targetSum, prefixSumFreq, count)

	prefixSumFreq[currentSum]--

}

func pathSum31(root *TreeNode, targetSum int) int {

	var res int

	h := make(map[int]int)

	var preorder func(node *TreeNode, currSum int)

	preorder = func(node *TreeNode, currSum int) {
		if node == nil {
			return
		}

		currSum += node.Val

		if currSum == targetSum {
			res++
		}

		res += h[currSum-targetSum]

		h[currSum]++

		preorder(node.Left, currSum)
		preorder(node.Right, currSum)

		h[currSum]--
	}

	preorder(root, 0)

	return res
}
