package main

/*
 * @lc app=leetcode id=1232 lang=golang
 *
 * [1232] Check If It Is a Straight Line
 *
 * https://leetcode.com/problems/check-if-it-is-a-straight-line/description/
 *
 * algorithms
 * Easy (47.27%)
 * Total Accepted:    85.3K
 * Total Submissions: 192.3K
 * Testcase Example:  '[[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]'
 *
 * You are given an array coordinates, coordinates[i] = [x, y], where [x, y]
 * represents the coordinate of a point. Check if these points make a straight
 * line in the XY plane.
 *
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= coordinates.length <= 1000
 * coordinates[i].length == 2
 * -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
 * coordinates contains no duplicate point.
 *
 *
 */
func checkStraightLine(coordinates [][]int) bool {
	n := len(coordinates)
	if n < 2 {
		return false
	}

	a := coordinates[0]
	b := coordinates[1]
	dx, dy := a[0]-b[0], a[1]-b[1]

	for i := 2; i < n; i++ {
		c := coordinates[i]
		gx, gy := a[0]-c[0], a[1]-c[1]

		if dx*gy != dy*gx {
			return false
		}
	}

	return true
}
