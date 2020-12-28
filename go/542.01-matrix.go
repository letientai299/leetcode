package main

/*
 * @lc app=leetcode id=542 lang=golang
 *
 * [542] 01 Matrix
 *
 * https://leetcode.com/problems/01-matrix/description/
 *
 * algorithms
 * Medium (37.59%)
 * Total Accepted:    108.3K
 * Total Submissions: 267.6K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * Given a matrix consists of 0 and 1, find the distance of the nearest 0 for
 * each cell.
 *
 * The distance between two adjacent cells is 1.
 *
 *
 *
 * Example 1:
 *
 *
 * Input:
 * [[0,0,0],
 * ⁠[0,1,0],
 * ⁠[0,0,0]]
 *
 * Output:
 * [[0,0,0],
 * [0,1,0],
 * [0,0,0]]
 *
 *
 * Example 2:
 *
 *
 * Input:
 * [[0,0,0],
 * ⁠[0,1,0],
 * ⁠[1,1,1]]
 *
 * Output:
 * [[0,0,0],
 * ⁠[0,1,0],
 * ⁠[1,2,1]]
 *
 *
 *
 *
 * Note:
 *
 *
 * The number of elements of the given matrix will not exceed 10,000.
 * There are at least one 0 in the given matrix.
 * The cells are adjacent in only four directions: up, down, left and right.
 *
 *
 */

// TODO (tai): TLE
func updateMatrix(mat [][]int) [][]int {
	// quick select: m*n*log(m*n)
	m := len(mat)
	if m == 0 {
		return nil
	}

	n := len(mat[0])
	if n == 0 {
		return nil
	}

	type point [2]int
	zeroes := make([]point, 0, m+n)
	for y, row := range mat {
		for x, c := range row {
			if c == 0 {
				zeroes = append(zeroes, point{x, y})
			}
		}
	}

	swap := func(i, j int) { zeroes[i], zeroes[j] = zeroes[j], zeroes[i] }
	find := func(i int) interface{} { return zeroes[i] }
	for y, row := range mat {
		for x, c := range row {
			if c != 0 {
				comp := func(a, b interface{}) int {
					p, q := a.(point), b.(point)
					d1 := abs(p[0]-x) + abs(p[1]-y)
					d2 := abs(q[0]-x) + abs(q[1]-y)
					if d1 == d2 {
						return 0
					}
					if d1 < d2 {
						return -1
					}
					return 1
				}

				p := quickSelectFunc(len(zeroes), 1, find, comp, swap).(point)
				mat[y][x] = abs(p[0]-x) + abs(p[1]-y)
			}
		}
	}

	return mat
}
