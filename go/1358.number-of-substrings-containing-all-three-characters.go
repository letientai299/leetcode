package main

// Number of Substrings Containing All Three Characters
//
// Medium
//
// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/
//
// Given a string `s` consisting only of characters _a_, _b_ and _c_.
//
// Return the number of substrings containing **at least** one occurrence of all
// these characters _a_, _b_ and _c_.
//
// **Example 1:**
//
// ```
// Input: s = "abcabc"
// Output: 10
// Explanation: The substrings containing at least one occurrence of the
// characters a, b and c are "abc", "abca", "abcab", "abcabc", "bca", "bcab",
// "bcabc", "cab", "cabc" and "abc" (again).
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "aaacb"
// Output: 3
// Explanation: The substrings containing at least one occurrence of the
// characters a, b and c are "aaacb", "aacb" and "acb".
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "abc"
// Output: 1
//
// ```
//
// **Constraints:**
//
// - `3 <= s.length <= 5 x 10^4`
// - `s` only consists of _a_, _b_ or _c_ characters.
func numberOfSubstrings(s string) int {
	// calculate number of sub strings don't contain all a, b, c.
	// First kind are those with length less than 3.
	// Second kind is counted using 2 runners.
	i, j, n := 0, 0, len(s)
	mask := 0
	invalids := n + (n - 1) // sub strings of 1 and 2 chars
	var lastSeen [3]int
	for ; j < n; j++ {
		x := int(s[j] - 'a')
		lastSeen[x] = j

		mask |= 1 << x
		if mask != 7 {
			continue
		}

		// s[i:j] doesn't contains all 3 chars.
		k := j - i - 2
		invalids += k * (k + 1) / 2
		y := int(s[j-1] - 'a')
		mask = (1 << x) | (1 << y)
		for z, v := range lastSeen {
			if z != x && z != y {
				i = v + 1
				break
			}
		}

		// remove duplicated counts
		k = j - i - 2
		invalids -= k * (k + 1) / 2
	}

	if mask != 7 {
		k := j - i - 2
		invalids += k * (k + 1) / 2
	}

	return n*(n+1)/2 - invalids
}
