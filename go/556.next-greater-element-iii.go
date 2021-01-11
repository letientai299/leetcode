package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=556 lang=golang
 *
 * [556] Next Greater Element III
 *
 * https://leetcode.com/problems/next-greater-element-iii/description/
 *
 * algorithms
 * Medium (30.71%)
 * Total Accepted:    49K
 * Total Submissions: 153K
 * Testcase Example:  '12'
 *
 * Given a positive integer n, find the smallest integer which has exactly the
 * same digits existing in the integer n and is greater in value than n. If no
 * such positive integer exists, return -1.
 *
 * Note that the returned integer should fit in 32-bit integer, if there is a
 * valid answer but it does not fit in 32-bit integer, return -1.
 *
 *
 * Example 1:
 * Input: n = 12
 * Output: 21
 * Example 2:
 * Input: n = 21
 * Output: -1
 *
 *
 * Constraints:
 * 1 <= n <= 2^31 - 1
 */

func nextGreaterElement(n int) int {
	if n < 10 {
		return -1
	}

	ds := digits(n)
	i := len(ds) - 2
	for ; i >= 0; i-- {
		if ds[i] < ds[i+1] {
			break
		}
	}

	if i == -1 {
		return -1
	}

	j := len(ds) - 1
	for j > i && ds[j] <= ds[i] {
		j--
	}

	ds[i], ds[j] = ds[j], ds[i]

	sort.Ints(ds[i+1:])
	v := int64(0)
	for _, x := range ds {
		v = 10*v + int64(x)
		if v > math.MaxInt32 {
			return -1
		}
	}
	return int(v)
}
