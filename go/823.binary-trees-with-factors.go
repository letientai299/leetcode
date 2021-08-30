package main

import "sort"

// Binary Trees With Factors
//
// Medium
//
// https://leetcode.com/problems/binary-trees-with-factors/
//
// Given an array of unique integers, `arr`, where each integer `arr[i]` is
// strictly greater than `1`.
//
// We make a binary tree using these integers, and each number may be used for
// any number of times. Each non-leaf node's value should be equal to the
// product of the values of its children.
//
// Return _the number of binary trees we can make_. The answer may be too large
// so return the answer **modulo** `109 + 7`.
//
// **Example 1:**
//
// ```
// [2], [4], [4, 2, 2]
// ```
//
// **Example 2:**
//
// ```
// [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2]
// ```
//
// **Constraints:**
//
// - `1 <= arr.length <= 1000`
// - `2 <= arr[i] <= 109`
// - All the values of `arr` are **unique**.
func numFactoredBinaryTrees(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 1
	}

	m := make(map[int]int)
	sort.Ints(arr)
	max := arr[n-1]

	for i := 0; i < n; i++ {
		x := arr[i]
		m[x]++
		if x*x <= max {
			m[x*x] = m[x] * m[x]
		}
		for j := i - 1; j >= 0; j-- {
			y := arr[j]
			if x*y <= max {
				m[x*y] += 2 * m[x] * m[y]
			}
		}
	}

	all := 0
	for _, x := range arr {
		if v, ok := m[x]; ok {
			all += v
			all %= 1e9 + 7
		}
	}
	return all
}
