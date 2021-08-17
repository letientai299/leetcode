package main

import "strings"

// Longest Substring Of All Vowels in Order
//
// Medium
//
// https://leetcode.com/problems/longest-substring-of-all-vowels-in-order/
//
// A string is considered **beautiful** if it satisfies the following
// conditions:
//
// - Each of the 5 English vowels (`'a'`, `'e'`, `'i'`, `'o'`, `'u'`) must
// appear **at least once** in it.
// - The letters must be sorted in **alphabetical order** (i.e. all `'a'`s
// before `'e'`s, all `'e'`s before `'i'`s, etc.).
//
// For example, strings `"aeiou"` and `"aaaaaaeiiiioou"` are considered
// **beautiful**, but `"uaeio"`, `"aeoiu"`, and `"aaaeeeooo"` are **not
// beautiful**.
//
// Given a string `word` consisting of English vowels, return _the **length of
// the longest beautiful substring** of_ `word` _. If no such substring exists,
// return_ `0`.
//
// A **substring** is a contiguous sequence of characters in a string.
//
// **Example 1:**
//
// ```
// Input: word = "aeiaaioaaaaeiiiiouuuooaauuaeiu"
// Output: 13
// Explanation: The longest beautiful substring in word is "aaaaeiiiiouuu" of
// length 13.
// ```
//
// **Example 2:**
//
// ```
// Input: word = "aeeeiiiioooauuuaeiou"
// Output: 5
// Explanation: The longest beautiful substring in word is "aeiou" of length 5.
//
// ```
//
// **Example 3:**
//
// ```
// Input: word = "a"
// Output: 0
// Explanation: There is no beautiful substring, so return 0.
//
// ```
//
// **Constraints:**
//
// - `1 <= word.length <= 5 * 105`
// - `word` consists of characters `'a'`, `'e'`, `'i'`, `'o'`, and `'u'`.
func longestBeautifulSubstring(word string) int {
	best := 0
	vowels := "aeiou"
	vi := 0
	start := -1
	for i := 0; i < len(word); i++ {
		b := word[i]
		x := strings.IndexByte(vowels, b)
		if start == -1 {
			if x == 0 {
				start = i
				vi = 0
			}
			continue
		}

		switch x {
		case vi:
			continue
		case vi + 1:
			vi++
		default:
			if vi == len(vowels)-1 {
				if start != -1 {
					cur := i - start
					if cur > best {
						best = cur
					}
				}
				start = -1
				vi = 0
			} else {
				start = -1
			}

			if start == -1 && x == 0 {
				start = i
				vi = 0
			}
		}
	}

	if vi == len(vowels)-1 {
		cur := len(word) - start
		if cur > best {
			best = cur
		}
	}

	return best
}
