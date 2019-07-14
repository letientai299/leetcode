package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=686 lang=golang
 *
 * [686] Repeated String Match
 *
 * https://leetcode.com/problems/repeated-string-match/description/
 *
 * algorithms
 * Easy (31.49%)
 * Total Accepted:    70.2K
 * Total Submissions: 222.9K
 * Testcase Example:  '"abcd"\n"cdabcdab"'
 *
 * Given two strings A and B, find the minimum number of times A has to be
 * repeated such that B is a substring of it. If no such solution, return -1.
 *
 * For example, with A = "abcd" and B = "cdabcdab".
 *
 * Return 3, because by repeating A three times (“abcdabcdabcd”), B is a
 * substring of it; and B is not a substring of A repeated two times
 * ("abcdabcd").
 *
 * Note:
 * The length of A and B will be between 1 and 10000.
 *
 */
func repeatedStringMatch(A string, B string) int {
	if strings.Contains(A, B) {
		return 1
	} else if strings.Contains(A+A, B) {
		return 2
	}

	// check the head
	i := strings.Index(B, A)
	if i == -1 {
		return -1
	}

	head := B[:i]
	if !strings.HasSuffix(A, head) {
		return -1
	}

	total := 1
	if head != "" {
		total++
	}

	B = B[i+len(A):]

	for len(B) >= len(A) {
		if strings.HasPrefix(B, A) {
			total++
			B = B[len(A):]
		} else {
			return -1
		}
	}

	if B == "" {
		return total
	}

	if strings.HasPrefix(A, B) {
		return total + 1
	} else {
		return -1
	}
}
