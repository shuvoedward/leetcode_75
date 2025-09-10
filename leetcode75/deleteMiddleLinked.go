package leetcode75

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteMiddle(head *ListNode) *ListNode {

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

/*

fast = 2*slow
prev = slow - 1

odd length = 5 (0-4)
Fast reaches at the end of the linked list. Not to nil
fast = 4
slow = 2
prev = 1

even legth = 4 (0-3)
Fast reaches to nil.
fast = 4 (nil)
slow = 2
prev = 1

For linked list with length of 2.

1 -> 2 -> nil
prev = 1,
slow = 2
fast = nil
middle = 2
prev.Next = slow.Next / 1 -> nil


Linked list with Odd length

1 -> 2 -> 3 -> nil
middle = 2
prev = 1
slow = 2
fast = 3

loop won't run.
prev.Next -> slow.Next / 1 -> 3 -> nil


For even list length greater than 2

1 -> 2 -> 3 -> 4 -> nil
middle = 3
prev = 1
slow = 2
fast = 3

loop 1:
prev = 2
slow = 3
fast = nil
loop ends

prev.Next = slow.Next / 1 -> 2 -> 4 -> nil
*/
