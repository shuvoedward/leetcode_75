package tree

/*
h
tree := TreeNode{
	Val : 1,
	Left : &TreeNode{ // go creates a struct TreeNode and assigns its address to Left
		Val : 2,
		Left : nil,
		Right: nil,
	}
	Right : &TreeNode{
		Val : 3,
		Left : nil,
		Right : nil,
	}

}

*/

func recursiveDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(recursiveDFS(root.Left), recursiveDFS(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Time complexity O(n)
Space complexity O(n)
*/

// type TreeNode struct {
//     Val   int
//     Left  *TreeNode
//     Right *TreeNode
// }

// Iterative Depth First Search
func maxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type NodeDepth struct {
		node  *TreeNode
		depth int
	}

	stack := []NodeDepth{{root, 1}}
	maxDepth := 0

	for len(stack) > 0 {
		// Pop the top element
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.node != nil {
			if current.depth > maxDepth {
				maxDepth = current.depth
			}
			// Push right and left children onto the stack with incremented depth
			stack = append(stack, NodeDepth{current.node.Left, current.depth + 1})
			stack = append(stack, NodeDepth{current.node.Right, current.depth + 1})
		}
	}

	return maxDepth
}

func maxDepthBFS(root *TreeNode) int {
	if root != nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue) // number of nodes at current level

		for range levelSize {
			current := queue[0]
			queue = queue[1:]

			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		depth++
	}

	return depth
}
