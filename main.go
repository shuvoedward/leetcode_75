package main

import "fmt"

func main() {
	// list := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 7,
	// 			Next: &ListNode{
	// 				Val:  4,
	// 				Next: nil,
	// 			},
	// 		},
	// 	},
	// }

	// sum := pairSum(list)
	// fmt.Println(sum)

	// spells := []int{5, 1, 3}
	// potions := []int{1, 2, 3, 4, 5}
	// success := 7

	// result := successfulPairsBruteForce(spells, potions, int64(success))
	result := minEatingSpeed([]int{3, 6, 7, 11}, 8)
	fmt.Println(result)

}
