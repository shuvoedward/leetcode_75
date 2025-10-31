package tree

import "fmt"

func LevelOrderBFS(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	type nodeInfo struct {
		node  *TreeNode
		level int
	}

	queue := []nodeInfo{{node: root, level: 0}}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n.level == len(result) {
			result = append(result, []int{})
		}

		result[n.level] = append(result[n.level], n.node.Val)

		if n.node.Left != nil {
			queue = append(queue, nodeInfo{node: n.node.Left, level: n.level + 1})
		}

		if n.node.Right != nil {
			queue = append(queue, nodeInfo{node: n.node.Right, level: n.level + 1})
		}
		fmt.Println("here")
	}

	return result
}

func levelOrderNeetcodeBFS(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		queueLen := len(queue)
		level := []int{}

		for range queueLen {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

func levelOrderDFS(root *TreeNode) [][]int {
	res := [][]int{}

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if len(res) == depth {
			res = append(res, []int{})
		}

		res[depth] = append(res[depth], node.Val)
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)

	}

	dfs(root, 0)
	return res
}
