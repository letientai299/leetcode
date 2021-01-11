package main

import "sort"

// medium
// https://leetcode.com/problems/search-in-rotated-sorted-array/

func search_33(nums []int, t int) int {
	if len(nums) == 0 {
		return -1
	}
	i := 0
	for ; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			break
		}
	}

	tmp := nums[i+1:]
	tmp = append(tmp, nums[:i+1]...)
	v := sort.SearchInts(tmp, t)
	if v == len(nums) || tmp[v] != t {
		return -1
	}
	return (v + i + 1) % (len(nums))
}
