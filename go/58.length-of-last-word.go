package main

/*
* @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.15%)
 * Total Accepted:    249.6K
 * Total Submissions: 776.1K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 *
 * Example:
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
*/
func lengthOfLastWord(s string) int {
	start := -1
	end := -1
	const sp = rune(' ')
	lastChar := sp
	i := 0
	for _, c := range s {
		if lastChar == sp && c != sp {
			start = i
		} else if lastChar != sp && c == sp {
			end = i
		}

		lastChar = c
		i++
	}

	if lastChar!=sp {
		end = i
	}

	return end - start
}
