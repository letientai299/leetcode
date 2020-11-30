package main

/*
 * @lc app=leetcode id=1109 lang=golang
 *
 * [1109] Corporate Flight Bookings
 *
 * https://leetcode.com/problems/corporate-flight-bookings/description/
 *
 * algorithms
 * Medium (49.71%)
 * Total Accepted:    21.6K
 * Total Submissions: 40.1K
 * Testcase Example:  '[[1,2,10],[2,3,20],[2,5,25]]\n5'
 *
 * There are n flights, and they are labeled from 1 to n.
 *
 * We have a list of flight bookings.  The i-th booking bookings[i] = [i, j, k]
 * means that we booked k seats from flights labeled i to j inclusive.
 *
 * Return an array answer of length n, representing the number of seats booked
 * on each flight in order of their label.
 *
 *
 * Example 1:
 *
 *
 * Input: bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
 * Output: [10,55,45,25,25]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= bookings.length <= 20000
 * 1 <= bookings[i][0] <= bookings[i][1] <= n <= 20000
 * 1 <= bookings[i][2] <= 10000
 *
 */
func corpFlightBookings(bookings [][]int, n int) []int {
	r := make([]int, n+1)
	for _, b := range bookings {
		r[b[0]-1] += b[2]
		r[b[1]] -= b[2]
	}

	for i := 1; i < len(r); i++ {
		r[i] += r[i-1]
	}
	return r
}
