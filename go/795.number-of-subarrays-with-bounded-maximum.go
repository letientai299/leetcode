package main

import "fmt"

/*
 * @lc app=leetcode id=795 lang=golang
 *
 * [795] Number of Subarrays with Bounded Maximum
 *
 * https://leetcode.com/problems/number-of-subarrays-with-bounded-maximum/description/
 *
 * algorithms
 * Medium (44.47%)
 * Total Accepted:    21.7K
 * Total Submissions: 46K
 * Testcase Example:  '[2,1,4,3]\n2\n3'
 *
 * We are given an array A of positive integers, and two positive integers L
 * and R (L <= R).
 *
 * Return the number of (contiguous, non-empty) subarrays such that the value
 * of the maximum array element in that subarray is at least L and at most R.
 *
 *
 * Example :
 * Input:
 * A = [2, 1, 4, 3]
 * L = 2
 * R = 3
 * Output: 3
 * Explanation: There are three subarrays that meet the requirements: [3], [2,
 * 1], [3].
 *
 *
 * Note:
 *
 *
 * L, R  and A[i] will be an integer in the range [0, 10^9].
 * The length of A will be in the range of [1, 50000].
 *
 *
 */

// TODO (tai): fix this
func numSubarrayBoundedMax(a []int, l int, r int) int {
	n := len(a)
	if n == 0 {
		return 0
	}
	within := func(x int) bool { return l <= x && x <= r }
	dp := make([][3]int, n)
	res := 0
	if within(a[0]) {
		res++
		dp[0] = [3]int{1, a[0], 1}
	} else {
		dp[0] = [3]int{0, a[0], 0}
	}

	for i := 1; i < n; i++ {
		x := a[i]
		prev := dp[i-1]

		ok := within(x)

		dp[i] = [3]int{0, x}
		if ok {
			res++
			dp[i][0] ++
			dp[i][2] = 1
		} else if x > prev[1] {
			continue
		}

		if prev[0] == 0 {
			continue
		}

		res += prev[0]
		dp[i][0] += prev[0]
		if x < prev[1] {
			dp[i][1] = prev[1]
		}
	}

	fmt.Println(l, r)
	for _, v := range a {
		fmt.Print(v, "\t")
	}
	fmt.Println()

	for i := 0; i < len(dp[0]); i++ {
		for _, p := range dp {
			fmt.Print(p[i], "\t")
		}
		fmt.Println()
	}

	return res
}
