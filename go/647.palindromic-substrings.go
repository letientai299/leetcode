package main

import "fmt"

/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 *
 * https://leetcode.com/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (58.50%)
 * Total Accepted:    230.1K
 * Total Submissions: 373.7K
 * Testcase Example:  '"abc"'
 *
 * Given a string, your task is to count how many palindromic substrings in
 * this string.
 *
 * The substrings with different start indexes or end indexes are counted as
 * different substrings even they consist of same characters.
 *
 * Example 1:
 *
 *
 * Input: "abc"
 * Output: 3
 * Explanation: Three palindromic strings: "a", "b", "c".
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "aaa"
 * Output: 6
 * Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
 *
 *
 *
 *
 * Note:
 *
 *
 * The input string length won't exceed 1000.
 *
 *
 *
 */

// 0ms, =))
func countSubstrings(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	indexes := make(map[int]struct{})
	same := 1
	cnt := 1

	for i := 1; i < n; i++ {
		cnt++
		if s[i] == s[i-1] {
			cnt += same
			same++
		} else {
			same = 1
		}

		next := make(map[int]struct{}, len(indexes))
		for j := range indexes {
			if s[j] != s[i] {
				ok := true
				for k := j + 1; k <= (i+j+1)/2; k++ {
					if s[k] != s[i+j+1-k] {
						ok = false
						break
					}
				}

				if ok {
					next[j] = struct{}{}
				}
				continue
			}

			cnt++
			if j > 0 {
				next[j-1] = struct{}{}
			}
		}

		if s[i] != s[i-1] {
			next[i-1] = struct{}{}
		}

		indexes = next
	}

	return cnt
}

// 24ms, 29.36%
func countSubstrings_slow(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	indexes := make(map[int]bool)
	indexes[0] = true

	cnt := 0
	for i := 1; i < n; i++ {
		next := make(map[int]bool, len(indexes))
		for j, sameChar := range indexes {
			if s[j] != s[i] {
				continue
			}

			cnt++
			if sameChar {
				next[j] = true
			}

			if j > 0 {
				next[j-1] = false
			}
		}

		if s[i] != s[i-1] {
			next[i-1] = false
		}

		next[i] = true
		indexes = next
	}

	return cnt + n
}

func main() {
	fmt.Println(countSubstrings_slow("abbabbba"))
}
