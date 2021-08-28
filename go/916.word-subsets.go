package main

// Word Subsets
//
// Medium
//
// https://leetcode.com/problems/word-subsets/
//
// You are given two string arrays `words1` and `words2`.
//
// A string `b` is a **subset** of string `a` if every letter in `b` occurs in
// `a` including multiplicity.
//
// - For example, `"wrr"` is a subset of `"warrior"` but is not a subset of
// `"world"`.
//
// A string `a` from `words1` is **universal** if for every string `b` in
// `words2`, `b` is a subset of `a`.
//
// Return an array of all the **universal** strings in `words1`. You may return
// the answer in **any order**.
//
// **Example 1:**
//
// ```
// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 =
// ["e","o"]
// Output: ["facebook","google","leetcode"]
//
// ```
//
// **Example 2:**
//
// ```
// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 =
// ["l","e"]
// Output: ["apple","google","leetcode"]
//
// ```
//
// **Example 3:**
//
// ```
// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 =
// ["e","oo"]
// Output: ["facebook","google"]
//
// ```
//
// **Example 4:**
//
// ```
// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 =
// ["lo","eo"]
// Output: ["google","leetcode"]
//
// ```
//
// **Example 5:**
//
// ```
// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 =
// ["ec","oc","ceo"]
// Output: ["facebook","leetcode"]
//
// ```
//
// **Constraints:**
//
// - `1 <= words1.length, words2.length <= 104`
// - `1 <= words1[i].length, words2[i].length <= 10`
// - `words1[i]` and `words2[i]` consist only of lowercase English letters.
// - All the strings of `words1` are **unique**.
func wordSubsets(words1 []string, words2 []string) []string {
	count := func(s string) [26]int {
		var r [26]int
		for _, c := range s {
			r[int(c-'a')]++
		}
		return r
	}

	b := count(words2[0])

	for _, w := range words2[1:] {
		other := count(w)
		for i, v := range b {
			if v < other[i] {
				b[i] = other[i]
			}
		}
	}

	i := 0
	for _, s := range words1 {
		k := 0
		a := count(s)
		for k < 26 {
			if a[k] < b[k] {
				break
			}
			k++
		}

		if k == 26 {
			words1[i] = s
			i++
		}
	}
	return words1[:i]
}

// Good for interview
