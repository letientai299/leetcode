package main

import "sort"

func minSubsequence(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	var res []int
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	s := 0
	for i := range nums {
		s += nums[i]
		res = append(res, nums[i])
		if s > total-s {
			break
		}
	}
	return res
}
