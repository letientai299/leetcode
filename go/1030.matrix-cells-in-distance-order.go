package main

import "sort"

/*
 * @lc app=leetcode id=1030 lang=golang
 *
 * [1030] Matrix Cells in Distance Order
 *
 * https://leetcode.com/problems/matrix-cells-in-distance-order/description/
 *
 * algorithms
 * Easy (64.39%)
 * Total Accepted:    27.2K
 * Total Submissions: 40.7K
 * Testcase Example:  '1\n2\n0\n0'
 *
 * We are given a matrix with R rows and C columns has cells with integer
 * coordinates (r, c), where 0 <= r < R and 0 <= c < C.
 *
 * Additionally, we are given a cell in that matrix with coordinates (r0, c0).
 *
 * Return the coordinates of all cells in the matrix, sorted by their distance
 * from (r0, c0) from smallest distance to largest distance.  Here, the
 * distance between two cells (r1, c1) and (r2, c2) is the Manhattan distance,
 * |r1 - r2| + |c1 - c2|.  (You may return the answer in any order that
 * satisfies this condition.)
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: R = 1, C = 2, r0 = 0, c0 = 0
 * Output: [[0,0],[0,1]]
 * Explanation: The distances from (r0, c0) to other cells are: [0,1]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: R = 2, C = 2, r0 = 0, c0 = 1
 * Output: [[0,1],[0,0],[1,1],[1,0]]
 * Explanation: The distances from (r0, c0) to other cells are: [0,1,1,2]
 * The answer [[0,1],[1,1],[0,0],[1,0]] would also be accepted as correct.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: R = 2, C = 3, r0 = 1, c0 = 2
 * Output: [[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
 * Explanation: The distances from (r0, c0) to other cells are: [0,1,1,2,2,3]
 * There are other answers that would also be accepted as correct, such as
 * [[1,2],[1,1],[0,2],[1,0],[0,1],[0,0]].
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= R <= 100
 * 1 <= C <= 100
 * 0 <= r0 < R
 * 0 <= c0 < C
 *
 *
 *
 *
 *
 */

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, 0, R*C)
	for x := 0; x < R; x++ {
		for y := 0; y < C; y++ {
			res = append(res, []int{x, y})
		}
	}

	abs := func(x int) int {
		if x > 0 {
			return x
		}
		return -x
	}

	dist := func(a, b []int) int {
		return abs(a[0]-b[0]) + abs(a[1]-b[1])
	}

	p := []int{r0, c0}

	sort.Slice(res, func(i, j int) bool {
		a := res[i]
		b := res[j]
		da, db := dist(a, p), dist(b, p)
		return da < db
	})

	return res
}
