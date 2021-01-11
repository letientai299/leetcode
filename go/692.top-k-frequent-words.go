package main

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode id=692 lang=golang
 *
 * [692] Top K Frequent Words
 *
 * https://leetcode.com/problems/top-k-frequent-words/description/
 *
 * algorithms
 * Medium (48.00%)
 * Total Accepted:    238.5K
 * Total Submissions: 451.7K
 * Testcase Example:  '["i", "love", "leetcode", "i", "love", "coding"]\n2'
 *
 * Given a non-empty list of words, return the k most frequent elements.
 * Your answer should be sorted by frequency from highest to lowest. If two
 * words have the same frequency, then the word with the lower alphabetical
 * order comes first.
 *
 * Example 1:
 *
 * Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
 * Output: ["i", "love"]
 * Explanation: "i" and "love" are the two most frequent words.
 * ⁠   Note that "i" comes before "love" due to a lower alphabetical order.
 *
 *
 *
 * Example 2:
 *
 * Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is",
 * "is"], k = 4
 * Output: ["the", "is", "sunny", "day"]
 * Explanation: "the", "is", "sunny" and "day" are the four most frequent
 * words,
 * ⁠   with the number of occurrence being 4, 3, 2 and 1 respectively.
 *
 *
 *
 * Note:
 *
 * You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
 * Input words contain only lowercase letters.
 *
 *
 *
 * Follow up:
 *
 * Try to solve it in O(n log k) time and O(n) extra space.
 *
 *
 */

func topKFrequent(words []string, k int) []string {
	cnt := make(map[string]int)
	uniq := make([]string, 0, (k+len(words))/2)
	for _, s := range words {
		if cnt[s] == 0 {
			uniq = append(uniq, s)
		}
		cnt[s]++
	}

	sort.Slice(uniq, func(i, j int) bool {
		a, b := uniq[i], uniq[j]
		x, y := cnt[a], cnt[b]
		if x > y {
			return true
		}

		if y > x {
			return false
		}

		return strings.Compare(a, b) < 0
	})

	n := len(uniq)
	if k > n {
		k = n
	}

	return uniq[:k]
}
