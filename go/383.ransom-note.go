package main

/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 *
 * https://leetcode.com/problems/ransom-note/description/
 *
 * algorithms
 * Easy (49.45%)
 * Total Accepted:    106.2K
 * Total Submissions: 214.8K
 * Testcase Example:  '"a"\n"b"'
 *
 *
 * Given an arbitrary ransom note string and another string containing letters
 * from all the magazines, write a function that will return true if the
 * ransom
 * note can be constructed from the magazines ; otherwise, it will return
 * false.
 *
 *
 * Each letter in the magazine string can only be used once in your ransom
 * note.
 *
 *
 * Note:
 * You may assume that both strings contain only lowercase letters.
 *
 *
 *
 * canConstruct("a", "b") -> false
 * canConstruct("aa", "ab") -> false
 * canConstruct("aa", "aab") -> true
 *
 *
 */
func canConstruct383(note string, magazine string) bool {
	m := make(map[int32]int, 128)
	for _, c := range magazine {
		m[c]++
	}

	for _, c := range note {
		m[c]--
		if m[c] < 0 {
			return false
		}
	}

	return true
}

// Good for interview
