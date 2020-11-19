package main

import "sort"

// medium

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] <= nums[i-1] {
			continue
		}

		j := i + 1
		for ; j < len(nums); j++ {
			if nums[j] > nums[i-1] && (j == len(nums)-1 || nums[j+1] <= nums[i-1]) {
				break
			}
		}

		if j == len(nums) {
			j = i
		}
		nums[j], nums[i-1] = nums[i-1], nums[j]
		sort.Ints(nums[i:])
		return
	}

	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	return
}
