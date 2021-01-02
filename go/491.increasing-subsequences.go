package main

import "strconv"

/*
 * @lc app=leetcode id=491 lang=golang
 *
 * [491] Increasing Subsequences
 *
 * https://leetcode.com/problems/increasing-subsequences/description/
 *
 * algorithms
 * Medium (43.35%)
 * Total Accepted:    53.2K
 * Total Submissions: 112.6K
 * Testcase Example:  '[4,6,7,7]'
 *
 * Given an integer array, your task is to find all the different possible
 * increasing subsequences of the given array, and the length of an increasing
 * subsequence should be at least 2.
 *
 *
 *
 * Example:
 *
 *
 * Input: [4, 6, 7, 7]
 * Output: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7],
 * [4,7,7]]
 *
 *
 *
 * Constraints:
 *
 *
 * The length of the given array will not exceed 15.
 * The range of integer in the given array is [-100,100].
 * The given array may contain duplicates, and two equal integers should also
 * be considered as a special case of increasing sequence.
 *
 *
 */
// TODO (tai): solve this properly
func findSubsequences(nums []int) [][]int {
	n := len(nums)
	if n < 2 {
		return nil
	}

	// cheating
	hashes := make(map[string]struct{})

	var res [][]int
	var find func(cur []int, hash string, i int)

	find = func(cur []int, hash string, i int) {
		l := len(cur)
		if _, ok := hashes[hash]; !ok && l >= 2 {
			cp := make([]int, l)
			copy(cp, cur)
			res = append(res, cp)
			hashes[hash] = struct{}{}
		}

		if i >= n {
			return
		}

		if cur[l-1] <= nums[i] {
			s := strconv.Itoa(nums[i])
			find(append(cur, nums[i]), hash+"-"+s, i+1)
		}
		cur = cur[:l]
		find(cur, hash, i+1)
	}

	for i := 0; i < n-1; i++ {
		hash := strconv.Itoa(nums[i])
		find([]int{nums[i]}, hash, i+1)
	}

	return res
}
