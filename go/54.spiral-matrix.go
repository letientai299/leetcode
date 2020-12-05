package main

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (32.06%)
 * Total Accepted:    431.7K
 * Total Submissions: 1.2M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given an m x n matrix, return all elements of the matrix in spiral order.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [1,2,3,6,9,8,7,4,5]
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 10
 * -100 <= matrix[i][j] <= 100
 *
 *
 */

func spiralOrder(mat [][]int) []int {
	i, j := 1, 0
	x, y := -1, 0
	m, n := len(mat), len(mat[0])
	r := make([]int, 0, m*n)
	const mask = 101
	for len(r) != m*n {
		x += i
		y += j
		for x < n && x >= 0 && y >= 0 && y < m && mat[y][x] != mask {
			r = append(r, mat[y][x])
			mat[y][x] = mask
			x += i
			y += j
		}
		x -= i
		y -= j
		i, j = -j, i
	}
	return r
}
