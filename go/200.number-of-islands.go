package main

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 *
 * https://leetcode.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (43.78%)
 * Total Accepted:    905.5K
 * Total Submissions: 1.9M
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * Given an m x n 2d grid map of '1's (land) and '0's (water), return the
 * number of islands.
 *
 * An island is surrounded by water and is formed by connecting adjacent lands
 * horizontally or vertically. You may assume all four edges of the grid are
 * all surrounded by water.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [
 * ⁠ ["1","1","1","1","0"],
 * ⁠ ["1","1","0","1","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","1","0","0"],
 * ⁠ ["0","0","0","1","1"]
 * ]
 * Output: 3
 *
 *
 *
 * Constraints:
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 300
 * grid[i][j] is '0' or '1'.
 *
 *
 */

// TODO: 0ms, but we're slower than others
func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}

	type segment []int // start, end, id
	var prevSegs []segment
	var cnt int

	id := 1
	for r := 0; r < rows; r++ {
		row := grid[r]
		cur := []int{-1, -1, 0}
		pre := byte('0')

		var segs []segment

		for c := 0; c < cols; c++ {
			if row[c] == pre {
				continue
			}

			if pre == '0' {
				cur[0] = c // found new segment
			} else {
				cur[1] = c - 1 // close current segment
				segs = append(segs, cur)
				cur = []int{-1, -1, 0}
			}
			pre = row[c]
		}

		if cur[1] == -1 && cur[0] != -1 {
			cur[1] = cols - 1
			segs = append(segs, cur)
		}

		connected := make(map[int]int)
		c, p := 0, 0
		for c < len(segs) && p < len(prevSegs) {
			prev := prevSegs[p]
			cur = segs[c]

			if !(prev[1] < cur[0] || cur[1] < prev[0]) {
				if cur[2] == 0 || cur[2] > prev[2] {
					old := cur[2]
					cur[2] = prev[2] // same id
					if old != 0 {
						for j := c - 1; j >= 0; j-- {
							if segs[j][2] == old {
								segs[j][2] = cur[2]
							}
						}
					}
				} else {
					old := prev[2]
					prev[2] = cur[2]
					for j := p + 1; j < len(prevSegs); j++ {
						if prevSegs[j][2] == old {
							prevSegs[j][2] = prev[2]
						}
					}
				}
				connected[prev[2]] = 1
			}

			if prev[1] < cur[1] {
				p++
			} else {
				if cur[2] == 0 {
					cur[2] = id
					id++
				}
				c++
			}
		}

		for p := 0; p < len(prevSegs); p++ {
			if connected[prevSegs[p][2]] == 0 {
				cnt++
				connected[prevSegs[p][2]]++
			}
		}

		for c < len(segs) {
			if segs[c][2] == 0 {
				segs[c][2] = id
				id++
			}
			c++
		}

		prevSegs = segs
	}

	candidates := make(map[int]int)
	for _, cur := range prevSegs {
		if candidates[cur[2]] == 0 {
			candidates[cur[2]]++
			cnt++
		}
	}

	return cnt
}
