package main

import "sort"

/*
 * @lc app=leetcode id=821 lang=golang
 *
 * [821] Shortest Distance to a Character
 *
 * https://leetcode.com/problems/shortest-distance-to-a-character/description/
 *
 * algorithms
 * Easy (64.83%)
 * Total Accepted:    72.6K
 * Total Submissions: 107.6K
 * Testcase Example:  '"loveleetcode"\n"e"'
 *
 * Given a string S and a character C, return an array of integers representing
 * the shortest distance from the character C in the string.
 *
 * Example 1:
 *
 *
 * Input: S = "loveleetcode", C = 'e'
 * Output: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
 *
 *
 *
 *
 * Note:
 *
 *
 * S string length is in [1, 10000].
 * C is a single character, and guaranteed to be in string S.
 * All letters in S and C are lowercase.
 *
 *
 */
func shortestToChar(s string, c byte) []int {
	var locs []int
	for i := range s {
		if s[i] == c {
			locs = append(locs, i)
		}
	}

	res := make([]int, 0, len(s))
	for i := range s {
		x := sort.SearchInts(locs, i)

		if x == 0 {
			res = append(res, locs[x]-i)
			continue
		}

		if x == len(locs) {
			res = append(res, i-locs[x-1])
			continue
		}

		if locs[x] == i {
			res = append(res, 0)
			continue
		}

		if locs[x]-i < i-locs[x-1] {
			res = append(res, locs[x]-i)
		} else {
			res = append(res, i-locs[x-1])
		}
	}

	return res
}
