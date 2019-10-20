package main

/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 *
 * https://leetcode.com/problems/is-subsequence/description/
 *
 * algorithms
 * Easy (47.61%)
 * Total Accepted:    112.1K
 * Total Submissions: 235.4K
 * Testcase Example:  '"abc"\n"ahbgdc"'
 *
 *
 * Given a string s and a string t, check if s is subsequence of t.
 *
 *
 *
 * You may assume that there is only lower case English letters in both s and
 * t. t is potentially a very long (length ~= 500,000) string, and s is a short
 * string (
 *
 *
 * A subsequence of a string is a new string which is formed from the original
 * string by deleting some (can be non) of the characters without disturbing
 * the relative positions of the remaining characters. (ie, "ace" is a
 * subsequence of "abcde" while "aec" is not).
 *
 *
 * Example 1:
 * s = "abc", t = "ahbgdc"
 *
 *
 * Return true.
 *
 *
 * Example 2:
 * s = "axc", t = "ahbgdc"
 *
 *
 * Return false.
 *
 *
 * Follow up:
 * If there are lots of incoming S, say S1, S2, ... , Sk where k >= 1B, and you
 * want to check one by one to see if T has its subsequence. In this scenario,
 * how would you change your code?
 *
 * Credits:Special thanks to @pbrother for adding this problem and creating all
 * test cases.
 */

/*this solution works fro the follow up questions but slow for the tests on leetcode */
/*
func isSubsequence(s string, t string) bool {
	ids := make(map[int32][]int, 0)

	for i, x := range t {
		ids[x] = append(ids[x], i)
	}

	cur := -1
	for _, x := range s {
		locations, ok := ids[x]
		if !ok {
			return false // can't found character x of s in t
		}

		// find the location smallest possible location of x that greater than cur
		j := sort.SearchInts(locations, cur)
		if j >= len(locations) {
			return false
		}

		if locations[j] == cur {
			j++ // need to skip this location
		}

		if j >= len(locations) {
			return false
		}

		cur = locations[j]
	}

	return true
}
*/

func isSubsequence(s string, t string) bool {
	ns := len(s)
	nt := len(t)
	is := 0
	it := 0

	for is < ns && it < nt {
		if s[is] != t[it] {
			it++
			continue
		}

		is++
		it++
	}

	return is == ns
}
