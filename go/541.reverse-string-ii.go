package main

/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 *
 * https://leetcode.com/problems/reverse-string-ii/description/
 *
 * algorithms
 * Easy (45.53%)
 * Total Accepted:    59.8K
 * Total Submissions: 131.5K
 * Testcase Example:  '"abcdefg"\n2'
 *
 *
 * Given a string and an integer k, you need to reverse the first k characters
 * for every 2k characters counting from the start of the string. If there are
 * less than k characters left, reverse all of them. If there are less than 2k
 * but greater than or equal to k characters, then reverse the first k
 * characters and left the other as original.
 *
 *
 * Example:
 *
 * Input: s = "abcdefg", k = 2
 * Output: "bacdfeg"
 *
 *
 *
 * Restrictions:
 *
 * ⁠The string consists of lower English letters only.
 * ⁠Length of the given string and k will in the range [1, 10000]
 *
 */
func reverseStr(s string, k int) string {
	reverse := func(bs []byte, start int) {
		end := start + k
		if end > len(bs) {
			end = len(bs)
		}

		n := end - start
		for i := 0; i < n/2; i++ {
			bs[start+i], bs[end-i-1] = bs[end-i-1], bs[start+i]
		}
	}

	bs := []byte(s)
	var i int
	for i = 1; 2*k*i < len(s); i += 1 {
		reverse(bs, 2*k*(i-1))
	}

	i--
	reverse(bs, 2*k*i)
	return string(bs)
}
