package main

import (
	"sort"
)

// Combination Sum IV
//
// Medium
//
// https://leetcode.com/problems/combination-sum-iv/
//
// Given an array of **distinct** integers `nums` and a target integer `target`,
// return _the number of possible combinations that add up to_ `target`.
//
// The answer is **guaranteed** to fit in a **32-bit** integer.
//
// **Example 1:**
//
// ```
// Input: nums = [1,2,3], target = 4
// Output: 7
// Explanation:
// The possible combination ways are:
// (1, 1, 1, 1)
// (1, 1, 2)
// (1, 2, 1)
// (1, 3)
// (2, 1, 1)
// (2, 2)
// (3, 1)
// Note that different sequences are counted as different combinations.
//
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [9], target = 3
// Output: 0
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 200`
// - `1 <= nums[i] <= 1000`
// - All the elements of `nums` are **unique**.
// - `1 <= target <= 1000`
//
// **Follow up:** What if negative numbers are allowed in the given array? How
// does it change the problem? What limitation we need to add to the question to
// allow negative numbers?
func combinationSum4(nums []int, target int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	calc := func(takes []int) int {
		r := 1
		i := 1
		for _, v := range takes {
			if v == 0 {
				continue
			}
			for j := 1; j <= v; j++ {
				r *= i
				i++
				r /= j
			}
		}
		return r
	}

	var walk func(i, cur int, takes []int) int
	walk = func(i, cur int, takes []int) int {
		if cur == 0 {
			return calc(takes)
		}

		if i >= len(nums) {
			return 0
		}

		total := 0
		step := cur / nums[i]
		for j := 0; j <= step; j++ {
			rem := cur - j*nums[i]
			takes[i] = j
			total += walk(i+1, rem, takes)
		}

		takes[i] = 0
		return total
	}

	takes := make([]int, len(nums))
	return walk(0, target, takes)
	// TODO (tai): TLE
}
