package main

/*
 * @lc app=leetcode id=316 lang=golang
 *
 * [316] Remove Duplicate Letters
 *
 * https://leetcode.com/problems/remove-duplicate-letters/description/
 *
 * algorithms
 * medium (33.79%)
 * Total Accepted:    107.2K
 * Total Submissions: 277.6K
 * Testcase Example:  '"bcabc"'
 *
 * Given a string s, remove duplicate letters so that every letter appears once
 * and only once. You must make sure your result is the smallest in
 * lexicographical order among all possible results.
 *
 * Note: This question is the same as 1081:
 * https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/
 *
 *
 * Example 1:
 *
 *
 * Input: s = "bcabc"
 * Output: "abc"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "cbacdcbc"
 * Output: "acdb"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^4
 * s consists of lowercase English letters.
 *
 *
 */
// TODO (tai): not done yet
func removeDuplicateLetters(s string) string {
	max := [27]byte{}
	for i, c := range s {
		v := max[byte(c-'a')]
		if v < byte(i) {
			max[byte(c-'a')] = byte(i)
		}
	}

	m := [27]byte{}
	var res []byte
	for k, x := range s {
		c := byte(x)
		i := m[c-'a']
		l := byte(len(res))
		if i >= l || res[i] != c {
			m[c-'a'] = l
			res = append(res, c)
			continue
		}

		change := false
		larger := -1
		for j := i + 1; j < l; j++ {
			if res[j] < c {
				change = true
			} else if larger == -1 {
				larger = int(j)
			}
		}

		if !change {
			continue
		}

		if larger != -1 && max[int(res[larger]-'a')] < byte(k) {
			continue
		}

		for j := i; j < l-1; j++ {
			res[j] = res[j+1]
			m[res[j]-'a']--
		}

		res[l-1] = c
		m[c-'a'] = l - 1
	}
	return string(res)
}
