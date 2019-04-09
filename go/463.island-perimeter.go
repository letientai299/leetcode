package main

/*
 * @lc app=leetcode id=463 lang=golang
 *
 * [463] Island Perimeter
 *
 * https://leetcode.com/problems/island-perimeter/description/
 *
 * algorithms
 * Easy (60.58%)
 * Total Accepted:    127.6K
 * Total Submissions: 210.6K
 * Testcase Example:  '[[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]'
 *
 * You are given a map in form of a two-dimensional integer grid where 1
 * represents land and 0 represents water.
 *
 * Grid cells are connected horizontally/vertically (not diagonally). The grid
 * is completely surrounded by water, and there is exactly one island (i.e.,
 * one or more connected land cells).
 *
 * The island doesn't have "lakes" (water inside that isn't connected to the
 * water around the island). One cell is a square with side length 1. The grid
 * is rectangular, width and height don't exceed 100. Determine the perimeter
 * of the island.
 *
 *
 *
 * Example:
 *
 *
 * Input:
 * [[0,1,0,0],
 * ⁠ [1,1,1,0],
 * ⁠ [0,1,0,0],
 * ⁠ [1,1,0,0]]
 *
 * Output: 16
 *
 * Explanation: The perimeter is the 16 yellow stripes in the image below:
 *
 *
 *
 *
 */
func islandPerimeter(grid [][]int) int {
	c := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == 0 {
				continue
			}

			faces := 4

			if j-1 >= 0 && row[j-1] == 1 {
				faces--
			}

			if j+1 < len(row) && row[j+1] == 1 {
				faces--
			}

			if i-1 >= 0 && grid[i-1][j] == 1 {
				faces--
			}

			if i+1 < len(grid) && grid[i+1][j] == 1 {
				faces--
			}

			c += faces
		}
	}
	return c
}
