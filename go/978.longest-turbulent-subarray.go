package main

/*
 * @lc app=leetcode id=978 lang=golang
 *
 * [978] Longest Turbulent Subarray
 *
 * https://leetcode.com/problems/longest-turbulent-subarray/description/
 *
 * algorithms
 * Medium (45.95%)
 * Total Accepted:    35.5K
 * Total Submissions: 76K
 * Testcase Example:  '[9,4,2,10,7,8,8,1,9]'
 *
 * Given an integer array arr, return the length of a maximum size turbulent
 * subarray of arr.
 *
 * A subarray is turbulent if the comparison sign flips between each adjacent
 * pair of elements in the subarray.
 *
 * More formally, a subarray [arr[i], arr[i + 1], ..., arr[j]] of arr is said
 * to be turbulent if and only if:
 *
 *
 * For i <= k < j:
 *
 *
 * arr[k] > arr[k + 1] when k is odd, and
 * arr[k] < arr[k + 1] when k is even.
 *
 *
 * Or, for i <= k < j:
 *
 * arr[k] > arr[k + 1] when k is even, and
 * arr[k] < arr[k + 1] when k is odd.
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: arr = [9,4,2,10,7,8,8,1,9]
 * Output: 5
 * Explanation: arr[1] > arr[2] < arr[3] > arr[4] < arr[5]
 *
 *
 * Example 2:
 *
 *
 * Input: arr = [4,8,12,16]
 * Output: 2
 *
 *
 * Example 3:
 *
 *
 * Input: arr = [100]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 4 * 10^4
 * 0 <= arr[i] <= 10^9
 *
 *
 */

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	if n < 2 {
		return 1
	}

	dp := 1
	if arr[1] != arr[0] {
		dp++
	}
	best := dp
	for i := 2; i < n; i++ {
		if arr[i] == arr[i-1] {
			dp = 1
			continue
		}

		if (arr[i] > arr[i-1] && arr[i-1] < arr[i-2]) ||
			(arr[i] < arr[i-1] && arr[i-1] > arr[i-2]) {
			dp++
		} else {
			dp = 2
		}

		if best < dp {
			best = dp
		}
	}

	return best
}
