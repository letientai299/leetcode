package main

import "strings"

/*
 * @lc app=leetcode id=1170 lang=golang
 *
 * [1170] Compare Strings by Frequency of the Smallest Character
 *
 * https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/description/
 *
 * algorithms
 * Easy (58.23%)
 * Total Accepted:    48.2K
 * Total Submissions: 81.2K
 * Testcase Example:  '["cbd"]\n["zaaaz"]'
 *
 * Let's define a function f(s) over a non-empty string s, which calculates the
 * frequency of the smallest character in s. For example, if s = "dcce" then
 * f(s) = 2 because the smallest character is "c" and its frequency is 2.
 *
 * Now, given string arrays queries and words, return an integer array answer,
 * where each answer[i] is the number of words such that f(queries[i]) < f(W),
 * where W is a word in words.
 *
 *
 * Example 1:
 *
 *
 * Input: queries = ["cbd"], words = ["zaaaz"]
 * Output: [1]
 * Explanation: On the first query we have f("cbd") = 1, f("zaaaz") = 3 so
 * f("cbd") < f("zaaaz").
 *
 *
 * Example 2:
 *
 *
 * Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
 * Output: [1,2]
 * Explanation: On the first query only f("bbb") < f("aaaa"). On the second
 * query both f("aaa") and f("aaaa") are both > f("cc").
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= queries.length <= 2000
 * 1 <= words.length <= 2000
 * 1 <= queries[i].length, words[i].length <= 10
 * queries[i][j], words[i][j] are English lowercase letters.
 *
 *
 */
func numSmallerByFrequency(queries []string, words []string) []int {
	f := func(s string) int {
		for i := 'a'; i <= 'z'; i++ {
			c := strings.Count(s, string(i))
			if c > 0 {
				return c
			}
		}
		return 0
	}
	m := make(map[int]int) // frequency -> count

	max := 0
	for _, w := range words {
		v := f(w)
		m[v]++
		if v > max {
			max = v
		}
	}

	res := make([]int, 0, len(queries))
	for _, q := range queries {
		c := f(q)
		n := 0
		for i := c + 1; i <= max; i++ {
			n += m[i]
		}
		res = append(res, n)
	}
	return res
}
