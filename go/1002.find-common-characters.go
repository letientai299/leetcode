package main

/*
 * @lc app=leetcode id=1002 lang=golang
 *
 * [1002] Find Common Characters
 *
 * https://leetcode.com/problems/find-common-characters/description/
 *
 * algorithms
 * Easy (66.05%)
 * Total Accepted:    85.1K
 * Total Submissions: 125.4K
 * Testcase Example:  '["bella","label","roller"]'
 *
 * Given an array A of strings made only from lowercase letters, return a list
 * of all characters that show up in all strings within the list (including
 * duplicates).  For example, if a character occurs 3 times in all strings but
 * not 4 times, you need to include that character three times in the final
 * answer.
 *
 * You may return the answer in any order.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: ["bella","label","roller"]
 * Output: ["e","l","l"]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: ["cool","lock","cook"]
 * Output: ["c","o"]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 100
 * 1 <= A[i].length <= 100
 * A[i][j] is a lowercase letter
 *
 *
 *
 */
func commonChars(a []string) []string {
	if len(a) < 2 {
		return a
	}

	m := make(map[byte]int)
	for _, c := range []byte(a[0]) {
		m[c]++
	}

	for i := 1; i < len(a); i++ {
		next := make(map[byte]int)
		for _, c := range []byte(a[i]) {
			if m[c] > 0 {
				m[c]--
				next[c]++
			}
		}
		m = next
	}

	var res []string
	for k, v := range m {
		for i := 0; i < v; i++ {
			res = append(res, string([]byte{k}))
		}
	}
	return res
}
