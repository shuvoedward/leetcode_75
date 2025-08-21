package main

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	level := 0
	maxSum := -1 << 63
	maxLevel := 0

	for len(queue) > 0 {
		length := len(queue)
		sum := 0

		level++

		for range length {
			node := queue[0]
			queue = queue[1:]

			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if sum > maxSum {
			maxSum = sum
			maxLevel = level
		}
	}
	return maxLevel
}
