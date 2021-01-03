package main

import "sort"

/*
 * @lc app=leetcode id=658 lang=golang
 *
 * [658] Find K Closest Elements
 *
 * https://leetcode.com/problems/find-k-closest-elements/description/
 *
 * algorithms
 * Medium (39.23%)
 * Total Accepted:    131.5K
 * Total Submissions: 315.6K
 * Testcase Example:  '[1,2,3,4,5]\n4\n3'
 *
 * Given a sorted integer array arr, two integers k and x, return the k closest
 * integers to x in the array. The result should also be sorted in ascending
 * order.
 *
 * An integer a is closer to x than an integer b if:
 *
 *
 * |a - x| < |b - x|, or
 * |a - x| == |b - x| and a < b
 *
 *
 *
 * Example 1:
 * Input: arr = [1,2,3,4,5], k = 4, x = 3
 * Output: [1,2,3,4]
 * Example 2:
 * Input: arr = [1,2,3,4,5], k = 4, x = -1
 * Output: [1,2,3,4]
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= arr.length
 * 1 <= arr.length <= 10^4
 * Absolute value of elements in the array and x will not exceed 10^4
 *
 *
 */
func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	if n <= k {
		return arr
	}

	if x <= arr[0] {
		return arr[:k]
	}

	if x >= arr[n-1] {
		return arr[n-k:]
	}

	i := sort.Search(n, func(i int) bool {
		return x <= arr[i]
	})

	if arr[i] > x && arr[i]-x >= x-arr[i-1] {
		i--
	}

	if i == 0 {
		return arr[:k]
	}

	j := i + 1

	if j == n {
		return arr[n-k:]
	}

	if j-i >= k {
		return arr[i:j]
	}

	if j-i == k-1 {
		return arr[i : j+1]
	}

	for j-i < k && j < n && i >= 1 {
		a := x - arr[i-1]
		b := arr[j] - x
		if a <= b {
			i--
			continue
		}
		j++
	}

	if i == 0 {
		return arr[:k]
	}

	if j == n {
		return arr[n-k:]
	}

	return arr[i:j]
}
