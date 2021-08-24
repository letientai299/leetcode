package main

import "math"

// Maximum Matrix Sum
//
// Medium
//
// https://leetcode.com/problems/maximum-matrix-sum/
//
// You are given an `n x n` integer `matrix`. You can do the following operation
// **any** number of times:
//
// - Choose any two **adjacent** elements of `matrix` and **multiply** each of
// them by `-1`.
//
// Two elements are considered **adjacent** if and only if they share a
// **border**.
//
// Your goal is to **maximize** the summation of the matrix's elements. Return
// _the **maximum** sum of the matrix's elements using the operation mentioned
// above._
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/07/16/pc79-q2ex1.png)
//
// ```
// Input: matrix = [[1,-1],[-1,1]]
// Output: 4
// Explanation: We can follow the following steps to reach sum equals 4:
// - Multiply the 2 elements in the first row by -1.
// - Multiply the 2 elements in the first column by -1.
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/07/16/pc79-q2ex2.png)
//
// ```
// Input: matrix = [[1,2,3],[-1,-2,-3],[1,2,3]]
// Output: 16
// Explanation: We can follow the following step to reach sum equals 16:
// - Multiply the 2 last elements in the second row by -1.
//
// ```
//
// **Constraints:**
//
// - `n == matrix.length == matrix[i].length`
// - `2 <= n <= 250`
// - `-105 <= matrix[i][j] <= 105`
func maxMatrixSum(matrix [][]int) int64 {
	pos := math.MaxInt32 // smallest positive
	neg := math.MinInt32 // largest negative
	negCnt := 0
	zeroCount := 0

	sum := int64(0)
	for _, row := range matrix {
		for _, v := range row {
			if v == 0 {
				zeroCount++
				continue
			}

			if v < 0 {
				if v > neg {
					neg = v
				}
				negCnt++
				sum += int64(-v)
				continue
			}

			if v < pos {
				pos = v
			}
			sum += int64(v)
		}
	}

	if negCnt%2 == 1 && zeroCount == 0 {
		if pos+neg < 0 {
			sum -= 2 * int64(pos)
		} else {
			sum += 2 * int64(neg)
		}
	}
	return sum
}
