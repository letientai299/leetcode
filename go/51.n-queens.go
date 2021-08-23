package main

// N-Queens
//
// Hard
//
// https://leetcode.com/problems/n-queens/
//
// The **n-queens** puzzle is the problem of placing `n` queens on an `n x n`
// chessboard such that no two queens attack each other.
//
// Given an integer `n`, return _all distinct solutions to the **n-queens
// puzzle**_. You may return the answer in **any order**.
//
// Each solution contains a distinct board configuration of the n-queens'
// placement, where `'Q'` and `'.'` both indicate a queen and an empty space,
// respectively.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/11/13/queens.jpg)
//
// ```
// Input: n = 4
// Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// Explanation: There exist two distinct solutions to the 4-queens puzzle as
// shown above
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 1
// Output: [["Q"]]
//
// ```
//
// **Constraints:**
//
// - `1 <= n <= 9`
func solveNQueens(n int) [][]string {
	var res [][]string

	var queens [][]int
	gen := func() []string {
		board := make([][]byte, n)
		for i := 0; i < n; i++ {
			board[i] = make([]byte, n)
			for j := 0; j < n; j++ {
				board[i][j] = '.'
			}
		}

		for _, cell := range queens {
			board[cell[1]][cell[0]] = 'Q'
		}

		var r []string
		for _, b := range board {
			r = append(r, string(b))
		}
		return r
	}

	var walk func(y int)
	walk = func(y int) {
		if y >= n {
			res = append(res, gen())
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
