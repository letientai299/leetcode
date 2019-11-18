package main

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 *
 * https://leetcode.com/problems/subsets/description/
 *
 * algorithms
 * Medium (55.75%)
 * Total Accepted:    443.8K
 * Total Submissions: 788.6K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a set of distinct integers, nums, return all possible subsets (the
 * power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: nums = [1,2,3]
 * Output:
 * [
 * ⁠ [3],
 *  [1],
 *  [2],
 *  [1,2,3],
 *  [1,3],
 *  [2,3],
 *  [1,2],
 *  []
 * ]
 *
 */
func subsets(nums []int) [][]int {
	n := len(nums)

	// pick n elements from the given sets
	var pick func(k int, start int) [][]int

	pick = func(k int, start int) [][]int {
		if k == 0 {
			return [][]int{{}}
		}

		res := make([][]int, 0)
		for i := start; i < n-k+1; i++ {
			subs := pick(k-1, i+1)
			for _, s := range subs {
				res = append(res, append(s, nums[i]))
			}
		}

		return res
	}

	res := make([][]int, 0, 1<<uint32(n))
	for i := 0; i <= n; i++ {
		res = append(res, pick(i, 0)...)
	}

	return res
}
