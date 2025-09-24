package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTreeBFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]

		currNode.Left, currNode.Right = currNode.Right, currNode.Left
		if currNode.Left != nil {
			queue = append(queue, currNode.Left)
		}
		if currNode.Right != nil {
			queue = append(queue, currNode.Right)
		}
	}
	return root
}

func InvertTreeDFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	InvertTreeDFS(root.Left)
	InvertTreeDFS(root.Right)

	return root

	/*
		Time Complexity: O(n)
		Space Complexity: O(n) for recursion stack
	*/
}

func InvertTree(root *TreeNode) *TreeNode {
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curNode.Left, curNode.Right = curNode.Right, curNode.Left
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
	}
	return root
}
