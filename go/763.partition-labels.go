package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 *
 * https://leetcode.com/problems/partition-labels/description/
 *
 * algorithms
 * Medium (72.97%)
 * Total Accepted:    206.1K
 * Total Submissions: 265.3K
 * Testcase Example:  '"ababcbacadefegdehijhklij"'
 *
 * A string S of lowercase English letters is given. We want to partition this
 * string into as many parts as possible so that each letter appears in at most
 * one part, and return a list of integers representing the size of these
 * parts.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: S = "ababcbacadefegdehijhklij"
 * Output: [9,7,8]
 * Explanation:
 * The partition is "ababcbaca", "defegde", "hijhklij".
 * This is a partition so that each letter appears in at most one part.
 * A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it
 * splits S into less parts.
 *
 *
 *
 *
 * Note:
 *
 *
 * S will have length in range [1, 500].
 * S will consist of lowercase English letters ('a' to 'z') only.
 *
 *
 *
 *
 */

func partitionLabels(s string) []int {
	parts := make([][2]int, 26)
	for i, c := range s {
		x := c - 'a'
		if parts[x][0] == 0 {
			parts[x][0] = i + 1
		}
		if parts[x][1] == 0 || parts[x][1] < i+1 {
			parts[x][1] = i + 1
		}
	}

	sort.Slice(parts, func(i, j int) bool {
		a, b := parts[i], parts[j]
		if a[0] < b[0] {
			return true
		}

		if a[0] == b[0] {
			return a[1] < b[1]
		}

		return false
	})

	i, j := 0, 0
	var res []int
	for ; i < len(parts) && j < len(parts); j++ {
		if parts[j][0] == 0 {
			continue
		}

		if parts[i][0] == 0 {
			parts[i] = parts[j]
			res = append(res, parts[i][1]-parts[i][0]+1)
			continue
		}

		if parts[i][1] > parts[j][0] {
			if parts[i][1] < parts[j][1] {
				parts[i][1] = parts[j][1]
			}
			if len(res) == 0 {
				res = append(res, parts[i][1]-parts[i][0]+1)
			} else {
				res[len(res)-1] = parts[i][1] - parts[i][0] + 1
			}
		} else {
			i++
			parts[i] = parts[j]
			res = append(res, parts[i][1]-parts[i][0]+1)
		}
	}

	return res
}
