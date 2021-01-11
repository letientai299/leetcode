package main

/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 *
 * https://leetcode.com/problems/hamming-distance/description/
 *
 * algorithms
 * Easy (70.16%)
 * Total Accepted:    227.8K
 * Total Submissions: 324.7K
 * Testcase Example:  '1\n4'
 *
 * The Hamming distance between two integers is the number of positions at
 * which the corresponding bits are different.
 *
 * Given two integers x and y, calculate the Hamming distance.
 *
 * Note:
 * 0 ≤ x, y < 231.
 *
 *
 * Example:
 *
 * Input: x = 1, y = 4
 *
 * Output: 2
 *
 * Explanation:
 * 1   (0 0 0 1)
 * 4   (0 1 0 0)
 * ⁠      ↑   ↑
 *
 * The above arrows point to positions where the corresponding bits are
 * different.
 *
 *
 */
func hammingDistance(x int, y int) int {
	v := x ^ y
	c := 0
	c = v - ((v >> 1) & 0x55555555)
	c = ((c >> 2) & 0x33333333) + (c & 0x33333333)
	c = ((c >> 4) + c) & 0x0F0F0F0F
	c = ((c >> 8) + c) & 0x00FF00FF
	c = ((c >> 16) + c) & 0x0000FFFF
	return c
}
