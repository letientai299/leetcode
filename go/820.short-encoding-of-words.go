package main

import "strings"

// Short Encoding of Words
//
// Medium
//
// https://leetcode.com/problems/short-encoding-of-words/
//
// A **valid encoding** of an array of `words` is any reference string `s` and
// array of indices `indices` such that:
//
// - `words.length == indices.length`
// - The reference string `s` ends with the `'#'` character.
// - For each index `indices[i]`, the **substring** of `s` starting from
// `indices[i]` and up to (but not including) the next `'#'` character is equal
// to `words[i]`.
//
// Given an array of `words`, return _the **length of the shortest reference
// string**_ `s` _possible of any **valid encoding** of_ `words` _._
//
// **Example 1:**
//
// ```
// "time#bell#" and indices = [0, 2, 5
// ```
//
// **Example 2:**
//
// ```
// Input: words = ["t"]
// Output: 2
// Explanation: A valid encoding would be s = "t#" and indices = [0].
//
// ```
//
// **Constraints:**
//
// - `1 <= words.length <= 2000`
// - `1 <= words[i].length <= 7`
// - `words[i]` consists of only lowercase letters.
func minimumLengthEncoding(words []string) int {
	groups := make(map[string]bool)

	for _, w := range words {
		ok := false
		for g, covered := range groups {
			if covered {
				continue
			}

			if strings.HasSuffix(g, w) {
				ok = true
				continue
			}

			if strings.HasSuffix(w, g) {
				groups[g] = true
			}
		}

		if !ok {
			groups[w] = false
		}
	}

	r := 0
	for k, covered := range groups {
		if !covered {
			r += 1 + len(k)
		}
	}
	return r
}
