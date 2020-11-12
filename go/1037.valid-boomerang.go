package main

import "math"

/*
 * @lc app=leetcode id=1037 lang=golang
 *
 * [1037] Valid Boomerang
 *
 * https://leetcode.com/problems/valid-boomerang/description/
 *
 * algorithms
 * Easy (37.41%)
 * Total Accepted:    20.1K
 * Total Submissions: 53.4K
 * Testcase Example:  '[[1,1],[2,3],[3,2]]'
 *
 * A boomerang is a set of 3 points that are all distinct and not in a straight
 * line.
 *
 * Given a list of three points in the plane, return whether these points are a
 * boomerang.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[1,1],[2,3],[3,2]]
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,1],[2,2],[3,3]]
 * Output: false
 *
 *
 *
 *
 * Note:
 *
 *
 * points.length == 3
 * points[i].length == 2
 * 0 <= points[i][j] <= 100
 *
 *
 *
 *
 *
 */
func isBoomerang(points [][]int) bool {
	pa := points[0]
	pb := points[1]
	pc := points[2]

	C := math.Sqrt(float64((pa[0]-pb[0])*(pa[0]-pb[0]) + (pa[1]-pb[1])*(pa[1]-pb[1])))
	A := math.Sqrt(float64((pc[0]-pb[0])*(pc[0]-pb[0]) + (pc[1]-pb[1])*(pc[1]-pb[1])))
	B := math.Sqrt(float64((pc[0]-pa[0])*(pc[0]-pa[0]) + (pc[1]-pa[1])*(pc[1]-pa[1])))

	return (A+B) > C && (B+C) > A && (C+A) > B
}
