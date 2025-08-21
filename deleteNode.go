package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		// key == root.Val

		// has one child or no child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		successor := findMin(root.Right)

		root.Val = successor.Val

		// delete successor
		root.Right = deleteNode(root.Right, successor.Val)

	}
	return root

}

func findMin(root *TreeNode) *TreeNode {
	if root.Left != nil {
		root = root.Left
	}
	return root
}
