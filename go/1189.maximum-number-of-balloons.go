package main

/*
 * @lc app=leetcode id=1189 lang=golang
 *
 * [1189] Maximum Number of Balloons
 *
 * https://leetcode.com/problems/maximum-number-of-balloons/description/
 *
 * algorithms
 * Easy (61.40%)
 * Total Accepted:    47.1K
 * Total Submissions: 76.2K
 * Testcase Example:  '"nlaebolko"'
 *
 * Given a string text, you want to use the characters of text to form as many
 * instances of the word "balloon" as possible.
 *
 * You can use each character in text at most once. Return the maximum number
 * of instances that can be formed.
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: text = "nlaebolko"
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: text = "loonbalxballpoon"
 * Output: 2
 *
 *
 * Example 3:
 *
 *
 * Input: text = "leetcode"
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= text.length <= 10^4
 * text consists of lower case English letters only.
 *
 */
func maxNumberOfBalloons(text string) int {
	base := map[byte]int{
		'b': 1,
		'a': 1,
		'l': 2,
		'o': 2,
		'n': 1,
	}
	m := make(map[byte]int)
	for _, x := range text {
		m[byte(x)]++
	}

	res := -1
	for k, v := range base {
		x, ok := m[k]
		if !ok {
			return 0
		}
		n := x / v
		if res == -1 || n < res {
			res = n
		}
	}

	if res == -1 {
		res = 0
	}

	return res
}
