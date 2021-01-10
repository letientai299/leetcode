package main

// medium
// 1679. Max Number of K-Sum Pairs
func maxOperations(nums []int, k int) int {
	m := make(map[int]int, len(nums))
	r := 0
	for _, v := range nums {
		if m[k-v] > 0 {
			r++
			m[k-v]--
		} else {
			m[v]++
		}
	}
	return r
}
