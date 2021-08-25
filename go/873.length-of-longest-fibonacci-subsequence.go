package main

// Length of Longest Fibonacci Subsequence
//
// Medium
//
// https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/
//
// A sequence `x1, x2, ..., xn` is _Fibonacci-like_ if:
//
// - `n >= 3`
// - `xi + xi+1 == xi+2` for all `i + 2 <= n`
//
// Given a **strictly increasing** array `arr` of positive integers forming a
// sequence, return _the **length** of the longest Fibonacci-like subsequence
// of_ `arr`. If one does not exist, return `0`.
//
// A **subsequence** is derived from another sequence `arr` by deleting any
// number of elements (including none) from `arr`, without changing the order of
// the remaining elements. For example, `[3, 5, 8]` is a subsequence of `[3, 4,
// 5, 6, 7, 8]`.
//
// **Example 1:**
//
// ```
// Input: arr = [1,2,3,4,5,6,7,8]
// Output: 5
// Explanation: The longest subsequence that is fibonacci-like: [1,2,3,5,8].
// ```
//
// **Example 2:**
//
// ```
// Input: arr = [1,3,7,11,12,14,18]
// Output: 3
// Explanation: The longest subsequence that is fibonacci-like: [1,11,12],
// [3,11,14] or [7,11,18].
// ```
//
// **Constraints:**
//
// - `3 <= arr.length <= 1000`
// - `1 <= arr[i] < arr[i + 1] <= 109`

// O(n^3) solution, fastest
func lenLongestFibSubseq(arr []int) int {
	m := make(map[int]int)
	for i, v := range arr {
		m[v] = i
	}
	best := 0

	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			x := i
			y := j
			count := 2
			for {
				if k, ok := m[arr[x]+arr[y]]; !ok {
					break
				} else {
					count++
					x, y = y, k
				}
			}

			if count > 2 && count > best {
				best = count
			}
		}
	}

	return best
}

/*
// DP solution, O(n^2) but is slower due to mem access
func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	m := make(map[int]int)
	for i, v := range arr {
		m[v] = i
	}

	best := 2
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			k, ok := m[arr[j]-arr[i]]
			if !ok || dp[i][k] == 0 {
				dp[j][i] = 2
				continue

			}

			dp[j][i] = 1 + dp[i][k]
			if best < dp[j][i] {
				best = dp[j][i]
			}
		}
	}

	if best == 2 {
		return 0
	}

	return best
}
*/
