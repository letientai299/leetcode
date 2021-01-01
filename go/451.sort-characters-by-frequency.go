package main

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 *
 * https://leetcode.com/problems/sort-characters-by-frequency/description/
 *
 * algorithms
 * Medium (57.93%)
 * Total Accepted:    230.7K
 * Total Submissions: 360.3K
 * Testcase Example:  '"tree"'
 *
 * Given a string, sort it in decreasing order based on the frequency of
 * characters.
 *
 * Example 1:
 *
 * Input:
 * "tree"
 *
 * Output:
 * "eert"
 *
 * Explanation:
 * 'e' appears twice while 'r' and 't' both appear once.
 * So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid
 * answer.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * "cccaaa"
 *
 * Output:
 * "cccaaa"
 *
 * Explanation:
 * Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
 * Note that "cacaca" is incorrect, as the same characters must be together.
 *
 *
 *
 * Example 3:
 *
 * Input:
 * "Aabb"
 *
 * Output:
 * "bbAa"
 *
 * Explanation:
 * "bbaA" is also a valid answer, but "Aabb" is incorrect.
 * Note that 'A' and 'a' are treated as two different characters.
 *
 *
 */
func frequencySort(s string) string {
	m := make(map[byte]int, 128)
	for _, c := range s {
		m[byte(c)] ++
	}

	r := make(map[int][]byte, len(m))
	arr := make([]int, 0, len(m))
	for k, v := range m {
		if len(r[v]) == 0 {
			arr = append(arr, v)
		}
		r[v] = append(r[v], k)
	}
	sort.Ints(arr)
	var sb strings.Builder
	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		for _, c := range r[v] {
			for j := 0; j < v; j++ {
				sb.WriteByte(c)
			}
		}
	}
	return sb.String()
}
