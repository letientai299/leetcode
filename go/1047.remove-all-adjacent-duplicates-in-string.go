package main

/*
 * @lc app=leetcode id=1047 lang=golang
 *
 * [1047] Remove All Adjacent Duplicates In String
 *
 * https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/description/
 *
 * algorithms
 * Easy (65.45%)
 * Total Accepted:    103.9K
 * Total Submissions: 149.2K
 * Testcase Example:  '"abbaca"'
 *
 * Given a string S of lowercase letters, a duplicate removal consists of
 * choosing two adjacent and equal letters, and removing them.
 *
 * We repeatedly make duplicate removals on S until we no longer can.
 *
 * Return the final string after all such duplicate removals have been made.
 * It is guaranteed the answer is unique.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "abbaca"
 * Output: "ca"
 * Explanation:
 * For example, in "abbaca" we could remove "bb" since the letters are adjacent
 * and equal, and this is the only possible move.  The result of this move is
 * that the string is "aaca", of which only "aa" is possible, so the final
 * string is "ca".
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= S.length <= 20000
 * S consists only of English lowercase letters.
 *
 */
func removeDuplicates(s string) string {
	r := make([]byte, 0, len(s))
	for _, x := range s {
		c := byte(x)
		if 0 == len(r) {
			r = append(r, c)
			continue
		}

		if c == r[len(r)-1] {
			r = r[0 : len(r)-1]
			continue
		}
		r = append(r, c)
	}
	return string(r)
}
