package main

/*
 * @lc app=leetcode id=520 lang=golang
 *
 * [520] Detect Capital
 *
 * https://leetcode.com/problems/detect-capital/description/
 *
 * algorithms
 * Easy (52.44%)
 * Total Accepted:    84.7K
 * Total Submissions: 161.5K
 * Testcase Example:  '"USA"'
 *
 * Given a word, you need to judge whether the usage of capitals in it is right
 * or not.
 *
 * We define the usage of capitals in a word to be right when one of the
 * following cases holds:
 *
 *
 * All letters in this word are capitals, like "USA".
 * All letters in this word are not capitals, like "leetcode".
 * Only the first letter in this word is capital, like "Google".
 *
 * Otherwise, we define that this word doesn't use capitals in a right way.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "USA"
 * Output: True
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "FlaG"
 * Output: False
 *
 *
 *
 *
 * Note: The input will be a non-empty word consisting of uppercase and
 * lowercase latin letters.
 *
 */
func detectCapitalUse(word string) bool {
	countUpperCase := 0
	lastUpper := -1
	for i, c := range word {
		if c <= int32('Z') {
			countUpperCase++
			lastUpper = i
		}
	}
	return countUpperCase == 0 || (countUpperCase == 1 && lastUpper == 0) || (countUpperCase == len(word))
}
