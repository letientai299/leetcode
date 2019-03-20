package main

import (
	"strings"
)

/*
* @lc app=leetcode id=125 lang=golang
*
* [125] Valid Palindrome
*
* https://leetcode.com/problems/valid-palindrome/description/
*
* algorithms
* Easy (30.37%)
* Total Accepted:    329.4K
* Total Submissions: 1.1M
* Testcase Example:  '"A man, a plan, a canal: Panama"'
*
* Given a string, determine if it is a palindrome, considering only
* alphanumeric characters and ignoring cases.
*
* Note: For the purpose of this problem, we define empty string as valid
* palindrome.
*
* Example 1:
*
*
* Input: "A man, a plan, a canal: Panama"
* Output: true
*
*
* Example 2:
*
*
* Input: "race a car"
* Output: false
*
*
 */

var alphanums = map[byte]struct{}{
	'0': {},
	'1': {},
	'2': {},
	'3': {},
	'4': {},
	'5': {},
	'6': {},
	'7': {},
	'8': {},
	'9': {},
	'a': {},
	'b': {},
	'c': {},
	'd': {},
	'e': {},
	'f': {},
	'g': {},
	'h': {},
	'i': {},
	'j': {},
	'k': {},
	'l': {},
	'm': {},
	'n': {},
	'o': {},
	'p': {},
	'q': {},
	'r': {},
	's': {},
	't': {},
	'u': {},
	'v': {},
	'w': {},
	'x': {},
	'y': {},
	'z': {},
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		if _, ok := alphanums[s[l]]; !ok {
			l++
			continue
		}
		if _, ok := alphanums[s[r]]; !ok {
			r--
			continue
		}

		if l >= r {
			break
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}
	return true
}
