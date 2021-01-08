package main

/*
 * @lc app=leetcode id=1020 lang=golang
 *
 * [1020] Number of Enclaves
 *
 * https://leetcode.com/problems/number-of-enclaves/description/
 *
 * algorithms
 * Medium (55.12%)
 * Total Accepted:    23.3K
 * Total Submissions: 39.7K
 * Testcase Example:  '[[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]'
 *
 * Given a 2D array A, each cell is 0 (representing sea) or 1 (representing
 * land)
 *
 * A move consists of walking from one land square 4-directionally to another
 * land square, or off the boundary of the grid.
 *
 * Return the number of land squares in the grid for which we cannot walk off
 * the boundary of the grid in any number of moves.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
 * Output: 3
 * Explanation:
 * There are three 1s that are enclosed by 0s, and one 1 that isn't enclosed
 * because its on the boundary.
 *
 * Example 2:
 *
 *
 * Input: [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
 * Output: 0
 * Explanation:
 * All 1s are either on the boundary or can reach the boundary.
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 500
 * 1 <= A[i].length <= 500
 * 0 <= A[i][j] <= 1
 * All rows have the same size.
 *
 */

func numEnclaves(mat [][]int) int {
	m := len(mat)
	if m == 0 {
		return 0
	}

	n := len(mat[0])
	if n == 0 {
		return 0
	}

	var clear func(y, x int)
	clear = func(y, x int) {
		if x < 0 || x >= n || y < 0 || y >= m || mat[y][x] == 0 {
			return
		}
		mat[y][x] = 0

		// the order of this matrix walking matter
		clear(y, x+1)
		clear(y, x-1)
		clear(y+1, x)
		clear(y-1, x)
	}

	// done this loop before the nest loop make it more cache friendly
	for i := 0; i < n; i++ {
		if mat[0][i] == 1 {
			clear(0, i)
		}

		if mat[m-1][i] == 1 {
			clear(m-1, i)
		}
	}

	for i := 0; i < m; i++ {
		if mat[i][0] == 1 {
			clear(i, 0)
		}

		if mat[i][n-1] == 1 {
			clear(i, n-1)
		}
	}

	c := 0
	for _, row := range mat {
		for _, v := range row {
			c += v
		}
	}
	return c
}
