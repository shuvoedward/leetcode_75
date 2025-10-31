package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

type DFS struct{}

// InOrder traversal
// left to right,
// left, then current node, right
// similar to array traversal [1, 2, 3, 4, 5]
func (d *DFS) InOrder(root *TreeNode) {
	if root == nil {
		return
	}

	d.InOrder(root.Left)
	fmt.Println(root.Val)
	d.InOrder(root.Right)

	// time complexity - O(n) - size of the tree - traverse every node

}

func (d *DFS) ReverseOrder(root *TreeNode) {
	if root == nil {
		return
	}

	d.ReverseOrder(root.Right)
	fmt.Println(root.Val)
	d.ReverseOrder(root.Left)
}

func (d *DFS) PreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	d.PreOrder(root.Left)
	d.PreOrder(root.Right)
}

func (d *DFS) PostOrder(root *TreeNode) {
	if root == nil {
		return
	}

	d.PostOrder(root.Left)
	d.PostOrder(root.Right)
	fmt.Println(root.Val)
}
