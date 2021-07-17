package main

// Determine Whether Matrix Can Be Obtained By Rotation
//
// Easy
//
// https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation/
//
// Given two `n x n` binary matrices `mat` and `target`, return `true` _if it is
// possible to make_ `mat` _equal to_ `target` _by **rotating**_ `mat` _in
// **90-degree increments**, or_ `false` _otherwise._
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/05/20/grid3.png)
//
// ```
// Input: mat = [[0,1],[1,0]], target = [[1,0],[0,1]]
// Output: true
// Explanation: We can rotate mat 90 degrees clockwise to make mat equal target.
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/05/20/grid4.png)
//
// ```
// Input: mat = [[0,1],[1,1]], target = [[1,0],[0,1]]
// Output: false
// Explanation: It is impossible to make mat equal to target by rotating mat.
//
// ```
//
// **Example 3:**
//
// ![](https://assets.leetcode.com/uploads/2021/05/26/grid4.png)
//
// ```
// Input: mat = [[0,0,0],[0,1,0],[1,1,1]], target = [[1,1,1],[0,1,0],[0,0,0]]
// Output: true
// Explanation: We can rotate mat 90 degrees clockwise two times to make mat
// equal target.
//
// ```
//
// **Constraints:**
//
// - `n == mat.length == target.length`
// - `n == mat[i].length == target[i].length`
// - `1 <= n <= 10`
// - `mat[i][j]` and `target[i][j]` are either `0` or `1`.
func findRotation(mat [][]int, target [][]int) bool {
	check := 4
	n := len(mat)
	for check > 0 {
		check--

		// rotate by fist swapping top and bottom rows
		for i := 0; i < n/2; i++ {
			mat[i], mat[n-1-i] = mat[n-1-i], mat[i]
		}

		// then flip diagonally
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
			}
		}

		// compare
		ok := true
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if mat[i][j] != target[i][j] {
					ok = false
				}
			}
		}

		if ok {
			return true
		}
	}

	return false
}
