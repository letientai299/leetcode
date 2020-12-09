package main

/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 *
 * https://leetcode.com/problems/search-a-2d-matrix-ii/description/
 *
 * algorithms
 * Medium (41.94%)
 * Total Accepted:    375.6K
 * Total Submissions: 859K
 * Testcase Example:  '[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n' +
  '5'
 *
 * Write an efficient algorithm that searches for a target value in an m x n
 * integer matrix. The matrix has the following properties:
 *
 *
 * Integers in each row are sorted in ascending from left to right.
 * Integers in each column are sorted in ascending from top to bottom.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: matrix =
 * [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],
 * target = 5
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: matrix =
 * [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],
 * target = 20
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= n, m <= 300
 * -10^9 <= matix[i][j] <= 10^9
 * All the integers in each row are sorted in ascending order.
 * All the integers in each column are sorted in ascending order.
 * -10^9 <= target <= 10^9
 *
 *
*/

func searchMatrix(mat [][]int, t int) bool {
	m := len(mat)
	if m == 0 {
		return false
	}
	n := len(mat[0])

	var find func(sx, ex, sy, ey int) bool
	find = func(sx, ex, sy, ey int) bool {
		if sx > ex || sy > ey {
			return false
		}

		if sy == ey {
			// single line
			for i := sx; i <= ex; i++ {
				if mat[sy][i] == t {
					return true
				}
			}
			return false
		}

		if sx == ex {
			// single line
			for i := sy; i <= ey; i++ {
				if mat[i][sx] == t {
					return true
				}
			}
			return false
		}

		cx := (sx + ex) / 2
		cy := (sy + ey) / 2
		v := mat[cy][cx]
		if t == v {
			return true
		}

		if t < v {
			return find(sx, cx, sy, cy) ||
				find(cx+1, ex, sy, cy-1) ||
				find(sx, cx-1, cy+1, ey)
		}

		return find(cx+1, ex, sy, cy) ||
			find(sx, cx, cy+1, ey) ||
			find(cx+1, ex, cy+1, ey)
	}

	return find(0, n-1, 0, m-1)
}
