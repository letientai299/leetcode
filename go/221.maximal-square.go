package main

// Maximal Square
//
// Medium
//
// https://leetcode.com/problems/maximal-square/
//
// Given an `m x n` binary `matrix` filled with `0`'s and `1`'s, _find the
// largest square containing only_ `1`'s _and return its area_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/11/26/max1grid.jpg)
//
// ```
// Input: matrix =
// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
// Output: 4
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2020/11/26/max2grid.jpg)
//
// ```
// Input: matrix = [["0","1"],["1","0"]]
// Output: 1
//
// ```
//
// **Example 3:**
//
// ```
// Input: matrix = [["0"]]
// Output: 0
//
// ```
//
// **Constraints:**
//
// - `m == matrix.length`
// - `n == matrix[i].length`
// - `1 <= m, n <= 300`
// - `matrix[i][j]` is `'0'` or `'1'`.
func maximalSquare(mat [][]byte) int {
	m := len(mat)
	n := len(mat[0])

	s := make([]int, n)

	best := 0
	for y := 0; y < m; y++ {
		if mat[y][0] == '1' {
			best = 1
		}
	}

	for x := 0; x < n; x++ {
		if mat[0][x] == '1' {
			s[x] = 1
			best = 1
		}
	}

	for y := 1; y < m; y++ {
		s[0] = int(mat[y][0] - '0')

		for x := 1; x < n; x++ {
			if mat[y][x] == '0' {
				s[x] = 0
				continue
			}

			a, b := s[x-1], s[x]
			side := 1
			if a != 0 && b != 0 {
				side = a
				if side > b {
					side = b
				}

				if mat[y-side][x-side] == '1' {
					side++
				}
			}

			s[x] = side
			if best < side {
				best = side
			}
		}
	}

	return best * best
}
