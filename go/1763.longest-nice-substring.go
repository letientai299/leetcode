package main

// Longest Nice Substring
//
// Easy
//
// https://leetcode.com/problems/longest-nice-substring/
//
// A string `s` is **nice** if, for every letter of the alphabet that `s`
// contains, it appears **both** in uppercase and lowercase. For example,
// `"abABB"` is nice because `'A'` and `'a'` appear, and `'B'` and `'b'` appear.
// However, `"abA"` is not because `'b'` appears, but `'B'` does not.
//
// Given a string `s`, return _the longest **substring** of `s` that is
// **nice**. If there are multiple, return the substring of the **earliest**
// occurrence. If there are none, return an empty string_.
//
// **Example 1:**
//
// ```
// Input: s = "YazaAay"
// Output: "aAa"
// Explanation: "aAa" is a nice string because 'A/a' is the only letter of the
// alphabet in s, and both 'A' and 'a' appear.
// "aAa" is the longest nice substring.
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "Bb"
// Output: "Bb"
// Explanation: "Bb" is a nice string because both 'B' and 'b' appear. The whole
// string is a substring.
// ```
//
// **Example 3:**
//
// ```
// Input: s = "c"
// Output: ""
// Explanation: There are no nice substrings.
// ```
//
// **Example 4:**
//
// ```
// Input: s = "dDzeE"
// Output: "dD"
// Explanation: Both "dD" and "eE" are the longest nice substrings.
// As there are multiple longest nice substrings, return "dD" since it occurs
// earlier.
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 100`
// - `s` consists of uppercase and lowercase English letters.
func longestNiceSubstring(s string) string {
	check := func(s string) [26]int {
		nices := [26]int{}
		for _, c := range s {
			if 'a' <= c && c <= 'z' {
				nices[c-'a'] |= 1
			} else {
				nices[c-'A'] |= 2
			}
		}
		return nices
	}

	nices := check(s)
	isNiceChar := func(i int) bool {
		c := s[i]
		isLower := 'a' <= c && c <= 'z'
		if isLower {
			return nices[c-'a'] == 3
		}
		return nices[c-'A'] == 3
	}

	isNiceString := func(i, j int) bool {
		v := check(s[i : j+1])
		for _, x := range v {
			if x != 0 && x != 3 {
				return false
			}
		}

		return true
	}

	res := ""
	for i := 0; i < len(s); i++ {
		if !isNiceChar(i) {
			continue
		}

		for j := i + 1; j < len(s) && isNiceChar(j); j++ {
			if !isNiceString(i, j) {
				continue
			}
			if len(res) < j-i+1 {
				res = s[i : j+1]
			}
		}
	}

	return res
}

// Good/Hard for interview
