package main

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 *
 * https://leetcode.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (51.28%)
 * Total Accepted:    312.8K
 * Total Submissions: 610K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t , write a function to determine if t is an anagram
 * of s.
 *
 * Example 1:
 *
 *
 * Input: s = "anagram", t = "nagaram"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "rat", t = "car"
 * Output: false
 *
 *
 * Note:
 * You may assume the string contains only lowercase alphabets.
 *
 * Follow up:
 * What if the inputs contain unicode characters? How would you adapt your
 * solution to such case?
 *
 */
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make([]int, 26)
	for i, x := range s {
		m[x-'a']++
		m[t[i]-'a']--
	}

	for _, i := range m {
		if i != 0 {
			return false
		}
	}
	return true
}
