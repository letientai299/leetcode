package main

// Count Square Submatrices with All Ones
//
// Medium
//
// https://leetcode.com/problems/count-square-submatrices-with-all-ones/
//
// Given a `m * n` matrix of ones and zeros, return how many **square**
// submatrices have all ones.
//
// **Example 1:**
//
// ```
// Input: matrix =
// [
//   [0,1,1,1],
//   [1,1,1,1],
//   [0,1,1,1]
// ]
// Output: 15
// Explanation:
// There are 10 squares of side 1.
// There are 4 squares of side 2.
// There is  1 square of side 3.
// Total number of squares = 10 + 4 + 1 = 15.
//
// ```
//
// **Example 2:**
//
// ```
// Input: matrix =
// [
//   [1,0,1],
//   [1,1,0],
//   [1,1,0]
// ]
// Output: 7
// Explanation:
// There are 6 squares of side 1.
// There is 1 square of side 2.
// Total number of squares = 6 + 1 = 7.
//
// ```
//
// **Constraints:**
//
// - `1 <= arr.length <= 300`
// - `1 <= arr[0].length <= 300`
// - `0 <= arr[i][j] <= 1`
func countSquares(matrix [][]int) int {
	res := 0
	my := len(matrix)
	mx := len(matrix[0])

	var walk func(x, y int)
	walk = func(x, y int) {
		n := 0
		for x+n < mx && y+n < my {
			if matrix[y+n][x+n] == 0 {
				return
			}

			for i := 0; i < n; i++ {
				if matrix[y+n][x+i] == 0 || matrix[y+i][x+n] == 0 {
					return
				}
			}

			res++
			n++
		}
	}

	for i := 0; i < my; i++ {
		for j := 0; j < mx; j++ {
			walk(j, i)
		}
	}
	return res
}
