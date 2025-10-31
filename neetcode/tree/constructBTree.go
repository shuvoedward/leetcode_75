package tree

import "math"

// Input: preorder = [1,2,3,4], inorder = [2,1,3,4]

// Output: [1,2,3,null,null,null,4]

/*
Preorder (Root -> Left -> Right):
The first element is always the root of the tree or subtree

Inorder (Left -> Root -> Right):
Elements before the root in the inorder belong to the left subtree, and elemnets
after it belong the right subtree.
*/

/*

Intuition Behind the Approach:
The main idea is that we process the preorder array to build
the tree top-down(root first), and use the inorder array to determine when
we've finished building a subtree.

*/

func BuildTreeOptimal(preorder, inorder []int) *TreeNode {

	preIdx, inIdx := 0, 0

	var dfs func(int) *TreeNode
	dfs = func(limit int) *TreeNode {
		if preIdx >= len(preorder) {
			return nil
		}
		if inorder[inIdx] == limit {
			inIdx++
			return nil
		}

		root := &TreeNode{Val: preorder[preIdx]}
		preIdx++

		root.Left = dfs(root.Val)
		root.Right = dfs(limit)

		return root
	}

	return dfs(math.MaxInt)
}

func BuildTreeHash(preorder, inorder []int) *TreeNode {
	indices := make(map[int]int)
	for i, val := range inorder {
		indices[val] = i
	}

	preIdx := 0

	var dfs func(int, int) *TreeNode
	dfs = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		rootVal := preorder[preIdx]
		preIdx++

		root := &TreeNode{Val: rootVal}
		mid := indices[rootVal]

		root.Left = dfs(left, mid-1)
		root.Right = dfs(mid+1, right)

		return root
	}

	return dfs(0, len(inorder)-1)
}
