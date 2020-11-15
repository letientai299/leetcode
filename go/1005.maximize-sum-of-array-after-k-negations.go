package main

import "sort"

/*
 * @lc app=leetcode id=1005 lang=golang
 *
 * [1005] Maximize Sum Of Array After K Negations
 *
 * https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/description/
 *
 * algorithms
 * Easy (50.30%)
 * Total Accepted:    32.4K
 * Total Submissions: 62.1K
 * Testcase Example:  '[4,2,3]\n1'
 *
 * Given an array A of integers, we must modify the array in the following way:
 * we choose an i and replace A[i] with -A[i], and we repeat this process K
 * times in total.  (We may choose the same index i multiple times.)
 *
 * Return the largest possible sum of the array after modifying it in this
 * way.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: A = [4,2,3], K = 1
 * Output: 5
 * Explanation: Choose indices (1,) and A becomes [4,-2,3].
 *
 *
 *
 * Example 2:
 *
 *
 * Input: A = [3,-1,0,2], K = 3
 * Output: 6
 * Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].
 *
 *
 *
 * Example 3:
 *
 *
 * Input: A = [2,-3,-1,5,-4], K = 2
 * Output: 13
 * Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 10000
 * 1 <= K <= 10000
 * -100 <= A[i] <= 100
 *
 *
 */

func largestSumAfterKNegations(a []int, k int) int {
	sort.Ints(a)
	i := 0
	for ; i < len(a) && k > 0 && a[i] < 0; i++ {
		a[i] = -a[i]
		k--
	}

	s := 0
	for _, x := range a {
		s += x
	}

	k %= 2
	if a[i] == 0 {
		k = 0
	}
	if k == 0 {
		return s
	}

	left := 0
	if i > 0 {
		left = a[i-1]
	}

	for i < len(a) && a[i] == 0 {
		i++
	}

	if i >= len(a) {
		s -= 2 * left
		return s
	}

	right := a[i]
	if left == 0 || right < left {
		s -= 2 * right
		return s
	}

	return s - 2*left
}
