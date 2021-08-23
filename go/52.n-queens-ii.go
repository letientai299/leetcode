package main

// N-Queens II
//
// Hard
//
// https://leetcode.com/problems/n-queens-ii/
//
// The **n-queens** puzzle is the problem of placing `n` queens on an `n x n`
// chessboard such that no two queens attack each other.
//
// Given an integer `n`, return _the number of distinct solutions to
// the **n-queens puzzle**_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/11/13/queens.jpg)
//
// ```
// Input: n = 4
// Output: 2
// Explanation: There are two distinct solutions to the 4-queens puzzle as
// shown.
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 1
// Output: 1
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 9`
func totalNQueens(n int) int {
	var res int
	var queens [][]int
	var walk func(y int)
	walk = func(y int) {
		if y >= n {
			res++
			return
		}

		for x := 0; x < n; x++ {
			ok := true
			for _, q := range queens {
				if q[0] == x ||
					(x-q[0] == y-q[1]) ||
					(q[0]-x == y-q[1]) {
					ok = false
					break
				}
			}

			if ok {
				queens = append(queens, []int{x, y})
				walk(y + 1)
				queens = queens[:y]
			}
		}
	}

	walk(0)
	return res
}
