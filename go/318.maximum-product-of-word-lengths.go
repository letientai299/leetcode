package main

import "sort"

/*
 * @lc app=leetcode id=318 lang=golang
 *
 * [318] Maximum Product of Word Lengths
 *
 * https://leetcode.com/problems/maximum-product-of-word-lengths/description/
 *
 * algorithms
 * Medium (49.65%)
 * Total Accepted:    103.9K
 * Total Submissions: 200.1K
 * Testcase Example:  '["abcw","baz","foo","bar","xtfn","abcdef"]'
 *
 * Given a string array words, find the maximum value of length(word[i]) *
 * length(word[j]) where the two words do not share common letters. You may
 * assume that each word will contain only lower case letters. If no such two
 * words exist, return 0.
 *
 * Example 1:
 *
 *
 * Input: ["abcw","baz","foo","bar","xtfn","abcdef"]
 * Output: 16
 * Explanation: The two words can be "abcw", "xtfn".
 *
 * Example 2:
 *
 *
 * Input: ["a","ab","abc","d","cd","bcd","abcd"]
 * Output: 4
 * Explanation: The two words can be "ab", "cd".
 *
 * Example 3:
 *
 *
 * Input: ["a","aa","aaa","aaaa"]
 * Output: 0
 * Explanation: No such pair of words.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= words.length <= 10^3
 * 0 <= words[i].length <= 10^3
 * words[i] consists only of lowercase English letters.
 *
 *
 */
func maxProduct(words []string) int {
	if len(words) < 2 {
		return 0
	}

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	calc := func(w string) int32 {
		t := int32(0)
		for _, c := range w {
			t |= 1 << (c - 'a')
		}
		return t
	}

	mask := make([]int32, len(words))
	for i, w := range words {
		mask[i] = calc(w)
	}

	best := 0
	n := len(words)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			v := len(words[i]) * len(words[j])
			if v < best {
				break
			}

			if mask[i]&mask[j] != 0 {
				continue
			}
			best = v
			break
		}
	}

	return best
}
