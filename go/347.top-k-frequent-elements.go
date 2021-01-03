package main

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 *
 * https://leetcode.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (57.62%)
 * Total Accepted:    511.8K
 * Total Submissions: 824.4K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * Given a non-empty array of integers, return the k most frequent elements.
 *
 * Example 1:
 *
 *
 * Input: nums = [1,1,1,2,2,3], k = 2
 * Output: [1,2]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1], k = 1
 * Output: [1]
 *
 *
 * Note:
 *
 *
 * You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
 * Your algorithm's time complexity must be better than O(n log n), where n is
 * the array's size.
 * It's guaranteed that the answer is unique, in other words the set of the top
 * k frequent elements is unique.
 * You can return the answer in any order.
 *
 *
 */
func topKFrequent347(nums []int, k int) []int {
	m := make(map[int]int)
	top := 0
	for _, v := range nums {
		m[v]++
		if top < m[v] {
			top = m[v]
		}
	}

	res := make([]int, 0, k)

	for k > 0 {
		next := 0
		for x, v := range m {
			if v == top {
				res = append(res, x)
				k--
			} else if v < top && v > next {
				next = v
			}
		}
		top = next
	}
	return res
}
