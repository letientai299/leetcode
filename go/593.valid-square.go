package main

/*
 * @lc app=leetcode id=593 lang=golang
 *
 * [593] Valid Square
 *
 * https://leetcode.com/problems/valid-square/description/
 *
 * algorithms
 * Medium (41.26%)
 * Total Accepted:    55.7K
 * Total Submissions: 128.5K
 * Testcase Example:  '[0,0]\n[1,1]\n[1,0]\n[0,1]'
 *
 * Given the coordinates of four points in 2D space p1, p2, p3 and p4, return
 * true if the four points construct a square.
 *
 * The coordinate of a point pi is represented as [xi, yi]. The input is not
 * given in any order.
 *
 * A valid square has four equal sides with positive length and four equal
 * angles (90-degree angles).
 *
 *
 * Example 1:
 *
 *
 * Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,12]
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input: p1 = [1,0], p2 = [-1,0], p3 = [0,1], p4 = [0,-1]
 * Output: true
 *
 *
 *
 * Constraints:
 *
 *
 * p1.length == p2.length == p3.length == p4.length == 2
 * -10^4 <= xi, yi <= 10^4
 *
 *
 */
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	// normalize other 3 points to a coordinate where p1 is (0,0)
	p2[0], p2[1] = p2[0]-p1[0], p2[1]-p1[1]
	p3[0], p3[1] = p3[0]-p1[0], p3[1]-p1[1]
	p4[0], p4[1] = p4[0]-p1[0], p4[1]-p1[1]

	// find square distance of other 3 points
	d2 := p2[0]*p2[0] + p2[1]*p2[1]
	d3 := p3[0]*p3[0] + p3[1]*p3[1]
	d4 := p4[0]*p4[0] + p4[1]*p4[1]

	sign := func(b, p []int) int { return (b[0])*(p[1]) - (b[1])*(p[0]) }
	orthogonal := func(a, b []int) bool { return a[0]*b[0]+a[1]*b[1] == 0 }

	// if they can be a square:
	// - 2 of above 3 distances must be equal,
	// - those 2 vectors make a right angle
	// - the last is equal to sum of other 2
	// - the last point must be between 2 vectors created by other 2 points
	return (orthogonal(p2, p3) && d2 == d3 && d4 == d2+d3 && sign(p2, p4) != sign(p3, p4)) ||
		(orthogonal(p3, p4) && d3 == d4 && d2 == d3+d4 && sign(p3, p2) != sign(p4, p2)) ||
		(orthogonal(p4, p2) && d2 == d4 && d3 == d2+d4 && sign(p4, p3) != sign(p2, p3))
}
