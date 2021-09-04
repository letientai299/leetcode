package main

// Substring with Concatenation of All Words
//
// Hard
//
// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
//
// You are given a string `s` and an array of strings `words` of **the same
// length**. Return all starting indices of substring(s) in `s` that is a
// concatenation of each word in `words` **exactly once**, **in any order**, and
// **without any intervening characters**.
//
// You can return the answer in **any order**.
//
// **Example 1:**
//
// ```
// Input: s = "barfoothefoobarman", words = ["foo","bar"]
// Output: [0,9]
// Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar"
// respectively.
// The output order does not matter, returning [9,0] is fine too.
//
// ```
//
// **Example 2:**
//
// ```
// Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
// Output: []
//
// ```
//
// **Example 3:**
//
// ```
// Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
// Output: [6,9,12]
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 104`
// - `s` consists of lower-case English letters.
// - `1 <= words.length <= 5000`
// - `1 <= words[i].length <= 30`
// - `words[i]` consists of lower-case English letters.
func findSubstring(s string, words []string) []int {
	var chars [26]int

	nw := len(words[0])
	total := len(words) * nw
	for _, w := range words {
		for _, c := range w {
			chars[int(c-'a')]--
		}
	}

	i := 0
	for ; i < total && i < len(s); i++ {
		chars[int(s[i]-'a')]++
	}

	diff := 0
	for _, v := range chars {
		if v != 0 {
			diff++
		}
	}

	// check whether the sub string contains all words,
	// and if so, return the index of the words in their appear orders.
	match := func(sub string) []int {
		j := 0
		var res []int
		seen := make(map[int]bool)
		for ; j < len(sub); j += nw {
			a := sub[j : j+nw]
			ok := false
			for i, w := range words {
				if !seen[i] && w == a {
					res = append(res, i)
					seen[i] = true
					ok = true
					break
				}
			}
			if !ok {
				return nil
			}
		}
		return res
	}

	var res []int

	for ; i <= len(s); i++ {
		if diff == 0 {
			// found potential match
			ids := match(s[i-total : i])
			if len(ids) != 0 {
				// collect the starting index
				res = append(res, i-total)

				// now, try to quick check the remaining of the string
				/*
					for i < len(s) {
						id := ids[0]
						w := words[id] // first word of the sub strings
						if i+nw <= len(s) && w == s[i:i+nw] {
							res = append(res, i-total+nw)
							ids = append(ids[1:], id)
							i += nw
						} else {
							break
						}
					}
				*/
			}
		}

		if i >= len(s) {
			break
		}

		pre := int(s[i-total] - 'a')
		if chars[pre] == 0 {
			diff++
		}
		chars[pre]--
		if chars[pre] == 0 {
			diff--
		}

		cur := int(s[i] - 'a')
		if chars[cur] == 0 {
			diff++
		}
		chars[cur]++
		if chars[cur] == 0 {
			diff--
		}
	}

	return res
}

// TODO (tai): slow, 3346 ms, 5.10%
