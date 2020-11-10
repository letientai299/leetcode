package main

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return []int{0}
	}

	ids := make(map[int][]int, len(nums))
	for i, v := range nums {
		ids[v] = append(ids[v], i)
	}

	sort.Ints(nums)
	smallerCount := 0
	prev := nums[0]

	res := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			continue
		}

		smallerCount += len(ids[prev])
		prev = nums[i]
		for _, x := range ids[prev] {
			res[x] = smallerCount
		}
	}
	return res
}
