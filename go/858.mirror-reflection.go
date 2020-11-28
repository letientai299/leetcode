package main

/*
 * @lc app=leetcode id=858 lang=golang
 *
 * [858] Mirror Reflection
 *
 * https://leetcode.com/problems/mirror-reflection/description/
 *
 * algorithms
 * Medium (52.41%)
 * Total Accepted:    22.7K
 * Total Submissions: 38.2K
 * Testcase Example:  '2\n1'
 *
 * There is a special square room with mirrors on each of the four walls.
 * Except for the southwest corner, there are receptors on each of the
 * remaining corners, numbered 0, 1, and 2.
 *
 * The square room has walls of length p, and a laser ray from the southwest
 * corner first meets the east wall at a distance q from the 0th receptor.
 *
 * Return the number of the receptor that the ray meets first.  (It is
 * guaranteed that the ray will meet a receptor eventually.)
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: p = 2, q = 1
 * Output: 2
 * Explanation: The ray meets receptor 2 the first time it gets reflected back
 * to the left wall.
 *
 *
 *
 * Note:
 *
 *
 * 1 <= p <= 1000
 * 0 <= q <= p
 *
 *
 *
 */
func mirrorReflection(p int, q int) int {
	left := false
	sig := 1
	x := q
	for x != p && x != 0 {
		x += sig * q
		left = !left
		if x > p {
			left = !left
			sig = -sig
			x = p - (x - p)
			continue
		}

		if x < 0 {
			left = !left
			x = -x
			sig = -sig
		}
	}

	if x == 0 {
		return 0
	}

	if left {
		return 2
	}

	return 1
}
