package main

/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 *
 * https://leetcode.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (42.31%)
 * Total Accepted:    188.6K
 * Total Submissions: 445.8K
 * Testcase Example:  '3'
 *
 * Given a non-negative index k where k ≤ 33, return the kth index row of the
 * Pascal's triangle.
 *
 * Note that the row index starts from 0.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 3
 * Output: [1,3,3,1]
 *
 *
 * Follow up:
 *
 * Could you optimize your algorithm to use only O(k) extra space?
 *
 */
func getRow(n int) []int {
	n += 1

	if n == 0 {
		return []int{1}
	}

	res := make([]int, n)
	res[0] = 1
	p := int64(1)
	q := int64(1)

	// compute the first half
	for i := 1; i < (n/2)+1; i++ {
		p *= int64(n - i)
		q *= int64(i)
		if p%q == 0 {
			p /= q
			q = 1
		}
		res[i] = int(p / q)
	}

	// copy the first half
	for i := n / 2; i < n; i++ {
		res[i] = res[n-1-i]
	}
	return res
}
