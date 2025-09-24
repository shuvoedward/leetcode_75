package leetcode75

func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		rightSide := 0
		length := len(queue)

		for range length {
			node := queue[0]
			queue = queue[1:]

			rightSide = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, rightSide)

	}
	return result
}
