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

func invertTreeBFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil // represents zero value
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// 1. swap it
		node.Left, node.Right = node.Right, node.Left

		// 2. check if nil, if not append
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}

	return root

}

// TODO: Time complexity for both of them is O(n)

func interativeDFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	interativeDFS(root.Left)
	interativeDFS(root.Right)

	return root
}

/*
func main(){
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{Val: 6},
		},
	}

	inverted := invertTree(tree)
}
*/
