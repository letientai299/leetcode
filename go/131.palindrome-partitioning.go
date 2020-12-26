package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 *
 * https://leetcode.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (43.63%)
 * Total Accepted:    279.9K
 * Total Submissions: 549.3K
 * Testcase Example:  '"aab"'
 *
 * Given a string s, partition s such that every substring of the partition is
 * a palindrome. Return all possible palindrome partitioning of s.
 *
 * A palindrome string is a string that reads the same backward as forward.
 *
 *
 * Example 1:
 * Input: s = "aab"
 * Output: [["a","a","b"],["aa","b"]]
 * Example 2:
 * Input: s = "a"
 * Output: [["a"]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 16
 * s contains only lowercase English letters.
 *
 *
 */
func partition(s string) [][]string {
	n := len(s)
	if n == 0 {
		return nil
	}

	m := make(map[int32][]int)
	for i, x := range s {
		m[x] = append(m[x], i)
	}

	isPalindromeStr := func(start, end int) bool {
		for start < end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}

	var res [][]string
	var cur []string

	var part func(start int)
	part = func(start int) {
		if start >= n {
			ss := make([]string, len(cur))
			copy(ss, cur)
			res = append(res, ss)
			return
		}

		l := len(cur)
		cur = append(cur, s[start:start+1])
		part(start + 1)
		cur = cur[:l]

		pos := m[int32(s[start])]
		i := sort.SearchInts(pos, start)
		for i++; i < len(pos); i++ {
			if isPalindromeStr(start+1, pos[i]-1) {
				pre := s[start : pos[i]+1]
				cur = append(cur, pre)
				part(pos[i] + 1)
				cur = cur[:l]
			}
		}
	}

	part(0)
	return res
}
