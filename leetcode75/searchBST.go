package main

func searchBST(root *TreeNode, val int) *TreeNode {
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		length := len(stack)

		for range length {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if node.Val == val {
				return node
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}

	return nil
}

func searchBSTOptimized(root *TreeNode, val int) *TreeNode {
	node := root
	for node != nil && node.Val != val {
		if val > node.Val {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return node
}
