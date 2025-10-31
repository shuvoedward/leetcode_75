package binarysearch

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	total := len(A) + len(B)
	half := (total + 1) / 2

	if len(B) < len(A) {
		A, B = B, A
	}

	l, r := 0, len(A)
	for l <= r { // l < r is not but l <= r because the last element, which is r. l = i +1, can equal r and the all the elements in A can be in the left partition
		i := (l + r) / 2 // start of right partition for A
		j := half - i    // j starts at the right partition for B

		Aleft := math.MinInt64
		if i > 0 { // if i 0th index, then because i is start of the right partition, there is no left. left stays -infinity
			Aleft = A[i-1]
		}
		// i can go out of index to the right. because, all the elements of A is to the left partitions.
		Aright := math.MaxInt64
		if i < len(A) { // out of bound check to the right
			Aright = A[i]
		}
		Bleft := math.MinInt64
		if j > 0 {
			Bleft = B[j-1] // the value before the right partition is max of left. its position is B[j-1]
		}
		Bright := math.MaxInt64
		if j < len(B) {
			Bright = B[j]
		}

		if Aleft <= Bright && Bleft <= Aright {
			if total%2 != 0 { // odd
				return float64(max(Aleft, Bleft)) // get the value that is end of the left partition. which in sorted array is the max in left partition.
			}
			return (float64(max(Aleft, Bleft)) + float64(min(Aright, Bright))) / 2.0
		} else if Aleft > Bright {
			r = i - 1
		} else {
			// Bleft > Aright
			// so decrease 1 from b and increase 1 from A
			// i is the right partition. l moves 1 to the right partition
			l = i + 1
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
