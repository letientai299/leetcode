package main

import "math"

// Shortest Bridge
//
// Medium
//
// https://leetcode.com/problems/shortest-bridge/
//
// In a given 2D binary array `grid`, there are two islands.  (An island is a
// 4-directionally connected group of `1`s not connected to any other 1s.)
//
// Now, we may change `0`s to `1`s so as to connect the two islands together to
// form 1 island.
//
// Return the smallest number of `0`s that must be flipped.  (It is guaranteed
// that the answer is at least 1.)
//
// **Example 1:**
//
// ```
// Input: grid = [[0,1],[1,0]]
// Output: 1
//
// ```
//
// **Example 2:**
//
// ```
// Input: grid = [[0,1,0],[0,0,0],[0,0,1]]
// Output: 2
//
// ```
//
// **Example 3:**
//
// ```
// Input: grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
// Output: 1
//
// ```
//
// **Constraints:**
//
// - `2 <= grid.length == grid[0].length <= 100`
// - `grid[i][j] == 0` or `grid[i][j] == 1`
func shortestBridge(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var islands [][]int // 2 slice of locations of 2 islands

	var collect func(x, y int, r *[]int)
	collect = func(x, y int, r *[]int) {
		grid[y][x] = 0
		*r = append(*r, x, y)
		if 0 <= x-1 && grid[y][x-1] == 1 {
			collect(x-1, y, r)
		}

		if n > x+1 && grid[y][x+1] == 1 {
			collect(x+1, y, r)
		}

		if 0 <= y-1 && grid[y-1][x] == 1 {
			collect(x, y-1, r)
		}

		if m > y+1 && grid[y+1][x] == 1 {
			collect(x, y+1, r)
		}
	}

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if grid[y][x] == 1 {
				var xys []int
				collect(x, y, &xys)
				islands = append(islands, xys)
			}
		}
	}

	a, b := islands[0], islands[1]
	best := math.MaxInt32
	for i := 0; i < len(a) && best > 1; i += 2 {
		x1, y1 := a[i], a[i+1]
		for j := 0; j < len(b) && best > 1; j += 2 {
			x2, y2 := b[j], b[j+1]
			d := abs(x2-x1) + abs(y2-y1) - 1
			if d < best {
				best = d
			}
		}
	}
	return best
}
