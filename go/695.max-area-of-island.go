package main

/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 *
 * https://leetcode.com/problems/max-area-of-island/description/
 *
 * algorithms
 * Medium (59.59%)
 * Total Accepted:    196.3K
 * Total Submissions: 305.9K
 * Testcase Example:  '[[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]'
 *
 * Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's
 * (representing land) connected 4-directionally (horizontal or vertical.) You
 * may assume all four edges of the grid are surrounded by water.
 *
 * Find the maximum area of an island in the given 2D array. (If there is no
 * island, the maximum area is 0.)
 *
 * Example 1:
 *
 *
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,1,1,0,1,0,0,0,0,0,0,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,0,1,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,1,1,0,0],
 * ⁠[0,0,0,0,0,0,0,0,0,0,1,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,0,0,0,0]]
 *
 * Given the above grid, return 6. Note the answer is not 11, because the
 * island must be connected 4-directionally.
 *
 * Example 2:
 *
 *
 * [[0,0,0,0,0,0,0,0]]
 * Given the above grid, return 0.
 *
 * Note: The length of each dimension in the given grid does not exceed 50.
 *
 */

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	best := 0
	var dfs func(x, y int, cur *int)
	dfs = func(x, y int, cur *int) {
		if x < 0 || x >= n || y < 0 || y >= m || grid[y][x] == 0 {
			return
		}
		*cur++
		grid[y][x] = 0
		dfs(x-1, y, cur)
		dfs(x+1, y, cur)
		dfs(x, y-1, cur)
		dfs(x, y+1, cur)
	}

	for y, row := range grid {
		for x, v := range row {
			if v == 0 {
				continue
			}

			cur := 0
			dfs(x, y, &cur)
			if best < cur {
				best = cur
			}
		}
	}

	return best
}
