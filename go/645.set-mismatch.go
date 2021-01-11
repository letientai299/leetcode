package main

/*
 * @lc app=leetcode id=645 lang=golang
 *
 * [645] Set Mismatch
 *
 * https://leetcode.com/problems/set-mismatch/description/
 *
 * algorithms
 * Easy (40.83%)
 * Total Accepted:    49.6K
 * Total Submissions: 121.4K
 * Testcase Example:  '[1,2,2,4]'
 *
 *
 * The set S originally contains numbers from 1 to n. But unfortunately, due to
 * the data error, one of the numbers in the set got duplicated to another
 * number in the set, which results in repetition of one number and loss of
 * another number.
 *
 *
 *
 * Given an array nums representing the data status of this set after the
 * error. Your task is to firstly find the number occurs twice and then find
 * the number that is missing. Return them in the form of an array.
 *
 *
 *
 * Example 1:
 *
 * Input: nums = [1,2,2,4]
 * Output: [2,3]
 *
 *
 *
 * Note:
 *
 * The given array size will in the range [2, 10000].
 * The given array's numbers won't have any order.
 *
 *
 */
func findErrorNums(nums []int) []int {
	// sometime the simplest solution works best.

	m := make([]int, len(nums))
	r := make([]int, 2)
	for _, n := range nums {
		if m[n-1] == 0 {
			m[n-1] = 1
		} else {
			r[0] = n
		}
	}
	for i, _ := range m {
		if m[i] == 0 {
			r[1] = i + 1
			break
		}
	}
	return r
}
