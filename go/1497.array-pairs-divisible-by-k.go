package main

// medium
// https://leetcode.com/problems/check-if-array-pairs-are-divisible-by-k/

func canArrange(a []int, k int) bool {
	m := make([]int, k)
	for _, v := range a {
		m[(v%k+k)%k]++
	}
	for i := 0; i <= k/2; i++ {
		if i == 0 || i == k-i {
			if m[i]%2 != 0 {
				return false
			}
		} else if m[i] != m[k-i] {
			return false
		}
	}
	return true
}
