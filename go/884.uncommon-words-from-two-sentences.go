package main

import "strings"

/*
 * @lc app=leetcode id=884 lang=golang
 *
 * [884] Uncommon Words from Two Sentences
 *
 * https://leetcode.com/problems/uncommon-words-from-two-sentences/description/
 *
 * algorithms
 * Easy (61.47%)
 * Total Accepted:    64K
 * Total Submissions: 100.5K
 * Testcase Example:  '"this apple is sweet"\n"this apple is sour"'
 *
 * We are given two sentences A and B.  (A sentence is a string of space
 * separated words.  Each word consists only of lowercase letters.)
 *
 * A word is uncommon if it appears exactly once in one of the sentences, and
 * does not appear in the other sentence.
 *
 * Return a list of all uncommon words.
 *
 * You may return the list in any order.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: A = "this apple is sweet", B = "this apple is sour"
 * Output: ["sweet","sour"]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: A = "apple apple", B = "banana"
 * Output: ["banana"]
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= A.length <= 200
 * 0 <= B.length <= 200
 * A and B both contain only spaces and lowercase letters.
 *
 *
 *
 *
 */

func uncommonFromSentences(a string, b string) []string {
	wa := strings.Split(a, " ")
	ma := make(map[string]int)
	for _, w := range wa {
		if len(w) == 0 {
			continue
		}
		ma[w]++
	}

	wb := strings.Split(b, " ")
	mb := make(map[string]int)
	for _, w := range wb {
		if len(w) == 0 {
			continue
		}
		mb[w]++
	}

	var res []string

	for w, v := range ma {
		if v != 1 {
			continue
		}
		if mb[w] == 0 {
			res = append(res, w)
		}
	}

	for w, v := range mb {
		if v != 1 {
			continue
		}
		if ma[w] == 0 {
			res = append(res, w)
		}
	}

	return res
}
