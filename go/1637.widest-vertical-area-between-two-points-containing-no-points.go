package main

import "sort"

// Widest Vertical Area Between Two Points Containing No Points
//
// Medium
//
// https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/
//
// Given `n``points` on a 2D plane where `points[i] = [xi, yi]`, Return _the
// **widest vertical area** between two points such that no points are inside
// the area._
//
// A **vertical area** is an area of fixed-width extending infinitely along the
// y-axis (i.e., infinite height). The **widest vertical area** is the one with
// the maximum width.
//
// Note that points **on the edge** of a vertical area **are not** considered
// included in the area.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2020/09/19/points3.png)​
//
// ```
// Input: points = [[8,7],[9,9],[7,4],[9,7]]
// Output: 1
// Explanation: Both the red and the blue area are optimal.
//
// ```
//
// **Example 2:**
//
// ```
// Input: points = [[3,1],[9,0],[1,0],[1,4],[5,3],[8,8]]
// Output: 3
//
// ```
//
// **Constraints:**
//
// - `n == points.length`
// - `2 <= n <= 105`
// - `points[i].length == 2`
// - `0 <= xi, yi <= 109`
func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	res := 0
	for i := 0; i < len(points)-1; i++ {
		a := points[i]
		b := points[i+1]
		if res < b[0]-a[0] {
			res = b[0] - a[0]
		}
	}
	return res
}
