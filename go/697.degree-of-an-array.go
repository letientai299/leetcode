package main

/*
 * @lc app=leetcode id=697 lang=golang
 *
 * [697] Degree of an Array
 *
 * https://leetcode.com/problems/degree-of-an-array/description/
 *
 * algorithms
 * Easy (50.76%)
 * Total Accepted:    53.3K
 * Total Submissions: 104.9K
 * Testcase Example:  '[1,2,2,3,1]'
 *
 * Given a non-empty array of non-negative integers nums, the degree of this
 * array is defined as the maximum frequency of any one of its elements.
 * Your task is to find the smallest possible length of a (contiguous) subarray
 * of nums, that has the same degree as nums.
 *
 * Example 1:
 *
 * Input: [1, 2, 2, 3, 1]
 * Output: 2
 * Explanation:
 * The input array has a degree of 2 because both elements 1 and 2 appear
 * twice.
 * Of the subarrays that have the same degree:
 * [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
 * The shortest length is 2. So return 2.
 *
 *
 *
 *
 * Example 2:
 *
 * Input: [1,2,2,3,1,4,2]
 * Output: 6
 *
 *
 *
 * Note:
 * nums.length will be between 1 and 50,000.
 * nums[i] will be an integer between 0 and 49,999.
 *
 */
func findShortestSubArray(nums []int) int {
	type stat struct{ first, freq int }

	// map of values and their number of appearance
	m := make(map[int]*stat)
	max := 0
	res := len(nums)
	for i, v := range nums {
		st, ok := m[v]
		if !ok {
			st = new(stat)
			st.first = i
		}
		st.freq++
		m[v] = st
		distance := i - st.first + 1

		if st.freq > max {
			max = st.freq
			res = distance
		} else if st.freq == max && res > distance {
			res = distance
		}
	}
	return res
}
