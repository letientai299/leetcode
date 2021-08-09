package main

import "sort"

// Minimum Deletions to Make Character Frequencies Unique
//
// Medium
//
// https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
//
// A string `s` is called **good** if there are no two different characters in
// `s` that have the same **frequency**.
//
// Given a string `s`, return _the **minimum** number of characters you need to
// delete to make_ `s` _**good**._
//
// The **frequency** of a character in a string is the number of times it
// appears in the string. For example, in the string `"aab"`, the **frequency**
// of `'a'` is `2`, while the **frequency** of `'b'` is `1`.
//
// **Example 1:**
//
// ```
// s
// ```
//
// **Example 2:**
//
// ```
// Input: s = "aaabbbcc"
// Output: 2
// Explanation: You can delete two 'b's resulting in the good string "aaabcc".
// Another way it to delete one 'b' and one 'c' resulting in the good string
// "aaabbc".
// ```
//
// **Example 3:**
//
// ```
// Input: s = "ceabaacb"
// Output: 2
// Explanation: You can delete both 'c's resulting in the good string "eabaab".
// Note that we only care about characters that are still in the string at the
// end (i.e. frequency of 0 is ignored).
//
// ```
//
// **Constraints:**
//
// - `1 <= s.length <= 105`
// - `s` contains only lowercase English letters.
func minDeletions(s string) int {
	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}
	sort.Ints(freq)
	res := 0
	pre := freq[25]
	for i := 24; i >= 0 && freq[i] != 0; i-- {
		cur := freq[i]
		if cur >= pre {
			if pre == 0 {
				res += cur
				cur = 0
			} else {
				res += cur - pre + 1
				cur = pre - 1
			}
		}
		pre = cur
	}

	return res
}

// Good for interview
