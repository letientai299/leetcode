package main

// Longest String Chain
//
// # Medium
//
// https://leetcode.com/problems/longest-string-chain
//
// You are given an array of `words` where each word consists of lowercase
// English letters.
//
// `wordA` is a **predecessor** of `wordB` if and only if we can insert
// **exactly one** letter anywhere in `wordA` **without changing the order of
// the other characters** to make it equal to `wordB`.
//
// - For example, `"abc"` is a **predecessor** of `"abac"`, while `"cba"` is not
// a **predecessor** of `"bcad"`.
//
// A **word chain** is a sequence of words `[word1, word2, ..., wordk]` with `k
// >= 1`, where `word1` is a **predecessor** of `word2`, `word2` is a
// **predecessor** of `word3`, and so on. A single word is trivially a **word
// chain** with `k == 1`.
//
// Return _the **length** of the **longest possible word chain** with words
// chosen from the given list of_ `words`.
//
// **Example 1:**
//
// ```
// Input: words = ["a","b","ba","bca","bda","bdca"]
// Output: 4
// Explanation: One of the longest word chains is ["a","ba","bda","bdca"].
//
// ```
//
// **Example 2:**
//
// ```
// Input: words = ["xbc","pcxbcf","xb","cxbc","pcxbc"]
// Output: 5
// Explanation: All the words can be put in a word chain ["xb", "xbc", "cxbc",
// "pcxbc", "pcxbcf"].
//
// ```
//
// **Example 3:**
//
// ```
// Input: words = ["abcd","dbqca"]
// Output: 1
// Explanation: The trivial word chain ["abcd"] is one of the longest word
// chains.
// ["abcd","dbqca"] is not a valid word chain because the ordering of the
// letters is changed.
//
// ```
//
// **Constraints:**
//
// - `1 <= words.length <= 1000`
// - `1 <= words[i].length <= 16`
// - `words[i]` only consists of lowercase English letters.
func longestStrChain(words []string) int {
	m := make([]map[string]int, 17) // max word length
	for _, w := range words {
		n := len(w)
		if m[n] == nil {
			m[n] = make(map[string]int)
		}
		m[n][w] = 1
	}

	res := 1
	for i := 2; i <= 16; i++ {
		for w, v := range m[i] {
			best := v
			for j := range w {
				pre := w[:j] + w[j+1:] // generate predecessor
				best = max(best, m[i-1][pre]+1)
			}
			m[i][w] = best

			res = max(best, res)
		}
	}

	return res
}
