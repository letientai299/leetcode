package main

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 *
 * https://leetcode.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (44.42%)
 * Total Accepted:    310K
 * Total Submissions: 643.9K
 * Testcase Example:  '[1,2,2]'
 *
 * Given a collection of integers that might contain duplicates, nums, return
 * all possible subsets (the power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: [1,2,2]
 * Output:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 *
 *
 */

func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	m := make(map[int]int, n)
	for j := 0; j < n; j++ {
		m[nums[j]]++
	}

	n = 0
	for i := range m {
		nums[n] = i
		n++
	}

	nums = nums[:n]
	res := make([][]int, 0, n*n)
	res = append(res, []int{})
	for _, x := range nums {
		l := len(res)
		for k := 0; k < l; k++ {
			a := res[k]
			for i := 1; i <= m[x]; i++ {
				b := make([]int, len(a), len(a)+i)
				copy(b, a)
				for j := 0; j < i; j++ {
					b = append(b, x)
				}
				res = append(res, b)
			}
		}
	}
	return res
}
