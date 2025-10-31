package tree

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// if any of the nodes are nil, no lca to find
	if root == nil || p == nil || q == nil {
		return nil
	}

	// max = both p and q are less than the current root. both nodes must be located in the left subtree
	if max(p.Val, q.Val) < root.Val {
		return LowestCommonAncestor(root.Left, p, q)
		// min = both p and q are greater than current root. both nodes must be located in the right subtree
	} else if min(p.Val, q.Val) > root.Val {
		return LowestCommonAncestor(root.Right, p, q)
		// p or q is on the left and the other is on the right, or one of them is equal to root
	} else {
		return root
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
