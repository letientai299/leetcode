package main

// Count Sub Islands
//
// Medium
//
// https://leetcode.com/problems/count-sub-islands/
//
// You are given two `m x n` binary matrices `grid1` and `grid2` containing only
// `0`'s (representing water) and `1`'s (representing land). An **island** is a
// group of `1`'s connected **4-directionally** (horizontal or vertical). Any
// cells outside of the grid are considered water cells.
//
// An island in `grid2` is considered a **sub-island** if there is an island in
// `grid1` that contains **all** the cells that make up **this** island in
// `grid2`.
//
// Return the _**number** of islands in_ `grid2` _that are considered
// **sub-islands**_.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/06/10/test1.png)
//
// ```
// Input: grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]],
// grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
// Output: 3
// Explanation: In the picture above, the grid on the left is grid1 and the grid
// on the right is grid2.
// The 1s colored red in grid2 are those considered to be part of a sub-island.
// There are three sub-islands.
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/06/03/testcasex2.png)
//
// ```
// Input: grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]],
// grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
// Output: 2
// Explanation: In the picture above, the grid on the left is grid1 and the grid
// on the right is grid2.
// The 1s colored red in grid2 are those considered to be part of a sub-island.
// There are two sub-islands.
//
// ```
//
// **Constraints:**
//
// - `m == grid1.length == grid2.length`
// - `n == grid1[i].length == grid2[i].length`
// - `1 <= m, n <= 500`
// - `grid1[i][j]` and `grid2[i][j]` are either `0` or `1`.
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m := len(grid1)
	n := len(grid1[0])
	res := 0

	var cover func(y, x int) bool
	cover = func(y, x int) bool {
		if x < 0 || y < 0 || x >= n || y >= m || grid2[y][x] == 0 {
			return true
		}

		ok := grid1[y][x] == 1
		grid2[y][x] = 0
		a := cover(y, x+1)
		b := cover(y, x-1)
		c := cover(y+1, x)
		d := cover(y-1, x)
		return a && b && c && d && ok
	}

	for y, row := range grid2 {
		for x, v := range row {
			if v == 1 && cover(y, x) {
				res++
			}
		}
	}

	return res
}
