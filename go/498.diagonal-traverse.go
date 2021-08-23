package main

/*
 * @lc app=leetcode id=498 lang=golang
 *
 * [498] Diagonal Traverse
 *
 * https://leetcode.com/problems/diagonal-traverse/description/
 *
 * algorithms
 * Medium (46.34%)
 * Total Accepted:    114.2K
 * Total Submissions: 228.1K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given a matrix of M x N elements (M rows, N columns), return all elements of
 * the matrix in diagonal order as shown in the below image.
 *
 *
 *
 * Example:
 *
 *
 * Input:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 *
 * Output:  [1,2,4,7,5,3,6,8,9]
 *
 * Explanation:
 *
 *
 *
 *
 *
 * Note:
 *
 * The total number of elements of the given matrix will not exceed 10,000.
 *
 */

func findDiagonalOrder498(mat [][]int) []int {
	m := len(mat)
	if m == 0 {
		return nil
	}

	if m == 1 {
		return mat[0]
	}
	n := len(mat[0])

	if n == 0 {
		return nil
	}

	res := make([]int, 0, m*n)
	up := true
	y, x := 0, 0
	dx, dy := 1, -1

	for len(res) < m*n {
		for 0 <= x && x < n && 0 <= y && y < m {
			res = append(res, mat[y][x])
			x += dx
			y += dy
		}

		up = !up
		dx, dy = dy, dx
		switch {
		case x < 0:
			x = 0
			if y >= m {
				y = m - 1
				x++
			}
		case y < 0:
			y = 0
			if x >= n {
				x = n - 1
				y++
			}
		case x >= n:
			x = n - 1
			y += 2
		case y >= m:
			y = m - 1
			x += 2
		}
	}

	return res
}
