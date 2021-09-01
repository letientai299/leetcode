package main

// Max Consecutive Ones III
//
// Medium
//
// https://leetcode.com/problems/max-consecutive-ones-iii/
//
// Given a binary array `nums` and an integer `k`, return _the maximum number of
// consecutive_ `1` _'s in the array if you can flip at most_ `k``0`'s.
//
// **Example 1:**
//
// ```
// Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
// Output: 6
// Explanation: [1,1,1,0,0,1,1,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
// ```
//
// **Example 2:**
//
// ```
// Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
// Output: 10
// Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
//
// ```
//
// **Constraints:**
//
// - `1 <= nums.length <= 105`
// - `nums[i]` is either `0` or `1`.
// - `0 <= k <= nums.length`
func longestOnes(nums []int, k int) int {
	best := 0
	ones := 0
	zeroes := 0
	l, r := 0, 0
	for ; r < len(nums); r++ {
		if nums[r] == 1 {
			ones += nums[r]
			continue
		}
		zeroes++

		if zeroes <= k {
			continue
		}

		if ones+k > best {
			best = ones + k
		}

		for l < r && nums[l] == 1 {
			ones--
			l++
		}

		l++
		zeroes--
	}

	if zeroes > k {
		zeroes = k
	}
	if ones+zeroes > best {
		best = ones + zeroes
	}

	return best
}
