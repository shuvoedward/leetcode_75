package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	currentNode := head

	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		digit := sum % 10
		carry = sum / 10

		currentNode.Next = &ListNode{Val: digit}
		currentNode = currentNode.Next
	}

	return head.Next
}
