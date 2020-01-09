package main

import "math"

/*
 * @lc app=leetcode id=243 lang=golang
 *
 * [243] Shortest Word Distance
 *
 * https://leetcode.com/problems/shortest-word-distance/description/
 *
 * algorithms
 * Easy (59.12%)
 * Total Accepted:    83.4K
 * Total Submissions: 140.3K
 * Testcase Example:  '["practice", "makes", "perfect", "coding", "makes"]\n"coding"\n"practice"'
 *
 * Given a list of words and two words word1 and word2, return the shortest
 * distance between these two words in the list.
 *
 * Example:
 * Assume that words = ["practice", "makes", "perfect", "coding", "makes"].
 *
 *
 * Input: word1 = “coding”, word2 = “practice”
 * Output: 3
 *
 *
 *
 * Input: word1 = "makes", word2 = "coding"
 * Output: 1
 *
 *
 * Note:
 * You may assume that word1 does not equal to word2, and word1 and word2 are
 * both in the list.
 *
 */
func shortestDistance(words []string, word1 string, word2 string) int {
	x, y := -1, -1
	result := math.MaxInt32
	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	for i, w := range words {
		if w == word1 {
			x = i
			if y != -1 {
				result = min(result, x-y)
			}
			continue
		}

		if w == word2 {
			y = i
			if x != -1 {
				result = min(result, y-x)
			}
			continue
		}

	}

	return result
}
