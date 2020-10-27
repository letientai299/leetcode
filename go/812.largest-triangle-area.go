package main

/*
 * @lc app=leetcode id=812 lang=golang
 *
 * [812] Largest Triangle Area
 *
 * https://leetcode.com/problems/largest-triangle-area/description/
 *
 * algorithms
 * Easy (57.07%)
 * Total Accepted:    23.4K
 * Total Submissions: 40K
 * Testcase Example:  '[[0,0],[0,1],[1,0],[0,2],[2,0]]'
 *
 * You have a list of points in the plane. Return the area of the largest
 * triangle that can be formed by any 3 of the points.
 *
 *
 * Example:
 * Input: points = [[0,0],[0,1],[1,0],[0,2],[2,0]]
 * Output: 2
 * Explanation:
 * The five points are show in the figure below. The red triangle is the
 * largest.
 *
 *
 *
 *
 * Notes:
 *
 *
 * 3 <= points.length <= 50.
 * No points will be duplicated.
 * -50 <= points[i][j] <= 50.
 * Answers within 10^-6 of the true value will be accepted as correct.
 *
 *
 *
 *
 */
func largestTriangleArea(points [][]int) float64 {
	area := func(a, b, c []int) float64 {
		x := float64(
			a[0]*(b[1]-c[1])+
				b[0]*(c[1]-a[1])+
				c[0]*(a[1]-b[1]),
		) / 2.0
		if x < 0 {
			return -x
		}
		return x
	}

	maxArea := float64(0)
	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				p2 := points[j]
				p3 := points[k]
				a := area(p1, p2, p3)
				if a > maxArea {
					maxArea = a
				}
			}
		}
	}
	return maxArea
}
