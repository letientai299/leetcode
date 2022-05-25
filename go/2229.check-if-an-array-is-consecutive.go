package main

// Check if an Array Is Consecutive
//
// Easy
//
// https://leetcode.com/problems/check-if-an-array-is-consecutive
func isConsecutive(nums []int) bool {
	// map, O(2n)
	m := make(map[int]struct{})
	max := nums[0]
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}

		if _, ok := m[n]; ok {
			// has duplicate
			return false
		}

		m[n] = struct{}{}
	}

	return max-min == len(nums)-1
	/*
		// O(nlog n), lower than map
		sort.Ints(nums)
		min := nums[0]
		for i, n := range nums {
			if min+i != n {
				return false
			}
		}
	*/

}
