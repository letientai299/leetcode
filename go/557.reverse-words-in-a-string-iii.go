package main

/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 *
 * https://leetcode.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (64.14%)
 * Total Accepted:    127.1K
 * Total Submissions: 198K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * Given a string, you need to reverse the order of characters in each word
 * within a sentence while still preserving whitespace and initial word order.
 *
 * Example 1:
 *
 * Input: "Let's take LeetCode contest"
 * Output: "s'teL ekat edoCteeL tsetnoc"
 *
 *
 *
 * Note:
 * In the string, each word is separated by single space and there will not be
 * any extra space in the string.
 *
 */
func reverseWords557(s string) string {
	buf := []byte(s)
	rev := func(buf []byte, left, right int) {
		n := right - left
		if n < 0 {
			return
		}
		for i := 0; i <= n/2; i++ {
			buf[i+left], buf[right-i] = buf[right-i], buf[left+i]
		}
	}
	left := 0
	for i := 1; i < len(buf); i++ {
		if buf[i] == ' ' && buf[i-1] != ' ' {
			rev(buf, left, i-1)
			continue
		}

		if buf[i] != ' ' && buf[i-1] == ' ' {
			left = i
		}
	}
	rev(buf, left, len(buf)-1)
	return string(buf)
}
