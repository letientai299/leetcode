package main

/*
 * @lc app=leetcode id=524 lang=golang
 *
 * [524] Longest Word in Dictionary through Deleting
 *
 * https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/description/
 *
 * algorithms
 * Medium (47.14%)
 * Total Accepted:    75.1K
 * Total Submissions: 154K
 * Testcase Example:  '"abpcplea"\n["ale","apple","monkey","plea"]'
 *
 *
 * Given a string and a string dictionary, find the longest string in the
 * dictionary that can be formed by deleting some characters of the given
 * string. If there are more than one possible results, return the longest word
 * with the smallest lexicographical order. If there is no possible result,
 * return the empty string.
 *
 * Example 1:
 *
 * Input:
 * s = "abpcplea", d = ["ale","apple","monkey","plea"]
 *
 * Output:
 * "apple"
 *
 *
 *
 *
 * Example 2:
 *
 * Input:
 * s = "abpcplea", d = ["a","b","c"]
 *
 * Output:
 * "a"
 *
 *
 *
 * Note:
 *
 * All the strings in the input will only contain lower-case letters.
 * The size of the dictionary won't exceed 1,000.
 * The length of all the strings in the input won't exceed 1,000.
 *
 *
 */

func findLongestWord(s string, d []string) string {
	best := ""
	for _, w := range d {
		if len(w) > len(s) || len(w) < len(best) {
			continue
		}

		i, j := 0, 0
		for i < len(s) && j < len(w) && len(s)-i >= len(w)-j {
			if s[i] == w[j] {
				j++
			}

			i++
		}

		if j != len(w) {
			continue
		}

		if len(best) < j || (len(best) == len(w) && best > w) {
			best = w
		}
	}

	return best
}
