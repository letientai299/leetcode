package main

/*
* @lc app=leetcode id=409 lang=golang
*
* [409] Longest Palindrome
*
* https://leetcode.com/problems/longest-palindrome/description/
*
* algorithms
* Easy (47.64%)
* Total Accepted:    91.9K
* Total Submissions: 192.8K
* Testcase Example:  '"abccccdd"'
*
* Given a string which consists of lowercase or uppercase letters, find the
* length of the longest palindromes that can be built with those letters.
*
* This is case sensitive, for example "Aa" is not considered a palindrome
* here.
*
* Note:
* Assume the length of given string will not exceed 1,010.
*
*
* Example:
*
* Input:
* "abccccdd"
*
* Output:
* 7
*
* Explanation:
* One longest palindrome that can be built is "dccaccd", whose length is 7.
*
*
 */
func longestPalindrome(s string) int {
	distribution := make(map[int32]int, 26*2)
	for _, c := range s {
		distribution[c]++
	}

	r := 0
	hasOdd := false
	for _, v := range distribution {
		if v%2 == 0 {
			r += v
		} else {
			hasOdd = true
			r += v - 1
		}
	}
	if hasOdd {
		return r + 1
	}
	return r
}
