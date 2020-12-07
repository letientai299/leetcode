package main

import (
	"math"
)

/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 *
 * https://leetcode.com/problems/triangle/description/
 *
 * algorithms
 * Medium (41.45%)
 * Total Accepted:    273.2K
 * Total Submissions: 605.1K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * Given a triangle, find the minimum path sum from top to bottom. Each step
 * you may move to adjacent numbers on the row below.
 *
 * For example, given the following triangle
 *
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
 *
 * Note:
 *
 * Bonus point if you are able to do this using only O(n) extra space, where n
 * is the total number of rows in the triangle.
 *
 */

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	best := math.MaxInt32
	for i := 1; i < len(triangle); i++ {
		pre := triangle[i-1]
		cur := triangle[i]
		for j, v := range cur {
			var x int
			if j >= len(pre) {
				x = v + pre[j-1]
			} else {
				x = v + pre[j]
				if j > 0 && v+pre[j-1] < x {
					x = v + pre[j-1]
				}
			}

			cur[j] = x
			if i == len(triangle)-1 && best > x {
				best = x
			}
		}
	}

	return best
}
