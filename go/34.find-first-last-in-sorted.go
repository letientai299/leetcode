package main

import "sort"

// medium
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	i := sort.SearchInts(nums, target)
	if i == len(nums) || nums[i] != target {
		return []int{-1, -1}
	}
	j := i + 1
	for j < len(nums) && nums[j] == target {
		j++
	}
	return []int{i, j - 1}
}
