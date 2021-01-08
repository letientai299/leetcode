package main

/*
 * @lc app=leetcode id=994 lang=golang
 *
 * [994] Rotting Oranges
 *
 * https://leetcode.com/problems/rotting-oranges/description/
 *
 * algorithms
 * Easy (46.34%)
 * Total Accepted:    185K
 * Total Submissions: 373.6K
 * Testcase Example:  '[[2,1,1],[1,1,0],[0,1,1]]'
 *
 * You are given an m x n grid where each cell can have one of three
 * values:
 *
 *
 * 0 representing an empty cell,
 * 1 representing a fresh orange, or
 * 2 representing a rotten orange.
 *
 *
 * Every minute, any fresh orange that is 4-directionally adjacent to a rotten
 * orange becomes rotten.
 *
 * Return the minimum number of minutes that must elapse until no cell has a
 * fresh orange. If this is impossible, return -1.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
 * Output: -1
 * Explanation: The orange in the bottom left corner (row 2, column 0) is never
 * rotten, because rotting only happens 4-directionally.
 *
 *
 * Example 3:
 *
 *
 * Input: grid = [[0,2]]
 * Output: 0
 * Explanation: Since there are already no fresh oranges at minute 0, the
 * answer is just 0.
 *
 *
 *
 * Constraints:
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 10
 * grid[i][j] is 0, 1, or 2.
 *
 *
 */

func orangesRotting(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	var dfs func(y, x int) int
	dfs = func(y, x int) int {
		if y < 0 || y >= m || x < 0 || x >= n || grid[y][x] == 0 {
			return -1
		}

		if grid[y][x] == 2 {
			return 0
		}

		grid[y][x] = 0
		v := dfs(y, x-1)

		if u := dfs(y-1, x); u != -1 && (u < v || v == -1) {
			v = u
		}

		if u := dfs(y+1, x); u != -1 && (u < v || v == -1) {
			v = u
		}

		if u := dfs(y, x+1); u != -1 && (u < v || v == -1) {
			v = u
		}

		grid[y][x] = 1

		if v == -1 {
			return -1
		}

		return 1 + v
	}

	best := 0
	for y, row := range grid {
		for x, v := range row {
			if v != 1 {
				continue
			}

			d := dfs(y, x)
			if d == -1 {
				return d
			}

			if d > best {
				best = d
			}
		}
	}

	return best
}
