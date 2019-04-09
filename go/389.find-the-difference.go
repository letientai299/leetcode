package main

/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 *
 * https://leetcode.com/problems/find-the-difference/description/
 *
 * algorithms
 * Easy (52.82%)
 * Total Accepted:    138.2K
 * Total Submissions: 261.7K
 * Testcase Example:  '"abcd"\n"abcde"'
 *
 *
 * Given two strings s and t which consist of only lowercase letters.
 *
 * String t is generated by random shuffling string s and then add one more
 * letter at a random position.
 *
 * Find the letter that was added in t.
 *
 * Example:
 *
 * Input:
 * s = "abcd"
 * t = "abcde"
 *
 * Output:
 * e
 *
 * Explanation:
 * 'e' is the letter that was added.
 *
 */
func findTheDifference(s string, t string) byte {
	m := make(map[byte]int, 128)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		m[t[i]]--
		if m[t[i]] == -1 {
			return t[i]
		}
	}
	return 0
}