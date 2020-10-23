package main

import (
	"sort"
	"strings"
	"unicode"
)

/*
 * @lc app=leetcode id=937 lang=golang
 *
 * [937] Reorder Data in Log Files
 *
 * https://leetcode.com/problems/reorder-data-in-log-files/description/
 *
 * algorithms
 * Easy (53.90%)
 * Total Accepted:    162.8K
 * Total Submissions: 299.5K
 * Testcase Example:  '["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"]'
 *
 * You have an array of logs.  Each log is a space delimited string of words.
 *
 * For each log, the first word in each log is an alphanumeric identifier.
 * Then, either:
 *
 *
 * Each word after the identifier will consist only of lowercase letters,
 * or;
 * Each word after the identifier will consist only of digits.
 *
 *
 * We will call these two varieties of logs letter-logs and digit-logs.  It is
 * guaranteed that each log has at least one word after its identifier.
 *
 * Reorder the logs so that all of the letter-logs come before any digit-log.
 * The letter-logs are ordered lexicographically ignoring identifier, with the
 * identifier used in case of ties.  The digit-logs should be put in their
 * original order.
 *
 * Return the final order of the logs.
 *
 *
 * Example 1:
 * Input: logs = ["dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit
 * dig","let3 art zero"]
 * Output: ["let1 art can","let3 art zero","let2 own kit dig","dig1 8 1 5
 * 1","dig2 3 6"]
 *
 *
 * Constraints:
 *
 *
 * 0 <= logs.length <= 100
 * 3 <= logs[i].length <= 100
 * logs[i] is guaranteed to have an identifier, and a word after the
 * identifier.
 *
 *
 */
func reorderLogFiles(logs []string) []string {
	digits := make([]string, 0, len(logs))
	res := make([]string, 0, len(logs))

	// log line index and location of first space in log
	m := make(map[string]int)

	for _, line := range logs {
		x := strings.Index(line, " ")
		i := x
		for ; i < len(line); i++ {
			if line[i] == ' ' {
				continue
			}

			if unicode.IsLetter(rune(line[i])) {
				m[line] = x
				res = append(res, line)
				break
			}
		}

		if i >= len(line) {
			digits = append(digits, line)
		}
	}

	sort.Slice(res, func(i, j int) bool {
		a, b := res[i], res[j]
		id1, s1 := a[:m[a]], a[m[a]+1:]
		id2, s2 := b[:m[b]], b[m[b]+1:]

		if s1 == s2 {
			return id1 < id2
		}

		return s1 < s2
	})
	res = append(res, digits...)
	return res
}
