package main

/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 *
 * https://leetcode.com/problems/sort-array-by-parity/description/
 *
 * algorithms
 * Easy (73.75%)
 * Total Accepted:    268.9K
 * Total Submissions: 359.2K
 * Testcase Example:  '[3,1,2,4]'
 *
 * Given an array A of non-negative integers, return an array consisting of all
 * the even elements of A, followed by all the odd elements of A.
 *
 * You may return any answer array that satisfies this condition.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [3,1,2,4]
 * Output: [2,4,3,1]
 * The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 5000
 * 0 <= A[i] <= 5000
 *
 *
 *
 */
func sortArrayByParity(a []int) []int {
	n := len(a)
	for x, y := 0, n-1; x < y; {
		for x < n && a[x]%2 == 0 {
			x++
		}

		for y >= 0 && a[y]%2 != 0 {
			y--
		}

		if x < y {
			a[x], a[y] = a[y], a[x]
			x++
			y--
		}
	}
	return a
}
