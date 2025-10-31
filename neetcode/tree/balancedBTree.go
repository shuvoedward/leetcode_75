package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
func isBalanced(root *TreeNode) bool {
	return dfsBalanced(root).balanced
}

type Result struct {
	balanced bool
	height   int
}

func dfsBalanced(root *TreeNode) Result {
	if root == nil {
		return Result{true, 0}
	}

	left := dfsBalanced(root.Left)
	right := dfsBalanced(root.Right)

	balanced := left.balanced && right.balanced &&
		abs(left.height-right.height) <= 1

	return Result{balanced: balanced, height: max(left.height, right.height)}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
