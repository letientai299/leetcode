package main

import "strings"

// Number of Valid Words in a Sentence
//
// Easy
//
// https://leetcode.com/problems/number-of-valid-words-in-a-sentence/
//
// A sentence consists of lowercase letters (`'a'` to `'z'`), digits (`'0'` to
// `'9'`), hyphens (`'-'`), punctuation marks (`'!'`, `'.'`, and `','`), and
// spaces (`' '`) only. Each sentence can be broken down into **one or more
// tokens** separated by one or more spaces `' '`.
//
// A token is a valid word if **all three** of the following are true:
//
// - It only contains lowercase letters, hyphens, and/or punctuation ( **no**
// digits).
// - There is **at most one** hyphen `'-'`. If present, it **must** be
// surrounded by lowercase characters (`"a-b"` is valid, but `"-ab"` and `"ab-"`
// are not valid).
// - There is **at most one** punctuation mark. If present, it **must** be at
// the **end** of the token (`"ab,"`, `"cd!"`, and `"."` are valid, but `"a!b"`
// and `"c.,"` are not valid).
//
// Examples of valid words include `"a-b."`, `"afad"`, `"ba-c"`, `"a!"`, and
// `"!"`.
//
// Given a string `sentence`, return _the **number** of valid words in_
// `sentence`.
//
// **Example 1:**
//
// ```
// Input: sentence = "cat and  dog"
// Output: 3
// Explanation: The valid words in the sentence are "cat", "and", and "dog".
//
// ```
//
// **Example 2:**
//
// ```
// Input: sentence = "!this  1-s b8d!"
// Output: 0
// Explanation: There are no valid words in the sentence.
// "!this" is invalid because it starts with a punctuation mark.
// "1-s" and "b8d" are invalid because they contain digits.
//
// ```
//
// **Example 3:**
//
// ```
// Input: sentence = "alice and  bob are playing stone-game10"
// Output: 5
// Explanation: The valid words in the sentence are "alice", "and", "bob",
// "are", and "playing".
// "stone-game10" is invalid because it contains digits.
//
// ```
//
// **Constraints:**
//
// - `1 <= sentence.length <= 1000`
// - `sentence` only contains lowercase English letters, digits, `' '`, `'-'`,
// `'!'`, `'.'`, and `','`.
// - There will be at least `1` token.
func countValidWords(s string) int {
	ws := strings.Split(s, " ")
	cnt := 0
	isChar := func(c rune) bool { return 'a' <= c && c <= 'z' }
	isDigit := func(c rune) bool { return '0' <= c && c <= '9' }
	isPunc := func(c rune) bool { return c == '!' || c == '.' || c == ',' }

	for _, w := range ws {
		if len(w) == 0 {
			continue
		}

		ok := true
		hyphen := false
		for i, c := range w {
			if isDigit(c) {
				ok = false
				break
			}

			if isPunc(c) {
				if i != len(w)-1 {
					ok = false
					break
				}
			}

			if c == '-' {
				if hyphen {
					ok = false
					break
				}

				if i < 1 || i >= len(w)-1 ||
					!(isChar(rune(w[i-1])) && isChar(rune(w[i+1]))) {
					ok = false
					break
				}
				hyphen = true
			}

		}

		if ok {
			cnt++
		}
	}
	return cnt
}
