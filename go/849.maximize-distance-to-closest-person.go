package main

/*
 * @lc app=leetcode id=849 lang=golang
 *
 * [849] Maximize Distance to Closest Person
 *
 * https://leetcode.com/problems/maximize-distance-to-closest-person/description/
 *
 * algorithms
 * Easy (41.82%)
 * Total Accepted:    96.2K
 * Total Submissions: 216.6K
 * Testcase Example:  '[1,0,0,0,1,0,1]'
 *
 * You are given an array representing a row of seats where seats[i] = 1
 * represents a person sitting in the i^th seat, and seats[i] = 0 represents
 * that the i^th seat is empty (0-indexed).
 *
 * There is at least one empty seat, and at least one person sitting.
 *
 * Alex wants to sit in the seat such that the distance between him and the
 * closest person to him is maximized. 
 *
 * Return that maximum distance to the closest person.
 *
 *
 * Example 1:
 *
 *
 * Input: seats = [1,0,0,0,1,0,1]
 * Output: 2
 * Explanation:
 * If Alex sits in the second open seat (i.e. seats[2]), then the closest
 * person has distance 2.
 * If Alex sits in any other open seat, the closest person has distance 1.
 * Thus, the maximum distance to the closest person is 2.
 *
 *
 * Example 2:
 *
 *
 * Input: seats = [1,0,0,0]
 * Output: 3
 * Explanation:
 * If Alex sits in the last seat (i.e. seats[3]), the closest person is 3 seats
 * away.
 * This is the maximum distance possible, so the answer is 3.
 *
 *
 * Example 3:
 *
 *
 * Input: seats = [0,1]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= seats.length <= 2 * 10^4
 * seats[i] is 0 or 1.
 * At least one seat is empty.
 * At least one seat is occupied.
 *
 *
 */
func maxDistToClosest(seats []int) int {
	pre := -1
	i := 0
	best := 0
	for ; i < len(seats) && seats[i] == 0; i++ {
		best++
	}
	pre = i
	for i++; i < len(seats); i++ {
		v := seats[i]
		if v != 0 {
			v = (i - pre) / 2
			if best < v {
				best = v
			}
			pre = i
		}
	}
	if pre != len(seats)-1 && best < len(seats)-pre {
		best = len(seats) - pre - 1
	}
	return best
}
