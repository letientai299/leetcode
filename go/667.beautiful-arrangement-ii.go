package main

// Beautiful Arrangement II
//
// Medium
//
// https://leetcode.com/problems/beautiful-arrangement-ii/
//
// Given two integers `n` and `k`, construct a list `answer` that contains `n`
// different positive integers ranging from `1` to `n` and obeys the following
// requirement:
//
// - Suppose this list is `answer = [a1, a2, a3, ... , an]`, then the list `[|a1
// - a2|, |a2 - a3|, |a3 - a4|, ... , |an-1 - an|]` has exactly `k` distinct
// integers.
//
// Return _the list_ `answer`. If there multiple valid answers, return **any of
// them**.
//
// **Example 1:**
//
// ```
// Input: n = 3, k = 1
// Output: [1,2,3]
// Explanation: The [1,2,3] has three different positive integers ranging from 1
// to 3, and the [1,1] has exactly 1 distinct integer: 1
//
// ```
//
// **Example 2:**
//
// ```
// Input: n = 3, k = 2
// Output: [1,3,2]
// Explanation: The [1,3,2] has three different positive integers ranging from 1
// to 3, and the [2,1] has exactly 2 distinct integers: 1 and 2.
//
// ```
//
// **Constraints:**
//
// - `1 <= k < n <= 104`
func constructArray(n int, k int) []int {
	res := make([]int, n)
	l, r := n, 1
	i := 0
	for i < n && k >= 1 {
		if i%2 == 0 {
			res[i] = l
			l--
		} else {
			res[i] = r
			r++
		}
		i++
		k--
	}

	sig := -1
	if i%2 == 0 {
		l = r
		sig = 1
	}
	for ; i < n; i++ {
		res[i] = l
		l += sig
	}

	return res
}
