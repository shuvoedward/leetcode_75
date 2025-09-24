package main

import (
	"fmt"
	"shuvoedward/leetcode/leetcode75"
)

func main() {
	// node1 := &leetcode.ListNode{Val: 2, Next: &leetcode.ListNode{Val: 4, Next: &leetcode.ListNode{Val: 3, Next: nil}}}
	// node2 := &leetcode.ListNode{Val: 5, Next: &leetcode.ListNode{Val: 6, Next: &leetcode.ListNode{Val: 8, Next: &leetcode.ListNode{Val: 2, Next: nil}}}}

	result := leetcode75.CanPlaceFlowers([]int{1, 0, 0}, 2)
	fmt.Println(result)

}
