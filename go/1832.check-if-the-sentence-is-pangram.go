package main

// Check if the Sentence Is Pangram
//
// Easy
//
// https://leetcode.com/problems/check-if-the-sentence-is-pangram/
//
// A **pangram** is a sentence where every letter of the English alphabet
// appears at least once.
//
// Given a string `sentence` containing only lowercase English letters,
// return`true` _if_ `sentence` _is a **pangram**, or_ `false` _otherwise._
//
// **Example 1:**
//
// ```
// Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
// Output: true
// Explanation: sentence contains at least one of every letter of the English
// alphabet.
//
// ```
//
// **Example 2:**
//
// ```
// Input: sentence = "leetcode"
// Output: false
//
// ```
//
// **Constraints:**
//
// - `1 <= sentence.length <= 1000`
// - `sentence` consists of lowercase English letters.
func checkIfPangram(sentence string) bool {
	want := 1<<26 - 1
	m := 0
	for _, c := range sentence {
		b := int(c - 'a')
		m |= 1 << b
		if m == want {
			return true
		}
	}
	return false
}
