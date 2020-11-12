package main

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	if nums[0] >= len(nums) {
		return len(nums)
	}

	for r := 1; r < len(nums); r++ {
		x := len(nums) - r
		if nums[r] >= x && nums[r-1] < x {
			return x
		}
	}
	return -1
}
