package main

/*
 * @lc app=leetcode id=986 lang=golang
 *
 * [986] Interval List Intersections
 *
 * https://leetcode.com/problems/interval-list-intersections/description/
 *
 * algorithms
 * Medium (65.52%)
 * Total Accepted:    146.1K
 * Total Submissions: 215.7K
 * Testcase Example:  '[[0,2],[5,10],[13,23],[24,25]]\n[[1,5],[8,12],[15,24],[25,26]]'
 *
 * Given two lists of closed intervals, each list of intervals is pairwise
 * disjoint and in sorted order.
 *
 * Return the intersection of these two interval lists.
 *
 * (Formally, a closed interval [a, b] (with a <= b) denotes the set of real
 * numbers x with a <= x <= b.  The intersection of two closed intervals is a
 * set of real numbers that is either empty, or can be represented as a closed
 * interval.  For example, the intersection of [1, 3] and [2, 4] is [2,
 * 3].)
 *
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: A = [[0,2],[5,10],[13,23],[24,25]], B =
 * [[1,5],[8,12],[15,24],[25,26]]
 * Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= A.length < 1000
 * 0 <= B.length < 1000
 * 0 <= A[i].start, A[i].end, B[i].start, B[i].end < 10^9
 *
 *
 *
 */
func intervalIntersection(a [][]int, b [][]int) [][]int {
	var res [][]int

	for x, y := 0, 0; x < len(a) && y < len(b); {
		ra := a[x]
		rb := b[y]

		// ra is completely before rb
		if ra[1] < rb[0] {
			x++
			continue
		}

		// rb is completely before ra
		if rb[1] < ra[0] {
			y++
			continue
		}

		var interval []int
		if ra[0] < rb[0] {
			interval = append(interval, rb[0])
		} else {
			interval = append(interval, ra[0])
		}

		if ra[1] < rb[1] {
			interval = append(interval, ra[1])
			x++
		} else {
			interval = append(interval, rb[1])
			y++
		}

		res = append(res, interval)
	}

	return res
}
