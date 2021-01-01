package main

import "sort"

/*
 * @lc app=leetcode id=436 lang=golang
 *
 * [436] Find Right Interval
 *
 * https://leetcode.com/problems/find-right-interval/description/
 *
 * algorithms
 * Medium (43.87%)
 * Total Accepted:    60K
 * Total Submissions: 124.1K
 * Testcase Example:  '[[1,2]]'
 *
 * You are given an array of intervals, where intervals[i] = [starti, endi] and
 * each starti is unique.
 *
 * The right interval for an interval i is an interval j such that startj >=
 * endi and startj is minimized.
 *
 * Return an array of right interval indices for each interval i. If no right
 * interval exists for interval i, then put -1 at index i.
 *
 *
 * Example 1:
 *
 *
 * Input: intervals = [[1,2]]
 * Output: [-1]
 * Explanation: There is only one interval in the collection, so it outputs
 * -1.
 *
 *
 * Example 2:
 *
 *
 * Input: intervals = [[3,4],[2,3],[1,2]]
 * Output: [-1,0,1]
 * Explanation: There is no right interval for [3,4].
 * The right interval for [2,3] is [3,4] since start0 = 3 is the smallest start
 * that is >= end1 = 3.
 * The right interval for [1,2] is [2,3] since start1 = 2 is the smallest start
 * that is >= end2 = 2.
 *
 *
 * Example 3:
 *
 *
 * Input: intervals = [[1,4],[2,3],[3,4]]
 * Output: [-1,2,-1]
 * Explanation: There is no right interval for [1,4] and [3,4].
 * The right interval for [2,3] is [3,4] since start2 = 3 is the smallest start
 * that is >= end1 = 3.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= intervals.length <= 2 * 10^4
 * intervals[i].length == 2
 * -10^6 <= starti <= endi <= 10^6
 * The start point of each interval is unique.
 *
 *
 */
func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	if n == 0 {
		return nil
	}

	starts := make([][2]int, n)
	for i, a := range intervals {
		starts[i] = [2]int{a[0], i}
	}

	sort.Slice(starts, func(i, j int) bool {
		a, b := starts[i], starts[j]
		return a[0]<b[0] || (a[0] == b[0] && a[1]<b[1])
	})

	res := make([]int, n)
	for i, a := range intervals {
		x := sort.Search(n, func(j int) bool {
			return a[1] <= starts[j][0]
		})
		if x == -1 || x == n {
			res[i] = -1
		} else {
			res[i] = starts[x][1]
		}
	}
	return res
}
