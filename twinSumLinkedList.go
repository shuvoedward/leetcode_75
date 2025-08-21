package main

func pairSum(head *ListNode) int {
	if head.Next.Next == nil {
		return head.Val + head.Next.Val
	}
	slow := head
	fast := head.Next
	maxSum := 0

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondList := slow.Next
	slow.Next = nil

	// reverse second linked list
	var prev *ListNode = nil

	for secondList != nil {
		temp := secondList.Next
		secondList.Next = prev
		prev = secondList
		secondList = temp
	}

	secondList = prev
	firstList := head
	for secondList != nil {
		sum := firstList.Val + secondList.Val
		if maxSum < sum {
			maxSum = sum
		}
		secondList = secondList.Next
		firstList = firstList.Next
	}
	return maxSum

}

func pairSum2(head *ListNode) int {
	// find middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse second part of linked list
	var pre *ListNode
	cur := slow
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	// find the max
	result := 0
	for pre != nil {
		result = max(result, head.Val+pre.Val)
		head = head.Next
		pre = pre.Next
	}
	return result
}
