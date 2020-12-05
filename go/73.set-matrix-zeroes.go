package main

/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 *
 * https://leetcode.com/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (41.38%)
 * Total Accepted:    371.8K
 * Total Submissions: 848K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * Given an m x n matrix. If an element is 0, set its entire row and column to
 * 0. Do it in-place.
 *
 * Follow up:
 *
 *
 * A straight forward solution using O(mn) space is probably a bad idea.
 * A simple improvement uses O(m + n) space, but still not the best
 * solution.
 * Could you devise a constant space solution?
 *
 *
 *
 * Example 1:
 *
 *
 * Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
 * Output: [[1,0,1],[0,0,0],[1,0,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
 * Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[0].length
 * 1 <= m, n <= 200
 * -2^31 <= matrix[i][j] <= 2^31 - 1
 */
func setZeroes(mat [][]int) {
	m := len(mat)
	n := len(mat[0])
	clearX := false
	clearY := false
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if mat[y][x] == 0 {
				if y == 0 {
					clearX = true
				}

				if x == 0 {
					clearY = true
				}

				if x != 0 && y != 0 {
					mat[0][x] = 0
					mat[y][0] = 0
				}
			}
		}
	}

	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			if mat[y][0] == 0 || mat[0][x] == 0 {
				mat[y][x] = 0
			}
		}
	}

	if clearX {
		for x := 0; x < n; x++ {
			mat[0][x] = 0
		}
	}

	if clearY {
		for y := 0; y < m; y++ {
			mat[y][0] = 0
		}
	}
}
