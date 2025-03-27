package main

// Row With Maximum Ones
//
// Easy
//
// https://leetcode.com/problems/row-with-maximum-ones
//
// Given a `m x n` binary matrix `mat`, find the **0-indexed** position of the
// row that contains the **maximum** count of **ones,** and the number of ones
// in that row.
//
// In case there are multiple rows that have the maximum count of ones, the row
// with the **smallest row number** should be selected.
//
// Return _an array containing the index of the row, and the number of ones in
// it._
//
// **Example 1:**
//
// ```
// )
// ```
//
// **Example 2:**
//
// ```
// (2)1
// ```
//
// **Example 3:**
//
// ```
// Input: mat = [[0,0],[1,1],[0,0]]
// Output: [1,2]
// Explanation: The row indexed 1 has the maximum count of ones (2). So the
// answer is [1,2].
//
// ```
//
// **Constraints:**
//
// - `m == mat.length`
// - `n == mat[i].length`
// - `1 <= m, n <= 100`
// - `mat[i][j]` is either `0` or `1`.
func rowAndMaximumOnes(mat [][]int) []int {
  m := len(mat)
  n := len(mat[0])
  best := 0
  id := 0

  for r := 0; r < m; r++ {
    cur := 0
    for c := 0; c < n; c++ {
      if mat[r][c] == 1 {
        cur++
      }
    }

    if cur > best {
      id = r
      best = cur
    }
  }

  return []int{id, best}
}
