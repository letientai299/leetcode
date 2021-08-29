package main

import "math"

// Minimum Falling Path Sum
//
// Medium
//
// https://leetcode.com/problems/minimum-falling-path-sum/
//
// Given an `n x n` array of integers `matrix`, return _the **minimum sum** of
// any **falling path** through_ `matrix`.
//
// A **falling path** starts at any element in the first row and chooses the
// element in the next row that is either directly below or diagonally
// left/right. Specifically, the next element from position `(row, col)` will be
// `(row + 1, col - 1)`, `(row + 1, col)`, or `(row + 1, col + 1)`.
//
// **Example 1:**
//
// ```
// Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]
// Output: 13
// Explanation: There are two falling paths with a minimum sum underlined below:
// [[2,1,3],      [[2,1,3],
//  [6,5,4],       [6,5,4],
//  [7,8,9]]       [7,8,9]]
//
// ```
//
// **Example 2:**
//
// ```
// Input: matrix = [[-19,57],[-40,-5]]
// Output: -59
// Explanation: The falling path with a minimum sum is underlined below:
// [[-19,57],
//  [-40,-5]]
//
// ```
//
// **Example 3:**
//
// ```
// Input: matrix = [[-48]]
// Output: -48
//
// ```
//
// **Constraints:**
//
// - `n == matrix.length`
// - `n == matrix[i].length`
// - `1 <= n <= 100`
// - `-100 <= matrix[i][j] <= 100`
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix[0])
	for i, row := range matrix[1:] {
		pre := matrix[i]
		for j := range row {
			a := pre[j]
			if j > 0 && a > pre[j-1] {
				a = pre[j-1]
			}

			if j < n-1 && a > pre[j+1] {
				a = pre[j+1]
			}

			row[j] += a
		}
	}

	best := math.MaxInt32
	for _, v := range matrix[n-1] {
		if v < best {
			best = v
		}
	}
	return best
}
