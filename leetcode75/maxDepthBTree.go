package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	Node  *TreeNode
	Depth int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []Pair{{Node: root, Depth: 1}}
	res := 0

	for len(stack) > 0 {
		currentPair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := currentPair.Node
		depth := currentPair.Depth

		if node != nil {
			res = max(res, depth)

			if node.Left != nil {
				stack = append(stack, Pair{Node: node.Left, Depth: depth + 1})
			}
			if node.Right != nil {
				stack = append(stack, Pair{Node: node.Right, Depth: depth + 1})
			}
		}
	}

	return res

}

func maxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := 1 + maxDepthDFS(root.Left)
	r := 1 + maxDepthDFS(root.Right)

	return max(l, r)
}

func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	levels := 0 // depth
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		length := len(queue)
		node := queue[0]
		queue = queue[1:]

		for i := 0; i < length; i++ {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levels++
	}
	return levels

}
