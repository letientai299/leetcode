package main

// Maximum Size Subarray Sum Equals k
//
// Medium
//
// https://leetcode.com/problems/maximum-size-subarray-sum-equals-k/
func maxSubArrayLen(nums []int, k int) int {
	m := make(map[int]int, len(nums)+1)
	s := 0
	m[0] = -1
	best := 0
	for i, n := range nums {
		s += n
		pre, ok := m[s-k]
		if ok {
			if best < i-pre {
				best = i - pre
			}
		}

		if _, ok := m[s]; !ok {
			m[s] = i
		}
	}

	if s == k {
		return len(nums)
	}

	return best
}
