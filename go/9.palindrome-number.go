package main

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (42.00%)
 * Total Accepted:    521.4K
 * Total Submissions: 1.2M
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 *
 * Example 1:
 *
 *
 * Input: 121
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 * Follow up:
 *
 * Coud you solve it without converting the integer to a string?
 *
 */
func isPalindrome_i(x int) bool {
	if x < 0 {
		return false
	}

	headbase := 1
	for {
		if headbase > x {
			headbase = headbase / 10
			break
		}

		headbase *= 10
	}

	for x > 0 {
		head := (x / headbase) % 10
		tail := x % 10
		if head != tail {
			return false
		}

		x = (x - head*headbase) / 10
		headbase /= 100
	}
	return true
}
