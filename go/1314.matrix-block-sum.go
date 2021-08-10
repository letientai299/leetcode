package main

// Matrix Block Sum
//
// Medium
//
// https://leetcode.com/problems/matrix-block-sum/
//
// Given a `m x n` matrix `mat` and an integer `k`, return _a matrix_ `answer`
// _where each_ `answer[i][j]` _is the sum of all elements_ `mat[r][c]` _for_:
//
// - `i - k <= r <= i + k,`
// - `j - k <= c <= j + k`, and
// - `(r, c)` is a valid position in the matrix.
//
// **Example 1:**
//
// ```
// Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 1
// Output: [[12,21,16],[27,45,33],[24,39,28]]
//
// ```
//
// **Example 2:**
//
// ```
// Input: mat = [[1,2,3],[4,5,6],[7,8,9]], k = 2
// Output: [[45,45,45],[45,45,45],[45,45,45]]
//
// ```
//
// **Constraints:**
//
// - `m == mat.length`
// - `n == mat[i].length`
// - `1 <= m, n, k <= 100`
// - `1 <= mat[i][j] <= 100`
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if x > 0 {
				mat[y][x] += mat[y][x-1]
			}
			if y > 0 {
				mat[y][x] += mat[y-1][x]
			}
			if x > 0 && y > 0 {
				mat[y][x] -= mat[y-1][x-1]
			}
		}
	}

	get := func(y, x int) int {
		if 0 > x || 0 > y {
			return 0
		}

		if m <= y {
			y = m - 1
		}

		if n <= x {
			x = n - 1
		}

		return mat[y][x]
	}

	res := make([][]int, 0, m)
	for y := 0; y < m; y++ {
		res = append(res, make([]int, n))
		for x := 0; x < n; x++ {
			top := y - k - 1
			bot := y + k
			left := x - k - 1
			right := x + k

			res[y][x] = get(bot, right) + get(top, left) - get(top, right) - get(bot, left)
		}
	}

	return res
}
