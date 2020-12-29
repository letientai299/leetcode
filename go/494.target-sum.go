package main

/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 *
 * https://leetcode.com/problems/target-sum/description/
 *
 * algorithms
 * Medium (46.01%)
 * Total Accepted:    204.3K
 * Total Submissions: 445.1K
 * Testcase Example:  '[1,1,1,1,1]\n3'
 *
 * You are given a list of non-negative integers, a1, a2, ..., an, and a
 * target, S. Now you have 2 symbols + and -. For each integer, you should
 * choose one from + and - as its new symbol.
 *
 * Find out how many ways to assign symbols to make sum of integers equal to
 * target S.
 *
 * Example 1:
 *
 *
 * Input: nums is [1, 1, 1, 1, 1], S is 3.
 * Output: 5
 * Explanation:
 *
 * -1+1+1+1+1 = 3
 * +1-1+1+1+1 = 3
 * +1+1-1+1+1 = 3
 * +1+1+1-1+1 = 3
 * +1+1+1+1-1 = 3
 *
 * There are 5 ways to assign symbols to make the sum of nums be target 3.
 *
 *
 *
 * Constraints:
 *
 *
 * The length of the given array is positive and will not exceed 20.
 * The sum of elements in the given array will not exceed 1000.
 * Your output answer is guaranteed to be fitted in a 32-bit integer.
 *
 *
 */
// TODO (tai): slow, 51%
func findTargetSumWays(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	i, j := 0, 0
	for ; j < len(nums); j++ {
		if nums[j] == 0 {
			continue
		}
		nums[i] = nums[j]
		i++
	}

	zeroes := n - i
	nums = nums[:i]

	if k < 0 {
		k = -k
	}

	all := sum(nums...)
	d := all - k

	if d == 0 {
		return 1<<zeroes
	}


	if d%2 != 0 || d < 0 {
		return 0
	}

	d /= 2

	res := 0
	var find func(i int, d int)
	find = func(i int, d int) {
		if d == 0 {
			res++
			return
		}

		if i >= len(nums) || d < 0 {
			return
		}

		find(i+1, d)
		if d >= nums[i] {
			find(i+1, d-nums[i])
		}
	}

	for i, v := range nums {
		find(i+1, d-v)
	}

	return res * (1 << zeroes)
}
