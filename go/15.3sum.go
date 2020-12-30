package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	return threeSumGeneral(nums, 0)
}

func threeSumGeneral(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for i := 0; i < n; {
		a := nums[i]
		if a > target {
			break // remaining number is larger than target, no need to check further
		}

		for j, k := i+1, n-1; j < k; {
			b := nums[j]
			c := nums[k]
			if a+b+c > target {
				k--
				continue
			}
			if a+b+c < target {
				j++
				continue
			}

			res = append(res, []int{a, b, c})
			for j > i+1 && j < k && nums[j] == b {
				j++
			}

			for k > j && nums[k] == c {
				k--
			}
		}

		for i < len(nums) && nums[i] == a {
			i++
		}
	}
	return res
}
