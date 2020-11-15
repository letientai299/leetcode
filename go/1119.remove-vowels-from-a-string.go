package main

import "strings"

/*
 * @lc app=leetcode id=1119 lang=golang
 *
 * [1119] Remove Vowels from a String
 *
 * https://leetcode.com/problems/remove-vowels-from-a-string/description/
 *
 * algorithms
 * Easy (88.09%)
 * Total Accepted:    48.9K
 * Total Submissions: 54.3K
 * Testcase Example:  '"leetcodeisacommunityforcoders"'
 *
 * Given a string S, remove the vowels 'a', 'e', 'i', 'o', and 'u' from it, and
 * return the new string.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "leetcodeisacommunityforcoders"
 * Output: "ltcdscmmntyfrcdrs"
 *
 *
 * Example 2:
 *
 *
 * Input: "aeiou"
 * Output: ""
 *
 *
 *
 *
 * Note:
 *
 *
 * S consists of lowercase English letters only.
 * 1 <= S.length <= 1000
 *
 *
 */

func removeVowels(s string) string {
	return strings.NewReplacer("a", "", "e", "", "i", "", "u", "", "o", "").Replace(s)
}
