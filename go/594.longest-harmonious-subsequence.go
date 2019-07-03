package main

/*
 * @lc app=leetcode id=594 lang=golang
 *
 * [594] Longest Harmonious Subsequence
 *
 * https://leetcode.com/problems/longest-harmonious-subsequence/description/
 *
 * algorithms
 * Easy (44.07%)
 * Total Accepted:    37.8K
 * Total Submissions: 85.8K
 * Testcase Example:  '[1,3,2,2,5,2,3,7]'
 *
 * We define a harmounious array as an array where the difference between its
 * maximum value and its minimum value is exactly 1.
 *
 * Now, given an integer array, you need to find the length of its longest
 * harmonious subsequence among all its possible subsequences.
 *
 * Example 1:
 *
 *
 * Input: [1,3,2,2,5,2,3,7]
 * Output: 5
 * Explanation: The longest harmonious subsequence is [3,2,2,2,3].
 *
 *
 *
 *
 * Note: The length of the input array will not exceed 20,000.
 *
 */
func findLHS(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	res := make(map[int]int, len(nums))
	for _, num := range nums {
		if _, exist := res[num]; exist {
			res[num]++
			continue
		}
		res[num] = 1
	}

	longest := 0
	for k, c := range res {
		if n, exist := res[k+1]; exist {
			if c+n > longest {
				longest = c + n
			}
		}
	}

	return longest
}
