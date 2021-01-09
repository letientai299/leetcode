package main

/*
 * @lc app=leetcode id=840 lang=golang
 *
 * [840] Magic Squares In Grid
 *
 * https://leetcode.com/problems/magic-squares-in-grid/description/
 *
 * algorithms
 * Easy (36.33%)
 * Total Accepted:    27.1K
 * Total Submissions: 71.8K
 * Testcase Example:  '[[4,3,8,4],[9,5,1,9],[2,7,6,2]]'
 *
 * A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to
 * 9 such that each row, column, and both diagonals all have the same sum.
 *
 * Given a row x col grid of integers, how many 3 x 3 "magic square" subgrids
 * are there?  (Each subgrid is contiguous).
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [[4,3,8,4],[9,5,1,9],[2,7,6,2]]
 * Output: 1
 * Explanation:
 * The following subgrid is a 3 x 3 magic square:
 *
 * while this one is not:
 *
 * In total, there is only one magic square inside the given grid.
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [[8]]
 * Output: 0
 *
 *
 * Example 3:
 *
 *
 * Input: grid = [[4,4],[3,3]]
 * Output: 0
 *
 *
 * Example 4:
 *
 *
 * Input: grid = [[4,7,8],[9,5,1],[2,3,6]]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * row == grid.length
 * col == grid[i].length
 * 1 <= row, col <= 10
 * 0 <= grid[i][j] <= 15
 *
 *
 */
func numMagicSquaresInside(grid [][]int) int {
	check := func(y, x int) bool {
		a, b, c := grid[y][x], grid[y][x+1], grid[y][x+2]
		y++
		d, e, f := grid[y][x], grid[y][x+1], grid[y][x+2]
		y++
		g, h, i := grid[y][x], grid[y][x+1], grid[y][x+2]
		if a > 9 || b > 9 || c > 9 || d > 9 || e > 9 || f > 9 || g > 9 || h > 9 || i > 9 ||
			a == b || b == c || c == d || d == e || e == f || f == g || g == h || h == i || i == a {
			return false
		}

		return a+b+c+d+e+f+g+h+i == 45 &&
			a+b+c == d+e+f &&
			d+e+f == g+h+i &&
			a+d+g == b+e+h &&
			b+e+h == c+f+i &&
			a+i == c+g
	}

	c := 0
	for y := 0; y < len(grid)-2; y++ {
		for x := 0; x < len(grid[y])-2; x++ {
			if check(y, x) {
				c++
			}
		}
	}
	return c
}
