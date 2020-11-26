package main

import "sort"

// medium

func largestDivisibleSubset(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	dp := make([]int, len(nums))
	pre := make([]int, len(nums))
	sort.Ints(nums)
	id := -1
	for i, v := range nums {
		dp[i] = 1
		pre[i] = -1
		for j := i - 1; j >= 0; j-- {
			if v%nums[j] != 0 {
				continue
			}

			if dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				pre[i] = j
			}
		}

		if id == -1 || dp[i] > dp[id] {
			id = i
		}
	}

	res := make([]int, 0, dp[id])

	for id != -1 {
		res = append(res, nums[id])
		id = pre[id]
	}

	return res
}
