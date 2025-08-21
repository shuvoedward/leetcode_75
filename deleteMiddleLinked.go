package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	prev := head
	slow := head.Next
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		prev = prev.Next
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = slow.Next

	return head

}
