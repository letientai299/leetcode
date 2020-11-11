package main

/*
 * @lc app=leetcode id=1184 lang=golang
 *
 * [1184] Distance Between Bus Stops
 *
 * https://leetcode.com/problems/distance-between-bus-stops/description/
 *
 * algorithms
 * Easy (55.70%)
 * Total Accepted:    26.8K
 * Total Submissions: 49.4K
 * Testcase Example:  '[1,2,3,4]\n0\n1'
 *
 * A bus has n stops numbered from 0 to n - 1 that form a circle. We know the
 * distance between all pairs of neighboring stops where distance[i] is the
 * distance between the stops number i and (i + 1) % n.
 *
 * The bus goes along both directions i.e. clockwise and counterclockwise.
 *
 * Return the shortest distance between the given start and destination
 * stops.
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: distance = [1,2,3,4], start = 0, destination = 1
 * Output: 1
 * Explanation: Distance between 0 and 1 is 1 or 9, minimum is 1.
 *
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: distance = [1,2,3,4], start = 0, destination = 2
 * Output: 3
 * Explanation: Distance between 0 and 2 is 3 or 7, minimum is 3.
 *
 *
 *
 *
 * Example 3:
 *
 *
 *
 *
 * Input: distance = [1,2,3,4], start = 0, destination = 3
 * Output: 4
 * Explanation: Distance between 0 and 3 is 6 or 4, minimum is 4.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 10^4
 * distance.length == n
 * 0 <= start, destination < n
 * 0 <= distance[i] <= 10^4
 *
 */
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	n := len(distance)
	a, b := start, destination
	if a > b {
		a, b = b, a
	}

	r, l := a, a-1

	forward := 0
	backward := 0
	for r != b || l != b {
		if r != b {
			forward += distance[r]
			r++
		}

		if l != b {
			if l < 0 {
				l = n - 1
			}
			backward += distance[l]
			if l != b {
				l--
				if l == b {
					backward += distance[l]
				}
			}
		}
	}

	if forward < backward {
		return forward
	}

	return backward
}
