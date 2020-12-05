package main

/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 *
 * https://leetcode.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (49.98%)
 * Total Accepted:    214.6K
 * Total Submissions: 386.9K
 * Testcase Example:  '3'
 *
 * Given a positive integer n, generate an n x n matrix filled with elements
 * from 1 to n^2 in spiral order.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 3
 * Output: [[1,2,3],[8,9,4],[7,6,5]]
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: [[1]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 20
 *
 *
 */
func generateMatrix(n int) [][]int {
	i, j := 1, 0
	x, y := -1, 0
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}
	v := 1
	for v <= n*n {
		x += i
		y += j
		for x < n && x >= 0 && y >= 0 && y < n && mat[y][x] == 0 {
			mat[y][x] = v
			v++
			x += i
			y += j
		}
		x -= i
		y -= j
		i, j = -j, i
	}
	return mat
}
