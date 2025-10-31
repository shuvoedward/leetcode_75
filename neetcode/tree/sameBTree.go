package tree

func dfsIsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil && p.Val == q.Val {
		return dfsIsSameTree(p.Left, q.Left) && dfsIsSameTree(p.Right, q.Right)
	}

	return false
}

func iterativeDfsIsSameTree(p *TreeNode, q *TreeNode) bool {
	type Pair struct {
		first, second *TreeNode
	}

	stack := []Pair{{p, q}}

	for len(stack) > 0 {
		lastIdx := len(stack) - 1
		node1, node2 := stack[lastIdx].first, stack[lastIdx].second
		stack = stack[:lastIdx]

		if node1 == nil && node2 == nil {
			continue
		}

		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}

		stack = append(stack, Pair{node1.Right, node2.Right})
		stack = append(stack, Pair{node1.Left, node2.Left})
	}

	return true
}

func bfsIsSameTree(p, q *TreeNode) bool {
	queue1 := []*TreeNode{p}
	queue2 := []*TreeNode{q}

	for len(queue1) > 0 && len(queue2) > 0 {
		for i := len(queue1); i > 0; i-- {
			nodeP := queue1[0]
			nodeQ := queue2[0]
			queue1 = queue1[1:]
			queue2 = queue2[1:]

			if nodeP == nil && nodeQ == nil {
				continue
			}

			if nodeP == nil || nodeQ == nil || nodeP.Val != nodeQ.Val {
				return false
			}

			queue1 = append(queue1, nodeP.Left, nodeP.Right)
			queue2 = append(queue2, nodeQ.Left, nodeQ.Right)
		}
	}

	return len(queue1) == 0 && len(queue2) == 0
}
