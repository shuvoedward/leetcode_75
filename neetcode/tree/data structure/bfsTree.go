package tree

// TODO: 1. Define a tree node struct type
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TODO: 2. Function to initialize tree node
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

// TODO: 3. To group methods together
type InsertAndRemove struct{}

// 4. Insert a new node and return the root of the binary search tree
func (ir *InsertAndRemove) Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return NewTreeNode(val)
		/*
			when intiailizing a new treenode with only value
				Val : val,
				Left : nil, // not initialized, zero value for pointer
				Right: nil,
		*/
	}

	if val > root.Val {
		root.Right = ir.Insert(root.Right, val)
	} else if val < root.Val {
		root.Left = ir.Insert(root.Left, val)
	}

	return root
	// worst case - visit one node every level of the tree, which is the hight h of the tree.
	// logn in a balanced tree

}

func (ir *InsertAndRemove) MinValueNode(root *TreeNode) *TreeNode {
	curr := root // curr and root both pointing to the same node in memory. both hold the same address
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr

}

func (ir *InsertAndRemove) Remove(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val > root.Val {
		root.Right = ir.Remove(root.Right, val)
	} else if val < root.Val {
		root.Left = ir.Remove(root.Left, val)
	} else {
		// found the value
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minNode := ir.MinValueNode(root.Right)
			root.Val = minNode.Val
			root.Right = ir.Remove(root.Right, minNode.Val)
		}

	}

}
