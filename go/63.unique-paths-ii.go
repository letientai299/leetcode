package main

// Unique Paths II
//
// Medium
//
// https://leetcode.com/problems/unique-paths-ii/
//
// A robot is located at the top-left corner of a `m x n` grid (marked 'Start'
// in the diagram below).
//
// The robot can only move either down or right at any point in time. The robot
// is trying to reach the bottom-right corner of the grid (marked 'Finish' in
// the diagram below).
//
// Now consider if some obstacles are added to the grids. How many unique paths
// would there be?
//
// An obstacle and space is marked as `1` and `0` respectively in the grid.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/11/04/robot1.jpg)
//
// ```
// Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
// Output: 2
// Explanation: There is one obstacle in the middle of the 3x3 grid above.
// There are two ways to reach the bottom-right corner:
// 1. Right -> Right -> Down -> Down
// 2. Down -> Down -> Right -> Right
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2020/11/04/robot2.jpg)
//
// ```
// Input: obstacleGrid = [[0,1],[0,0]]
// Output: 1
//
// ```
//
// **Constraints:**
//
// - `m == obstacleGrid.length`
// - `n == obstacleGrid[i].length`
// - `1 <= m, n <= 100`
// - `obstacleGrid[i][j]` is `0` or `1`.
func uniquePathsWithObstacles(grid [][]int) int {
	if grid[0][0] == 1 {
		return 0 // WTF
	}

	m := len(grid)
	n := len(grid[0])

	// top row
	grid[0][0] = 1
	for i := 1; i < n; i++ {
		if grid[0][i] == 1 {
			grid[0][i] = -1 // obstacle
			continue
		}

		if grid[0][i-1] > 0 {
			grid[0][i] = 1
		}
	}

	// left columns
	for i := 1; i < m; i++ {
		if grid[i][0] == 1 {
			grid[i][0] = -1 // obstacle
			continue
		}

		if grid[i-1][0] > 0 {
			grid[i][0] = 1
		}
	}

	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			if grid[y][x] == 1 {
				grid[y][x] = -1
				continue
			}

			up := grid[y-1][x]
			if up > 0 {
				grid[y][x] += up
			}

			left := grid[y][x-1]
			if left > 0 {
				grid[y][x] += left
			}
		}
	}

	v := grid[m-1][n-1]
	if v < 0 {
		return 0
	}

	return v
}
