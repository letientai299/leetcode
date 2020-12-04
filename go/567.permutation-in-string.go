package main

/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 *
 * https://leetcode.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (39.57%)
 * Total Accepted:    161.2K
 * Total Submissions: 362.4K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * Given two strings s1 and s2, write a function to return true if s2 contains
 * the permutation of s1. In other words, one of the first string's
 * permutations is the substring of the second string.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: s1 = "ab" s2 = "eidbaooo"
 * Output: True
 * Explanation: s2 contains one permutation of s1 ("ba").
 *
 *
 * Example 2:
 *
 *
 * Input:s1= "ab" s2 = "eidboaoo"
 * Output: False
 *
 *
 *
 * Constraints:
 *
 *
 * The input strings only contain lower case letters.
 * The length of both given strings is in range [1, 10,000].
 *
 *
 */

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	m := [26]int{}
	for i := range s1 {
		m[s1[i]-'a']--
		m[s2[i]-'a']++
	}

	diff := 0
	for _, v := range m {
		if v != 0 {
			diff++
		}
	}

	if diff == 0 {
		return true
	}

	for i := 1; i <= len(s2)-len(s1); i++ {
		if m[s2[i-1]-'a'] == 0 {
			diff++
		}
		m[s2[i-1]-'a']--
		if m[s2[i-1]-'a'] == 0 {
			diff--
		}

		if m[s2[i+len(s1)-1]-'a'] == 0 {
			diff++
		}
		m[s2[i+len(s1)-1]-'a']++
		if m[s2[i+len(s1)-1]-'a'] == 0 {
			diff--
		}

		if diff == 0 {
			return true
		}
	}

	return false
}
