package main

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode id=1065 lang=golang
 *
 * [1065] Index Pairs of a String
 *
 * https://leetcode.com/problems/index-pairs-of-a-string/description/
 *
 * algorithms
 * Easy (58.02%)
 * Total Accepted:    8.9K
 * Total Submissions: 14.5K
 * Testcase Example:  '"thestoryofleetcodeandme"\n["story","fleet","leetcode"]'
 *
 * Given a text string and words (a list of strings), return all index pairs
 * [i, j] so that the substring text[i]...text[j] is in the list of words.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: text = "thestoryofleetcodeandme", words =
 * ["story","fleet","leetcode"]
 * Output: [[3,7],[9,13],[10,17]]
 *
 *
 * Example 2:
 *
 *
 * Input: text = "ababa", words = ["aba","ab"]
 * Output: [[0,1],[0,2],[2,3],[2,4]]
 * Explanation:
 * Notice that matches can overlap, see "aba" is found in [0,2] and [2,4].
 *
 *
 *
 *
 * Note:
 *
 *
 * All strings contains only lowercase English letters.
 * It's guaranteed that all strings in words are different.
 * 1 <= text.length <= 100
 * 1 <= words.length <= 20
 * 1 <= words[i].length <= 50
 * Return the pairs [i,j] in sorted order (i.e. sort them by their first
 * coordinate in case of ties sort them by their second coordinate).
 *
 */

func indexPairs(text string, words []string) [][]int {
	var res [][]int
	for _, w := range words {
		start := 0
		i := strings.Index(text[start:], w)
		for i != -1 {
			res = append(res, []int{i + start, i + start + len(w) - 1})
			start += i + 1
			i = strings.Index(text[start:], w)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		a := res[i]
		b := res[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		}

		return a[0] < b[0]
	})
	return res
}
