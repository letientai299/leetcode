package main

/*
 * @lc app=leetcode id=941 lang=golang
 *
 * [941] Valid Mountain Array
 *
 * https://leetcode.com/problems/valid-mountain-array/description/
 *
 * algorithms
 * Easy (35.50%)
 * Total Accepted:    93.2K
 * Total Submissions: 288.6K
 * Testcase Example:  '[2,1]'
 *
 * Given an array of integers arr, return true if and only if it is a valid
 * mountain array.
 *
 * Recall that arr is a mountain array if and only if:
 *
 *
 * arr.length >= 3
 * There exists some i with 0 < i < arr.length - 1 such that:
 *
 * arr[0] < arr[1] < ... < arr[i - 1] < A[i]
 * arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
 *
 *
 *
 *
 *
 * Example 1:
 * Input: arr = [2,1]
 * Output: false
 * Example 2:
 * Input: arr = [3,5,5]
 * Output: false
 * Example 3:
 * Input: arr = [0,3,2,1]
 * Output: true
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 10^4
 * 0 <= arr[i] <= 10^4
 *
 *
 */
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	down := false
	hasTip := false
	for i := 1; i < len(arr)-1; i++ {
		a, b, c := arr[i-1], arr[i], arr[i+1]
		if a == b || b == c {
			return false
		}

		if a < b && b > c {
			hasTip = true
			down = true
			continue
		}

		if (a < b || b < c) && down {
			return false
		}

		if (a > b || b > c) && !down {
			return false
		}
	}

	return down && hasTip
}
